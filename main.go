package main

import (
	"fmt"
	"log"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/gobroker/brokers"
	"github.com/pmoura-dev/gobroker/middleware"
	"github.com/pmoura-dev/hauto.device-gateway/clients/transaction_service"
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

	transaction_service.Setup("http://localhost:8080")

	mqttConfigurations, err := configuration.GetDeviceMQTTConfigurations()
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

	addActionConsumer(server, turnOnActionQueue, actions.TurnOn)
	addActionConsumer(server, turnOffActionQueue, actions.TurnOff)
	addActionConsumer(server, setRGBColorActionQueue, actions.SetRGBColor)

	// listeners
	for deviceID, deviceConfig := range mqttConfigurations {

		for property, topic := range deviceConfig.Listeners {
			switch property {
			case "state":
				queue := fmt.Sprintf("state.%d.device-gateway.queue", deviceID)
				broker.AddQueue(queue).Bind(defaultTopicExchange, topic)
				addListenerConsumer(server, queue, listeners.State).
					AddParam("device_id", deviceID)
			case "power":
				queue := fmt.Sprintf("power.%d.device-gateway.queue", deviceID)
				broker.AddQueue(queue).Bind(defaultTopicExchange, topic)
				addListenerConsumer(server, queue, listeners.Power).
					AddParam("device_id", deviceID)
			case "energy":
				queue := fmt.Sprintf("energy.%d.device-gateway.queue", deviceID)
				broker.AddQueue(queue).Bind(defaultTopicExchange, topic)
				addListenerConsumer(server, queue, listeners.Energy).
					AddParam("device_id", deviceID)
			}
		}
	}

	fmt.Println("Service is running")
	if err := server.Run("amqp://guest:guest@localhost:5672"); err != nil {
		log.Fatal(err)
	}
}

func addActionConsumer(server *gobroker.Server, queue string, handler gobroker.ConsumerHandlerFunc) *gobroker.Consumer {
	return server.AddConsumer(queue, handler).AddParam("consumer_type", "action")
}

func addListenerConsumer(server *gobroker.Server, queue string, handler gobroker.ConsumerHandlerFunc) *gobroker.Consumer {
	return server.AddConsumer(queue, handler).AddParam("consumer_type", "listener")
}
