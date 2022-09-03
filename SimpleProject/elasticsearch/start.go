package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"math/rand"
	"strconv"
)

type User struct {
	ID   string
	Name string
	Age  int32
}

var (
	client *elasticsearch.Client
)

func main() {
	var err error
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "changeme",
	}

	client, err = elasticsearch.NewClient(cfg)
	//if err != nil {
	//	log.Fatalf("Error creating the cliente %s", err)
	//}
	res, err := client.Info()

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(elasticsearch.Version)
	log.Println(res)

	//var wg sync.WaitGroup
	//go func() {
	//	wg.Add(1)
	//	elasticInsertloop("Thread 01", 1000)
	//}()
	//go func() {
	//	wg.Add(1)
	//	elasticInsertloop("Thread 02", 1000)
	//}()
	//go func() {
	//	wg.Add(1)
	//	elasticInsertloop("Thread 03", 1000)
	//}()
	//go func() {
	//	wg.Add(1)
	//	elasticInsertloop("Thread 04", 1000)
	//}()
	//wg.Wait()

	user := User{
		Name: "Teste",
		Age:  int32(rand.Intn(100)),
	}
	idUser, err := Index(user)
	if err != nil {
		panic(err)
	}

	res, err = GetUser(idUser)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func elasticInsertloop(name string, iterations int) {
	for i := 0; i < iterations; i++ {
		user := User{
			Name: name,
			Age:  int32(rand.Intn(100)),
		}

		Index(user)
	}
}

func Index(user User) (idUser string, err error) {
	id := rand.Intn(100)
	idUser = strconv.Itoa(id)

	requestBytes, err := json.Marshal(user)
	if err != nil {
		return "", err
	}

	cfg := esapi.IndexRequest{
		DocumentID: idUser,
		Index:      "users",
		Body:       bytes.NewReader(requestBytes),
		Refresh:    "true",
	}

	res, err := cfg.Do(context.Background(), client)
	if err != nil {
		return "", err
	}

	fmt.Println(res)

	return idUser, nil
}

func GetUser(id string) (*esapi.Response, error) {
	cfg := esapi.GetRequest{
		Index:      "users",
		DocumentID: id,
	}

	res, err := cfg.Do(context.Background(), client)
	if err != nil {
		return nil, err
	}

	return res, nil
}
