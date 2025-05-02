/*
  Base.ino - Main code for the API sender.
  SmartSensor - January 4, 2022.
  Last notes:
    Implements mDNS for testing purpose
  TODO :
    Reimplement OTA
    Support of specific changes (eg: if I want to change only the delay I can)
*/

// Import for network base
#if defined(ESP32)
#include <WiFi.h>
#include <WebServer.h>
#include <WiFiMulti.h>
WiFiMulti wifiMulti;
#elif defined(ESP8266)
#include <ESP8266WiFi.h>
#include <ESP8266WebServer.h>
#include <ESP8266WiFiMulti.h>
ESP8266WiFiMulti wifiMulti;
#endif

// Import for OTA update
#include <ArduinoOTA.h>

// Import for mDNS name
// #include <ESP8266mDNS.h>

// InfluxDB
#include <InfluxDbClient.h>
#include <InfluxDbCloud.h>
InfluxDBClient *influxdb_client;

// DHT controller
#include <DHT.h>
DHT *dht;

// Import JSON support
#include <ArduinoJson.h>

// Thread controller
#include <Thread.h>
Thread *senderThread;

// API server class init
ESP8266WebServer *server;

// Point d'entré
Point *point;

// Template for Golang
// SSID of WiFi AP
#define WIFI_SSID "WiFi_SSID"
// WiFi password
#define WIFI_PASSWORD "password"
// API port
#define API_PORT 5543

// LED pin
#define LED_PIN 4

// API username and password
#define USERNAME "admin"
#define PASSWORD "Test1234"

// Save var
char *globalTimezones[4] = { "PST8PDT", "EST5EDT", "JST-9", "CET-1CEST,M3.5.0,M10.5.0/3" };
char *currentTimezone = "CET-1CEST,M3.5.0,M10.5.0/3";

String sensorName = "ESP_room_xx";
String sensorVersion = "0.0.3";

int sensorDelay = 60000;
int sensorOldDelay = 60000;
String sensorType = "DHT22";
int sensorPin = 5;
bool isHum = true;
bool isTemp = true;
bool isAllowed = false;

String influxdb_url;
String influxdb_token;
String influxdb_org;
String influxdb_bucket;
bool influxdb_setup = false;

// Setup a few things
void setup() {
  // Setting up information LED
  pinMode(LED_PIN, OUTPUT);

  // Class instances
  senderThread = new Thread();
  server = new ESP8266WebServer(API_PORT);
  point = new Point(sensorName);

  Serial.begin(9600);
  while (!Serial) {
    ;  // Waiting for serial port connection
  }

  // WiFi setup
  WiFi.mode(WIFI_STA);

  WiFi.hostname(sensorName);

  wifiMulti.addAP(WIFI_SSID, WIFI_PASSWORD);

  Serial.print("Connexion au WiFi");
  digitalWrite(LED_PIN, HIGH);
  while (wifiMulti.run() != WL_CONNECTED) {
    Serial.print(".");
    delay(100);
  }
  digitalWrite(LED_PIN, LOW);
  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(WIFI_SSID);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  timeSync(currentTimezone, "pool.ntp.org", "time.nis.gov");

  Serial.println();

  /*
  // Starting OTA support updates
  // Port defaults to 8266
  ArduinoOTA.setPort(8266);

  // Hostname defaults to esp8266-[ChipID]
  ArduinoOTA.setHostname(SENSORNAME);

  // No authentication by default
  ArduinoOTA.setPassword(PASSWORD);

  // Password can be set with it's md5 value as well
  // MD5(admin) = 21232f297a57a5a743894a0e4a801fc3
  // ArduinoOTA.setPasswordHash("21232f297a57a5a743894a0e4a801fc3");

  ArduinoOTA.onStart([]() {
    String type;
    if (ArduinoOTA.getCommand() == U_FLASH) {
      type = "sketch";
    } else { // U_FS
      type = "filesystem";
    }

    // NOTE: if updating FS this would be the place to unmount FS using FS.end()
    Serial.println("Start updating " + type);
  });
  ArduinoOTA.onEnd([]() {
    Serial.println("\nEnd");
  });
  ArduinoOTA.onProgress([](unsigned int progress, unsigned int total) {
    Serial.printf("Progress: %u%%\r", (progress / (total / 100)));
  });
  ArduinoOTA.onError([](ota_error_t error) {
    Serial.printf("Error[%u]: ", error);
    if (error == OTA_AUTH_ERROR) {
      Serial.println("Auth Failed");
    } else if (error == OTA_BEGIN_ERROR) {
      Serial.println("Begin Failed");
    } else if (error == OTA_CONNECT_ERROR) {
      Serial.println("Connect Failed");
    } else if (error == OTA_RECEIVE_ERROR) {
      Serial.println("Receive Failed");
    } else if (error == OTA_END_ERROR) {
      Serial.println("End Failed");
    }
  });
  ArduinoOTA.begin();
  */

  // Setting up the name of the device
  // if (MDNS.begin("test")) {
  //   Serial.println("MDNS responder started");
  // }

  // API config
  restServerRouting();
  // Set function to call if 404 http error
  server->onNotFound(handleNotFound);
  // API startup
  server->begin();
  Serial.print("HTTP server started on port ");
  Serial.println(API_PORT);

  dht = new DHT(sensorPin, DHT22);
  dht->begin();

  influxdb_client = new InfluxDBClient();

  senderThread->enabled = true;      // Default enabled value is true
  senderThread->setInterval(sensorDelay);  // Setts the wanted interval to be 10ms
  // This will set the callback of the Thread: "What should I run"?
  senderThread->onRun(senderHandler);  // callback_function is the name of the function
}

