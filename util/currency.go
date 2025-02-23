package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	RON = "RON"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, RON:
		return true
	}
	return false
}
