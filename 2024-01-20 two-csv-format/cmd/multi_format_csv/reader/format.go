package reader

type FigureType string

const (
	FigureType01 = FigureType("figure01")
	FigureType02 = FigureType("figure02")
)

var FigureTypes = map[string]FigureType{
	FigureType01.String(): FigureType01,
	FigureType02.String(): FigureType02,
}

func (f FigureType) Validate() bool {
	_, ok := FigureTypes[f.String()]
	return ok
}

func (f FigureType) String() string {
	return string(f)
}
