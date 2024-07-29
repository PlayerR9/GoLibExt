package Document

/*
// DocumentPrinter is a type that represents a printer for a document.
type DocumentPrinter struct {
	// The name of the document.
	Name string

	// The document to print.
	Doc *DocumentBuilder

	// IfEmpty is the string to print if the document is empty.
	IfEmpty string
}

// FString is a helper function that prints a document.
//
// Format:
//
//	<name>:
//	  	<document>
//
// or
//
//	<name>: <ifEmpty>
//
// Parameters:
//   - trav: The traversor to use for printing.
//
// Returns:
//   - error: An error if the printing fails.
//
// Behaviors:
//   - If the document is nil, the function prints the IfEmpty string.
//   - If the document is not nil, the function prints the document.
func (dp *DocumentPrinter) FString(trav *ffs.Traversor) error {
	err := trav.AppendStrings(dp.Name, ":")
	if err != nil {
		return err
	}

	if dp.Doc == nil {
		trav.AppendRune(' ')

		err := trav.AppendString(dp.IfEmpty)
		if err != nil {
			return err
		}

		trav.AcceptHalfLine()
	} else {
		trav.AcceptHalfLine()

		err = dp.Doc.FString(trav.IncreaseIndent(1))
		if err != nil {
			return err
		}
	}

	return nil
}

// NewDocumentPrinter is a function that creates a new document printer.
//
// Parameters:
//   - name: The name of the document.
//   - doc: The document to print.
//   - ifEmpty: The string to print if the document is empty.
//
// Returns:
//   - *DocumentPrinter: A pointer to the new document printer.
func NewDocumentPrinter(name string, doc *DocumentBuilder, ifEmpty string) *DocumentPrinter {
	return &DocumentPrinter{
		Name:    name,
		Doc:     doc,
		IfEmpty: ifEmpty,
	}
}
*/
