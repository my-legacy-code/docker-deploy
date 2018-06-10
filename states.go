package main

import "fmt"

const (
	Running     Status = "running"
	Deploying   Status = "deploying"
	Stopped     Status = "stopped"
	Initialized Status = "initialized"
)

type Status string

type AppState struct {
	ServiceStates ServiceStates
	Clients map[string] Client
}

type ServiceStates map[string]*Service

func initAppState(serviceConfig configs) *AppState {
	appState := new(AppState)
	appState.ServiceStates = initServiceState(serviceConfig)
	appState.Clients = make(map[string]Client)
	return appState
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
	UpdateServiceState WSMessageType = "update_service_state"
)  

type WSMessageType string

type WSMessage struct {
	Type WSMessageType `json:"type"`
	Body interface{} `json:"body"`
}

func sendInitialServiceStates(username string, appState *AppState)  {
	message := WSMessage{
		Type: InitialServiceStates,
		Body: appState.ServiceStates,

	}
	appState.Clients[username].Conn.WriteJSON(message)
}

func updateServiceState(service *Service, appState *AppState)  {
	message := WSMessage{
		Type: UpdateServiceState,
		Body: service,
		
	}
	for username, client := range appState.Clients {
		log(fmt.Sprintf("Pushing new service state to %s", username))
		client.Conn.WriteJSON(message)
	}
}