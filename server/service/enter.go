package service

import (
	"github.com/Luxii44/library/server/service/example"
	"github.com/Luxii44/library/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
