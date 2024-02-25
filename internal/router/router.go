package router

import (
	"app/main/utils"
	"net/http"

	"github.com/gin-gonic/gin"

	cors "github.com/rs/cors/wrapper/gin"
)

type Router struct {
	engine *gin.Engine

	ip string
}

func InitRouter(ip string) Router {

	router := Router{
		engine: gin.Default(),
		ip:     ip,
	}

	router.engine.Use(cors.Default())
	router.setupRouting()

	return router
}

func (r *Router) setupRouting() {

	r.engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ping report service",
		})
	})
}

func (r *Router) Run() {
	env := utils.Env()

	r.engine.RunTLS(r.ip,
		env.Value("SSL_CERT_PATH"),
		env.Value("SSL_KEY_PATH"))
}
