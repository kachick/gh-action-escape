package normalizer

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerateDelimiter(t *testing.T) {
	generator := &Base64DelimiterGenerator{}
	delimiterLength := generator.EncodedLength(ByteSizeFromGitHubDoc)
	got, err := generator.Generate(ByteSizeFromGitHubDoc)
	if err != nil {
		t.Fatalf("delimiter generation failed: %v", err)
	}
	if len(got) != delimiterLength {
		t.Errorf("generated wrong string %q", got)
	}
	for i := 0; i < 1000; i++ {
		last, err := generator.Generate(ByteSizeFromGitHubDoc)
		if err != nil {
			t.Fatalf("delimiter generation failed: %v", err)
		}
		if got == last {
			t.Errorf("conflicted in short loop %q", last)
		}
		if len(last) != delimiterLength {
			t.Errorf("generated wrong string %q", last)
		}
	}
}

type fixedDelimiterGenerator struct {
	DelimiterGeneratorable

	Fixed string
}

func (f *fixedDelimiterGenerator) Generate(byteSize int) (string, error) {
	return f.Fixed, nil
}

func TestNormalizer(t *testing.T) {
	const example = `{
		"productPath": "out/depop-0.0.0.11.zip",
		"includesPaths": [
			"manifest.json",
			"README.md",
			"LICENSE",
			"assets/icon-sadness-star.png",
			"static/style.css",
			"static/options.css",
			"static/options.html",
			"dist/options.js",
			"dist/primer.css",
			"dist/main.js"
		]
	}`
	nr := DefaultNormalizer()
	_, err := nr.Normalize("report_multiline_json", example, ByteSizeFromGitHubDoc)
	if err != nil {
		t.Fatalf("normalization failed: %v", err)
	}

	nr = &Normalizer{
		DelimiterGenerator: &fixedDelimiterGenerator{Fixed: "QtrsiRp5MOu48A4v9J9z"},
	}

	got, err := nr.Normalize("report_multiline_json", example, ByteSizeFromGitHubDoc)
	if err != nil {
		t.Fatalf("normalization failed: %v", err)
	}

	const want = `report_multiline_json<<QtrsiRp5MOu48A4v9J9z
{
		"productPath": "out/depop-0.0.0.11.zip",
		"includesPaths": [
			"manifest.json",
			"README.md",
			"LICENSE",
			"assets/icon-sadness-star.png",
			"static/style.css",
			"static/options.css",
			"static/options.html",
			"dist/options.js",
			"dist/primer.css",
			"dist/main.js"
		]
	}
QtrsiRp5MOu48A4v9J9z`

	diff := cmp.Diff(got, want)
	if diff != "" {
		t.Errorf(diff)
	}

	nr = &Normalizer{
		DelimiterGenerator: &fixedDelimiterGenerator{Fixed: "depop"},
	}
	_, err = nr.Normalize("report_multiline_json", example, ByteSizeFromGitHubDoc)
	if !errors.Is(err, ErrConflict) {
		t.Errorf("can not detect conflict: %v", err)
	}
}
