package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"github.com/sovrinbloc/kairos/config"
	"github.com/sovrinbloc/kairos/router"
	"github.com/gin-contrib/gzip"
)

func init() {
	config.LoadEnv()
}

func main() {
	fmt.Println("gin.Version: ", gin.Version)
	if config.ServerConfig != config.TestMode {
		gin.SetMode(gin.ReleaseMode)
		// Disable Console Color, you don't need console color when writing the logs to file.
		gin.DisableConsoleColor()
		// Logging to a file.
		logFile, err := os.OpenFile(config.GetLogPath(), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf(err.Error())
			os.Exit(-1)
		}
		gin.DefaultWriter = io.MultiWriter(logFile)
	}

	// Creates a router without any middleware by default
	app := gin.Default()
	PrepareMiddleware(app)

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())
	//app.Use(cors.Default())




	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())

	router.Route(app)

	app.Run(":" + fmt.Sprintf("%s", config.ServerPort))
}

func PrepareMiddleware(r *gin.Engine) {
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(CORSMiddleware())
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	}
}

