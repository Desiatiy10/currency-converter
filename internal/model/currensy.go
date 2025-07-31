package model

type Currency struct {
	code string
	rate float64
	name string
	symbol string
}

func (c *Currency) GetCode() string {
    return c.code
}

func (c *Currency) GetRate() float64 {
	return c.rate
}	

func (c *Currency) GetName() string {
	return c.name
}

func (c *Currency) GetSymbol() string {
	return c.symbol
}