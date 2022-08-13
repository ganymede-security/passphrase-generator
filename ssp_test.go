package passphrasegenerator

import (
	"log"
	"testing"
)

func TestPowerSet(t *testing.T) {
	arr, _ := memoize(opts)
	//arr := []int32 { 5, 10, 15 }

	log.Print("Arr: ", arr)
	//log.Print("Target: ", opts.withoutModifiers())

	subset := powerSet(arr)
	total := int32(0)

	for _, v := range subset {
		if sumArr(v) == opts.withoutModifiers() {
			log.Print("Success", v)
			for _, v := range v {
				total = total + v
			}
			log.Print("Total ", total)
		}
	}

	//log.Print(subset)
}


func TestSubSetSum(t *testing.T) {
	arr, _ := memoize(opts)

	subset, sum, err := SubSetSum(arr, opts)
	if err != nil {
		log.Print("Error: ", err)
		t.Fail()
	} else if sum != opts.withoutModifiers() {
		t.Fail()
		log.Print("Incorrect Sum reached: ", sum)
		log.Print("Expected: ", opts.withoutModifiers())
	} else {
		log.Print("Success: ", subset, sum) 
	}
}

func BenchmarkSubSetSum(b *testing.B) {
	for i:= 0; i < b.N; i++ {
		arr, _ := memoize(opts)

		subset, sum, err := SubSetSum(arr, opts)
		if err != nil {
			log.Print("Error: ", err)
		} else if sum != opts.withoutModifiers() {
			log.Print("Incorrect Sum reached: ", sum)
			log.Print("Expected: ", opts.withoutModifiers())
			log.Print("Subset: ", subset)
		}
	}
}