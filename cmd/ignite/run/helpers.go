package run

import (
	api "github.com/weaveworks/ignite/pkg/apis/ignite"
	"github.com/weaveworks/ignite/pkg/client"
	"github.com/weaveworks/ignite/pkg/filter"
)

// TODO: This
func getVMForMatch(vmMatch string) (*api.VM, error) {
	return client.VMs().Find(filter.NewIDNameFilter(vmMatch))
}

// TODO: This
func getVMsForMatches(vmMatches []string) ([]*api.VM, error) {
	allVMs := make([]*api.VM, 0, len(vmMatches))
	for _, match := range vmMatches {
		vm, err := getVMForMatch(match)
		if err != nil {
			return nil, err
		}
		allVMs = append(allVMs, vm)
	}
	return allVMs, nil
}

func getAllVMs() ([]*api.VM, error) {
	return client.VMs().FindAll(filter.NewAllFilter())
}
