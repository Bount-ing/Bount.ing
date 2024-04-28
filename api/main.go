package main

import (
    "open-bounties/api/api"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Setup route group for the API
    apiRoutes := router.Group("/api")
    {
        api.SetupRoutes(apiRoutes)
    }

    router.Run(":8080")  // Listen and serve on 0.0.0.0:8080
}

