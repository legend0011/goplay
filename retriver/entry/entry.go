package main

import (
	"fmt"
	"hello/retriver/mock"
)

type Retriver interface {
	Get(url string) string
}

type Poster interface {
	Post(url string) bool
}

type RetriverPoster interface {
	Retriver
	Poster
}

func download(r Retriver) string {
	return r.Get("http://www.imooc.com")
}

func post(p Poster, url string) {
	if p.Post(url) {
		fmt.Println("great job! posted!")
	} else {
		fmt.Println("bad job... post failed!")
	}
}

func session(url string, rp RetriverPoster) bool {
	fmt.Printf("start a session on %s!\n", url)
	post(rp, url)
	rs := download(rp)
	fmt.Printf("session ended. rp.Head = %s", rs)
	return true
}

func main() {
	// var r Retriver
	// fmt.Printf("1 %T, %v\n", r, r)
	// r = mock.Retriver{Content: "mock content"}
	// fmt.Printf("2 %T, %v\n", r, r)

	// r = &ihttp.Retriver{}
	// fmt.Printf("3 %T, %v\n", r, r)

	// rs := download(r)
	// fmt.Println("get result => ", rs)

	// var p Poster
	// p = &mock.Poster{}
	// post(p)

	var rp RetriverPoster
	rp = &mock.RetriverPoster{}

	fmt.Println(rp) // call rp's String()

	session("www.google.com", rp)

	fmt.Println(rp) // call rp's String()
}
