package routes

import (
    "github.com/gin-gonic/gin"
    "open-bounties-api/controllers"
    "open-bounties-api/middleware"
    "open-bounties-api/services"
    "open-bounties-api/models"
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
    "log"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Middleware that applies to all routes
    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    dsn := "host=db user=user password=password dbname=bountydb port=5432 sslmode=disable TimeZone=Europe/Paris"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    db.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Initialize UserService with the database connection
    userService := services.NewUserService(db)

    userController := controllers.NewUserController(userService)
    loginController := controllers.NewLoginController(userService)

    // Versioning API
    v1 := r.Group("/api/v1")
    {
        // Public routes
        public := v1.Group("/")
        {
            public.POST("/login", loginController.Login)
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
/*
            // Issue routes
            issueRoutes := authorized.Group("/issues")
            {
                issueRoutes.GET("/", controllers.GetAllIssues)
                issueRoutes.POST("/", controllers.CreateIssue)
                issueRoutes.GET("/:id", controllers.GetIssue)
                issueRoutes.PUT("/:id", controllers.UpdateIssue)
                issueRoutes.DELETE("/:id", controllers.DeleteIssue)
            }

            // Repository routes
            repoRoutes := authorized.Group("/repositories")
            {
                repoRoutes.GET("/", controllers.GetAllRepositories)
                repoRoutes.POST("/", controllers.CreateRepository)
                repoRoutes.GET("/:id", controllers.GetRepository)
                repoRoutes.PUT("/:id", controllers.UpdateRepository)
                repoRoutes.DELETE("/:id", controllers.DeleteRepository)
            }

            // Bounty routes
            bountyRoutes := authorized.Group("/bounties")
            {
                bountyRoutes.GET("/", controllers.GetAllBounties)
                bountyRoutes.POST("/", controllers.AddBounty)
                bountyRoutes.GET("/:id", controllers.GetBounty)
                bountyRoutes.PUT("/:id", controllers.UpdateBounty)
                bountyRoutes.DELETE("/:id", controllers.DeleteBounty)
            }

            // Organization routes
            orgRoutes := authorized.Group("/organizations")
            {
                orgRoutes.GET("/", controllers.GetAllOrganizations)
                orgRoutes.POST("/", controllers.CreateOrganization)
                orgRoutes.GET("/:id", controllers.GetOrganization)
                orgRoutes.PUT("/:id", controllers.UpdateOrganization)
                orgRoutes.DELETE("/:id", controllers.DeleteOrganization)
            }
	    */
        }
    }

    return r
}

