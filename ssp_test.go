package passphrasegenerator

import (
	"log"
	"testing"
)

func TestPowerSet(t *testing.T) {
	for i:=0; i < 10000; i++ {
		arr, _ := memoize(opts)
		target := opts.withoutModifiers()
	
		log.Print("Arr: ", arr)
	
		subset := powerSet(arr, target)
	
		for _, v := range subset {
			total := int32(0)
			if sumArr(v) == opts.withoutModifiers() {
				log.Print("Success", v)
				for _, v := range v {
					total = total + v
				}
				log.Print("Total ", total)
			}
		}
	}
}


func TestSubSetSum(t *testing.T) {
	for i:=0; i < 10000; i++ {
		arr, _ := memoize(opts)
		target := opts.withoutModifiers()
	
		subset, err := SubSetSum(arr, target)
		if err != nil {
			log.Print("Error: ", err)
			t.Fail()
		} else if sumArr(subset) != target {
			t.Fail()
			log.Print("Incorrect Sum reached: ", sumArr(subset))
			log.Print("Expected: ", target)
			log.Print("Success: ", subset, sumArr(subset)) 
		}
	}
}

func BenchmarkSubSetSum(b *testing.B) {
	target := opts.withoutModifiers()
	arr, _ := memoize(opts)

	for i:= 0; i < b.N; i++ {
		subset, err := SubSetSum(arr, target)
		if err != nil {
			log.Print("Error: ", err)
			log.Print("Subset: ", subset)
		}
	}
}