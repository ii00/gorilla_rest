package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "cocoa",
		Price: 1.00,
		SKU:   "abs-fdf-gfg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
