package passphrasegenerator

import (
	"fmt"
	"math/rand"
	"strings"
)

func NewPassword(opts Options) (string, error) {
	var password strings.Builder

	// Special Character set
	for i := 0; i < opts.SpecialChars; i++ {
		rand, err := getRandNum(0, len(specialCharSet))
		if err != nil {
			return "", fmt.Errorf("NewPassword: Error generating random special character")
		}
		password.WriteString(string(specialCharSet[rand]))
	}
	// Set number
	for i := 0; i < opts.Numbers; i++ {
		rand, err := getRandNum(0, len(numberSet))
		if err != nil {
			return "", fmt.Errorf("NewPassword: Error generating random number")
		}
		password.WriteString(string(numberSet[rand]))
	}
	
	//Set uppercase
	// for i := 0; i < opts.ChangeCase; i++ {
	// 	rand, err := getRandNum(0, len(numberSet))
	// 	if err != nil {
	// 		return "", fmt.Errorf("NewPassword: Error generating random special character")
	// 	}
	// 	password.WriteString(string(numberSet[rand]))
	// }

	remaining := opts.withoutModifiers()

	for i := 0; i < int(remaining); i++ {
		rand, err := getRandNum(0, len(lowerCharSet))
		if err != nil {
			return "", fmt.Errorf("NewPassword: Error generating random character")
		}
		password.WriteString(string(lowerCharSet[rand]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune), nil
}
