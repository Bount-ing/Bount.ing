package routes

import (
	"log"
	"open-bounties-api/controllers"
	"open-bounties-api/middleware"
	"open-bounties-api/models"
	"open-bounties-api/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Middleware that applies to all routes
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	dsn := "host=db user=user password=password dbname=bountydb port=5432 sslmode=disable TimeZone=Europe/Paris"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Organization{})
	db.AutoMigrate(&models.Repository{})
	db.AutoMigrate(&models.Issue{})
	db.AutoMigrate(&models.Bounty{})
	db.AutoMigrate(&models.Claim{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize UserService with the database connection
	userService := services.NewUserService(db)
	bountyService := services.NewBountyService(db)

	// Initialize controllers
	loginController := controllers.NewLoginController(userService)
	userController := controllers.NewUserController(userService)
	bountyController := controllers.NewBountyController(db, bountyService)

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowWildcard: true,
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"*",
		},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Versioning API
	v1 := r.Group("/api/v1")
	{
		// Public routes
		public := v1.Group("/")
		{
			public.GET("/oauth/github/callback", loginController.GithubCallback)
			public.POST("/register", userController.RegisterUser)
		}

		// Routes that require authentication
		authorized := v1.Group("/")
		authorized.Use(middleware.AuthorizeJWT())
		{
			// User routes
			userRoutes := authorized.Group("/users")
			{
				userRoutes.GET("/:id", userController.GetUser)
				userRoutes.PUT("/:id", userController.UpdateUser)
				userRoutes.DELETE("/:id", userController.DeleteUser)
			}

			// Bounty routes
			bountyRoutes := authorized.Group("/bounties")
			{
				bountyRoutes.OPTIONS("/", func(c *gin.Context) {
					c.Header("Access-Control-Allow-Origin", "*")
					c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
					c.Header("Access-Control-Allow-Headers", "Content-Type")
					c.Header("Content-Type", "application/json")
					c.JSON(200, nil)
				})
				bountyRoutes.GET("/", bountyController.GetAllBounties)
				bountyRoutes.POST("/", bountyController.CreateBounty)
				bountyRoutes.GET("/:id", bountyController.GetBounty)
				bountyRoutes.PUT("/:id", bountyController.UpdateBounty)
				bountyRoutes.DELETE("/:id", bountyController.DeleteBounty)
			}

		}
	}

	return r
}
