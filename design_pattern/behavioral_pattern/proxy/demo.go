package proxy

import "fmt"

type Seller interface {
	sell(name string)
}

type Station struct {
	stock int
}

func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("Station %s sold one ticket, remaining %d \n", name, station.stock)
	} else {
		fmt.Println("Sold out")
	}

}

type StationProxy struct {
	station *Station
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("Proxy %s sold one ticket, remaining %d \n", name, proxy.station.stock)
	} else {
		fmt.Println("Sold out")
	}
}