void loop() {
  // Listen for API input everytime
  server->handleClient();

  // MDNS.update();

  // Listen for update everytime
  //ArduinoOTA.handle();

  // Check WiFi status and reconnect if needed
  if ((WiFi.RSSI() == 0) || (wifiMulti.run() != WL_CONNECTED)) {
    digitalWrite(LED_PIN, HIGH);  // turn the LED on (HIGH is the voltage level)
    Serial.println("WiFi lost... reconnecting");
  } else {
    digitalWrite(LED_PIN, LOW);  // turn the LED off by making the voltage LOW
  }

  // First check if our Thread should be run
  if (senderThread->shouldRun()) {
    // Yes, the Thread should run, let's run it
    senderThread->run();
  }

  // Update timer depending on the sensor delay
  if (sensorOldDelay != sensorDelay) {
    sensorOldDelay = sensorDelay;
    senderThread->setInterval(sensorDelay);
  }

}

// Handler
void restServerRouting() {
  /*
    // On root page show a little welcom HTLM page
    server->on("/html", HTTP_GET, []() {
    server->send(200, F("text/html"),
                F("<h1>Welcome to the REST Web Server</h1>"));
    });
  */

  // Routing for pages
  server->on(F("/"), HTTP_GET, getSettings);

  server->on(F("/network"), HTTP_GET, getNetworkSettings);

  server->on(F("/system"), HTTP_GET, getSystemSettings);
  server->on(F("/system"), HTTP_POST, setSystemSettings);

  server->on(F("/sensor"), HTTP_GET, getSensorSettings);
  server->on(F("/sensor"), HTTP_POST, setSensorSettings);

  server->on(F("/influxdb"), HTTP_GET, getInfluxDBSettings);
  server->on(F("/influxdb"), HTTP_POST, setInfluxDBSettings);
}

// secure asks for authentication on every pages
void secure() {
  server->sendHeader("Connection", "close");
  if (!server->authenticate(USERNAME, PASSWORD)) {
    return server->requestAuthentication(DIGEST_AUTH, "Login Required", "Authentication Failed");
  }
}

// senderHandler is the function called by the thread senderThread
void senderHandler() {
  if (influxdb_setup) {
    if (influxdb_client->validateConnection()) {
      if (isAllowed) {
        float h = dht->readHumidity();
        float t = dht->readTemperature();

        if (isnan(h) && isnan(t)) {
          Serial.println("Erreur de lecture");
        } else {
          if (isTemp) {
            digitalWrite(LED_PIN, HIGH);  // turn the LED on (HIGH is the voltage level)
            Serial.println("Temp allowed");
            sender("Temperature", t);
            digitalWrite(LED_PIN, LOW);  // turn the LED off by making the voltage LOW
          }
          delay(1);
          if (isHum) {
            digitalWrite(LED_PIN, HIGH);  // turn the LED on (HIGH is the voltage level)
            Serial.println("Hum allowed");
            sender("Humidity", h);
            digitalWrite(LED_PIN, LOW);  // turn the LED off by making the voltage LOW
          }
        }
      } else {
        Serial.println("Sensor not allowed to write");
      }
    } else {
      Serial.println("InfluxDB connection failed");
    }
  } else {
    Serial.println("InfluxDB not configured");
  }
  Serial.println("Pause...");
}

// sender sends content (value) with a specific info (value)
void sender(String info, float content) {
  // Vide les champs précendant en gardant le point utilisé. Les tags restes présent
  point->clearFields();

  // Enregistre la valeur dans le point
  point->addField(info, content);

  // Affiche sur le port série le contenu de ce qui va être envoyé
  Serial.print("Ecriture: ");
  Serial.println(point->toLineProtocol());

  // Si il n'y a pas de WiFi, on essai de se reconnecter
  if ((WiFi.RSSI() == 0) && (wifiMulti.run() != WL_CONNECTED)) {
    Serial.println("Connexion WiFi perdu");
  }

  // On écrit (on envoie) le point dans la base de donnée
  if (!influxdb_client->writePoint(*point)) {
    Serial.print("Ecriture echouer InfluxDB: ");
    Serial.println(influxdb_client->getLastErrorMessage());
  }
}

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
      if (postObj.containsKey("sensor_name")) {
        // Update the name of the sensor

        sensorName = doc["sensor_name"].as<String>();

        WiFi.hostname(sensorName);
        //ArduinoOTA.setHostname(sensorName);

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
  doc["username"] = USERNAME;
  doc["status"] = "running";

  doc["current_timezone"] = currentTimezone;

  JsonArray timezones = doc.createNestedArray("timezones");
  for (int index = 0; index <= 3; index++) {
    timezones.add(globalTimezones[index]);
  }

  doc["version"] = sensorVersion;

  // TODO :
  // Implement last error message

  if (server->arg("all") == "true") {
    doc["chipId"] = ESP.getChipId();
    doc["flashChipId"] = ESP.getFlashChipId();
    doc["flashChipSize"] = ESP.getFlashChipSize();
    doc["flashChipRealSize"] = ESP.getFlashChipRealSize();
    doc["freeHeap"] = ESP.getFreeHeap();
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

  /*
    String message = "File Not Found\n\n";
    message += "URI: ";
    message += server->uri();
    message += "\nMethod: ";
    message += (server->method() == HTTP_GET) ? "GET" : "POST";
    message += "\nArguments: ";
    message += server->args();
    message += "\n";
    for (uint8_t i = 0; i < server->args(); i++) {
    message += " " + server->argName(i) + ": " + server->arg(i) + "\n";
    }
    server->send(404, "text/plain", message);
  */
}
