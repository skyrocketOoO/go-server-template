package restapi

import (
	"web-server-template/internal/controller/rest"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Binding(r *gin.Engine, d *rest.RestController) {
	r.GET("/ping", d.Ping)
	r.GET("/healthy", d.Healthy)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
