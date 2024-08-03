package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/professordaehn/hello-go/morestrings"
)

func main() {
	strToReverse := "!oG ,olleH"
	result := morestrings.ReverseRunes(strToReverse)
	fmt.Println(result)
	fmt.Println(cmp.Diff(strToReverse, result))
}
