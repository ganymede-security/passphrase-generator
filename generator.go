package passphrasegenerator

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

var (
	wordList = processFile()
)

type Generator struct {
	Wordlist map[int]string
}

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
	lowerCharSet string = "abcdedfghijklmnopqrst"
 	upperCharSet string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
 	specialCharSet string = "!@#$%&*{}[]?<>"
 	numberSet string = "0123456789"
	allCharSet string = (lowerCharSet + upperCharSet + specialCharSet + numberSet)
)

func New(opts Options) string {
	wds, msk := newPhrase(opts)
	s := new(wds, msk, opts)
	return s
}

// NewFromMask generates random strings for all parts of a mask except the
// words. 
func new(w words, m mask, opts Options) (string) {
	var builder strings.Builder
	// ensures the index value of the words list stays constant
	// with the mask index
	keeper := 0
	lenNums := len(numberSet) - 1
	lenSpecs := len(specialCharSet) - 1

	for i:=0; i < len(m); {
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
				log.Fatal(length, keeper)
			}
			builder.WriteString(str)
		}
	}
	return builder.String()
}

// ParseMask returns the word lengths as positive ints and negative
// integers representing locations of separators
func ParseMask(m mask, opts Options) (string) {
	var builder strings.Builder
	nums := 0
	sC := 0
	var lengths []int32

	for _, v := range m {
		switch v {
		case PG_NUMBER: 
			n, err := getRandNum(0, len(numberSet) - 1)
			if err != nil {
				log.Print(err)
			}
			builder.WriteByte(numberSet[n])
			nums++ 
		case PG_SEPARATOR:
			builder.WriteString(opts.Separator)
			lengths = append(lengths, -(v))
		case PG_SPEC_CHAR:
			n, err := getRandNum(0, len(numberSet) - 1)
			if err != nil {
				log.Print(err)
			}
			builder.WriteByte(specialCharSet[n])
			sC++
		default:
			if v < 0 {
				lengths = append(lengths, countTrailingBits(v))

				size := countTrailingBits(v)

				string, err := getRandString(wordList, int(size))
				if err != nil {
					log.Print(err)
				}
				builder.WriteString(string)
			} else {
				log.Printf("unrecognized value: %v", v)
			}
		}	
	}
	return builder.String()
}

// On init seed the random number generator
func Init() {
	rand.Seed(time.Now().UnixMicro())
}
