package main

import (
	"fmt"
	"log"

	"github.com/pmoura-dev/gobroker"
	"github.com/pmoura-dev/gobroker/brokers"
	"github.com/pmoura-dev/gobroker/middleware"
	"github.com/pmoura-dev/hauto.device-gateway/api"
	"github.com/pmoura-dev/hauto.device-gateway/api/action_requests/handlers"
	listeners_handlers "github.com/pmoura-dev/hauto.device-gateway/api/device_listeners/handlers"
	"github.com/pmoura-dev/hauto.device-gateway/config"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/clients/transaction_service"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/memory"
)

func setup() (*gobroker.Server, error) {
	broker := brokers.NewRabbitMQBroker()
	server := gobroker.NewServer(broker)

	broker.AddExchange(api.ExchangeDevices)

	broker.AddQueue(api.QueueActionRequestedTurnOn).Bind(api.ExchangeDevices, api.TopicActionRequestedTurnOn)
	broker.AddQueue(api.QueueActionRequestedTurnOff).Bind(api.ExchangeDevices, api.TopicActionRequestedTurnOff)

	server.AddConsumer(api.QueueActionRequestedTurnOn, handlers.TurnOn)
	server.AddConsumer(api.QueueActionRequestedTurnOff, handlers.TurnOff)

	manager, err := memory.GetDeviceManagerInstance()
	if err != nil {
		return nil, err
	}

	for id, controller := range manager.GetControllers() {
		if controller == nil {
			continue
		}

		listenerController, ok := controller.(controllers.StateListener)
		if !ok {
			continue
		}

		queue := fmt.Sprintf(api.QueueStateUpdatedFormat, id)
		broker.AddQueue(queue).Bind(api.ExchangeDefaultTopic, listenerController.GetStateTopic())
		server.AddConsumer(queue, listeners_handlers.StateUpdated).
			AddParam("device_id", id).
			AddParam("controller", listenerController)
	}

	server.Use(middleware.Logging)

	return server, err
}

func main() {

	transactionServiceConfig := config.GetTransactionServiceConfig()
	transaction_service.Setup(
		fmt.Sprintf("http://%s:%s",
			transactionServiceConfig.Host,
			transactionServiceConfig.Port,
		),
	)

	rabbitMQConfig := config.GetRabbitMQConfig()

	server, err := setup()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run(
		fmt.Sprintf("amqp://%s:%s@%s:%s",
			rabbitMQConfig.User,
			rabbitMQConfig.Password,
			rabbitMQConfig.Host,
			rabbitMQConfig.Port,
		),
	)
	if err != nil {
		log.Fatal(err)
	}
}
