package main

import (
	"net/http"

	"github.com/siroger/gowebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3100", nil)
}
