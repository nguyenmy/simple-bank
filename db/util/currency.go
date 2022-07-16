package util

const (
	EUR = "EUR"
	USD = "USD"
	CAD = "CAD"
)

func IsCurrencySupported(currency string) bool {
	switch currency {
	case EUR, USD, CAD:
		return true
	default:
		return false
	}
}
