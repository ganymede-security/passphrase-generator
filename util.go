package passphrasegenerator

import (
	"bytes"
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	mrand "math/rand"
	"sort"
	"time"
)

const (
	FILE_NAME = "./words.txt"
)

// ProcessFile is a helper function to read a word file, returning
// the items as a list in memory sorted by the length of the word.
func processFile() (s1 map[int][]string) {
	f, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	doc := bytes.Split(f, []byte{'\n'})

	sort.Slice(doc, func(i, j int) bool {
		return len(doc[i]) < len(doc[j])
	})

	//create map of available words
	wordsMap := make(map[int][]string)

	for _, v := range doc {
		key := len(v)
		wordsMap[key] = append(wordsMap[key], string(v))
	}

	delete(wordsMap, 0) // Delete EOF 0 value key

	return wordsMap
}

// getRandString is a helper function that takes in a map containing
// the words from a word list and returns a randomly selected string
// from the list of a specific length.
func getRandString(m map[int][]string, wl int) (string, error) {
	noOfItems := len(m[wl]) - 1
	if noOfItems <= 1 {
		err := fmt.Errorf("util: No items of size %d in list", noOfItems)
		return "", err
	}
	random, err := getRandNum(2, noOfItems)
	if err != nil {
		panic(err)
	}
	words := m[wl]

	return words[random], nil
}

// getRandNum returns a cryptographically secure
// random number, requires the maximum desired number that should be returned.
//
// In the event of an error the function panics
func getRandNum(min int, max int) (int32, error) {
	var src mrand.Source64
	src = &cSrc{}
	src.Seed(time.Now().UnixMicro())
	rnd := mrand.New(src)
	var err error
	var randNum int

	for int(randNum) <= min {
		randNum = rnd.Intn(max + 1)
		if err != nil {
			return -1, err
		}
	}

	return int32(randNum), nil
}

// CryptoShuffle is an implementation of the Fisher-Yates shuffle
func CryptoShuffle(arr []int32) []int32 {
	var rn int32
	arrLen := len(arr) - 2

	for i := int32(0); int(i) < arrLen; i++ {
		rn, _ = getRandNum(0, arrLen)
		arr[i], arr[rn] = arr[rn], arr[i]
	}
	return arr
}

// Helper function to count the number of trailing unset bits in
// a negative int32
func countTrailingBits(n int32) int32 {
	count := int32(0)
	for n != 0 {
		n &= (n - 1)
		count++
	}
	return 32 - count
}

func reverse(b int32) int32 {
	b = (b&0xF0)>>4 | (b&0x0F)<<4
	b = (b&0xCC)>>2 | (b&0x33)<<2
	b = (b&0xAA)>>1 | (b&0x55)<<1
	return b
}

type cSrc struct{}

func (s cSrc) Seed(seed int64) {}

func (s cSrc) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cSrc) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}

func (opts Options) withoutModifiers() int32 {
	return int32(opts.PhraseLength - opts.Numbers - opts.SpecialChars)
}