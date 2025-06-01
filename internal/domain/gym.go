package domain

type Gym struct {
	ID             int64
	Name           string
	Location       string
	TrainersCount  int16
	DumbbellsCount int16
	HasBathhouse   bool
}
