package passphrasegenerator

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	//wordList = processFile()
	wordList = newWordMap
)

type Options struct {
	// Specify the desired maximum length of any single word.
	// Shorter words are easier to remember. Default word list
	// has a cap of 15
	MaxWordLength int
	// Desired Length of the phrase.
	PhraseLength int
	// Minimum number of special characters to
	// include in a generated phrase.
	SpecialChars int
	// Minimum number of numbers to include
	// in a generated phrase.
	Numbers int
	// Whether the letter case should be changed
	// in a generated phrase
	ChangeCase bool
	// The separator to separate different words.
	//
	// If no separator is desired enter "" for the value.
	Separator string
}

const (
	lowerCharSet   string = "abcdedfghijklmnopqrstuvwxyz"
	upperCharSet   string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet string = "!@#$%&*{}[]?<>"
	numberSet      string = "0123456789"
	allCharSet     string = (lowerCharSet + upperCharSet + specialCharSet + numberSet)
)

// New returns a newly generated passphrase.
func New(opts Options) (string, error) {
	if opts.withoutModifiers() < 4 {
		return "", fmt.Errorf("generator: invalid options. Phrase modifiers and word length incompatible")
	}
	wds, msk := newPhrase(opts)
	s := new(wds, msk, opts)
	return s, nil
}

// NewWithEntropy returns a new passphrase as well as a float32 containing
// the total entropy of the generated phrase.
func NewWithEntropy(opts Options) (string, float64) {
	wds, msk := newPhrase(opts)
	s := new(wds, msk, opts)
	entr := calculateEntropy(msk)
	return s, entr
}

// NewFromMask generates random strings for all parts of a mask except the
// words.
func new(w words, m mask, opts Options) string {
	var builder strings.Builder
	// ensures the index value of the words list stays constant
	// with the mask index
	keeper := 0
	lenNums := len(numberSet) - 1
	lenSpecs := len(specialCharSet) - 1

	for i := 0; i < len(m); {
		switch m[i] {
		case PG_NUMBER:
			i++
			index, err := getRandNum(1, lenNums)
			if err != nil {
				log.Printf("pg_num: %v", err)
			}
			builder.WriteByte(numberSet[index])
		case PG_SPEC_CHAR:
			i++
			index, err := getRandNum(1, lenSpecs)
			if err != nil {
				log.Printf("pg_spec_char: %v", err)
			}
			builder.WriteByte(specialCharSet[index])
		case PG_SEPARATOR:
			i++
			builder.WriteString(opts.Separator)
		case PG_WORD:
			i++
			length := int(w[keeper])
			keeper++

			str, err := getRandString(wordList, length)
			if err != nil {
				log.Printf("pg_word: %v", err)
				log.Fatal("Regular Word: ", length)
			}

			builder.WriteString(str)
		case PG_LAST_WORD:
			i++
			length := int(w[keeper])
			keeper++

			str, err := getRandString(wordList, length)
			if err != nil {
				log.Printf("pg_last_word: %v\n", err)
				log.Fatalf("Failed with length: %d\tAt Index: %d\n", length, keeper)
			}
			builder.WriteString(str)
		}
	}
	// TODO: Find a way to force specific words to be uppercase
	var end string

	if opts.ChangeCase {
		end = cases.Title(language.English, cases.Compact).String(builder.String())
	} else {
		end = builder.String()
	}

	return end
}
