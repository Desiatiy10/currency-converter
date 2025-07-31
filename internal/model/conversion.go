package model

type Conversion struct {
	amount float64
	fromCurrency *Currency
	toCurrency *Currency
	result float64
}

func (c *Conversion) GetAmount() float64 {
	return c.amount
}

func (c *Conversion) GetFromCurrency() *Currency {
	return c.fromCurrency
}

func (c *Conversion) GetToCurrency() *Currency {
	return c.toCurrency
}

func (c *Conversion) GetResult() float64 {
	return c.result
}