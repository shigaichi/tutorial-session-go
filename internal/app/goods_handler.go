package app

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/csrf"
	"github.com/shigaichi/tutorial-session-go/internal/domain/service"
	log "github.com/shigaichi/tutorial-session-go/internal/logger"
	"github.com/shigaichi/tutorial-session-go/internal/middleware"
	"github.com/shigaichi/tutorial-session-go/internal/view"
	"go.uber.org/zap"
)

type Goods interface {
	ShowGoods(w http.ResponseWriter, r *http.Request)
}

type GoodsHandler struct {
	gs service.GoodsService
	cs service.CategoryService
}

func NewGoodsHandler(gs service.GoodsService, cs service.CategoryService) GoodsHandler {
	return GoodsHandler{gs: gs, cs: cs}
}

func (h GoodsHandler) ShowGoods(w http.ResponseWriter, r *http.Request) {
	categoryId, err := strconv.Atoi(r.FormValue("categoryId"))
	if err != nil {
		categoryId = 1
	}

	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.FormValue("pageSize"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	goods, count, err := h.gs.FindByCategoryId(categoryId, page, pageSize)
	if err != nil {
		h.handleInternalError(w, err)
		return
	}

	categories, err := h.cs.FindAll()
	if err != nil {
		h.handleInternalError(w, err)
		return
	}

	store := middleware.SessionStore
	session, err := store.Get(r, middleware.SessionName)
	if err != nil {
		h.handleInternalError(w, err)
		return
	}

	email := session.Values["email"]

	p := view.Page{
		Total:      int(count),
		Page:       page,
		Size:       pageSize,
		CurrentURL: r.URL.Path,
	}

	t := template.Must(
		template.New("template.gohtml").Funcs(
			template.FuncMap{"minus": view.Minus, "plus": view.Plus, "sequence": view.Sequence},
		).ParseFiles("templates/layout/template.gohtml", "templates/goods/showGoods.gohtml", "templates/pgnt/fragment.gohtml", "templates/layout/footer.gohtml"),
	)

	if err := t.Execute(w, map[string]interface{}{csrf.TemplateTag: csrf.TemplateField(r), "title": "Item List Page", "email": email, "categories": categories, "goods": goods, "page": p}); err != nil {
		h.handleInternalError(w, err)
		return
	}
}

func (h GoodsHandler) handleInternalError(w http.ResponseWriter, err error) {
	log.Error("goods_handler", zap.Error(err))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
