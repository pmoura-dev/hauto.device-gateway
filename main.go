package main

import (
	"log"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/gobroker/brokers"
	"github.com/pmoura-dev/gobroker/middleware"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/handlers"
)

const (
	actionsExchange = "actions"

	turnOnActionQueue      = "turn_on.action.device-gateway.queue"
	turnOffActionQueue     = "turn_off.action.device-gateway.queue"
	setRGBColorActionQueue = "set_rgb_color.action.device-gateway.queue"
)

func main() {
	broker := brokers.NewRabbitMQBroker()
	broker.AddExchange(actionsExchange)
	broker.AddQueue(turnOnActionQueue).Bind(actionsExchange, "turn_on.action")
	broker.AddQueue(turnOffActionQueue).Bind(actionsExchange, "turn_off.action")
	broker.AddQueue(setRGBColorActionQueue).Bind(actionsExchange, "set_rgb_color.action")

	server := gobroker.NewServer(broker)

	server.Use(middleware.Logging)
	server.Use(controllers.GetController)

	server.AddConsumer(turnOnActionQueue, handlers.TurnOn)
	server.AddConsumer(turnOffActionQueue, handlers.TurnOff)
	server.AddConsumer(setRGBColorActionQueue, handlers.SetRGBColor)

	if err := server.Run("amqp://guest:guest@localhost:5672"); err != nil {
		log.Fatal(err)
	}
}
