/**
  IoT - SmartSensor
  Point d'accès WiFi
  Affiche sur le port série les adresses IP et MAC des périphériques connecté
  15 Novembre 2021
**/

#include <WiFi.h>
#include "esp_wifi.h"

// On défini le SSID du réseau ainsi que le mot de passe requis
#define ssid "IoT-SmartSensor"
#define passphrase "iotroot9"

void setup() {
  // On démmare l'échange sur le port série à la vitesse 115200 bit/s (bps)
  Serial.begin(115200);
  while (!Serial) {
    ; // En attente de connexion du port série
  }
  Serial.println();

  // Configuration du point d'accès
  Serial.print("Configuration du point d'acces... ");
  Serial.println(WiFi.softAP(ssid, passphrase) ? "Ok\n" : "Erreur!\n");

  // On affiche sur le port série l'IP de l'ESP32
  Serial.print("IP de l'ESP32 : ");
  Serial.println(WiFi.softAPIP());

  // On affiche sur le port série le SSID et le mot de passe utilisé
  Serial.println("SSID : " ssid);
  Serial.println("Password : " passphrase);
}

void loop() {
  wifi_sta_list_t wifi_sta_list;
  tcpip_adapter_sta_list_t adapter_sta_list;

  memset(&wifi_sta_list, 0, sizeof(wifi_sta_list));
  memset(&adapter_sta_list, 0, sizeof(adapter_sta_list));

  esp_wifi_ap_get_sta_list(&wifi_sta_list);
  tcpip_adapter_get_sta_list(&wifi_sta_list, &adapter_sta_list);

  // Pour chaque element dans la liste on affiche les informations
  for (int i = 0; i < adapter_sta_list.num; i++) {

    tcpip_adapter_sta_info_t station = adapter_sta_list.sta[i];

    Serial.print("station nr ");
    Serial.println(i);

    Serial.print("MAC: ");

    // On parcours les éléments de l'adresse MAC
    for (int i = 0; i < 6; i++) {

      Serial.printf("%02X", station.mac[i]);
      if (i < 5)Serial.print(":");
    }

    Serial.print("\nIP: ");
    Serial.println(ip4addr_ntoa(&(station.ip)));
  }

  Serial.println();
  Serial.println("-------------------------");
  Serial.println();

  delay(5000);
}