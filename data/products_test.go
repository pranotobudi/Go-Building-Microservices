package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "bud",
		Price: 2.0,
		SKU:   "abc-abc-abc",
	}
	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
