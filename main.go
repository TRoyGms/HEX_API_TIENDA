package main

import (
	"demo/src/database"
	productDeps "demo/src/products/dependencies"
	serviceDeps "demo/src/services/dependencies"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// Conectar con la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Could not establish database connection:", err)
	}
	defer db.Close()

	// Inicializar Gin
	r := gin.Default()

	// Configuración de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Cargar dependencias de productos
	productsDeps := productDeps.NewProductsDependencies(db)
	productsDeps.Execute(r)

	// Cargar dependencias de servicios
	servicesDeps := serviceDeps.NewServicesDependencies(db)
	servicesDeps.Execute(r)

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}
