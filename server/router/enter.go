package router

import (
	"github.com/Luxii44/library/server/router/example"
	"github.com/Luxii44/library/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
