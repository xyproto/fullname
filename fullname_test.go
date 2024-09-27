package fullname

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	fullName := Get()
	if fullName == "" {
		t.Fail()
	}
	fmt.Println(fullName)
}
