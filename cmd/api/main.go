package main

import (
	"fmt"
	"net/http"

	"github.com/dayli-fragrance/api/internal/application/usecase"
	"github.com/dayli-fragrance/api/internal/infrastructure/database"
	"github.com/dayli-fragrance/api/internal/presentation/handler"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	var err = godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar .env:", err)
		return
	}

	var connection, connectionErr = database.NewPostgresConnection()
	if connectionErr != nil {
		fmt.Println("Erro ao conectar no banco:", connectionErr)
		return
	}
	defer connection.Close(nil)

	fmt.Println("Banco conectado!")

	var productRepository = database.NewProductRepositoryPostgres(connection)
	var getProductsUseCase = usecase.NewGetProductsUseCase(productRepository)
	var productHandler = handler.NewProductHandler(getProductsUseCase)

	var router = chi.NewRouter()

	router.Get("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	router.Get("/api/products", productHandler.GetAll)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", router)
}
