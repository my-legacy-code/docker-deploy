package main

import (
	"log"
)

const (
	Running     Status = "running"
	Deploying   Status = "deploying"
	Stopped     Status = "stopped"
	Initialized Status = "initialized"
)

type Status string

type AppState struct {
	ServiceStates ServiceStates
	Clients       map[string]Client
}

type ServiceStates map[string]*Service

func initAppState(serviceConfig configs, errLogger *log.Logger) *AppState {
	appState := new(AppState)
	appState.ServiceStates = initServiceState(serviceConfig)

	for imageName, serviceState := range appState.ServiceStates {
		containerIds, err := getContainerIds(imageName)
		if err != nil {
			errLogger.Println(err)
			continue
		}

		if len(containerIds) < 1 {
			continue
		}

		isRunning, err := isContainerRunning(containerIds[0])
		if err != nil {
			errLogger.Println(err)
			continue
		}
		serviceState.Status = boolToStatus(isRunning)
	}

	appState.Clients = make(map[string]Client)
	return appState
}

func boolToStatus(isRunning bool) Status {
	if isRunning {
		return Running
	}
	return Stopped
}

func initServiceState(configs configs) ServiceStates {
	serviceStates := make(ServiceStates)

	for k, v := range configs {
		serviceStates[k] = &Service{
			Status:        Initialized,
			ServiceConfig: v,
		}
	}

	return serviceStates
}

const (
	InitialServiceStates WSMessageType = "initial_service_states"
	UpdateServiceState   WSMessageType = "update_service_state"
)

type WSMessageType string

type WSMessage struct {
	Type WSMessageType `json:"type"`
	Body interface{}   `json:"body"`
}

func sendInitialServiceStates(username string, appState *AppState) {
	message := WSMessage{
		Type: InitialServiceStates,
		Body: appState.ServiceStates,
	}
	appState.Clients[username].Conn.WriteJSON(message)
}

func updateServiceState(service *Service, appState *AppState) {
	message := WSMessage{
		Type: UpdateServiceState,
		Body: service,
	}
	for username, client := range appState.Clients {
		log.Printf("Pushing new service state to %s", username)
		client.Conn.WriteJSON(message)
	}
}
