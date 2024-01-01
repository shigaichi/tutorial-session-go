package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/glebarez/sqlite"
	"github.com/shigaichi/tutorial-session-go/internal/app"
	"github.com/shigaichi/tutorial-session-go/internal/domain/repository"
	"github.com/shigaichi/tutorial-session-go/internal/domain/service"
	log "github.com/shigaichi/tutorial-session-go/internal/logger"
	"github.com/shigaichi/tutorial-session-go/internal/migrate"
	"github.com/shigaichi/tutorial-session-go/internal/route"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func main() {
	err := startServer()
	if err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}

func startServer() error {
	db, err := gorm.Open(sqlite.Open("db/db.sqlite3"), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, "failed to open db")
	}
	err = migrate.Migrate(db)
	if err != nil {
		return errors.Wrap(err, "failed to migration")
	}

	ar := repository.NewAccountRepositoryImpl(db)
	as := service.NewAccountServiceImpl(ar)
	lh := app.NewLoginHandler(as)
	ach := app.NewAccountCreteHandler(as)

	a, err := route.NewInitRoute(lh, ach).InitRouting()
	if err != nil {
		return errors.Wrap(err, "failed to init routing")
	}
	srv := http.Server{
		Addr:              ":8080",
		Handler:           a,
		ReadHeaderTimeout: 3 * time.Minute,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				log.Info("shutting down server", zap.Error(err))
			} else {
				log.Fatal("failed to start server", zap.Error(err))
			}
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to shutdown server")
	}
	log.Info("Shutting down")
	return nil
}
