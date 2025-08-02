package main

import (
	"fmt"
	"learnpack/src/currency-converter/repository"
	"learnpack/src/currency-converter/service"
	"time"
)

func main() {

	go service.ProcessEntities()

	time.Sleep(2 * time.Second)

	currencies := repository.GetAllCurrencies()
	conversions := repository.GetAllConversions()

	fmt.Println("\nСохраненные валюты:")
	for _, currency := range currencies {
		fmt.Printf("Код: %s, Курс: %f, Название: %s, Символ: %s\n",
			currency.GetCode(),
			currency.GetRate(),
			currency.GetName(),
			currency.GetSymbol())
	}

	fmt.Println("\nСохраненные конвертации:")
	for _, conversion := range conversions {
		fmt.Printf("Сумма: %f, Из: %s, В: %s, Результат: %f\n\n",
			conversion.GetAmount(),
			conversion.GetFromCurrency().GetCode(),
			conversion.GetToCurrency().GetCode(),
			conversion.GetResult())
	}
}