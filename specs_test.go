package apispecs

import (
	"fmt"
	"testing"

	"github.com/grokify/swaggman/openapi3"
	"github.com/grokify/swaggman/openapi3/validate"
)

var specTests = []struct {
	filepath string
	title    string
}{
	{"medium/medium-api-v1.0.0_openapi-v3.0.yaml", "Medium API"},
}

// TestSpecs test reading specs.
func TestSpecs(t *testing.T) {
	for _, tt := range specTests {
		spec, err := openapi3.ReadFile(tt.filepath, true)
		if err != nil {
			t.Errorf("openapi3.ReadFile('%s', true) Error [%s]", tt.filepath, err)
		}
		if tt.title != spec.Info.Title {
			t.Errorf("openapi3.ReadFile('%s', true) Want [%s] Got [%s]", tt.filepath, tt.title, spec.Info.Title)
		}
		_, err = validate.ValidateMore(spec)
		if err != nil {
			t.Errorf("validate.ValidateMore(spec) Error [%s]", err.Error())
		} else {
			fmt.Printf("VALID [%s]\n", tt.filepath)
		}
		fmt.Printf("TITLE [%s]\n", spec.Info.Title)
	}
}
