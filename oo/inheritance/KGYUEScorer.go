package inheritance

type KGYUEScorer struct {
	KGScorer
}

func (s *KGYUEScorer) NameOfScorer() string {
	return "KGYUEScorer"
}
