package passphrasegenerator

import (
	"log"
	"testing"
)

var (
	opts = Options{
		MaxWordLength: 15,
		PhraseLength:  32,
		SpecialChars:  1,
		Numbers:       1,
		ChangeCase:    true,
		Separator:     "-",
	}
)

var expectedLength int32 = int32(opts.PhraseLength - opts.Numbers - opts.SpecialChars)

func TestGenWords(t *testing.T) {
	var wrds words
	var err error
	for i:=0; i<1000000;i++ {
		wrds, err = genWords(opts)
		if err != nil {
			t.Fail()
			log.Print(err)
		}
		for i := 0; i < len(wrds); i++ {
			if int(wrds[i]) > opts.MaxWordLength {
				log.Print(wrds)
				t.Fatal()
			}
		}
	}

	
	total := 0

	for i := 0; i < len(wrds); i++ {
		total += int(wrds[i])
		if int(wrds[i]) > opts.MaxWordLength {
			t.Fail()
			log.Print(wrds)
		}
		if int(wrds[i]) < 4 {
			log.Print(wrds)
			t.Fatal()
		}
	}

	if total != int(expectedLength) {
		t.Fail()
		log.Printf("Expected: %d\tGot: %d\t", expectedLength, total)
	}

	log.Print(wrds)
}

func TestGenMask(t *testing.T) {
	_, mask := genMask(opts)
	req := 0

	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case PG_WORD:
			continue
		case PG_SPEC_CHAR:
			continue
		case PG_NUMBER:
			continue
		case PG_SEPARATOR:
			t.Fail()
		case PG_LAST_WORD:
			req++
		default:
			t.Fail()
		}
	}

	if req != 1 {
		t.Fail()
		log.Print(req)
	}

	log.Print(mask)
}

func TestAddSeparators(t *testing.T) {
	w, m := genMask(opts)
	req := 0
	wl := 0
	sl := 0

	words, mask := addSeparators(w, m, opts)

	for i := 0; i < len(mask); i++ {
		switch mask[i] {
		case PG_WORD:
			wl++
			continue
		case PG_SPEC_CHAR:
			continue
		case PG_NUMBER:
			continue
		case PG_SEPARATOR:
			sl++
		case PG_LAST_WORD:
			req++
		default:
			log.Printf("Failure: %d\n", i)
			t.Fail()
		}
	}

	if sl != wl {
		t.Fail()
		log.Printf("Expected: %d\tGot: %d\t", wl, sl)
	}

	if req != 1 {
		t.Fail()
		log.Printf("Expected: %d\tGot: %d\t", 1, req)
	}

	log.Print(words)
	log.Print(mask)
}

func TestNewPhrase(t *testing.T) {
	rw, rm := newPhrase(opts)
	mods := opts.Numbers + opts.SpecialChars
	newest := make([]int32, 0)

	total := int32(0)

	for _, v := range rw {
		total += v
		newest = append(newest, v - 1)
	}

	for i:=0; i<10000;i++ {
		list, _ := newPhrase(opts)

		for _, v := range list {
			if v > int32(opts.MaxWordLength) {
				t.Fail()
				log.Print(v)
			}
		}
	}

	flags := 0

	for i:=0; i < len(rm);i++ {
			if rm[i] == PG_SEPARATOR {
				flags++
			}
	}

	log.Print(flags)

	log.Print(total + int32(flags) + int32(mods))

	log.Print(rw)
	log.Print(rm)
}

func TestEntropy(t *testing.T) {
	_, mask := newPhrase(opts)

	sum := calculateEntropy(mask)

	
	log.Print(sum)
	log.Print(mask)
}

func BenchmarkNewPhrase(b *testing.B) {

	total := 0

	for i:= 0; i < b.N; i++ {
		words, _ := genWords(opts)
		if words != nil {
			total++
		}
	}
	result = total
}

func TestNew(t *testing.T) {
	
	words := []int32{9, 5, 11, 5}
	mask := []int32{8, 4, 1, 8, 4, 8, 4, 2, 16}

	phrase := new(words, mask, opts)

	log.Print(phrase)

}