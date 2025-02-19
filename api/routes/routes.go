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
	public.POST("/refresh", handlers.RefreshToken)
	public.POST("/reset-password-request", handlers.RequestPasswordReset)
	public.POST("/reset-password", handlers.ResetPassword)

	needAuth.GET("/oauth/:host", handlers.GetOAuthState)
	public.GET("/oauth/github/callback", handlers.OAuthGithubCallback)
	public.GET("/oauth/stripe/callback", handlers.OAuthStripeCallback)
	public.POST("/register", handlers.CreateUser)

	user := needAuth.Group("/users")

	// Self user routes
	user.GET("/me", handlers.GetCurrentUserProfileInfo)
	user.PUT("/me/profile", handlers.UpdateCurrentUserProfileInfo)

	adminBounties := needAdmin.Group("/bounties")
	adminBounties.GET("/unconfirmed", handlers.GetAllUnconfirmedBounties)
	adminBounties.PUT("/finalize/:id", handlers.FinalizeBounty)

	public.GET("/claims", handlers.GetClaims)
	public.GET("/claims/:id", handlers.GetClaim)
	needAuth.POST("/claims", handlers.ClaimBounty)
	needAuth.POST("/claims/:id/approve", handlers.ApproveClaim)
	needAdmin.PUT("/claims/:id", handlers.UpdateClaim)
	needAdmin.DELETE("/claims/:id", handlers.DeleteClaim)

	public.GET("/bounties/", handlers.GetBounties)
	public.GET("/bounties/:id", handlers.GetBounty)
	needAuth.POST("/bounties", handlers.CreateBounty)
	needAdmin.PUT("/bounties/:id", handlers.UpdateBounty)
	needAuth.DELETE("/bounties/:id", handlers.DeleteBounty)

	public.GET("/public-bounties-issue", handlers.GetPublicBountiesByIssue)

	public.GET("/issues", handlers.GetIssues)
	public.GET("/issues/:id", handlers.GetIssue)
	public.GET("/issues/:id/bounties", handlers.GetIssueBounties)
	public.GET("/issues-by-url/*url", handlers.GetIssueByUrl)
	needAuth.POST("/issues", handlers.CreateIssue)
	needAuth.POST("/import-issue-by-url/:url", handlers.CreateIssueFromURL)
	needAdmin.PUT("/issues/:id", handlers.UpdateIssue)
	needAdmin.DELETE("/issues/:id", handlers.DeleteIssue)

	needAuth.GET("/users/me/issues", handlers.GetMyIssues)

	public.GET("/repositories", handlers.GetRepositories)
	public.GET("/repositories/:id", handlers.GetRepository)
	needAuth.POST("/repositories", handlers.CreateRepository)
	needAdmin.PUT("/repositories/:id", handlers.UpdateRepository)
	needAdmin.DELETE("/repositories/:id", handlers.DeleteRepository)

	public.GET("/hosts", handlers.GetHosts)
	public.GET("/hosts/:id", handlers.GetHost)
	needAdmin.POST("/hosts", handlers.CreateHost)
	needAdmin.PUT("/hosts/:id", handlers.UpdateHost)
	needAdmin.DELETE("/hosts/:id", handlers.DeleteHost)

	public.GET("/organizations", handlers.GetOrganizations)
	public.GET("/organizations/:id", handlers.GetOrganization)
	needAuth.POST("/organizations", handlers.CreateOrganization)
	needAdmin.PUT("/organizations/:id", handlers.UpdateOrganization)
	needAdmin.DELETE("/organizations/:id", handlers.DeleteOrganization)

	needAuth.POST("/create-setup-intent", handlers.CreateSetupIntent)
	needAuth.POST("/confirm-setup", handlers.ConfirmSetup)
	needAuth.GET("/payment-methods", handlers.GetPaymentMethods)
	needAuth.DELETE("/payment-methods/:id", handlers.RemovePaymentMethod)

	return r
}
