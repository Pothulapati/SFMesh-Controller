package sfmeshutil

import (
	mesh "github.com/Azure/azure-sdk-for-go/services/preview/servicefabricmesh/mgmt/2018-09-01-preview/servicefabricmesh"
	"github.com/pothulapati/sfmesh-controller/pkg/apis/sfmesh/v1alpha1"
)

func CovertApplication(application v1alpha1.Application) (*mesh.ApplicationResourceDescription, error) {

	appResourceDescription := mesh.ApplicationResourceDescription{
		Name: &application.Name,
	}
	return &appResourceDescription, nil
}

func ConvertNetwork(network v1alpha1.Network) (*mesh.NetworkResourceDescription, error) {

	basicNetworkProperties := mesh.LocalNetworkResourceProperties{
		NetworkAddressPrefix: &network.Spec.NetworkAddressPrefix,
		Description:          &network.Spec.Description,
	}
	networkRD := mesh.NetworkResourceDescription{
		Name:       &network.Name,
		Location:   &network.Spec.Location,
		Tags:       network.Spec.Tags,
		Properties: &basicNetworkProperties,
	}
	return &networkRD, nil
}
