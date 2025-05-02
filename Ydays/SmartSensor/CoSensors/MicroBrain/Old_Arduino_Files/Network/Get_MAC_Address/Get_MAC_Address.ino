/**
  IoT - SmartSensor
  Affiche sur le port série l'adresse MAC de la puce WiFi
  15 Novembre 2021
**/

#ifdef ESP32
#include <WiFi.h>
#else
#include <ESP8266WiFi.h>
#endif

void setup() {
  Serial.begin(115200);
  while (!Serial) {
    ; // En attente de connexion du port série
  }

  // On affiche sur le port série l'adresse MAC de la puce WiFi du microcontrolleur
  Serial.println();
  Serial.print("Adresse MAC: ");
  Serial.println(WiFi.macAddress());
}

void loop() {

}
