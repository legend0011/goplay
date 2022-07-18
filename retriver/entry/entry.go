package main

import (
	"fmt"
	"hello/retriver/ihttp"
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

	rp := &mock.RetriverPoster{}
	session("www.google.com", rp)

	fmt.Println(rp) // call rp's String()

	fmt.Println("Try type assertion ----------------------")
	// want to check what's behind an interface
	var rpArr []RetriverPoster
	mockRP := &mock.RetriverPoster{}
	realRP := &ihttp.RetriverPoster{}

	rpArr = append(rpArr, mockRP, realRP)

	for i, item := range rpArr {
		fmt.Printf("[%d] %v", i, item) // call rp's String()
		if _, ok := item.(*mock.RetriverPoster); ok {
			fmt.Printf("[%d] is mocked!\n", i)
		} else if _, ok := item.(*ihttp.RetriverPoster); ok {
			fmt.Printf("[%d] is real!\n", i)
		} else {
			fmt.Printf("[%d] unknown rp\n", i)
		}
	}

	fmt.Println("Try type switch ----------------------")

	for i, item := range rpArr {
		switch item.(type) {
		case *mock.RetriverPoster:
			fmt.Printf("[%d] is mocked!\n", i)
		case *ihttp.RetriverPoster:
			fmt.Printf("[%d] is real!\n", i)
		default:
			fmt.Printf("[%d] unknown rp\n", i)
		}
	}

}
