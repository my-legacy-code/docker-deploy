package main

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
