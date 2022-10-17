package bleve

import (
	"fmt"
	"testing"
)

func init() {
	err := Init("../../greet.blv")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestSearchGreet(t *testing.T) {
	re, err := SearchGreet("早安")
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(re[0].Author)
}
