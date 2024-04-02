package dfa

import (
	"fmt"
	"testing"
)

func TestBuild(test *testing.T) {
	t := Tire{}
	t.ReplaceHold("*")
	t.Build([]string{"hell", "fuck", "ass crack"})
	//words := strings.Split("fuckyou"," ")
	//for i,w := range words {
	//	words[i] = t.Filter(w)
	//}
	fw := t.Filter("fuck you Hello")
	//fmt.Println(strings.Join(words," "))
	fmt.Println(fw)

}
