// File for testing internal functions

package passphrasegenerator


import (
	"log"
	"math/rand"
	"testing"
	"time"
)

var result int

func TestProcessFile(t *testing.T) {
	mapped := processFile()

	for k, _ := range mapped {
		log.Printf("Key: %d\tNum of items: %d\t", k, len(mapped[k]))
	}

	if len(mapped) == 0 {
		t.Fail()
	}
}

func BenchmarkProcessFile(b *testing.B) {
	total := 0

	for i:= 0; i < b.N; i++ {
		str := processFile()
		if str != nil {
			total++
		}
	}
	result = total
}

func TestGetRandString(t *testing.T) {
	total := 0
	words := processFile()

	// i starts at 2 because there are no strings with a length of 1
	for i := 2; i < len(words); i++ {
		st, err := getRandString(words, i)
		if err != nil {
			t.Fatalf("%v", err)
		}
		if len(st) != i {
			log.Printf("Expected: %d\t Got: %d\n", i, len(st))
			t.Fail()
		}
		total++
	}

	result = total
}

func BenchmarkGetRandString(b *testing.B) {
	total := 0
	words := processFile()
	mapped := make(map[string][]int)

	for i := 0; i < b.N; i++ {
		str, _ := getRandString(words, 6) // key 6 has the most items to search through
		mapped[str] = append(mapped[str], len(str))
		total ++ 
	}

	min := float32(2147483647)
	max := float32(0)

	for _, v := range mapped {
		if float32(len(v)) < min {
			min = float32(len(v))
		}
		if float32(len(v)) >= max {
			max = float32(len(v))
		}
	}

	log.Printf("Difference: %f\t", min/max)
	result = total
}

func TestCryptoShuffle(t *testing.T) {
	mask, _ := newPhrase(opts)
	log.Print(mask)
	var newMask []int32

	logResults := make(map[int]int32)

	for k, v := range mask {
		
		logResults[k] = v
	}

	counter := int32(0)

	for i:= 0; i < 100000; i++ {
		newMask = CryptoShuffle(mask)
		for k, _ := range newMask {
			counter++
			logResults[k] =  counter
		}
	}

	min := float32(2147483647)
	max := float32(0)

	for k, v := range logResults {
		log.Printf("Number: %d\tTimes Encountered: %d\t", k, v)
		if float32(v) < min {
			min = float32(v)
		}
		if float32(v) >= max {
			max = float32(v)
		}
	}
	log.Print(newMask)

	log.Printf("Difference: %f\t", min/max)
	
}

func BenchmarkCryptoShuffle(b *testing.B) {
	rand.Seed(time.Now().UnixMicro())

	mask, _ := newPhrase(opts)

	tracker := make(map[int][]int32)

	for k, v := range mask {
		tracker[k] = append(tracker[k], v)
	}

	for i:= 0;i < b.N; i++ {
		arr := CryptoShuffle(mask)
		for k, v := range arr {
			tracker[k] = append(tracker[k], v)
		}
	}

	total := 0

	for k, v := range tracker {
		log.Printf("Number: %d\tAmount: %d\t\n", k, len(v))
		total += len(v)
	}

	log.Printf("Average: %d\n", (len(tracker)/total))

}

func BenchmarkCrypto(b *testing.B) {
	random := rand.New(&cSrc{})
	for n := 0; n < b.N; n++ {
		result = random.Intn(7919)
	}
}

func BenchmarkIntn(b *testing.B) {
	tracker := make(map[int][]int32)

	for i :=0; i < b.N; i++ {
		n, _ := getRandNum(1, 15)
		tracker[int(n)] = append(tracker[int(n)], int32(n))
	}

	total := 0

	for k, v := range tracker {
		log.Printf("Number: %d\tAmount: %d\t\n", k, len(v))
		total += len(v)
	}

}
