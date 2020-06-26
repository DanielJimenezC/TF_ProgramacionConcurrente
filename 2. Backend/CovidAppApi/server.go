package main

import (
	"fmt"
	"net/http"

	controller "./api/controllers"
	router "./config/router"
)

var (
	httpRouter     router.Router              = router.MuxRouter()
	userController controller.IUserController = controller.UserController()
)

func main() {

	const port string = ":5000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is up and running...")
	})

	httpRouter.GET("/api/users", userController.GetAll)
	httpRouter.GET("/api/users/{id}", userController.GetByID)
	httpRouter.PUT("/api/users/{id}", userController.Update)
	httpRouter.POST("/api/users/signup", userController.SignUp)
	httpRouter.DELETE("/api/users/{id}", userController.Delete)
	httpRouter.POST("/api/users/login", userController.Login)
	httpRouter.SERVE(port)
}
