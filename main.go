package main

import (
	"fmt"
	"log"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/gobroker/brokers"
	"github.com/pmoura-dev/gobroker/middleware"
	"github.com/pmoura-dev/hauto.device-gateway/configuration"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/handlers/actions"
	"github.com/pmoura-dev/hauto.device-gateway/handlers/listeners"
)

const (
	devicesExchange      = "devices"
	defaultTopicExchange = "amq.topic"

	turnOnActionQueue      = "turn_on.action.device-gateway.queue"
	turnOffActionQueue     = "turn_off.action.device-gateway.queue"
	setRGBColorActionQueue = "set_rgb_color.action.device-gateway.queue"
)

func main() {
	listenersConfig, err := configuration.LoadDeviceListeners()
	if err != nil {
		log.Fatal(err)
	}

	broker := brokers.NewRabbitMQBroker()
	broker.AddExchange(devicesExchange)
	broker.AddQueue(turnOnActionQueue).Bind(devicesExchange, "turn_on.action")
	broker.AddQueue(turnOffActionQueue).Bind(devicesExchange, "turn_off.action")
	broker.AddQueue(setRGBColorActionQueue).Bind(devicesExchange, "set_rgb_color.action")

	server := gobroker.NewServer(broker)

	server.Use(middleware.Logging)
	server.Use(controllers.GetController)

	server.AddConsumer(turnOnActionQueue, actions.TurnOn)
	server.AddConsumer(turnOffActionQueue, actions.TurnOff)
	server.AddConsumer(setRGBColorActionQueue, actions.SetRGBColor)

	// listeners
	for _, l := range listenersConfig.Devices {
		for _, t := range l.Topics {
			switch t.Type {
			case "status":
				queue := fmt.Sprintf("status.%s.device-gateway.queue", l.NaturalID)
				broker.AddQueue(queue).Bind(defaultTopicExchange, t.Name)
				server.AddConsumer(queue, listeners.Status).AddParam("natural_id", l.NaturalID)
			case "power":
				queue := fmt.Sprintf("power.%s.device-gateway.queue", l.NaturalID)
				broker.AddQueue(queue).Bind(defaultTopicExchange, t.Name)
				server.AddConsumer(queue, listeners.Power).AddParam("natural_id", l.NaturalID)
			case "energy":
				queue := fmt.Sprintf("energy.%s.device-gateway.queue", l.NaturalID)
				broker.AddQueue(queue).Bind(defaultTopicExchange, t.Name)
				server.AddConsumer(queue, listeners.Energy).AddParam("natural_id", l.NaturalID)
			}
		}
	}

	if err := server.Run("amqp://guest:guest@localhost:5672"); err != nil {
		log.Fatal(err)
	}
}
