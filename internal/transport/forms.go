package transport

type createSportsmanForm struct {
	Name      string  `form:"name"`
	BirthDate string  `form:"birth_date"`
	HeightCm  string  `form:"height_cm"`
	WeightKg  string  `form:"weight_kg"`
	ClubID    int64   `form:"club_id"`
	SportIDs  []int64 `form:"sport_ids"`
}

type updateSportsmanForm struct {
	Name      string  `form:"name"`
	BirthDate string  `form:"birth_date"`
	HeightCm  string  `form:"height_cm"`
	WeightKg  string  `form:"weight_kg"`
	ClubID    int64   `form:"club_id"`
	SportIDs  []int64 `form:"sport_ids"`
}
