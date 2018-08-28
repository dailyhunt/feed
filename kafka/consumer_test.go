package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"testing"
)

func TestNewKafkaConsumer(t *testing.T) {
	// Test Only on local
	//Call()
}

func Call() {

	// Initialize Kafka
	done := make(chan struct{})
	conf := NewConsumerConfig("127.0.0.1:9092", "test1", "gid-0031", 2)
	kvChan, err := NewConsumer(conf, done)

	if err != nil {
		panic(err)
	}

	// consume messages, watch signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	count := 0

	for {
		select {
		case m, ok := <-kvChan:
			if ok {
				fmt.Println("Count ", count, "   Consumed -> ", m.GetStringValue())
				count = count + 1
				if count > 3 {
					done <- struct{}{}
				}
			} else {
				return
			}
		case <-signals:
			done <- struct{}{}
		}

	}

}
