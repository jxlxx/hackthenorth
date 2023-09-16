package main

import (
	"encoding/json"
	"log"
	"time"
)

func (b *Bot) CreateTrader() {
	password := MustGetEnv("NATS_PASSWORD")
	t := Trader{Name: b.name, Password: password}
	bytes, err := json.Marshal(t)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := b.nc.Request(SubjCreateTrader, bytes, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	player := Trader{}
	if err := json.Unmarshal(resp.Data, &player); err != nil {
		log.Fatalln(err)
	}
	b.ID = player.ID
}

func (b *Bot) GetSecurities() []string {
	s := []string{}
	resp, err := b.nc.Request(SubjGetSecurities, nil, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(resp.Data, &s); err != nil {
		log.Fatalln(err)
	}
	return s
}

func (b *Bot) GetExchanges() []string {
	s := []string{}
	resp, err := b.nc.Request(SubjGetExchanges, nil, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(resp.Data, &s); err != nil {
		log.Fatalln(err)
	}
	return s
}

func (b *Bot) GetOrders() []Order {
	s := []Order{}
	p, _ := json.Marshal(b.ID)
	resp, err := b.nc.Request(SubjGetOrders, p, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(resp.Data, &s); err != nil {
		log.Fatalln(err)
	}
	return s
}

func (b *Bot) GetTransactions() []Transaction {
	s := []Transaction{}
	p, _ := json.Marshal(b.ID)
	resp, err := b.nc.Request(SubjGetTransactions, p, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(resp.Data, &s); err != nil {
		log.Fatalln(err)
	}
	return s
}

func (b *Bot) GetPortfolio() Portfolio {
	s := Portfolio{}
	p, _ := json.Marshal(b.ID)
	resp, err := b.nc.Request(SubjGetPortfolio, p, time.Second)
	if err != nil {
		log.Fatalln(err)
	}
	if err := json.Unmarshal(resp.Data, &s); err != nil {
		log.Fatalln(err)
	}
	return s
}
