package main

import (
	"fmt"
	"log"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/gobroker/brokers"
	"github.com/pmoura-dev/gobroker/middleware"
	"github.com/pmoura-dev/hauto.device-gateway/clients/transaction_service"
	"github.com/pmoura-dev/hauto.device-gateway/config"
	"github.com/pmoura-dev/hauto.device-gateway/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/handlers/actions"
	"github.com/pmoura-dev/hauto.device-gateway/handlers/listeners"
	"github.com/pmoura-dev/hauto.device-gateway/mqtt_configuration"
)

const (
	devicesExchange      = "devices"
	defaultTopicExchange = "amq.topic"

	turnOnActionQueue      = "turn_on.action.device-gateway.queue"
	turnOffActionQueue     = "turn_off.action.device-gateway.queue"
	setRGBColorActionQueue = "set_rgb_color.action.device-gateway.queue"

	setHeatingThresholdTemperatureActionQueue = "set_heating_threshold_temperature.action.device-gateway.queue"
	setCoolingThresholdTemperatureActionQueue = "set_cooling_threshold_temperature.action.device-gateway.queue"
	setHeaterCoolerModeActionQueue            = "set_heater_cooler_mode.action.device-gateway.queue"
)

func main() {

	transactionServiceConfig := config.GetTransactionServiceConfig()
	transaction_service.Setup(
		fmt.Sprintf("http://%s:%s",
			transactionServiceConfig.Host,
			transactionServiceConfig.Port,
		),
	)

	mqttConfigurations, err := mqtt_configuration.GetDeviceMQTTConfigurations()
	if err != nil {
		log.Fatal(err)
	}

	broker := brokers.NewRabbitMQBroker()
	broker.AddExchange(devicesExchange)
	broker.AddQueue(turnOnActionQueue).Bind(devicesExchange, "turn_on.action")
	broker.AddQueue(turnOffActionQueue).Bind(devicesExchange, "turn_off.action")
	broker.AddQueue(setRGBColorActionQueue).Bind(devicesExchange, "set_rgb_color.action")
	broker.AddQueue(setHeatingThresholdTemperatureActionQueue).
		Bind(devicesExchange, "set_heating_threshold_temperature.action")
	broker.AddQueue(setCoolingThresholdTemperatureActionQueue).
		Bind(devicesExchange, "set_cooling_threshold_temperature.action")
	broker.AddQueue(setHeaterCoolerModeActionQueue).
		Bind(devicesExchange, "set_heater_cooler_mode.action")

	server := gobroker.NewServer(broker)

	server.Use(middleware.Logging)
	server.Use(controllers.GetController)

	addActionConsumer(server, turnOnActionQueue, actions.TurnOn)
	addActionConsumer(server, turnOffActionQueue, actions.TurnOff)
	addActionConsumer(server, setRGBColorActionQueue, actions.SetRGBColor)
	addActionConsumer(server, setHeatingThresholdTemperatureActionQueue, actions.SetHeatingThresholdTemperature)
	addActionConsumer(server, setCoolingThresholdTemperatureActionQueue, actions.SetCoolingThresholdTemperature)
	addActionConsumer(server, setHeaterCoolerModeActionQueue, actions.SetHeaterCoolerMode)

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

	rabbitMQConfig := config.GetRabbitMQConfig()
	if err := server.Run(
		fmt.Sprintf("amqp://%s:%s@%s:%s",
			rabbitMQConfig.User,
			rabbitMQConfig.Password,
			rabbitMQConfig.Host,
			rabbitMQConfig.Port,
		),
	); err != nil {
		log.Fatal(err)
	}
}

func addActionConsumer(server *gobroker.Server, queue string, handler gobroker.ConsumerHandlerFunc) *gobroker.Consumer {
	return server.AddConsumer(queue, handler).AddParam("consumer_type", "action")
}

func addListenerConsumer(server *gobroker.Server, queue string, handler gobroker.ConsumerHandlerFunc) *gobroker.Consumer {
	return server.AddConsumer(queue, handler).AddParam("consumer_type", "listener")
}
