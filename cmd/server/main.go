package main

import (
	"jwt-auth-crud/internal/config"
	"jwt-auth-crud/internal/database"
	"jwt-auth-crud/internal/handlers"
	"jwt-auth-crud/internal/middleware"
	"jwt-auth-crud/internal/repositories"
	"jwt-auth-crud/internal/services"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Load config
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	productService := services.NewProductService(productRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)

	// Initialize middleware
	jwtMiddleware := middleware.NewJWTMiddleware(cfg.JWTSecret)

	// Setup router
	r := gin.Default()

	r.Use(middleware.RateLimiterMiddleware("10-M"))

	// Auth routes
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		
		// Apply rate limiter only to login endpoint
		auth.POST("/login", 
			middleware.RateLimiterMiddleware("10-M"), // 10 requests per minute
			authHandler.Login)
	}

	// Protected product routes
	api := r.Group("/api")
	api.Use(jwtMiddleware.ValidateToken())
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.POST("", productHandler.CreateProduct)
			products.GET("/all", productHandler.GetAllProducts)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}
	}

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
