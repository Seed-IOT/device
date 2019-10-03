package web

import (
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (srv *server) routes() http.Handler {
	//Declare web routing table at here.

	url := ginSwagger.URL("http://localhost:8070/swagger/doc.json") // The url pointing to API definition
	srv.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	srv.router.POST("protocol", srv.CreateProtocol)

	return srv.router
}
