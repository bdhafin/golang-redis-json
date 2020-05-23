package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v7"
)

// Car is
type Car struct {
	Color string
	Type  string
}

func main() {
	fmt.Println("Hallo World")
	ExampleClient()
}

// ExampleClient is
func ExampleClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	car := &Car{
		Color: "Red",
		Type:  "SUV",
	}

	jsonData, err := json.Marshal(car)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = client.Set("testjson", jsonData, 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("testjson").Result()
	if err != nil {
		panic(err)
	}

	carUnmarshal := &Car{}
	err = json.Unmarshal([]byte(val), carUnmarshal)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Color", carUnmarshal.Color)
	fmt.Println("Type", carUnmarshal.Type)
}
