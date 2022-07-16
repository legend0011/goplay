package main

import (
	"bufio"
	"fmt"
	"io"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

/* simple version read
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small,
	// 	need impl struct to cache last time generated number
	return strings.NewReader(s).Read(p)
}
*/

func (g intGen) Read(p []byte) (n int, err error) {
	next := g() // var next int = g()
	if next > 1000 {
		return 0, io.EOF
	}

	s := fmt.Sprintf("%d\n", next)
	n = copy(p, s)
	// fmt.Println(string(p) + " " + s)
	return n, nil
}

func printReaderContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	//f := fib()  // like a generator, each call will generate a fib number
	//fmt.Printf("=> %d\n", f())
	//fmt.Printf("=> %d\n", f())
	//fmt.Printf("=> %d\n", f())
	//fmt.Printf("=> %d\n", f())

	var i intGen = fib()
	printReaderContents(i)
}
