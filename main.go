package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {
	bot := Bot{
		name: "super bot",
		nc:   Connect(),
	}

	bot.CreateTrader()

	securities := bot.GetSecurities()
	fmt.Println(securities)
	exchanges := bot.GetExchanges()
	fmt.Println(exchanges)
	transactions := bot.GetTransactions()
	fmt.Println(transactions)
	orders := bot.GetOrders()
	fmt.Println(orders)

	subject := NewOrder{
		SecurityID: securities[0],
		ExchangeID: exchanges[1],
		Qty:        1,
		OrderType:  BuyLimit,
	}

	payload := OrderPayload{
		TraderID: bot.ID,
		Price:    500 * 100, // $500
	}

	js, err := bot.nc.JetStream()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = js.Publish(subject.String(), payload.Bytes())
	if err != nil {
		log.Fatalln(err)
	}

	orders = bot.GetOrders()
	fmt.Println(orders)

	porfolio := bot.GetPortfolio()
	fmt.Println(porfolio)

}

func Connect() *nats.Conn {
	url := MustGetEnv("NATS_URL")
	username := MustGetEnv("NATS_USER")
	password := MustGetEnv("NATS_PASSWORD")
	nc, err := nats.Connect(url, nats.UserInfo(username, password))
	if err != nil {
		log.Fatalln(err)
	}
	return nc
}

func JetStreamConnect() nats.JetStreamContext {
	nc := Connect()
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalln(err)
	}
	return js
}
func MustGetEnv(key string) string {
	value, success := os.LookupEnv(key)
	if !success {
		log.Fatalln("failed to get: ", key)
	}
	return value
}
