package main

import (
	"encoding/json"
	"net/http"
)

type Order struct {
	Price int64 `json:"price"`
	Num   int64 `json:"num"`
}

// GetAveragePricePerStore 每家店的人均价
//func GetAveragePricePerStore(storeName string) (int64, error) {
//	res, err := http.Get("https://liwenzhou.com/api/orders?storeName=" + storeName)
//	if err != nil {
//		return 0, err
//	}
//	defer res.Body.Close()
//
//	var orders []Order
//	if err := json.NewDecoder(res.Body).Decode(&orders); err != nil {
//		return 0, err
//	}
//
//	if len(orders) == 0 {
//		return 0, nil
//	}
//
//	var (
//		p int64
//		n int64
//	)
//
//	for _, order := range orders {
//		p += order.Price
//		n += order.Num
//	}
//
//	return p / n, nil
//}

// OrderInfoGetter 订单信息提供者
type OrderInfoGetter interface {
	GetOrders(string) ([]Order, error)
}

// HttpApi HTTP API类型
type HttpApi struct{}

// GetOrders 通过HTTP请求获取订单数据的方法
func (a HttpApi) GetOrders(storeName string) ([]Order, error) {
	res, err := http.Get("https://liwenzhou.com/api/orders?storeName=" + storeName)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var orders []Order
	if err := json.NewDecoder(res.Body).Decode(&orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// GetAveragePricePerStore 每家店的人均价
func GetAveragePricePerStore(getter OrderInfoGetter, storeName string) (int64, error) {
	orders, err := getter.GetOrders(storeName)
	if err != nil {
		return 0, err
	}

	if len(orders) == 0 {
		return 0, nil
	}

	var (
		p int64
		n int64
	)

	for _, order := range orders {
		p += order.Price
		n += order.Num
	}

	return p / n, nil
}
