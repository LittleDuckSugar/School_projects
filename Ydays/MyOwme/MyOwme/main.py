#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# Own import
from Display import Display
from Game import Game   # Display object
from Weather import Weather   # Weather object
from STT import SpeechToText   # STT object
# from Sensor import DHT    # Sensor object
from Audio import Audio   #Import Audio object
from TTS import synthese
from ShiFuMi import ShiFuMi
from RCT import RCT
from Config import Config
from RPG import RPG


# System import
from threading import *   # Multithreading
from time import sleep
from datetime import timedelta,datetime

def newCity(cities, city, country):
    cities[city] = Weather(city, country)
    return cities

def delCity(cities, city):
    del cities[city]

cities = {}

# Instanciation
# Sensors
#sensor1 = DHT(4)
#sensor2 = DHT(17)

# Display
#nextion = Display("/dev/ttyUSB0")

# Voice recognition
STT = SpeechToText()


#RCT
rct = RCT()
trad = {"Monday":"lundi","Tuesday":"mardi","wednesday":"mercredi","Thursday":"jeudi","Friday":"vendredi","Saturday":"samedi","Sunday":"dimanche"}

# Music controller
controller = Audio()

# ShiFuMi game
shifumi = ShiFuMi()

#jeu calcul mental
game = Game()
# nextion.sendData("music.pandp", "pic", nextion.musicPics["pause"])

#Config
config = Config()


#RPG

rpg = RPG()


# Weather
cities = newCity(cities, config.city, config.country)


# nextion.updateWeather(cities.get(city))
# nextion.sendData("home.piece1", "txt", "\"Chambre:\"")
# nextion.sendData("home.piece2", "txt", "\"Salon:\"")



def DisplayHandler():
    while True:
        if (nextion.ser.is_open):
            receive = nextion.readData(5)
            if receive == "meteo":
                nextion.updateWeather(cities.get(config.city))
            elif receive == "pandp":
                controller.MusicPandP()
            elif receive == "music":
                current = str(timedelta(milliseconds=controller.Music.get_pos())).split(".")[0]
                if current[:2] == "0:":
                    current=current[2:]
                nextion.sendData("music.min", "val", current[:2])
                nextion.sendData("music.sec", "val", current[3:])


def DHTHandler():
    while True:
        nextion.sendData("home.localtemp1", "txt", "\""+str(sensor1.ReturnTemp())+"°C\"")
        nextion.sendData("home.localtemp2", "txt", "\""+str(sensor2.ReturnTemp())+"°C\"")
        nextion.sendData("home.localhum1", "txt", "\""+str(sensor1.ReturnHum())+"%\"")
        nextion.sendData("home.localhum2", "txt", "\""+str(sensor2.ReturnHum())+"%\"")
        sleep(60)


#displayThread=Thread(target=DisplayHandler) # creating Thread for the display handler
#DHTThread=Thread(target=DHTHandler) # creating Thread for the sensor handler

#displayThread.start()
#DHTThread.start()

def commandexec(command):
    key , arg = rct.Rec(command)
    if command == "mets à jour la météo":
        nextion.updateWeather(cities.get(config.city))
        pass
    elif command == "mets à jour l'heure":
        nextion.updateTime()
        pass
    elif command == "pause":
        controller.MusicPandP()
    elif command == "quelle est la température de la chambre":
        synthese("La température de la chambre est de "+str(sensor1.ReturnTemp())+"°C.")
        sleep(2)
        controller.SystemPlay("return.wav")
    elif command == "quelle est la température du salon":
        synthese("La température du salon est de "+str(sensor2.ReturnTemp())+"°C.")
        sleep(2)
        controller.SystemPlay("return.wav")
    elif command == "quel est le taux moyen d'humidité":
        calc = (sensor2.ReturnHum() + sensor1.ReturnHum())/2
        synthese("Le taux d'humidité moyen est de "+str(calc)+"%.")
        sleep(2)
        controller.SystemPlay("return.wav")
    elif key=="temp":
        if arg == "aujourd'hui":
            now = datetime.now()
            arg=trad[now.strftime("%A")]
        elif arg =="demain":
            now= datetime.datetime.now() + datetime.timedelta(days=1)
            arg=trad[now.strftime("%A")]
        cities[config.city].specificDay(arg)
        sleep(2)
        controller.SystemPlay("return.wav")
    elif key=="music":
        musiclen = controller.MusicPlay("Scream.wav")
        total = str(timedelta(seconds=musiclen.get_length())).split(".")[0]
        if total[:2] == "0:":
            total=total[2:]
        nextion.sendData("music.end", "txt", "\""+total+"\"")
        nextion.sendData("music.total", "val", str(musiclen.get_length()).split(".")[0])
    elif (key=="jeu" and arg == "rpg") or key == "rpg":
        rpg.Play()
    elif key=="jeu":
        shifumi.Play()
    elif key== "calcul":
        game.Start()
    else:
        print("Commande non reconnu")
        synthese("Commande non reconnu.")
        sleep(2)
        controller.SystemPlay("return.wav")

#nextion.updateWeather(cities[city])


# Main thread
run = True
while run:
    try:
        command = STT.STT(False)
        if command != None:
            commandexec(command)
    except KeyboardInterrupt:
        print("Bye")
        run = False