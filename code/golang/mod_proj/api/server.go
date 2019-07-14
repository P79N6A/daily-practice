package apis

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func Serve(port int) {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	LoadApiRouters("/api/v1", router)

	log.Printf("[API] API SITE is listening on port: %d\n", port)

	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Shutdown Gracefully BEGIN
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server Shutdown: %v", err)
	}

	log.Println("Server exiting.")
	// Shutdown Gracefully END
}
