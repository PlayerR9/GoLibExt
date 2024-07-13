package Document

import (
	"fmt"

	fss "github.com/PlayerR9/MyGoLib/Display/Section"
	ut "github.com/PlayerR9/MyGoLib/Units/Tray"
	ue "github.com/PlayerR9/MyGoLib/Units/errors"
)

// DocumentViewer is a type that represents a document.
type DocumentViewer struct {
	*ut.DynamicTray[*Page, *RenderedPage]
}

/*
func (d *Document) RenderAllPages() error {
	allRenderedPages := make([]*RenderedPage, 0)

	page, err := d.pages.Read()
	if err != nil {
		return fmt.Errorf("could not read page: %w", err)
	}

	renders, err := page.View()
	if err != nil {
		return fmt.Errorf("could not view page: %w", err)
	}

	currRender := NewRenderedPage(renders[0], 1, 1)

	allRenderedPages = append(allRenderedPages, currRender)

	for i, todo := range renders[1:] {
		otherRender := NewRenderedPage(todo, 1, i+2)

		allRenderedPages = append(allRenderedPages, otherRender)
	}

	// TODO: Finish rendering the pages.

	return nil
}
*/

// MakeDocument creates a new document.
//
// Parameters:
//   - rawPages: The raw pages to add to the document.
//
// Returns:
//   - *Document: A pointer to the newly created document.
func MakeDocument(rawPages [][][][]string) (*DocumentViewer, error) {
	var pages []*Page

	for _, rawPage := range rawPages {
		page := NewPage()

		for i, rawSection := range rawPage {
			mlt := new(fss.MultilineText)

			err := mlt.FromTextBlock(rawSection)
			if err != nil {
				return nil, ue.NewErrAt(i, "section", err)
			}

			page.AddSection(mlt)
		}

		pages = append(pages, page)
	}

	iter := ut.NewDynamicTray(
		pages,
		PageRenderTransition,
	)

	return &DocumentViewer{iter}, nil
}

/*

// String returns the string representation of the document.
//
// Returns:
//   - string: The string representation of the document.
func (d *Document) String() string {
	return strings.Join(d.lines, "\n")
}

func (d *Document) Tmp() []string {
	return d.lines
}

// FString returns the formatted string representation of the document.
//
// Parameters:
//   - indentLevel: The level of indentation.
//
// Returns:
//   - []string: The formatted string representation of the document.
func (d *Document) FString(trav *ffs.Traversor) error {
	err := trav.AddLines(d.lines)
	if err != nil {
		return err
	}

	return nil
}

// NewDocument creates a new document.
//
// Parameters:
//   - sentences: The sentences to add to the document.
//
// Returns:
//   - *Document: A pointer to the newly created document.
//
// Behaviors:
//   - The sentences are separated by a space and on the same line.
func NewDocument(sentences ...string) *Document {
	d := &Document{
		lines: make([]string, 0),
	}

	d.AddLine(sentences...)

	return d
}

// AddLine adds sentences to the document separated by a space.
// The line is split by the newline character.
//
// Parameters:
//   - line: The line to add.
//
// Returns:
//   - *Document: A pointer to the document. This allows for chaining.
//
// Example:
//   - AddLine("Hello,", "world!")
//   - AddLine("This is a sentence.")
func (d *Document) AddLine(sentences ...string) *Document {
	if len(sentences) == 0 {
		return d
	}

	var builder strings.Builder

	if strings.HasSuffix(sentences[0], "\n") {
		builder.WriteString(strings.TrimSuffix(sentences[0], "\n"))

		d.lines = append(d.lines, strings.Split(builder.String(), "\n")...)
	} else {
		builder.WriteString(sentences[0])
	}

	for _, sentence := range sentences[1:] {
		if strings.HasSuffix(sentence, "\n") {
			builder.WriteRune(' ')
			builder.WriteString(strings.TrimSuffix(sentence, "\n"))

			d.lines = append(d.lines, strings.Split(builder.String(), "\n")...)
		} else {
			builder.WriteRune(' ')
			builder.WriteString(sentence)
		}
	}

	if builder.Len() != 0 {
		d.lines = append(d.lines, strings.Split(builder.String(), "\n")...)
	}

	return d
}

*/

func PageRenderTransition(page *Page) *ut.SimpleTray[*RenderedPage] {
	renders, err := page.View()
	if err != nil {
		panic(fmt.Errorf("could not view page: %w", err))
	}

	return ut.NewSimpleTray(renders)
}
