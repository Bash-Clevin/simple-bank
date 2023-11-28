package util

const (
	USD = "USD"
	EUR = "EUR"
	KES = "KES"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, KES:
		return true
	}
	return false
}
