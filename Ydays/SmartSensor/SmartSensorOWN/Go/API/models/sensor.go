package models

import (
	"fmt"
	"net"
	"strings"
)

type Sensor struct {
	SystemSettings  SystemSettings  `json:"system_settings"`
	SensorSettings  SensorSettings  `json:"sensor_settings"`
	InfluxDBConfig  InfluxDBConfig  `json:"influxdb_config"`
	NetworkSettings NetworkSettings `json:"network_settings"`
}

type SystemSettings struct {
	SensorName        string   `json:"sensor_name"`
	RoomName          string   `json:"room_name"`
	Position          string   `json:"position"`
	Etage             int8     `json:"etage"`
	Username          string   `json:"username"`
	Status            string   `json:"status"`
	CurrentTimezone   string   `json:"current_timezone"`
	Timezones         []string `json:"timezones"`
	Version           string   `json:"version"`
	ChipId            uint     `json:"chip_id"`
	FlashChipId       uint     `json:"flash_chip_id"`
	FlashChipSize     uint     `json:"flash_chip_size"`
	FlashChipRealSize uint     `json:"flash_chip_real_size"`
	FreeHeap          uint     `json:"free_heap"`
}

type SensorSettings struct {
	SensorType  string  `json:"sensor_type"`
	SensorPin   uint8   `json:"sensor_pin"`
	SensorDelay uint    `json:"sensor_delay"`
	IsHum       bool    `json:"is_hum"`
	IsTemp      bool    `json:"is_temp"`
	IsAllowed   bool    `json:"is_allowed"`
	CurrentTemp float64 `json:"current_temp"`
	CurrentHum  float64 `json:"current_hum"`
}

type InfluxDBConfig struct {
	InfluxDBURL    string `json:"influxdb_url"`
	InfluxDBToken  string `json:"influxdb_token"`
	InfluxDBOrg    string `json:"influxdb_org"`
	InfluxDBBucket string `json:"influxdb_bucket"`
	Status         string `json:"status"`
	IsConnected    bool   `json:"is_connected"`
}

type NetworkSettings struct {
	SSID          string `json:"ssid"`
	IP            string `json:"ip"`
	APIPort       uint16 `json:"api_port"`
	MAC           string `json:"mac"`
	Gateway       string `json:"gateway"`
	SubnetMask    string `json:"subnet_mask"`
	SignalStrengh int    `json:"signal_strengh"`
}

type SensorInput struct {
	SystemSettingsInput SystemSettingsInput `json:"system_settings" binding:"required"`
	SensorSettingsInput SensorSettingsInput `json:"sensor_settings" binding:"required"`
	InfluxDBConfigInput InfluxDBConfigInput `json:"influxdb_config" binding:"required"`
}

type SystemSettingsInput struct {
	SensorName      string `json:"sensor_name" binding:"required"`
	RoomName        string `json:"room_name" binding:"required"`
	Position        string `json:"position" binding:"required"`
	Etage           int8   `json:"etage"`
	CurrentTimezone string `json:"current_timezone" binding:"required"`
}

type SensorSettingsInput struct {
	SensorDelay uint `json:"sensor_delay" binding:"required"`
	IsHum       bool `json:"is_hum"`
	IsTemp      bool `json:"is_temp"`
	IsAllowed   bool `json:"is_allowed"`
}

type InfluxDBConfigInput struct {
	InfluxDBURL    string `json:"influxdb_url" binding:"required"`
	InfluxDBToken  string `json:"influxdb_token" binding:"required"`
	InfluxDBOrg    string `json:"influxdb_org" binding:"required"`
	InfluxDBBucket string `json:"influxdb_bucket" binding:"required"`
}

/*
func (sensor Sensor) String() string {
	return fmt.Sprintf("%s : SensorType:%s, Temperature:%v, Humidity:%v, Position:%v, IPAddress:%s, MACAddress:%s, Delay:%v, Allowed:%v", sensor.Name, sensor.SensorType, sensor.Temperature, sensor.Humidity, sensor.Position, sensor.IPAddress, sensor.MACAddress, sensor.Delay, sensor.Allowed)
}
*/
// CheckSensorsInNetwork try to match a password with the MAC Address of connected device
func CheckSensorsInNetwork(keypass string, currentNetwork map[string]net.IP) bool {

	for key, ip := range currentNetwork {
		key = strings.ReplaceAll(key, ":", "")
		if key == keypass {
			fmt.Println("Capteur trouvé avec l'adresse IP suivante " + ip.String())
			return true
		}
	}
	return false
}
