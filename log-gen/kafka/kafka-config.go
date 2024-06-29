package kafka

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
)

//Create Producer Instance
func KafkaProducer(broker []string, config *sarama.Config, data []byte) {

	producer, err := sarama.NewSyncProducer(broker, config)
	if err != nil {
		log.Fatal("Error creating Kafka producer:", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal("Error closing Kafka producer:", err)
		}
	}()
	//Kafka topic name
	topic := "codespotify-topic"
	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data), //CH for changing Encoding
	}

	// Send the message to Kafka
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		fmt.Println("Error sending message to Kafka:", err)
	}

	fmt.Printf("Message sent to topic %s, partition %d, offset %d\n", topic, partition, offset)

}

//Exit The Producer Gracefully
func GracefullyExit() {

	flagChan := make(chan bool, 1)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	select {
	case <-sigchan:
		fmt.Println("Interrupting .....")
	case <-flagChan:
		fmt.Println("Interrupting .....")
	}

}
