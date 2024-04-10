package memory

import (
	"reflect"
	"sync"

	"github.com/pmoura-dev/hauto.device-gateway/pkg/clients/transaction_service"
	controllers2 "github.com/pmoura-dev/hauto.device-gateway/pkg/controllers"
	"github.com/pmoura-dev/hauto.device-gateway/pkg/states"
)

var deviceManager *DeviceManager = nil

var lock = &sync.Mutex{}

type DeviceManager struct {
	states      map[int]states.State
	controllers map[int]controllers2.Controller

	lock *sync.Mutex
}

func GetDeviceManagerInstance() (*DeviceManager, error) {
	if deviceManager == nil {
		lock.Lock()
		defer lock.Unlock()

		if deviceManager == nil {
			var err error
			deviceManager, err = newDeviceManager()
			if err != nil {
				return nil, err
			}
		}
	}

	return deviceManager, nil
}

func newDeviceManager() (*DeviceManager, error) {
	devicesWithState, err := transaction_service.GetAllDeviceDetails()
	if err != nil {
		return nil, err
	}

	statesMap := map[int]states.State{}
	controllersMap := map[int]controllers2.Controller{}
	for _, d := range devicesWithState {
		statesMap[d.ID] = d.State
		controllersMap[d.ID] = createController(d.Controller, d.NaturalID)
	}

	return &DeviceManager{
		states:      statesMap,
		controllers: controllersMap,
		lock:        &sync.Mutex{},
	}, nil
}

func (m *DeviceManager) UpdateState(deviceID int, newState states.State) (bool, states.State, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	var err error
	var updated bool
	currentState := m.states[deviceID]

	singleParamState, ok := newState.(states.SingleParamState)
	if ok {
		updated, err = currentState.UpdateSingleValue(singleParamState.Name, singleParamState.Value)
		if err != nil {
			return false, nil, err
		}

		m.states[deviceID] = currentState
		return updated, m.states[deviceID], nil
	}

	if reflect.DeepEqual(currentState, newState) {
		return false, nil, nil
	}

	m.states[deviceID] = newState
	return true, m.states[deviceID], nil
}

func (m *DeviceManager) GetControllers() map[int]controllers2.Controller {
	return m.controllers
}

func (m *DeviceManager) GetController(deviceID int) controllers2.Controller {
	return m.controllers[deviceID]
}

func createController(controller string, naturalID string) controllers2.Controller {
	switch controller {
	case controllers2.ShellyColorBulbControllerName:
		return controllers2.NewShellyColorBulbController(naturalID)
	case controllers2.HisenseACControllerName:
		return controllers2.NewHisenseACController(naturalID)
	}

	return nil
}
