package main

const (
	Running     Status = "running"
	Deploying   Status = "deploying"
	Stopped     Status = "stopped"
	Initialized Status = "initialized"
)

type Status string

type serviceStates map[string]Service

func initServiceState(configs configs) serviceStates {

	serviceStates := make(serviceStates)

	for k, v := range configs {
		serviceStates[k] = Service{
			Status:        Initialized,
			ServiceConfig: v,
		}
	}

	return serviceStates
}
