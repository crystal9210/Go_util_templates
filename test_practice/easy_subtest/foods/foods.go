package foods

type Apple struct {
	Variety string
}

func NewApple(variety string) *Apple {
	return &Apple{Variety: variety}
}

func (a *Apple) String() string {
	return a.Variety
}
