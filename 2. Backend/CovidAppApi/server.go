package main

import (
	"fmt"
	"net/http"

	router "./Router"
	controller "./controllers"
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
	httpRouter.POST("/api/users", userController.Create)
	httpRouter.SERVE(port)
}
