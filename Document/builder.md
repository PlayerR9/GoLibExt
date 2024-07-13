package Document

import (
	fss "github.com/PlayerR9/MyGoLib/Display/Section"
	fsp "github.com/PlayerR9/MyGoLib/Formatting/FString"
)

// Builder is a type that represents a builder for creating formatted strings.
type Builder struct {
	// printer is the printer used by the builder.
	printer *fsp.StdPrinter

	// Traversor is the traversor used by the builder.
	*fsp.Traversor
}

// NewBuilder creates a new builder with the default configuration.
//
// Returns:
//   - *Builder: A pointer to the new builder.
func NewBuilder() *Builder {
	printer := fsp.NewStdPrinterFromConfig(
		fsp.NewIndentConfig("   ", 0),
	)

	return &Builder{
		printer:   printer,
		Traversor: printer.TraversorOf(),
	}
}

// str -> line -> section -> page -> document

// Apply is a function that applies the element to the builder.
//
// Parameters:
//   - elem: The element to apply.
//
// Returns:
//   - error: An error if the application fails.
func (b *Builder) Apply(elem fsp.FStringer) error {
	return elem.FString(b.Traversor)
}

func (d *Builder) AppendSection(section fss.Sectioner) {

}

func (d *Builder) AppendSections(sections []fss.Sectioner) {

}

func (d *Builder) AddPage(page *Page) {

}

func (d *Builder) AddPages(pages []*Page) {

}

// Build is a function that builds the document.
//
// Returns:
//   - *Document: The document.
//   - error: An error if the building fails.
func (b *Builder) Build() (*DocumentViewer, error) {
	rawPages := b.printer.GetPages()

	b.Reset()

	return MakeDocument(rawPages)
}

// Reset is a function that resets the builder.
func (b *Builder) Reset() {
	b.printer.Clean()
	b.Traversor.Clean()

	b.printer = fsp.NewStdPrinterFromConfig(
		fsp.NewIndentConfig("   ", 0),
	)

	b.Traversor = b.printer.TraversorOf()
}
