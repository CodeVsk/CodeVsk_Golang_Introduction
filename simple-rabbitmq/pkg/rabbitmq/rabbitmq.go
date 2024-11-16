package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel) {
	conn, err := amqp.Dial("amqp://dev:dev@localhost:5672");
	if err != nil {
		panic("Failed to connect in RabbitMQ.")
	}
	
	ch, err := conn.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}

	return ch
}

func Publish(ch *amqp.Channel, queueName string, message []byte) (error){
	_, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = ch.Publish("", 
		queueName, 
		false, 
		false, 
		amqp.Publishing{
			ContentType: "text/plain",
			Body: message,
		})
		if err != nil {
			return err
		}

	return nil
}

func Consume(ch *amqp.Channel, queueName string, callback chan<- amqp.Delivery) error {
	_, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queueName, 
		"", 
		false, //Don't use in production
		false, 
		false, 
		false, 
		nil,
	)

	if(err != nil) {
		return err
	}

	for msg := range msgs {
		callback <- msg
	}

	return nil
}