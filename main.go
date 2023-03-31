package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

	// Serve static files
    router.Static("/assets", "./frontend/assets")
    router.StaticFile("/", "./frontend/index.html")

    router.GET("/calculate", func(c *gin.Context) {

		//lattitude of NYC
		lat := 52.37
		//longitude of NYC
		lng := 4.89



        // Add your coordinate calculation logic here
        // ...

        c.JSON(http.StatusOK, gin.H{
            "latitude":  lat,
            "longitude": lng,
        })
    })

    router.Run(":8080")
}
