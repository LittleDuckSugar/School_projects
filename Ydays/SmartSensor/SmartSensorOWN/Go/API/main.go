package main

import (
	// Models (types)
	"SmartSensorServer/controllers"
	"SmartSensorServer/models"
	"net/http"

	// Go build in

	// gin is a framework for routing pages
	// official gin documentation https://gin-gonic.com/docs/
	"github.com/gin-gonic/gin"
)

var myPlace models.Place = models.Place{Name: "Ynov - Demo", Rooms: nil}

// Should delete init() when implementing missing functions are done
// init add 2 rooms to the place with 3 sensors total
func init() {
	demo()
}

func demo() {
	// Must replace this with a network scan
	ESP_Salon_01 := models.Sensor{
		SystemSettings: models.SystemSettings{
			SensorName:        "ESP_Salon_01",
			RoomName:          "Salon",
			Position:          "Sol",
			Etage:             1,
			Username:          "admin",
			Status:            "running",
			CurrentTimezone:   "CET-1CEST,M3.5.0,M10.5.0/3",
			Timezones:         []string{"CET-1CEST,M3.5.0,M10.5.0/3", "PST8PDT", "EST5EDT", "JST-9"},
			Version:           "0.0.3",
			ChipId:            42,
			FlashChipId:       42,
			FlashChipSize:     42,
			FlashChipRealSize: 42,
			FreeHeap:          42},
		SensorSettings: models.SensorSettings{
			SensorType:  "DHT22",
			SensorPin:   5,
			SensorDelay: 60,
			IsHum:       true,
			IsTemp:      true,
			IsAllowed:   true,
			CurrentHum:  46.4,
			CurrentTemp: 20.1},
		InfluxDBConfig: models.InfluxDBConfig{
			InfluxDBURL:    "http://192.168.1.3:8086",
			InfluxDBToken:  "aToken",
			InfluxDBOrg:    "Ynov",
			InfluxDBBucket: "Paris",
			Status:         "Configured and connected",
			IsConnected:    true},
		NetworkSettings: models.NetworkSettings{
			SSID:          "WiFI",
			IP:            "192.168.1.100",
			APIPort:       5565,
			MAC:           "ef:re:sd:gf:az",
			Gateway:       "192.168.1.255",
			SubnetMask:    "255.255.255.255",
			SignalStrengh: -32}}

	ESP_Salon_02 := models.Sensor{
		SystemSettings: models.SystemSettings{
			SensorName:        "ESP_Salon_02",
			RoomName:          "Salon",
			Position:          "Fenêtre nord",
			Etage:             1,
			Username:          "admin",
			Status:            "running",
			CurrentTimezone:   "CET-1CEST,M3.5.0,M10.5.0/3",
			Timezones:         []string{"CET-1CEST,M3.5.0,M10.5.0/3", "PST8PDT", "EST5EDT", "JST-9"},
			Version:           "0.0.3",
			ChipId:            42,
			FlashChipId:       42,
			FlashChipSize:     42,
			FlashChipRealSize: 42,
			FreeHeap:          42},
		SensorSettings: models.SensorSettings{
			SensorType:  "DHT22",
			SensorPin:   5,
			SensorDelay: 60,
			IsHum:       true,
			IsTemp:      true,
			IsAllowed:   true,
			CurrentHum:  52.6,
			CurrentTemp: 17.2},
		InfluxDBConfig: models.InfluxDBConfig{
			InfluxDBURL:    "http://192.168.1.3:8086",
			InfluxDBToken:  "aToken",
			InfluxDBOrg:    "Ynov",
			InfluxDBBucket: "Paris",
			Status:         "Configured and connected",
			IsConnected:    true},
		NetworkSettings: models.NetworkSettings{
			SSID:          "WiFI",
			IP:            "192.168.1.100",
			APIPort:       5565,
			MAC:           "ef:re:sd:gf:az",
			Gateway:       "192.168.1.255",
			SubnetMask:    "255.255.255.255",
			SignalStrengh: -32}}

	ESP_Chambre_01 := models.Sensor{
		SystemSettings: models.SystemSettings{
			SensorName:        "ESP_Chambre_01",
			RoomName:          "Chambre",
			Position:          "Fenêtre nord",
			Etage:             1,
			Username:          "admin",
			Status:            "running",
			CurrentTimezone:   "CET-1CEST,M3.5.0,M10.5.0/3",
			Timezones:         []string{"CET-1CEST,M3.5.0,M10.5.0/3", "PST8PDT", "EST5EDT", "JST-9"},
			Version:           "0.0.3",
			ChipId:            42,
			FlashChipId:       42,
			FlashChipSize:     42,
			FlashChipRealSize: 42,
			FreeHeap:          42},
		SensorSettings: models.SensorSettings{
			SensorType:  "DHT22",
			SensorPin:   5,
			SensorDelay: 60,
			IsHum:       true,
			IsTemp:      true,
			IsAllowed:   true,
			CurrentHum:  52.6,
			CurrentTemp: 17.2},
		InfluxDBConfig: models.InfluxDBConfig{
			InfluxDBURL:    "http://192.168.1.3:8086",
			InfluxDBToken:  "aToken",
			InfluxDBOrg:    "Ynov",
			InfluxDBBucket: "Paris",
			Status:         "Configured and connected",
			IsConnected:    true},
		NetworkSettings: models.NetworkSettings{
			SSID:          "WiFI",
			IP:            "192.168.1.100",
			APIPort:       5565,
			MAC:           "ef:re:sd:gf:az",
			Gateway:       "192.168.1.255",
			SubnetMask:    "255.255.255.255",
			SignalStrengh: -32}}

	// Créer les pièces avec les capteurs
	myPlace.AddRoom(models.CreateNewRoom("Salon", 20))
	myPlace.AddRoom(models.CreateNewRoom("Chambre", 10))

	myPlace.Rooms["Salon"].AddSensor(&ESP_Salon_01)
	myPlace.Rooms["Salon"].AddSensor(&ESP_Salon_02)
	myPlace.Rooms["Chambre"].AddSensor(&ESP_Chambre_01)

	// Provide repo to controllers
	controllers.Place = myPlace
}

func main() {
	// Tells to gin if we are in a dev environment or not
	gin.SetMode(gin.DebugMode)
	// gin.SetMode(gin.ReleaseMode)

	// Tells to gin to force color in shell
	gin.ForceConsoleColor()

	router := gin.Default()

	// router.Static("/static", "./static") for static path

	router.GET("/", controllers.GetRoot)

	router.GET("/room/:roomName", controllers.GetRoom)
	router.DELETE("/room/:roomName", controllers.DeleteRoom)
	router.PATCH("/room/:roomName", controllers.PatchRoom)

	router.POST("/room", controllers.PostRoom)

	router.GET("/sensor/:sensorName", controllers.GetSensorBySensorName)
	router.PATCH("/sensor/:sensorName", controllers.PatchSensor)

	// Missing a POST on sensor for editing sensor params

	router.GET("/reset", func(c *gin.Context) {
		demo()
		c.Redirect(http.StatusSeeOther, "http://localhost:8080")
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run()
	router.Run(":3000")
}
