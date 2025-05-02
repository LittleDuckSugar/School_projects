#include "DHT.h"
#define DHT_Broche 5
#define DHTTYPE DHT22

DHT dht(DHT_Broche, DHTTYPE);

//******* SETUP **************
void setup() {
  Serial.begin(9600);
  Serial.println("DHT22 test!");  dht.begin();
}
// ******* LOOP *************
void loop() {
  delay(2000); // Remarque : la lecture de la température et humidité nécessite 250 ms
  float Humidite = dht.readHumidity();
  float Temperature = dht.readTemperature();
  if (isnan(Humidite) || isnan(Temperature)) // isnan : nombre ? || : OU
  {
    Serial.println(" Echec de lecture des données du capteur DHT !");
    return;
  }
  // Facultatif : commande d’une action par relais si Humidité > 60% ou Température > 40°C

  Serial.print("Humidité: ");
  Serial.print(Humidite);
  Serial.print(" %\t");
  Serial.print("Température: ");
  Serial.print(Temperature);
  Serial.println(" *C ");
}
