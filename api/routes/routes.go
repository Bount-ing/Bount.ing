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
	public.GET("/bounties/", handlers.GetBounties)

	user := needAuth.Group("/users")
	user.GET("/me", handlers.GetCurrentUser)
	user.POST("/stripe", handlers.ConnectStripe)

	bounty := needAuth.Group("/bounties")
	bounty.POST("/", handlers.CreateBounty)
	bounty.GET("/:id", handlers.GetBounty)
	bounty.PUT("/:id", handlers.UpdateBounty)
	bounty.DELETE("/:id", handlers.DeleteBounty)

	adminBounties := needAdmin.Group("/bounties")
	adminBounties.GET("/unconfirmed", handlers.GetAllUnconfirmedBounties)
	adminBounties.PUT("/finalize/:id", handlers.FinalizeBounty)

	return r
}
