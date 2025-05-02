/*
  OpenFirm.ino - Main code for the API sender.
  CoSensor - March 29, 2023.
  Last notes:
    Clean code
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
// Test TODO on ESP32 and ESP8266
WebServer *server;

// Point d'entré
Point *point;

// SSID of WiFi AP
#define WIFI_SSID "WIFI_SSID"
// WiFi password
#define WIFI_PASSWORD "WIFI_PASSWORD"
// API port
#define API_PORT 5543

// LED pin
#define LED_PIN 4

// API username and password
#define USERNAME "admin"
#define PASSWORD "password"

// Save var
char *globalTimezones[4] = { "PST8PDT", "EST5EDT", "JST-9", "CET-1CEST,M3.5.0,M10.5.0/3" };
char *currentTimezone = "CET-1CEST,M3.5.0,M10.5.0/3";

String sensorName = "ESP_room_xx";
String sensorVersion = "0.1.0";

String roomName = "undefined";
String pos = "";
int etage = 0;

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
  server = new WebServer(API_PORT);
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

  senderThread->enabled = true;            // Default enabled value is true
  senderThread->setInterval(sensorDelay);  // Setts the wanted interval to be 10ms
  // This will set the callback of the Thread: "What should I run"?
  senderThread->onRun(senderHandler);  // callback_function is the name of the function
}

void loop() {
  // Listen for API input everytime
  server->handleClient();

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


