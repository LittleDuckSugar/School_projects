/*
  API_routes.ino - Functions called depending on /path asked.
  CoSensor - August 25, 2023.
  Last notes:
    Clean code
*/

/* POST Mapping
    /system
    /sensor
    /influxdb
*/

// /system POST
void setSystemSettings() {
  Serial.println("/sensor POST called");

  secure();

  String postBody = server->arg("plain");

  DynamicJsonDocument doc(512);
  DeserializationError error = deserializeJson(doc, postBody);
  if (error) {
    // if the file didn't open, print an error:
    String msg = error.c_str();

    server->send(400, F("text/html"),
                 "Error in parsin json body! <br>" + msg);

  } else {
    JsonObject postObj = doc.as<JsonObject>();

    if (server->method() == HTTP_POST) {
      if (postObj.containsKey("sensor_name") && postObj.containsKey("room_name") && postObj.containsKey("position") && postObj.containsKey("etage")) {
        // Update the name of the sensor

        sensorName = doc["sensor_name"].as<String>();
        roomName = doc["room_name"].as<String>();
        pos = doc["position"].as<String>();
        etage = doc["etage"].as<int>();

        WiFi.hostname(sensorName);

        // Response
        generateSystemSettings();

      } else {
        DynamicJsonDocument doc(512);
        doc["status"] = "No data found, or incorrect!";

        String buf;
        serializeJson(doc, buf);
        server->send(400, F("application/json"), buf);
      }
    }
  }
}

// /sensor POST
void setSensorSettings() {
  Serial.println("/sensor POST called");

  secure();

  String postBody = server->arg("plain");

  DynamicJsonDocument doc(512);
  DeserializationError error = deserializeJson(doc, postBody);
  if (error) {
    // if the file didn't open, print an error:
    String msg = error.c_str();

    server->send(400, F("text/html"),
                 "Error in parsin json body! <br>" + msg);

  } else {
    JsonObject postObj = doc.as<JsonObject>();

    if (server->method() == HTTP_POST) {
      if (postObj.containsKey("delay") && postObj.containsKey("is_hum") && postObj.containsKey("is_temp") && postObj.containsKey("is_allowed")) {
        // Here store data or doing operation

        sensorDelay = doc["delay"];
        isHum = doc["is_hum"];
        isTemp = doc["is_temp"];
        isAllowed = doc["is_allowed"];

        // Response
        generateSensorSettings();

      } else {
        DynamicJsonDocument doc(512);
        doc["status"] = "No data found, or incorrect!";

        String buf;
        serializeJson(doc, buf);
        server->send(400, F("application/json"), buf);
      }
    }
  }
}

// /influxdb POST
void setInfluxDBSettings() {
  Serial.println("/influxdb POST called");

  secure();

  String postBody = server->arg("plain");

  DynamicJsonDocument doc(512);
  DeserializationError error = deserializeJson(doc, postBody);
  if (error) {
    // if the file didn't open, print an error:
    String msg = error.c_str();

    server->send(400, F("text/html"),
                 "Error in parsin json body! <br>" + msg);

  } else {
    JsonObject postObj = doc.as<JsonObject>();

    if (server->method() == HTTP_POST) {
      if (postObj.containsKey("influxdb_url") && postObj.containsKey("influxdb_token") && postObj.containsKey("influxdb_org") && postObj.containsKey("influxdb_bucket")) {

        influxdb_setup = true;

        // Saving incomming data
        influxdb_url = doc["influxdb_url"].as<String>();
        influxdb_token = doc["influxdb_token"].as<String>();
        influxdb_org = doc["influxdb_org"].as<String>();
        influxdb_bucket = doc["influxdb_bucket"].as<String>();

        influxdb_client->setConnectionParams(doc["influxdb_url"], doc["influxdb_org"], doc["influxdb_bucket"], doc["influxdb_token"], InfluxDbCloud2CACert);

        // Response
        generateInfluxDBSettings();
      } else {
        DynamicJsonDocument doc(512);
        doc["status"] = "No data found, or incorrect!";

        String buf;
        serializeJson(doc, buf);
        server->send(400, F("application/json"), buf);
      }
    }
  }
}

/* Get Mapping
    /sensor
    /influxdb
    /
    /system
    /network
*/

// /sensor GET
void getSensorSettings() {
  Serial.println("/sensor GET called");

  secure();

  generateSensorSettings();
}

// /influxdb GET
void getInfluxDBSettings() {
  Serial.println("/influxdb GET called");

  secure();

  generateInfluxDBSettings();
}

