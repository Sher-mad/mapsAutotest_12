package stats

import (
	"github.com/Sher-mad/mapsAutotest_12/pkg/bank/types"
)

func FirterBuCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}
