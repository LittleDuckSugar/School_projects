package models

type Sensor struct {
	SystemSettings  SystemSettings  `json:"system_settings"`
	SensorSettings  SensorSettings  `json:"sensor_settings"`
	InfluxDBConfig  InfluxDBConfig  `json:"influxdb_config"`
	NetworkSettings NetworkSettings `json:"network_settings"`
	RoomAvailable   []string
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

type SensorPost struct {
	SystemSettings SystemSettingsPost `json:"system_settings"`
	SensorSettings SensorSettingsPost `json:"sensor_settings"`
	InfluxDBConfig InfluxDBConfigPost `json:"influxdb_config"`
}

type SystemSettingsPost struct {
	SensorName      string `json:"sensor_name"`
	RoomName        string `json:"room_name"`
	Position        string `json:"position"`
	Etage           int8   `json:"etage"`
	CurrentTimezone string `json:"current_timezone"`
}

type SensorSettingsPost struct {
	SensorDelay uint `json:"sensor_delay"`
	IsHum       bool `json:"is_hum"`
	IsTemp      bool `json:"is_temp"`
	IsAllowed   bool `json:"is_allowed"`
}

type InfluxDBConfigPost struct {
	InfluxDBURL    string `json:"influxdb_url"`
	InfluxDBToken  string `json:"influxdb_token"`
	InfluxDBOrg    string `json:"influxdb_org"`
	InfluxDBBucket string `json:"influxdb_bucket"`
}

/*
func (sensor Sensor) String() string {
	return fmt.Sprintf("%s : SensorType:%s, Temperature:%v, Humidity:%v, Position:%v, IPAddress:%s, MACAddress:%s, Delay:%v, Allowed:%v", sensor.Name, sensor.SensorType, sensor.Temperature, sensor.Humidity, sensor.Position, sensor.IPAddress, sensor.MACAddress, sensor.Delay, sensor.Allowed)
}
*/