// / (root) GET
void getSettings() {
  Serial.println("/ GET called");

  secure();

  String path = "http://" + WiFi.localIP().toString() + ":" + API_PORT;

  DynamicJsonDocument doc(512);

  // Show every settings available
  doc["system"] = path + "/system";
  doc["sensor"] = path + "/sensor";
  doc["network"] = path + "/network";
  doc["influxdb"] = path + "/influxdb";

  String buf;
  serializeJson(doc, buf);
  server->send(200, F("application/json"), buf);
}

// /system GET
void getSystemSettings() {
  Serial.println("/system GET called");

  secure();

  generateSystemSettings();
}

// /network GET
void getNetworkSettings() {
  Serial.println("/network GET called");

  secure();

  DynamicJsonDocument doc(512);

  // Base info
  doc["ssid"] = WIFI_SSID;
  doc["ip"] = WiFi.localIP().toString();
  doc["api_port"] = API_PORT;
  doc["mac"] = WiFi.macAddress();

  // More info
  doc["gateway"] = WiFi.gatewayIP().toString();
  doc["subnet_mask"] = WiFi.subnetMask().toString();
  doc["signal_strengh"] = WiFi.RSSI();

  String buf;
  serializeJson(doc, buf);
  server->send(200, F("application/json"), buf);
}

/* GET generators
    /influxdb
    /sensor
    /system
*/

// generateInfluxDBSettings generate the return of /influxdb
void generateInfluxDBSettings() {
  DynamicJsonDocument doc(512);

  if (influxdb_setup) {
    // Save database config
    doc["influxdb_url"] = influxdb_url;
    doc["influxdb_token"] = "configured";
    doc["influxdb_org"] = influxdb_org;
    doc["influxdb_bucket"] = influxdb_bucket;

    // Remember if InfluxDB server is connected
    bool linked = influxdb_client->validateConnection();

    // Save 'true' if linked
    doc["is_connected"] = linked;

    if (linked) {
      doc["status"] = "Setup done and connection working";
    } else {
      doc["status"] = "Setup done but connection failed";
    }
  } else {
    doc["status"] = "Waiting to be configure";
  }

  String buf;
  serializeJson(doc, buf);
  server->send(200, F("application/json"), buf);
}

// generateSensorSettings generate the return of /sensor
void generateSensorSettings() {
  DynamicJsonDocument doc(512);

  // Get sensor config
  doc["sensor"] = sensorType;
  doc["pin"] = sensorPin;
  doc["delay"] = sensorDelay;
  doc["is_hum"] = isHum;
  doc["is_temp"] = isTemp;
  doc["is_allowed"] = isAllowed;

  // Get current value of temperature and humidity
  doc["current_hum"] = dht->readHumidity();
  doc["current_temp"] = dht->readTemperature();

  String buf;
  serializeJson(doc, buf);
  server->send(200, F("application/json"), buf);
}

// generateSystemSettings generate the return of /system
void generateSystemSettings() {
  DynamicJsonDocument doc(512);

  doc["sensor_name"] = sensorName;
  doc["room_name"] = roomName;
  doc["position"] = pos;
  doc["etage"] = etage;
  doc["username"] = USERNAME;
  doc["status"] = "running";

  doc["current_timezone"] = currentTimezone;

  JsonArray timezones = doc.createNestedArray("timezones");
  for (int index = 0; index <= 3; index++) {
    timezones.add(globalTimezones[index]);
  }

  doc["version"] = sensorVersion;

  if (server->arg("all") == "true") {
    doc["chipModel"] = ESP.getChipModel();
    doc["flashChipMode"] = ESP.getFlashChipMode();
    doc["flashChipSize"] = ESP.getFlashChipSize();
    doc["sketchSize"] = ESP.getSketchSize();
    doc["freeSketchSpace"] = ESP.getFreeSketchSpace();
    doc["sdkVersion"] = ESP.getSdkVersion();
    doc["chipRevision"] = ESP.getChipRevision();
  }

  String buf;
  serializeJson(doc, buf);
  server->send(200, F("application/json"), buf);
}

/* Server 404 error
************************
*/

// 404 NOT FOUND
void handleNotFound() {
  Serial.println(server->uri() + " was called but don't exist");

  DynamicJsonDocument doc(512);

  doc["error"] = "There is nothing here.";

  String buf;
  serializeJson(doc, buf);
  server->send(404, F("application/json"), buf);
}
