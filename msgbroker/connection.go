package msgbroker

import (
	"fmt"

	config "github.com/ItsArul/gomicro/db"
	"github.com/ItsArul/gomicro/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

var channel *amqp.Channel

func GetConnection() *amqp.Channel {
	conf := config.GetConfig()

	logon := fmt.Sprintf("amqp://%s:%s@%s:%s",
		conf.RabbitMQ.User,
		conf.RabbitMQ.Password,
		conf.RabbitMQ.Host,
		conf.RabbitMQ.Port,
	)

	conn, err := amqp.Dial(logon)
	if err != nil {
		utils.PanicIfError(err)
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		utils.PanicIfError(err)
	}

	defer ch.Close()

	channel = ch
	return ch
}
