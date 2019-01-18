package sfmeshutil

import (
	"github.com/pothulapati/sfmesh-controller/pkg/apis/sfmesh/v1alpha1"
	mesh "github.com/Azure/azure-sdk-for-go/services/preview/servicefabricmesh/mgmt/2018-09-01-preview/servicefabricmesh"
	)

func CovertApplication(application v1alpha1.Application) (*mesh.ApplicationResourceDescription, error) {

	appResourceDescription := mesh.ApplicationResourceDescription{
		Name: &application.Name,

	}
	return &appResourceDescription, nil
}
