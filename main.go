package main

import (
	"demo/src/database"
	productDeps "demo/src/products/dependencies"
	serviceDeps "demo/src/services/dependencies"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Could not establish database connection:", err)
	}
	defer db.Close()

	r := gin.Default()

	// Cargar dependencias de productos
	productsDeps := productDeps.NewProductsDependencies(db)
	productsDeps.Execute(r)

	// Cargar dependencias de servicios
	servicesDeps := serviceDeps.NewServicesDependencies(db)
	servicesDeps.Execute(r)

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}
