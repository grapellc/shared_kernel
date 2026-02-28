package messaging

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v3/pkg/amqp"
	"github.com/spf13/viper"
)

var (
	Subscriber *amqp.Subscriber
	Publisher  *amqp.Publisher

	Logger watermill.LoggerAdapter
)

func NewPubSub() {
	Logger = watermill.NewStdLogger(false, false)

	amqpConfig := amqp.NewDurablePubSubConfig(viper.GetString("amqp.uri"), amqp.GenerateExchangeNameTopicName)

	sub, err := amqp.NewSubscriber(amqpConfig, Logger)
	if err != nil {
		log.Fatal(err)
	}
	Subscriber = sub

	pub, err := amqp.NewPublisher(amqpConfig, Logger)
	if err != nil {
		log.Fatal(err)
	}
	Publisher = pub
}

func Close() {
	if Subscriber != nil {
		Subscriber.Close()
	}
	if Publisher != nil {
		Publisher.Close()
	}
}
