package sfmeshutil

import (
	mesh "github.com/Azure/azure-sdk-for-go/services/preview/servicefabricmesh/mgmt/2018-09-01-preview/servicefabricmesh"
	"github.com/pothulapati/sfmesh-controller/pkg/apis/sfmesh/v1alpha1"
)

func CovertApplication(application v1alpha1.Application) (*mesh.ApplicationResourceDescription, error) {

	var MeshServices []mesh.ServiceResourceDescription
	for _, service := range application.Spec.Services {

		//Get All Code Packages for this service
		var codepkgs []mesh.ContainerCodePackageProperties

		for _, codepkg := range service.CodePackages {
			var meshEndpoints []mesh.EndpointProperties
			for _, endpoint := range codepkg.EndPoints {
				meshEndpoint := mesh.EndpointProperties{
					Name: &endpoint.Name,
					Port: &endpoint.Port,
				}

				meshEndpoints = append(meshEndpoints, meshEndpoint)
			}

			meshcodepkg := mesh.ContainerCodePackageProperties{
				Name:      &codepkg.Name,
				Image:     &codepkg.Image,
				Endpoints: &meshEndpoints,
			}
			codepkgs = append(codepkgs, meshcodepkg)
		}

		meshService := mesh.ServiceResourceDescription{
			Name: &service.Name,

			ServiceResourceProperties: &mesh.ServiceResourceProperties{
				OsType:       mesh.Linux,
				CodePackages: &codepkgs,
				ReplicaCount: &service.ReplicaCount,
			},
		}
		MeshServices = append(MeshServices, meshService)
	}

	applicationRP := &mesh.ApplicationResourceProperties{
		Description: &application.Spec.Description,
		Services:    &MeshServices,
	}

	appResourceDescription := mesh.ApplicationResourceDescription{
		Name:                          &application.Name,
		Location:                      &application.Spec.Location,
		ApplicationResourceProperties: applicationRP,
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
