package main

import (
	"fmt"
	"net/http"

	"SmartSensorClient/controllers"
	"SmartSensorClient/repository"
)

/*
	This client calls the go API
	This client should not talk directly to sensor API directly
*/

// var myPlace models.Place = models.Place{Name: "Ynov - Demo", Rooms: make(map[string]*models.Room)}

func init() {
	// Must replace this with a network scan
	// DHT22_Salon := models.Sensor{Name: "ESP_Salon_01", SensorType: "DHT22", Temperature: true, Humidity: true, Position: "Mi hauteur", IPAddress: "192.168.1.28", MACAddress: "80:7D:3A:F3:9A:E0", Delay: 60, Allowed: true}
	// DHT22_Salon2 := models.Sensor{Name: "ESP_Salon_02", SensorType: "DHT22", Temperature: true, Humidity: true, Position: "Mi hauteur", IPAddress: "192.168.1.28", MACAddress: "80:7D:3A:F3:9A:E0", Delay: 60, Allowed: true}
	// DHT22_Chambre := models.Sensor{Name: "ESP_Chambre_01", SensorType: "DHT22", Temperature: true, Humidity: true, Position: "Mi hauteur", IPAddress: "192.168.1.29", MACAddress: "3C:71:BF:44:A7:5C", Delay: 60, Allowed: true}

	// // Créer les pièces avec les capteurs
	// Salon := models.Room{Name: "Salon", Sensors: make(map[string]*models.Sensor), Area: 20}
	// myPlace.Rooms["Salon"] = &Salon

	// Chambre := models.Room{Name: "Chambre", Sensors: make(map[string]*models.Sensor), Area: 10}
	// myPlace.Rooms["Chambre"] = &Chambre

	// myPlace.Rooms["Salon"].Sensors["ESP_Salon_01"] = &DHT22_Salon
	// myPlace.Rooms["Salon"].Sensors["ESP_Salon_02"] = &DHT22_Salon2
	// myPlace.Rooms["Chambre"].Sensors["ESP_Chambre_01"] = &DHT22_Chambre

	// // Provide repo to controllers
	// controllers.Place = myPlace
	// fmt.Println(controllers.Place)

	repository.URL = "http://localhost:3000"
}

func main() {
	fmt.Println("Client go")
	fmt.Println("Server is starting...")

	// Client part
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))) // give acces to every static files (css, js, images)
	http.HandleFunc("/", controllers.RootHandler)                                                // handle root page
	http.HandleFunc("/room/", controllers.RoomHandler)                                           // handle room page

	http.HandleFunc("/deleteRoom/", controllers.DeleteRoom) // handle delete room call
	http.HandleFunc("/updateRoom/", controllers.UpdateRoom) // handle update room call
	http.HandleFunc("/createRoom/", controllers.CreateRoom) // handle create room call

	http.HandleFunc("/settings", controllers.SettingsHandler) // handle settings

	http.HandleFunc("/sensor/", controllers.SensorHandler)      // handle sensor page
	http.HandleFunc("/updateSensor/", controllers.UpdateSensor) // handle update sensor call

	fmt.Printf("Starting server at port 8080\n")
	fmt.Println("Go on http://localhost:8080")

	fmt.Printf("\nTo shutdown the server and exit the code hit \"crtl+C\"\n")

	if err := http.ListenAndServe(":8080", nil); err != nil { // Launches the server on port 8080 (if port 8080 is not busy)
		error.Error(err)
	}
}
