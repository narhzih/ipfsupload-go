package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Api struct {
	Router         *gin.Engine
	Logger         zerolog.Logger
	BaseRouteGroup *gin.RouterGroup
}

func NewApi(routeGroupString string, logger zerolog.Logger) *Api {
	r := gin.Default()
	init := &Api{
		Router:         r,
		Logger:         logger,
		BaseRouteGroup: r.Group(routeGroupString),
	}
	init.useCORsMiddleware()
	init.bootRoutes()
	return init
}

func (api *Api) useCORsMiddleware() {
	// In case there's possibility for a more complicated cors configuration
	// in the future
	api.Router.Use(cors.Default())
}

func (api *Api) bootRoutes() {
	api.loadUploadRoutes()
}
