package main

import (
    "fmt"
    "net/http"
)

type Currency struct {
    id int
    name string
    min_size string
}

type Currencies struct {
    data []Currency
}

func main() {
    fmt.Println("Hello world");
    currency1 := Currency {
        id: 1,
        name: "BTC",
        min_size: "12",
    }
    currency2 := Currency {
        id: 2,
        name: "ETH",
        min_size: "244",
    }
    currencies := Currencies {
        data: []Currency {currency1, currency2},
    }
    fmt.Printf("Currency 1: %d, %s, %s\n", currency1.id, currency1.name, currency1.min_size)
    fmt.Println("Currency 2: ", currency2);
    fmt.Println("Currencies: ", currencies);
}
