package passphrasegenerator_test

import (
	"log"
	"testing"

	pg "github.com/ganymede-security/passphrase-generator"
)

var (
	opts = pg.Options{
		MaxWordLength: 15,
		PhraseLength:  32,
		SpecialChars:  1,
		Numbers:       1,
		ChangeCase:    true,
		Separator:     "-",
	}
)

var results int

func TestNew(t *testing.T) {

	var phrase string

	for i := 0; i < 100000; i++ {
		phrase = pg.New(opts)
		if len(phrase) != opts.PhraseLength {
			log.Printf("Failed with length: %d", len(phrase))
			t.Fail()
		}
	}

	log.Print(phrase)
}

func BenchmarkNew(b *testing.B) {
	total := 0

	for i := 0; i < b.N; i++ {
		pg.New(opts)
		total++
	}

	results = total
}
