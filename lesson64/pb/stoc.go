package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

type Stock struct {
	Name  string    `json:"name"`
	Time  time.Time `json:"time"`
	Price int       `json:"price"`
}

var companies = []string{"BMW", "", "dodi", "Amazon", "Facebook"}

func main() {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()
	for {
		for _, company := range companies {
			time.Sleep(time.Second)
			stock := Stock{
				Name:  company,
				Time:  time.Now(),
				Price: rand.Intn(201),
			}
			stockJSON, err := json.Marshal(stock)
			if err != nil {
				log.Printf("JSON marshalling error: %v", err)
				continue
			}
			err = rdb.Publish(ctx, "stock_updates", stockJSON).Err()
			if err != nil {
				log.Printf("Redis publish error: %v", err)
			}
			checkAndSaveExtremes(ctx, rdb, stock)
		}
	}
}

func checkAndSaveExtremes(ctx context.Context, rdb *redis.Client, stock Stock) {
	highKey := fmt.Sprintf("%s_high", stock.Name)
	lowKey := fmt.Sprintf("%s_low", stock.Name)
	currentHigh, err := rdb.Get(ctx, highKey).Int()
	if err == redis.Nil || stock.Price > currentHigh {
		err = rdb.Set(ctx, highKey, stock.Price, 0).Err()
		if err != nil {
			log.Printf("Error setting high value: %v", err)
		}
	}
	currentLow, err := rdb.Get(ctx, lowKey).Int()
	if err == redis.Nil || stock.Price < currentLow {
		err = rdb.Set(ctx, lowKey, stock.Price, 0).Err()
		if err != nil {
			log.Printf("Error setting low value: %v", err)
		}
	}
}
