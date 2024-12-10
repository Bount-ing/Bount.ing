package routes

import (
	"log"
	"os"
	"time"

	"github.com/bount-ing/bount.ing/api/auth"
	"github.com/bount-ing/bount.ing/api/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use()
	log.Println("Setting up routes")
	log.Println("APP_BASE_URL: ", os.Getenv("APP_BASE_URL"))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("APP_BASE_URL")},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-CSRF-Token", "Accept", "Origin", "Cache-Control", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1 := r.Group("/v1")

	// Auth protected routes
	needAuth := v1.Group("/", auth.AuthRequired())
	needAdmin := v1.Group("/", auth.AdminRequired())

	webhooks := v1.Group("/webhooks")
	public := v1.Group("/")

	github := webhooks.Group("/github")
	github.POST("/repos/:repo_id", handlers.IncomingRepositoryWebhook)

	public.POST("/signup", handlers.CreateUser)
	public.POST("/signup/verif/:code", handlers.ValidateUserCode)
	public.POST("/signup/password", handlers.CreateUserPassword)
	public.POST("/signin", handlers.Signin)

	public.GET("/oauth/github/callback", handlers.OAuthGithubCallback)
	public.GET("/oauth/stripe/callback", handlers.OAuthStripeCallback)
	public.POST("/register", handlers.CreateUser)

	user := needAuth.Group("/users")
	user.GET("/me", handlers.GetCurrentUser)
	user.POST("/stripe", handlers.ConnectStripe)

	adminBounties := needAdmin.Group("/bounties")
	adminBounties.GET("/unconfirmed", handlers.GetAllUnconfirmedBounties)
	adminBounties.PUT("/finalize/:id", handlers.FinalizeBounty)

	public.GET("/claims", handlers.GetClaims)
	public.GET("/claims/:id", handlers.GetClaim)
	public.POST("/claims", handlers.CreateClaim)
	public.PUT("/claims/:id", handlers.UpdateClaim)
	public.DELETE("/claims/:id", handlers.DeleteClaim)

	public.GET("/bounties/", handlers.GetBounties)
	public.GET("/bounties/:id", handlers.GetBounty)
	public.POST("/bounties", handlers.CreateBounty)
	public.PUT("/bounties/:id", handlers.UpdateBounty)
	public.DELETE("/bounties/:id", handlers.DeleteBounty)

	public.GET("/issues", handlers.GetIssues)
	public.GET("/issues/:id", handlers.GetIssue)
	public.POST("/issues", handlers.CreateIssue)
	public.PUT("/issues/:id", handlers.UpdateIssue)
	public.DELETE("/issues/:id", handlers.DeleteIssue)

	public.GET("/repositories", handlers.GetRepositories)
	public.GET("/repositories/:id", handlers.GetRepository)
	public.POST("/repositories", handlers.CreateRepository)
	public.PUT("/repositories/:id", handlers.UpdateRepository)
	public.DELETE("/repositories/:id", handlers.DeleteRepository)

	public.GET("/hosts", handlers.GetHosts)
	public.GET("/hosts/:id", handlers.GetHost)
	public.POST("/hosts", handlers.CreateHost)
	public.PUT("/hosts/:id", handlers.UpdateHost)
	public.DELETE("/hosts/:id", handlers.DeleteHost)

	public.GET("/organizations", handlers.GetOrganizations)
	public.GET("/organizations/:id", handlers.GetOrganization)
	public.POST("/organizations", handlers.CreateOrganization)
	public.PUT("/organizations/:id", handlers.UpdateOrganization)
	public.DELETE("/organizations/:id", handlers.DeleteOrganization)

	return r
}
