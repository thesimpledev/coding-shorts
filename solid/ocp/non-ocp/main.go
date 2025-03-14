package discount

func GetDiscount(customerType string, price float64) float64 {
	switch customerType {
	case "regular":
		return price * 0.1
	case "premium":
		return price * 0.2
	default:
		return 0
	}
}
