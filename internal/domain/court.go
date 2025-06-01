package domain

type Court struct {
	ID        int64
	Name      string
	Location  string
	WidthCm   int64
	LengthCm  int64
	IsOutdoor bool
}
