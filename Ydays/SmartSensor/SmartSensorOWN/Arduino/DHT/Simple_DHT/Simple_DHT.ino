/**
  IoT - SmartSensor
  Communication avec un capteur DHT22
  Affiche sur le port série la température et le taux d'humidité relative
  15 Novembre 2021
**/

// On importe la bibliothèque requise pour le capteur
#include "DHT.h"

// DHTPIN 5 --> PIN 1 physique
#define DHTPIN 5

// Instanciation du capteur
DHT dht(DHTPIN, DHT22);

void setup() {
  // On démmare l'échange sur le port série à la vitesse 9600 bit/s (bps)
  Serial.begin(9600);
  while (!Serial) {
    ; // En attente de connexion du port série
  }

  dht.begin();
}


void loop() {
  // Pause de 2 secondes
  delay(2000);

  // Lecture de l'humidité et de la température
  float h = dht.readHumidity();
  float t = dht.readTemperature();

  // Si au moins l'une des lectures échou alors on affiche "Erreur de lecture" sur le port série
  if (isnan(h) || isnan(t)) {
    Serial.println("Erreur de lecture");
  } else { // Sinon on affiche sur l'écran les valeurs
    Serial.print("Humidite: ");
    Serial.print(h);
    Serial.print(" %\t");
    Serial.print("Temperature: ");
    Serial.print(t);
    Serial.println(" *C");
  }
}
