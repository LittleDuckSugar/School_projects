package controllers

import (
	"SmartSensorServer/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetSensorByRoomName provide /room/roomName/sensor/sensorName with a sensor choosen from a room name (GET)
func GetSensorBySensorName(c *gin.Context) {
	found := false
	sensorFound := models.Sensor{}

	for _, room := range Place.Rooms {
		for _, sensor := range room.Sensors {
			if sensor.SystemSettings.SensorName == c.Params.ByName("sensorName") {
				found = true
				sensorFound = *sensor
			}
		}
	}

	if found {
		c.JSON(http.StatusOK, sensorFound)
	} else {
		c.String(http.StatusNotFound, "Sensor not found")
	}
}

// PatchSensor provide /sensor/sensorName to patch an existing sensor
func PatchSensor(c *gin.Context) {

	// Validation of input
	var input models.SensorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sensorName := c.Params.ByName("sensorName")

	if currentRoom, found := Place.FindSensor(sensorName); found {
		// Sensor exist

		// Saving current room
		sensorToUpdate := currentRoom.Sensors[sensorName]
		sensorSaved := sensorToUpdate
		Place.Rooms[currentRoom.Name].DeleteSensor(sensorName)

		sensorToUpdate.SystemSettings.Position = input.SystemSettingsInput.Position
		sensorToUpdate.SystemSettings.Etage = input.SystemSettingsInput.Etage
		sensorToUpdate.SystemSettings.CurrentTimezone = input.SystemSettingsInput.CurrentTimezone

		sensorToUpdate.SensorSettings.SensorDelay = input.SensorSettingsInput.SensorDelay
		sensorToUpdate.SensorSettings.IsHum = input.SensorSettingsInput.IsHum
		sensorToUpdate.SensorSettings.IsTemp = input.SensorSettingsInput.IsTemp
		sensorToUpdate.SensorSettings.IsAllowed = input.SensorSettingsInput.IsAllowed

		sensorToUpdate.InfluxDBConfig.InfluxDBURL = input.InfluxDBConfigInput.InfluxDBURL
		sensorToUpdate.InfluxDBConfig.InfluxDBToken = input.InfluxDBConfigInput.InfluxDBToken
		sensorToUpdate.InfluxDBConfig.InfluxDBOrg = input.InfluxDBConfigInput.InfluxDBOrg
		sensorToUpdate.InfluxDBConfig.InfluxDBBucket = input.InfluxDBConfigInput.InfluxDBBucket

		if sensorSaved.SystemSettings.SensorName != input.SystemSettingsInput.SensorName {
			// The sensor name is different
			if _, found := Place.FindSensor(input.SystemSettingsInput.SensorName); !found {
				// There is no existing sensor with this name --> Name available
				sensorToUpdate.SystemSettings.SensorName = input.SystemSettingsInput.SensorName
			} else {
				// New name not available
				Place.Rooms[sensorSaved.SystemSettings.RoomName].AddSensor(sensorSaved)
				c.String(http.StatusNotFound, "Sensor '%s' already exist", input.SystemSettingsInput.SensorName)
				return
			}
		}

		// Registering into the selected room
		if currentRoom.Name != input.SystemSettingsInput.RoomName {
			// New room name found
			if _, found := Place.Rooms[input.SystemSettingsInput.RoomName]; found {
				// The new room name exist
				sensorToUpdate.SystemSettings.RoomName = input.SystemSettingsInput.RoomName
				Place.Rooms[input.SystemSettingsInput.RoomName].AddSensor(sensorToUpdate)
				c.JSON(http.StatusOK, gin.H{"status": sensorName + " updated"})
			} else {
				// The new room name don't exists
				Place.Rooms[sensorSaved.SystemSettings.RoomName].AddSensor(sensorSaved)
				c.String(http.StatusNotFound, "New room '%s' not found, changes not applied", input.SystemSettingsInput.RoomName)
			}
		} else {
			// Adding the sensor back to the same room he came from
			Place.Rooms[sensorToUpdate.SystemSettings.RoomName].AddSensor(sensorToUpdate)
			c.JSON(http.StatusOK, gin.H{"status": sensorName + " updated"})
		}

	} else {
		// Capteur non trouvé
		c.String(http.StatusNotFound, "Sensor '%s' not found", sensorName)
	}
}
