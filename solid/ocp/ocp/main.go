package discount

type Customer interface {
	GetDiscount(price float64) float64
}

type RegularCustomer struct{}

func (r RegularCustomer) GetDiscount(price float64) float64 {
	return price * 0.1
}

type PremiumCustomer struct{}

func (p PremiumCustomer) GetDiscount(price float64) float64 {
	return price * 0.2
}

func CalculateDiscount(c Customer, price float64) float64 {
	return c.GetDiscount(price)
}
