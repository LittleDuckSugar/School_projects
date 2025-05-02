#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# DEPRECATED ! Do not use this local sensor code
# See InfluxDB query call code

import Adafruit_DHT

class DHT:

    def __init__(self, pin):
        self.DHT_SENSOR = Adafruit_DHT.DHT22
        self.DHT_PIN = pin

    def ReturnTemp(self):
        return self.ReadSensor()[0]
        
    def ReturnHum(self):
        return self.ReadSensor()[1]
        
    def ReadSensor(self):
        self.humidity, self.temperature = Adafruit_DHT.read_retry(self.DHT_SENSOR, self.DHT_PIN)
        return round(self.temperature,1), round(self.humidity,1)