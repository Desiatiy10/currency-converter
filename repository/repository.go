package repository

import (
	"fmt"
	"learnpack/src/currency-converter/internal/model"
)

var currencies []*model.Currency = make([]*model.Currency, 0)
var conversions []*model.Conversion = make([]*model.Conversion, 0)

func StoreEntity(entity model.Entity) {
	switch v := entity.(type) {
	case *model.Currency:
		currencies = append(currencies, v)
	case *model.Conversion:
		conversions = append(conversions, v)
	default:
		panic(fmt.Sprintf("Неизвестный тип сущности: %T", v))
	}
}

func GetAllCurrencies() []*model.Currency {
	return currencies
}

func GetAllConversions() []*model.Conversion {
	return conversions
}