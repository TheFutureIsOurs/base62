package base62

import (
	"fmt"
	"testing"
)

func TestCode(t *testing.T) {
	num := 264688217293324297
	str := Encode(int64(num))
	fmt.Println(str)
	res, err := Decode(str)
	fmt.Println(res, err)
}
