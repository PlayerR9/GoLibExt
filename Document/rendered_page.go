package Document

import (
	fss "github.com/PlayerR9/display/Displays/Section"
)

// RenderedPage is a type that represents a page of a document.
type RenderedPage struct {
	// render is the render of the page.
	render *fss.Render

	// pageNumber is the page number of the page.
	pageNumber int

	// subPageNumber is the sub-page number of the page.
	subPageNumber int
}

func NewRenderedPage(render *fss.Render, pageNumber, subPageNumber int) *RenderedPage {
	return &RenderedPage{
		render:        render,
		pageNumber:    pageNumber,
		subPageNumber: subPageNumber,
	}
}

/*
func (p *RenderedPage) View() ([]*fss.Render, error) {
	renders, err := p.sections[0].ApplyRender(p.width, p.height)
	if err != nil {
		return nil, fmt.Errorf("could not apply render: %w", err)
	}

	switch len(renders) {
	case 0:
		// No renders were created.
		return nil, errors.New("no renders were created")
	case 1:
		// The page can be rendered in one go.

		// TODO: Render the page.

		return nil, nil
	default:
		// The page needs to be rendered in parts.

		// TODO: Render the first page.

		return renders[1:], nil
	}
}
*/
