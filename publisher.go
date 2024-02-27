package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func readJSONFromFile(filePath string) (map[string]interface{}, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func getRandomAttribute(data map[string]interface{}) (string, interface{}) {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	rand.Seed(time.Now().UnixNano())
	randomKey := keys[rand.Intn(len(keys))]
	return randomKey, data[randomKey]
}

func main() {
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("go_publisher")

	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer client.Disconnect(250)

	jsonData, err := readJSONFromFile("dados.json")
	if err != nil {
		panic(err)
	}

	for {
		key, value := getRandomAttribute(jsonData)
		attributeData := map[string]interface{}{
			key: value,
		}
		data, err := json.Marshal(attributeData)
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			continue
		}

		token := client.Publish("test/topic", 0, false, data)
		token.Wait()
		fmt.Printf("Published %s: %v\n", key, value)
		time.Sleep(2 * time.Second)
	}
}
