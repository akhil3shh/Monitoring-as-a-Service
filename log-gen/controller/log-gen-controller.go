// package controller

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"

// 	"github.com/IBM/sarama"
// 	"github.com/Pratham-Karmalkar/kafka"
// 	"github.com/Pratham-Karmalkar/models"
// 	"github.com/gorilla/mux"
// )

// var (
// 	generationFlag bool
// 	stopChannel    chan struct{}
// )

// func StartGeneration(w http.ResponseWriter, r *http.Request) {
// 	param := mux.Vars(r)
// 	sig := param["stat"]
// 	brokerList := []string{"localhost:29092"}
// 	fmt.Println(sig)
// 	//Kafka producer configuration
// 	config := sarama.NewConfig()
// 	config.Producer.Return.Successes = true
// 	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to acknowledge the message
// 	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
// 	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
// 	i := false
// 	//done := make(chan bool)
// 	if sig == "true" {
// 		i = true
// 	} else {
// 		i = false
// 	}
// 	go func() {

// 		timer := time.NewTicker(time.Duration(time.Second * 3))

// 		for {
// 			select {
// 			case <-timer.C:

// 				//Produce and send Message
// 				log := &models.Log{}
// 				fmt.Println("i: ", i)
// 				data, err := json.Marshal(log.GenerateLog(i))
// 				if err != nil {
// 					panic(err)
// 				}
// 				kafka.KafkaProducer(brokerList, config, data)
// 				//NOTE: Does not encode JSON Data, if needed to change in kafka-config.go->KafkaProducer to 'CH' marked location
// 				// case <-done:
// 				// 	return
// 				if !i {
// 					return
// 				}
// 			}
// 		}

// 	}()
// 	w.WriteHeader(202)

// }

package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/IBM/sarama"
	"github.com/Pratham-Karmalkar/kafka"
	"github.com/Pratham-Karmalkar/models"
	"github.com/gorilla/mux"
)

var (
	generationFlag bool
	stopChannel    chan struct{}
	brokerList     []string
	config         *sarama.Config
)

func StartGeneration(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	sig := param["stat"]

	if sig == "start" {
		// Start generation
		startGeneration()
		w.WriteHeader(http.StatusAccepted)
	} else if sig == "stop" {
		// Stop generation
		stopGeneration()
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func startGeneration() {
	brokerList = []string{"localhost:29092"}
	// Kafka producer configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to acknowledge the message
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	if !generationFlag {
		generationFlag = true
		stopChannel = make(chan struct{})

		go func() {
			timer := time.NewTicker(time.Duration(time.Second * 3))

			for {
				select {
				case <-timer.C:
					log := &models.Log{}
					fmt.Println("Generating log...")
					data, err := json.Marshal(log.GenerateLog(generationFlag))
					if err != nil {
						panic(err)
					}
					kafka.KafkaProducer(brokerList, config, data)

				case <-stopChannel:
					fmt.Println("Stopping generation...")
					generationFlag = false
					return
				}
			}
		}()
	}
}

func stopGeneration() {
	if generationFlag {
		close(stopChannel)
	}
}

// // func StopGeneration(w http.ResponseWriter, r *http.Request) {
// // 	brokerList := []string{"localhost:29092"}

// // 	//Kafka producer configuration
// // 	config := sarama.NewConfig()
// // 	config.Producer.Return.Successes = true
// // 	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to acknowledge the message
// // 	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
// // 	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms

// // 	kafka.KafkaProducer(brokerList, config, false)

// // 	w.WriteHeader(202)

// // }
