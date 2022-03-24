package main

import (
	"log"

	nomics "github.com/dariubs/gonomics"
)

func main() {
	nb := nomics.NewNomics("your-api-key")
	ctr := nomics.CurrenciesTickerRequest{PerPage: "100"}

	data, err := nb.CurrenciesTicker(&ctr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s", data)

}
