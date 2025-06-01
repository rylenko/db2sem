package domain

type Stadium struct {
	ID            int64
	Name          string
	Location      string
	WidthCm       int64
	LengthCm      int64
	MaxSpectators int16
	IsOutdoor     bool
	Coating       string
}
