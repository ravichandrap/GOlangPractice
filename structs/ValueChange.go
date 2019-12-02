package main

import "fmt"

type Price struct {
	Currency string  `json:"currency"`
	Price    float64 `json:"price"`
	Type     string  `json:"type"`
}

func main() {
	cfp := []Price{{Currency:"EUR-1", Price:100, Type:"price-1"},{Currency:"EUR-2", Price:8000, Type:"Price-2"}}
	rp := []Price{{Currency:"EUR", Price:2, Type:"price"},{Currency:"EUR", Price:2, Type:""}}
	replacePrices(rp, cfp)
	fmt.Println(rp)

	c := Price{"INR", 20,"Price-c" }
	r := Price{"EUR", 0, "Price-r"}
	replacePricesValue(c, &r)
	fmt.Println(r)
}

func replacePrices(rp []Price, cfp []Price) {
	for i := range rp {
		rp[i].Price = cfp[i].Price
		rp[i].Currency = cfp[i].Currency
	}
}

func replacePricesValue(c Price, r *Price)  {

	(*r).Price = c.Price
	(*r).Currency = c.Currency
	(*r).Price = c.Price
}
