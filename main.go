package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/subosito/gotenv"
	"net/http"
	"os"
)

type Cotizacion struct {
	DolarCompra float32 `json:"dolarCompra"`
	DolarVenta  float32 `json:"dollarVenta"`
	EuroCompra  float32 `json:"euroCompra"`
	EuroVenta   float32 `json:"euroVenta"`
}

func init() {
	_ = gotenv.Load(".env")
}

func main() {

	//CRON FOR WEBHOOK SEND NOTIFICATION
	CronNotificationWebHook()

	//WEB SERVER AND START LISTENING Cotizacion
	WebServerCotizaion()

}

func CronNotificationWebHook() {
	c := cron.New()
	c.AddFunc(os.Getenv("CRON_NOTIFICATION"), func() {
		cotiz, _ := CotizacionDolarEuroBNAScraping()
		err := SendNotification(cotiz)
		if err != nil {
			fmt.Println(err)
		}
	})
	//IF THE VARIABLE ENVIROMENT IS TRUE, THE CRON WILL BE STARTED
	if os.Getenv("CRON_NOTIFICATION_ENABLED") == "true" && os.Getenv("WEBHOOK_ENDPOINT") != "" {
		c.Start()
		fmt.Println("CRON WEBHOOK STARTED")
	}
}

func WebServerCotizaion() {
	r := gin.Default()
	r.GET("/cotizacion", func(c *gin.Context) {
		cotiz, err := CotizacionDolarEuroBNAScraping()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusOK, cotiz)
		}
	})
	r.Run(":" + os.Getenv("API_SERVER_PORT"))
}

// SendNotification SEND POST NOTIFICATION TO WEBHOOK
func SendNotification(cotiz Cotizacion) error {
	payload, err := json.Marshal(cotiz)
	req, err := http.NewRequest("POST", os.Getenv("WEBHOOK_ENDPOINT"), bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
