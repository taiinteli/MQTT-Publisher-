package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"
	"os"
	"reflect"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func TestReadJSONFromFile_Success(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "test*.json")
	if err != nil {
		t.Fatalf("Unable to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	expected := map[string]interface{}{
		"key": "value",
	}
	bytes, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("Unable to marshal JSON: %v", err)
	}
	if _, err := tmpFile.Write(bytes); err != nil {
		t.Fatalf("Unable to write to temporary file: %v", err)
	}
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("Unable to close temporary file: %v", err)
	}

	result, err := readJSONFromFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("readJSONFromFile returned an error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected result to be %v, got %v", expected, result)
	}
}

func TestReadJSONFromFile_Failure(t *testing.T) {
	_, err := readJSONFromFile("nonexistent.json")
	if err == nil {
		t.Errorf("Expected error for nonexistent file, got nil")
	}
}

func TestPublishRate(t *testing.T) {
    jsonData, err := readJSONFromFile("dados.json")
    if err != nil {
        t.Fatalf("Erro ao ler o arquivo JSON: %v", err)
    }

    opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
    opts.SetClientID("go_publisher")

    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }
    defer client.Disconnect(250)

    ch := make(chan struct{})
    
    client.Subscribe("test/topic", 0, func(client MQTT.Client, msg MQTT.Message) {
        ch <- struct{}{}
    })

    go func() {
        for {
            key, _ := getRandomAttribute(jsonData)
            client.Publish("test/topic", 0, false, key)
            time.Sleep(2 * time.Second)
        }
    }()

    tolerance := 100 * time.Millisecond
    expectedInterval := 2 * time.Second
    timer := time.NewTimer(expectedInterval + tolerance)
    defer timer.Stop()

    for i := 0; i < 5; i++ {
        select {
        case <-ch:
            if !timer.Stop() {
                <-timer.C
            }
            timer.Reset(expectedInterval + tolerance)
        case <-timer.C:
            t.Errorf("Mensagem %d não foi publicada dentro do intervalo esperado", i+1)
        }
    }
}
