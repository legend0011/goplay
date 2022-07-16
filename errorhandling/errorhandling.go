package main

import (
	"bufio"
	"fmt"
	"os"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func printFile(name string) {
	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		// if you know some error type, list it out. otherwise print unknown error
		if pathErr, ok := err.(*os.PathError); ok {
			fmt.Println(pathErr.Err)
		} else {
			fmt.Println("unknown error")
		}
	}

	writer := bufio.NewWriter(file)
	defer writer.Flush() // this run before file.Close()

	f := fib()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f()) // write something into writer(impl Write)
	}
}

func main() {
	printFile("fib.txt")
}
