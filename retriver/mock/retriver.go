package mock

type Retriver struct {
	Content string
}

// func (r *Retriver) Get(url string) string {
// 	return r.Content
// }

func (r Retriver) Get(url string) string {
	return r.Content
}
