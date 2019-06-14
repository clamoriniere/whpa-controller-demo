package controller

import (
	"github.com/datadog/whpa/pkg/controller/watermarkhorizontalpodautoscaler"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, watermarkhorizontalpodautoscaler.Add)
}
