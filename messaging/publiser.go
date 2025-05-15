package rabbitmq

import (
	"context"
	"filesender/config"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

// NewPublisher cria uma nova instância do publisher RabbitMQ
func NewPublisher(config *config.Config) (*Publisher, error) {
	conn, err := amqp.Dial(config.RabbitMQURI)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"file_queue", // nome
		true,         // durável
		false,        // delete when unused
		false,        // exclusiva
		false,        // no-wait
		nil,          // argumentos
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, err
	}

	return &Publisher{
		conn:    conn,
		channel: ch,
		queue:   q.Name,
	}, nil
}

// PublishFile envia um arquivo como binário para a fila
func (p *Publisher) PublishFile(fileData []byte, filename string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return p.channel.PublishWithContext(
		ctx,
		"",      // exchange
		p.queue, // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/octet-stream",
			Body:        fileData,
			Headers: amqp.Table{
				"filename": filename,
			},
		},
	)
}

// Close fecha a conexão com o RabbitMQ
func (p *Publisher) Close() error {
	if err := p.channel.Close(); err != nil {
		return err
	}
	return p.conn.Close()
}
