package viewutils

import (
	"fmt"
	"strings"

	"db2sem/internal/transport/models"
)

func JoinSportsWithComma(sports []models.Sport) string {
	var builder strings.Builder

	for i, sport := range sports {
		if i+1 == len(sports) {
			fmt.Fprintf(&builder, "%s.", sport.Name)
			continue
		}

		fmt.Fprintf(&builder, "%s, ", sport.Name)
	}

	return builder.String()
}

func ContainsSport(sports []models.Sport, sport models.Sport) bool {
	for _, item := range sports {
		if item.ID == sport.ID {
			return true
		}
	}

	return false
}

func ContainsClub(clubs []models.Club, club models.Club) bool {
	for _, item := range clubs {
		if item.ID == club.ID {
			return true
		}
	}

	return false
}
