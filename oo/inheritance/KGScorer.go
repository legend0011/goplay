package inheritance

type KGScorer struct {
}

func (s *KGScorer) NameOfScorer() string {
	return "KGScorer"
}
