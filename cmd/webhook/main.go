package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/tomaszkiewicz/prometheus-alertmanager-facebook-messenger-webhook/pkg/template"
	"log"
	"net/http"
)

func webhook(c *gin.Context) {
	defer c.Request.Body.Close()

	var data template.Data
	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("Alerts: GroupLabels=%v, CommonLabels=%v", data.GroupLabels, data.CommonLabels)

	for _, alert := range data.Alerts {
		log.Printf("Alert:\n%v", alert)
	}

	c.JSON(http.StatusOK, struct{}{})
}

func main() {
	viper.SetDefault("http-port", 8079)

	r := setupRouter()
	listenAddress := fmt.Sprintf(":%d", viper.GetInt("http-port"))
	log.Printf("listening on: %s", listenAddress)
	if err := r.Run(listenAddress); err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct{}{})
	})
	r.GET("/webhook", webhook)

	return r
}
