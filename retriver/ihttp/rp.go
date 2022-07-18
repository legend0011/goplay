package ihttp

import "fmt"

type RetriverPoster struct {
	Head    string
	Content map[string]string
}

func (rp *RetriverPoster) Post(url string) bool {
	rp.Head = url
	return true
}

func (rp *RetriverPoster) Get(url string) string {
	return rp.Head
}

// implement fmt.Stringer
func (rp *RetriverPoster) String() string {
	return fmt.Sprintf("Real Jack...... here\n RP.Head=%s, RP.Content=%s\n",
		rp.Head, rp.Content["key"])
}
