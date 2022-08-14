package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/prasetyodavid/go-stack/config"
	"github.com/prasetyodavid/go-stack/models"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

type RabbitServiceImpl struct {
	rabbitCollection *mongo.Collection
	ctx              context.Context
}

func NewRabbitService(rabbitCollection *mongo.Collection, ctx context.Context) RabbitService {
	return &RabbitServiceImpl{rabbitCollection, ctx}
}

func (p *RabbitServiceImpl) CreatePublisherRabbit(rabbit *models.CreateRabbitRequest) (*models.DBRabbit, error) {
	rabbit.CreateAt = time.Now()
	rabbit.UpdatedAt = rabbit.CreateAt

	var newRabbit *models.DBRabbit

	config, err := config.LoadConfig(".")

	conn, err := amqp.Dial(config.RabbitmqUri)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"golang-queue", //name
		false,          // durable
		false,          //delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		panic(err)
	}

	body, err := json.Marshal(rabbit)
	if err != nil {
		fmt.Println(err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immadiate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(body),
		})
	if err != nil {
		panic(err)
	}

	fmt.Println("sent publish")

	return newRabbit, nil
}

func (p *RabbitServiceImpl) CreateConsumerRabbit(rabbit *models.CreateRabbitRequest) (*models.DBRabbit, error) {
	rabbit.CreateAt = time.Now()
	rabbit.UpdatedAt = rabbit.CreateAt

	var newRabbit *models.DBRabbit

	config, err := config.LoadConfig(".")

	conn, err := amqp.Dial(config.RabbitmqUri)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"golang-queue", //name
		false,          // durable
		false,          //delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	msg, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msg {
			log.Printf("received as message: %s", d.Body)
		}
	}()
	log.Printf("waiting for message. to exit press CTRL+C")
	<-forever

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBRabbit{}, err
		}
		return nil, err
	}

	return newRabbit, nil
}
