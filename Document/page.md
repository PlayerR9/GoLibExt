package Document

import (
	"errors"
	"fmt"

	ffs "github.com/PlayerR9/MyGoLib/Formatting/FString"
	fss "github.com/PlayerR9/display/Displays/Section"
)

const (
	// DefaultPageWidth is the default width of a page.
	DefaultPageWidth int = 80

	// DefaultPageHeight is the default height of a page.
	DefaultPageHeight int = 24
)

// Page is a type that represents a page of a document.
type Page struct {
	// sections are the sections of the page.
	sections []fss.Sectioner

	// width is the width of the page.
	width int

	// height is the height of the page.
	height int
}

// FString implements the FStringer interface.
func (p *Page) FString(trav *ffs.Traversor) error {
	if trav == nil {
		return nil
	}

	return nil
}

func NewPage() *Page {
	return &Page{
		sections: make([]fss.Sectioner, 0),
		width:    DefaultPageWidth,
		height:   DefaultPageHeight,
	}
}

func (p *Page) AddSection(section fss.Sectioner) *Page {
	p.sections = append(p.sections, section)

	return p
}

func (p *Page) View() ([]*RenderedPage, error) {
	renders, err := p.sections[0].ApplyRender(p.width, p.height)
	if err != nil {
		return nil, fmt.Errorf("could not apply render: %w", err)
	}

	if len(renders) == 0 {
		// No renders were created.
		return nil, errors.New("no renders were created")
	}

	// FIXME: This is a temporary fix.

	sol := make([]*RenderedPage, 0, len(renders))

	for i, render := range renders {
		sol = append(sol, NewRenderedPage(render, 1, i+1))
	}

	return sol, nil
}
