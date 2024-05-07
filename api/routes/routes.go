package routes

import (
	"log"
	"open-bounties-api/controllers"
	"open-bounties-api/middleware"
	"open-bounties-api/models"
	"open-bounties-api/services"

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
	db.AutoMigrate(&models.Bounty{})
	db.AutoMigrate(&models.Issue{})
	db.AutoMigrate(&models.Repository{})
	db.AutoMigrate(&models.Organization{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Initialize UserService with the database connection
	userService := services.NewUserService(db)
	bountyService := services.NewBountyService(db)
	issueService := services.NewIssueService(db)
	repositoryService := services.NewRepositoryService(db)
	organizationService := services.NewOrganizationService(db)

	// Initialize controllers
	loginController := controllers.NewLoginController(userService)
	userController := controllers.NewUserController(userService)
	bountyController := controllers.NewBountyController(bountyService)
	issueController := controllers.NewIssueController(issueService)
	repositoryController := controllers.NewRepositoryController(repositoryService)
	organizationController := controllers.NewOrganizationController(organizationService)

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

			// Bounty routes
			bountyRoutes := authorized.Group("/bounties")
			{
				bountyRoutes.GET("/", bountyController.GetAllBounties)
				bountyRoutes.POST("/", bountyController.CreateBounty)
				bountyRoutes.GET("/:id", bountyController.GetBounty)
				bountyRoutes.PUT("/:id", bountyController.UpdateBounty)
				bountyRoutes.DELETE("/:id", bountyController.DeleteBounty)
			}

			// Issue routes
			issueRoutes := authorized.Group("/issues")
			{
				issueRoutes.GET("/", issueController.GetAllIssues)
				issueRoutes.POST("/", issueController.CreateIssue)
				issueRoutes.GET("/:id", issueController.GetIssue)
				issueRoutes.PUT("/:id", issueController.UpdateIssue)
				issueRoutes.DELETE("/:id", issueController.DeleteIssue)
			}

			// Repository routes
			repositoryRoutes := authorized.Group("/repositories")
			{
				repositoryRoutes.GET("/", repositoryController.GetAllRepositories)
				repositoryRoutes.POST("/", repositoryController.CreateRepository)
				repositoryRoutes.GET("/:id", repositoryController.GetRepository)
				repositoryRoutes.PUT("/:id", repositoryController.UpdateRepository)
				repositoryRoutes.DELETE("/:id", repositoryController.DeleteRepository)
			}

			// Organization routes
			organizationRoutes := authorized.Group("/organizations")
			{
				organizationRoutes.GET("/", organizationController.GetAllOrganizations)
				organizationRoutes.POST("/", organizationController.CreateOrganization)
				organizationRoutes.GET("/:id", organizationController.GetOrganization)
				organizationRoutes.PUT("/:id", organizationController.UpdateOrganization)
				organizationRoutes.DELETE("/:id", organizationController.DeleteOrganization)
			}
		}
	}

	return r
}