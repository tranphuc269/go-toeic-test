package routers

import (
	v1 "english_exam_go/application/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	ac *v1.AuthController
}

func (ar AuthRouter) routing(e *gin.RouterGroup) {
	//TODO implement me
	be := e.Group("/auth")
	{
		be.POST("/login", ar.ac.Login())
		be.POST("/register", ar.ac.Register())
	}
}

func CreateAuthRouter(ac *v1.AuthController) Router {
	return AuthRouter{ac: ac}
}
