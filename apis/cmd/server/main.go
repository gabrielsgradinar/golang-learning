package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabrielsgradinar/golang-learning/apis/configs"
	_ "github.com/gabrielsgradinar/golang-learning/apis/docs"
	"github.com/gabrielsgradinar/golang-learning/apis/internal/entity"
	"github.com/gabrielsgradinar/golang-learning/apis/internal/infra/database"
	"github.com/gabrielsgradinar/golang-learning/apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Go API
// @version         1.0
// @description     Product API with authentication

// @host      localhost:8000
// @BasePath  /
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))
	r.Use(middleware.WithValue("JWTExpiresIn", configs.JWTExpiresIn))
	fmt.Println(configs.JWTExpiresIn)
	// r.Use(LogRequest)
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/generate-token", userHandler.GetJWT)
	})

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8000/docs/doc.json")))

	http.ListenAndServe(":8000", r)
}

// client -> request -> http.Handler -> middleware -> http.Handler

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
