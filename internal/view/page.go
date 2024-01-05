package view

import "math"

type Page struct {
	// 合計レコード数
	Total int
	// 現在のページ
	Page int
	// 1ページに含まれるレコード数
	Size int
	// 現在のパス
	CurrentURL string
}

func (p Page) HasPrevious() bool {
	return p.Page > 1
}

func (p Page) HasNext() bool {
	return p.Page+1 < p.GetTotalPages()
}

func (p Page) IsFirst() bool {
	return !p.HasPrevious()
}

func (p Page) IsLast() bool {
	return !p.HasNext()
}

func (p Page) GetTotalPages() int {
	return int(math.Ceil(float64(p.Total) / float64(p.Size)))
}
