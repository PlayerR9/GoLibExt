package Document

import (
	"testing"
)

func TestFString(t *testing.T) {
	builder := NewBuilder()

	err := builder.AddLine("Hello")
	if err != nil {
		t.Fatalf("Expected no error, but got %s", err.Error())
	}

	err = builder.AddLine("World")
	if err != nil {
		t.Fatalf("Expected no error, but got %s", err.Error())
	}

	_, err = builder.Build()
	if err != nil {
		t.Fatalf("Expected no error, but got %s", err.Error())
	}

	/*
		expected := make([]string, 0)

		for _, line := range trav.GetLines() {
			expected = append(expected, line.GetLines()...)
		}

		var builder Builder

		builder.SetSeparator(nil)

		actual := builder.Build().Apply(doc.Tmp())

		if len(actual) != len(expected) {
			// DEBUG: Show the expected and actual values.
			fmt.Println("Expected:", expected)

			fmt.Println("Actual:", actual)

			t.Errorf("Expected: %v, Actual: %v", expected, actual)
		}

		for i := range actual {
			if actual[i] != expected[i] {
				t.Errorf("Expected: %v, Actual: %v", expected, actual)
			}
		}
	*/
}
