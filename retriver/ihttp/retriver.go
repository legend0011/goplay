package ihttp

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type Retriver struct {
}

func (r *Retriver) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	rs, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs[:100])
	return string(rs)
}
