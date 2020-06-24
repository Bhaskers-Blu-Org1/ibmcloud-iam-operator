package controller

import (
	"github.ibm.com/seed/ibmcloud-iam-operator/pkg/controller/accessgroup"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, accessgroup.Add)
}
