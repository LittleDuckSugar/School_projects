/**
  IoT - SmartSensor
  Configuration
  26 Novembre 2021
**/

// Import pour gestion reseau
#if defined(ESP32)
#include <WiFi.h>
#include <WebServer.h>
#elif defined(ESP8266)
#include <ESP8266WiFi.h>
#include <ESP8266WebServer.h>
#endif

// Import pour JSON
#include <ArduinoJson.h>

#include <DHT.h>

#define ssid "IoT-SmartSensor"
#define password "iotroot9"

#define DHTPIN 5

#define Port 8080

// Instantiation
DHT dht(DHTPIN, DHT22);
ESP8266WebServer server(Port);

void setup() {
  Serial.begin(115200);
  while (!Serial) {
    ; // En attente de connexion du port série
  }
  WiFi.mode(WIFI_STA);
  WiFi.begin(ssid, password);
  Serial.println("");

  // Wait for connection
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("");
  Serial.print("Connected to ");
  Serial.println(ssid);
  Serial.print("IP address: ");
  Serial.println(WiFi.localIP());

  // Set server routing
  restServerRouting();
  // Set not found response
  server.onNotFound(handleNotFound);
  // Start server
  server.begin();
  Serial.print("HTTP server started on port ");
  Serial.println(Port);

  dht.begin();
}

// Route Hello World GET
void getHelloWord() {
  DynamicJsonDocument doc(512);
  doc["name"] = "Hello world";

  String buf;
  serializeJson(doc, buf);
  server.send(200, "application/json", buf);
}

void getSettings() {
  DynamicJsonDocument doc(512);
  doc["ip"] = WiFi.localIP().toString();
  doc["gw"] = WiFi.gatewayIP().toString();
  doc["nm"] = WiFi.subnetMask().toString();

  if (server.arg("signalStrength") == "true") {
    doc["signalStrengh"] = WiFi.RSSI();
  }

  if (server.arg("chipInfo") == "true") {
    doc["chipId"] = ESP.getChipId();
    doc["flashChipId"] = ESP.getFlashChipId();
    doc["flashChipSize"] = ESP.getFlashChipSize();
    doc["flashChipRealSize"] = ESP.getFlashChipRealSize();
  }
  if (server.arg("freeHeap") == "true") {
    doc["freeHeap"] = ESP.getFreeHeap();
  }

  String buf;
  serializeJson(doc, buf);
  server.send(200, F("application/json"), buf);
}

void postFormulaire() {
  String postBody = server.arg("plain");
  Serial.println(postBody);

  DynamicJsonDocument doc(512);
  DeserializationError error = deserializeJson(doc, postBody);
  if (error) {
    // if the file didn't open, print an error:
    Serial.print(F("Error parsing JSON "));
    Serial.println(error.c_str());

    String msg = error.c_str();

    server.send(400, F("text/html"),
                "Error in parsin json body! <br>" + msg);

  } else {
    JsonObject postObj = doc.as<JsonObject>();

    Serial.print(F("HTTP Method: "));
    Serial.println(server.method());

    if (server.method() == HTTP_POST) {
      if (postObj.containsKey("name") && postObj.containsKey("type")) {

        Serial.println(F("done."));

        // Here store data or doing operation
        String firstname = postObj["name"];
        Serial.println(firstname);


        // Create the response
        // To get the status of the result you can get the http status so
        // this part can be unusefully
        DynamicJsonDocument doc(512);
        doc["status"] = "OK";

        Serial.print(F("Stream..."));
        String buf;
        serializeJson(doc, buf);

        server.send(201, F("application/json"), buf);
        Serial.print(F("done."));

      } else {
        DynamicJsonDocument doc(512);
        doc["status"] = "KO";
        doc["message"] = F("No data found, or incorrect!");

        Serial.print(F("Stream..."));
        String buf;
        serializeJson(doc, buf);

        server.send(400, F("application/json"), buf);
        Serial.print(F("done."));
      }
    }
  }
}

// Route de temperature
void getTemperature() {
  String reponse = "{\"temperature\": ";
  reponse.concat(readTemp());
  reponse.concat("}");
  server.send(200, "text/json", reponse);
}


float readTemp() {
  float t = dht.readTemperature();
  if (isnan(t)) {
    Serial.println("Erreur de lecture");
    return 404.00;
  }
  return t;
}

float readHum() {
  float h = dht.readHumidity();
  if (isnan(h)) {
    Serial.println("Erreur de lecture");
    return 404.00;
  }
  return h;
}

// Route Humidite GET
void getHumidite() {
  String reponse = "{\"humidite\": ";
  reponse.concat(readHum());
  reponse.concat("}");
  server.send(200, "text/json", reponse);
}

// Route Humidite et temperature GET
void getAll() {
  String reponse = "{\"humidite\": ";
  reponse.concat(readHum());
  reponse.concat(",\"temperature\" : ");
  reponse.concat(readTemp());
  reponse.concat("}");
  server.send(200, "text/json", reponse);
}

// Handler
void restServerRouting() {
  server.on("/", HTTP_GET, []() {
    server.send(200, F("text/html"),
                F("Welcome to the REST Web Server"));
  });
  server.on(F("/helloWorld"), HTTP_GET, getHelloWord);
  server.on(F("/all"), HTTP_GET, getAll);
  server.on(F("/temperature"), HTTP_GET, getTemperature);
  server.on(F("/humidite"), HTTP_GET, getHumidite);
  server.on(F("/settings"), HTTP_GET, getSettings);
  server.on(F("/formulaire"), HTTP_POST, postFormulaire);
}

// Manage not found URL
void handleNotFound() {
  String message = "File Not Found\n\n";
  message += "URI: ";
  message += server.uri();
  message += "\nMethod: ";
  message += (server.method() == HTTP_GET) ? "GET" : "POST";
  message += "\nArguments: ";
  message += server.args();
  message += "\n";
  for (uint8_t i = 0; i < server.args(); i++) {
    message += " " + server.argName(i) + ": " + server.arg(i) + "\n";
  }
  server.send(404, "text/plain", message);
}


void loop() {
  server.handleClient();
}
