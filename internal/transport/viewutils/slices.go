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

func JoinStringsWithComma(strs []string) string {
	var builder strings.Builder

	for i, str := range strs {
		if i+1 == len(strs) {
			fmt.Fprintf(&builder, "%s.", str)
			continue
		}

		fmt.Fprintf(&builder, "%s, ", str)
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
