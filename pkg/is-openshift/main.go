package is_openshift

import (
	"fmt"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/rest"
)

// IsOpenShift checks if the client is connected to an OpenShift server
func IsOpenShift(config *rest.Config) (bool, error) {
	dcl, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return false, fmt.Errorf("there was an error while checking if connected to an OpenShift cluster: %w", err)
	}

	apiList, err := dcl.ServerGroups()
	if err != nil {
		return false, fmt.Errorf("there was an error while checking if connected to an OpenShift cluster: %w", err)
	}

	apiGroups := apiList.Groups
	for i := 0; i < len(apiGroups); i++ {
		if apiGroups[i].Name == "config.openshift.io" {
			return true, nil
		}
	}
	return false, nil
}