package mock

type Poster struct {
	Head    string
	Content map[string]string
}

func (p *Poster) Post(url string) bool {
	return false
}
