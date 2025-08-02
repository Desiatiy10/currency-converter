package service

import (
	"learnpack/src/currency-converter/internal/model"
	"learnpack/src/currency-converter/repository"
)

func ProcessEntities() {
    currencyUSD := model.Currency{
        Code:    "USD",
        Rate:    1.0,
        Name:    "US Dollar",
        Symbol:  "$",
    }
    
    currencyEUR := model.Currency{
        Code:    "EUR",
        Rate:    0.85,
        Name:    "Euro",
        Symbol:  "€",
    }
    
    conversion := model.Conversion{
        Amount:      100,
        FromCurrency: &currencyUSD,
        ToCurrency:   &currencyEUR,
        Result:      85,
    }
    
    repository.StoreEntity(&currencyUSD)
    repository.StoreEntity(&currencyEUR)
    repository.StoreEntity(&conversion)
}