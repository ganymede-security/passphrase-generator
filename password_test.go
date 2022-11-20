package passphrasegenerator

import (
	"fmt"
	"testing"
)

func TestNewPassword(t *testing.T) {
	pass, err := NewPassword(opts)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(pass)
}
