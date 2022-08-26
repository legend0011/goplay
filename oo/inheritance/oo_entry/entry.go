package main

import (
	"fmt"
	"hello/oo/inheritance"
)

func main() {
	//kgScorer := oo.KGScorer{}
	//fmt.Printf("get a %s!", kgScorer.NameOfScorer())

	var arr []inheritance.Scorer
	arr = append(arr, &inheritance.KGScorer{})
	arr = append(arr, &inheritance.KGCNScorer{})
	arr = append(arr, &inheritance.KGYUEScorer{})

	for i, item := range arr {
		fmt.Printf("[%d] get a %s!\n", i, item.NameOfScorer())
	}
}
