package main

import "fmt"

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("ERRRR occured:", err)
		} else {
			panic(r)
		}
	}()
	b := 0
	a := 5 / b
	fmt.Println(a)
	return
}
func main() {
	tryRecover()
}
