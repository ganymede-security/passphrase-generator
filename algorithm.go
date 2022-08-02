package passphrasegenerator

import (
	"log"
	"math"
)

// The mask represents an array of bits representing a passphrase
// No matter the order the items are arranged, in the combined
// Length should equal the requested passphrase length
type mask []int32

// Phrase modifiers
const (
	PG_SPEC_CHAR int32 = 1 << iota
	
	PG_NUMBER

	PG_SEPARATOR

	PG_WORD
	// Only the last word is given this value
	PG_LAST_WORD
	// TODO: Add a bitmask representing the lengths of each word
	PG_WORD_LENGTH
)

// Words are an array of ints representing the lengths of each string 
// that will be used in the phrase. 
type words []int32

// NewPhrase returns a shuffled list of words and a mask ready for
// generating a passphrase
func newPhrase(opts Options) (words, mask) {
	ws := make([]int32, 0)
	w, m := genMask(opts)
	_, rm := addSeparators(w, m, opts)
	for i:= 0; i < len(w)-1; i++ {
		ws = append(ws, w[i] - 1)
	}
	ws = append(ws, w[len(w)-1])
	return ws, rm
}

// Since separators will be added to each item except the last, their addition
// does not count towards the total length. However, the lengths and number
// of words must be known before separators can be added
func addSeparators(w words, m mask, opts Options) (words, mask) {
	last := w[len(w) - 1]
	rt := make([]int32, 0)
	rtm := make([]int32, 0)
	sep := int32(len(opts.Separator))

	for i:=0; i < len(w) - 1; i++ {
		rt = append(rt, w[i] - sep)
		rt = append(rt, sep)
	}

	rt = append(rt, last)

	for i:=0; i < len(m); i++ {
		switch m[i] {
		case PG_NUMBER:
			rtm = append(rtm, PG_NUMBER)

		case PG_SPEC_CHAR:
			rtm = append(rtm, PG_SPEC_CHAR)

		case PG_WORD:
			rtm = append(rtm, PG_WORD)
			rtm = append(rtm, PG_SEPARATOR)

		case PG_LAST_WORD:
			rtm = append(rtm, PG_LAST_WORD)

		}
	}

	return rt, rtm
}

// GetNumSizes returns an array containing the lengths of each word in the phrase.
// Requires an input of the desired total phrase length, and desired numbers of
// special characters, numbers, and whether the case should be changed
func genWords(opts Options) (w words, err error) {
	rl := int32(opts.PhraseLength - opts.SpecialChars - opts.Numbers)
	words := make([]int32, 0)
	var lw int32 = 0

	for i := int32(0); i != rl; {
		var rn int32 = 0
		for rn < 1 {
			rn, _ = getRandNum(len(opts.Separator)+2, opts.MaxWordLength - 1)
		}
		if rn+i > rl {
			lw = rl - i
			for lw == 0 || lw == 1 {
				_, words = words[len(words)-1], words[:len(words)-1]
				nlw := int32(0)
				for _, v := range words {
					if v > 0 {
						nlw += v
					} else {
						nlw += (-v)
					}
				}
				i = nlw
				lw = rl - i
			}
			words = append(words, lw)
			i += lw
		} else {
			words = append(words, rn)
			i += rn
		}
	}
	return words, nil
}

// genMask creates a phrase mask representing the different types of items
// a phrase will need to generate.  
func genMask(opts Options) (words, mask) {
	n := opts.Numbers
	sC := opts.SpecialChars

	phraseMask := mask{}

	for i := 0; i < n; i++ {
		phraseMask = append(phraseMask, PG_NUMBER)
	}

	for i := 0; i < sC; i++ {
		phraseMask = append(phraseMask, PG_SPEC_CHAR)
	}

	wds, err := genWords(opts)
	if err != nil {
		log.Print(err)
	}

	for i := 0; i < len(wds) - 1; i++ {
		phraseMask = append(phraseMask, PG_WORD)
	}
	// Append last word
	phraseMask = append(phraseMask, PG_LAST_WORD)

	shuffled := CryptoShuffle(phraseMask)

	return wds, shuffled
}

// calculateEntropy calculates the total entropy of a generated passphrase
// mask. Returns the calculated value as a float
func calculateEntropy(m mask) float64 {
	var sum float64

	for i:= 0; i<len(m); i++ {

		switch m[i] {
		case PG_NUMBER:
			totalNums := float64(10)

			sum += math.Log2(totalNums)
		case PG_SPEC_CHAR:
			totalSpecs := float64(14)

			sum += math.Log2(totalSpecs)
		case PG_WORD:
			totalWords := float64(9464) // Total words in the list

			sum += math.Log2(totalWords)
		case PG_LAST_WORD:
			totalWords := float64(9464) // Total words in the list

			sum += math.Log2(totalWords)
		case PG_SEPARATOR:
			//TODO: Test whether this makes sense
			totalSeparators := float64(104)

			sum += math.Log2(totalSeparators)
		default:
			continue
		}
	}
	return sum
}
