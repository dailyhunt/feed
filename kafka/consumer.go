package kafka

import (
	"fmt"
	"github.com/bsm/sarama-cluster"
	"github.com/dailyhunt/feed/util"
	"github.com/pkg/errors"
	"strings"
	"time"
)

const (
	Space = " "
	Comma = ","
)

type ConsumerConfig struct {
	brokers      string
	topics       string
	groupId      string
	chanCapacity int
}

func (config ConsumerConfig) OK() (bool, error) {
	if config.brokers == "" {
		return false, errors.New("Broker list should not be empty")
	}

	if config.topics == "" {
		return false, errors.New("Topic list should not be empty")
	}

	if config.groupId == "" {
		return false, errors.New("Consumer group id should not be empty")
	}

	if config.chanCapacity < 1 {
		return false, errors.New("Channel capacity can not be less than 1")
	}

	return true, nil

}

func NewConsumerConfig(brokers, topics, groupId string, chanCapacity int) ConsumerConfig {
	return ConsumerConfig{
		brokers:      brokers,
		topics:       topics,
		groupId:      groupId,
		chanCapacity: chanCapacity,
	}
}

// TODO: LOGGING
func NewConsumer(config ConsumerConfig, done <-chan struct{}) (outKvChan <-chan *util.KV, err error) {
	if ok, err := config.OK(); !ok {
		return nil, err
	}

	var consumer *cluster.Consumer
	if consumer, err = initKafkaConsumer(config); err != nil {
		return nil, err
	}

	kvChan := make(chan *util.KV, config.chanCapacity)
	go consumeMessages(consumer, kvChan, done)

	return kvChan, nil

}
func consumeMessages(consumer *cluster.Consumer, kvs chan *util.KV, done <-chan struct{}) {
	defer closeConsumer(consumer)
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				kvs <- util.NewKV(msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case <-done:
			fmt.Println("Done signal received . Will stop consuming messages.")
			time.Sleep(2 * time.Second) // Give some time for committing offset
			close(kvs)
			return
		}
	}
}
func closeConsumer(consumer *cluster.Consumer) {
	fmt.Println("Closing consumer")
	consumer.Close()
}

func initKafkaConsumer(config ConsumerConfig) (*cluster.Consumer, error) {
	clusterConfig := cluster.NewConfig()
	brokers := strings.Split(strings.Trim(config.brokers, Space), Comma)
	topics := strings.Split(strings.Trim(config.topics, Space), Comma)

	fmt.Println(fmt.Sprintf("Brokers = %v", brokers))
	fmt.Println(fmt.Sprintf("Topics  = %v", topics))

	consumer, err := cluster.NewConsumer(brokers, config.groupId, topics, clusterConfig)
	if err != nil {
		return nil, err
	}

	go consumeErrors(consumer)
	go consumeNotification(consumer)

	return consumer, nil
}

func consumeNotification(consumer *cluster.Consumer) {
	for ntf := range consumer.Notifications() {
		// TODO: To log
		fmt.Println(fmt.Sprintf("Rebalanced: %+v", ntf))
	}
}

func consumeErrors(consumer *cluster.Consumer) {
	for err := range consumer.Errors() {
		// TODO: To log
		fmt.Println(fmt.Sprintf("Error: %s", err.Error()))
	}
}
