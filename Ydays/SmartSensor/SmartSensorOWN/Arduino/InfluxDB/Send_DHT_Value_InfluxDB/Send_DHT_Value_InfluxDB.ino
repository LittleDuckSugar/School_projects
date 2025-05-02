/**
  IoT - SmartSensor
  Envoie sur la base de données la valeur de la température et la valeur de l'humidité
  15 Novembre 2021
**/

#if defined(ESP32)
#include <WiFiMulti.h>
WiFiMulti wifiMulti;
#elif defined(ESP8266)
#include <ESP8266WiFiMulti.h>
ESP8266WiFiMulti wifiMulti;
#endif

#include <InfluxDbClient.h>
#include <InfluxDbCloud.h>

// SSID du point d'accès WiFi
#define WIFI_SSID "IoT-SmartSensor"
// Mot de passe du réseau WiFi
#define WIFI_PASSWORD "iotroot9"
// InfluxDB v2 server url, e.g. https://eu-central-1-1.aws.cloud2.influxdata.com (Use: InfluxDB UI -> Load Data -> Client Libraries)
#define INFLUXDB_URL "http://192.168.4.x:8086"
// InfluxDB v2 server ou cloud API authentication token (Use: InfluxDB UI -> Data -> Tokens -> <select token>)
#define INFLUXDB_TOKEN "_HTahEBwHcTOxAWgVZub9SEzTcg3aJCztnmPtqYGFPIPmMsHH4oqdX5MhWTRG7Nr7GRLSI0X8xc4uWSMdTxa3w=="
// InfluxDB v2 organization id (Use: InfluxDB UI -> User -> About -> Common Ids )
#define INFLUXDB_ORG "SmartSensor"
// InfluxDB v2 bucket name (Use: InfluxDB UI ->  Data -> Buckets)
#define INFLUXDB_BUCKET "Home"

// Set timezone string according to https://www.gnu.org/software/libc/manual/html_node/TZ-Variable.html
// Examples:
//  Pacific Time: "PST8PDT"
//  Eastern: "EST5EDT"
//  Japanesse: "JST-9"
//  Central Europe: "CET-1CEST,M3.5.0,M10.5.0/3"
#define TZ_INFO "CET-1CEST,M3.5.0,M10.5.0/3"

// Instanciation du client InfluxDB avec ces parametres dont le certificat InfluxCloud
InfluxDBClient client(INFLUXDB_URL, INFLUXDB_ORG, INFLUXDB_BUCKET, INFLUXDB_TOKEN, InfluxDbCloud2CACert);

// Point d'entré
Point sensor("DHT22");

#include "DHT.h"
#define DHTPIN 5

DHT dht(DHTPIN, DHT22);

void setup() {
  Serial.begin(115200);
  while (!Serial) {
    ; // En attente de connexion du port série
  }

  // Configuration du WiFi
  WiFi.mode(WIFI_STA);
  wifiMulti.addAP(WIFI_SSID, WIFI_PASSWORD);

  Serial.print("Connexion au WiFi");
  while (wifiMulti.run() != WL_CONNECTED) {
    Serial.print(".");
    delay(100);
  }
  Serial.println();

  // Ajout de tags (tag name, value)
  sensor.addTag("Room", "Salon");

  // Synchronisation du temps OBLIGATOIRE
  // Requis pour la validation des certificats de sécurité
  // Synchronisation et affichage sur le port série du temps
  // Pour une synchronisation plus rapide chercher le serveur NTP le plus proche: https://www.pool.ntp.org/zone/
  timeSync(TZ_INFO, "pool.ntp.org", "time.nis.gov");

  // Vérification de la connexion au serveur InfluxDB
  if (client.validateConnection()) {
    Serial.print("Connexion a InfluxDB valide: ");
    Serial.println(client.getServerUrl());
  } else {
    Serial.print("Connexion a InfluxDB echouer: ");
    Serial.println(client.getLastErrorMessage());
  }

  dht.begin();
}

// Fonction sender permet d'envoyer des informations a la base de données
void sender(String info, float content) {
  // Vide les champs précendant en gardant le point utilisé. Les tags restes présent
  sensor.clearFields();

  // Enregistre la valeur dans le point
  sensor.addField(info, content);

  // Affiche sur le port série le contenu de ce qui va être envoyé
  Serial.print("Ecriture: ");
  Serial.println(sensor.toLineProtocol());

  // Si il n'y a pas de WiFi, on essai de se reconnecter
  if ((WiFi.RSSI() == 0) && (wifiMulti.run() != WL_CONNECTED)) {
    Serial.println("Connexion WiFi perdu");
  }

  // On écrit (on envoie) le point dans la base de donnée
  if (!client.writePoint(sensor)) {
    Serial.print("Ecriture echouer InfluxDB: ");
    Serial.println(client.getLastErrorMessage());
  }
}


void loop() {
  float h = dht.readHumidity();
  float t = dht.readTemperature();

  if (isnan(h) && isnan(t)) {
    Serial.println("Erreur de lecture");
  } else {
    sender("Temperature", t);
    delay(1);
    sender("Humidity", h);
  }

  // Pause de 60s (1 minute)
  Serial.println("Pause de 60s");
  delay(60000);
}
