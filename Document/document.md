package Document

import (
	ffs "github.com/PlayerR9/MyGoLib/Formatting/FString"
)

type Document struct {
	pages []*Page
}

func (d *Document) FString(trav *ffs.Traversor) error {
	for _, page := range d.pages {
		err := page.FString(trav)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDocument() *Document {
	return &Document{
		pages: make([]*Page, 0),
	}
}

func (d *Document) AddPage(page *Page) {
	d.pages = append(d.pages, page)
}
