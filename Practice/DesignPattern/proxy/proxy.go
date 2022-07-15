package proxy

import "fmt"

// 代理模式可以为另一个对象提供一个替身或者占位符，以控制对这个对象的访问

type Seller interface {
	sell(name string)
}

// Station
// @Description: 火车站
type Station struct {
	stock int
}

// sell
// @Description: 售票
// @receiver station
// @param name
func (station *Station) sell(name string) {
	if station.stock > 0 {
		station.stock--
		fmt.Printf("代理处理中：%s 买了一张票，剩余：%d \n", name, station.stock)
	} else {
		fmt.Println("out of stock...")
	}
}

// StationProxy
// @Description: 火车站代理售票点
type StationProxy struct {
	station *Station
}

func (proxy *StationProxy) sell(name string) {
	if proxy.station.stock > 0 {
		proxy.station.stock--
		fmt.Printf("代理处理中：%s 买了一张票，剩余：%d \n", name, proxy.station.stock)
	} else {
		fmt.Println("out of stock...")
	}
}
