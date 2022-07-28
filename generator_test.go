package passphrasegenerator_test

import (
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

var expectedLength = opts.PhraseLength - opts.Numbers - opts.SpecialChars

var results int

func TestNew(t *testing.T) {

	for i := 0; i < 1000; i++ {
		/*phrase := */ pg.New(opts)

	}

	//log.Print(phrase)

}

func BenchmarkNew(b *testing.B) {
	total := 0

	for i := 0; i < b.N; i++ {
		pg.New(opts)
		total++
	}

	results = total
}
