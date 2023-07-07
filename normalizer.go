package normalizer

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"strings"
)

// https://docs.github.com/en/actions/using-workflows/workflow-commands-for-github-actions#example-of-a-multiline-string
const ByteSizeFromGitHubDoc = 15

type DelimiterGeneratorable interface {
	Generate(byteSize int) (string, error)
}

type Normalizer struct {
	DelimiterGenerator DelimiterGeneratorable
}

type Base64DelimiterGenerator struct {
}

var (
	ErrConflict = errors.New("generated delimiter has been conflicted with given input")
)

func (b *Base64DelimiterGenerator) Generate(byteSize int) (string, error) {
	bytes := make([]byte, byteSize)
	_, err := rand.Read(bytes)

	return base64.StdEncoding.EncodeToString(bytes), err
}

func (b *Base64DelimiterGenerator) EncodedLength(byteSize int) int {
	// https://stackoverflow.com/questions/13378815/base64-length-calculation
	return 4 * byteSize / 3
}

func (n *Normalizer) Normalize(name string, value string, byteSize int) (string, error) {
	const attemptLimit = 100
	for i := 0; i < attemptLimit; i++ {
		delimiter, err := n.DelimiterGenerator.Generate(byteSize)
		if err != nil {
			return "", err
		}

		if !strings.Contains(value, delimiter) {
			return name + "<<" + delimiter + "\n" + value + "\n" + delimiter, nil
		}
	}

	return "", ErrConflict
}

func DefaultNormalizer() *Normalizer {
	return &Normalizer{DelimiterGenerator: &Base64DelimiterGenerator{}}
}
