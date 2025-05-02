#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from itertools import count
from jsonfile import ReadJson,WriteJson ,ExistFile
import datetime
from TTS import synthese
from Audio import Audio
from STT import SpeechToText 
from time import sleep

PathConfFile = "data/config.json"


class Config:
    '''
    gere la configuration de l'appareille
    '''
    def __init__(self):
        if ExistFile(PathConfFile):
            self.__load(ReadJson(PathConfFile))
        else:
            self.conf()

    def __str__(self) -> str:
        return "L'utilisateur s'appelle " + self.name + " , il parle " + self.lang + ", il vient de " + self.city + " en " + self.country + " et est né le " + self._anniv

    # anniversaire format datetime

    @property
    def anniv(self):
        '''donne l'anniversaire au format datetime'''
        t = self._anniv.split("/")
        return datetime.datetime(int(t[2]), int(t[1]), int(t[0]))

    @anniv.setter
    def anniv(self, value: datetime.datetime):
        '''permet de set l'anniversaire au format datetime'''
        t = value.date()
        self._anniv = str(t.day) + "/" + str(t.month) + "/" + str(t.year)

    def __load(self, data):
        self.city = data["city"]
        self.lang = data["lang"]
        self.name = data["name"]
        self.country = data["country"]
        self._anniv = data["anniv"]

    def conf(self):
        '''permet la configuration du profile de l'utilisateur'''
        controller = Audio()
        STT = SpeechToText()
        run = True
        self.lang = None
        while self.lang != "fr" and self.lang != "us":
            print("Lang (fr | us):")
            self.lang = input()
        if self.lang == "fr":
            self.lang = "fr-FR"
            countrytxt = 'Entrer le code de votre pays (France = FR) :'
            nametxt = "Entrer Votre Prénom :"
            citytxt = "Entrer le nom de votre ville :"
            annivtxt = ("Quel est ton jour de naissance",
                        "Quel est ton mois de naissance", "Quel est ton année de naissance")
        else:
            self.lang = "en-US"
            countrytxt = ''
            nametxt = ""
            citytxt = ""
            annivtxt = ("", "", "")
        self.country =""
        while len(self.country) != 2:
            print(countrytxt)
            self.country = input()
        etape = 0
        intro = [citytxt, nametxt,
                 annivtxt[0], annivtxt[1], annivtxt[2]]
        temp = []
        while run:
            synthese(intro[etape])
            sleep(2)
            controller.SystemPlay("return.wav")
            command = STT.STT(False,True)
            #print(intro[etape])
            #command = input("")
            if command != None:
                if etape == 0:
                    self.city = command
                elif etape == 1:
                    self.name = command
                elif etape == 2:
                    try:
                        if 0 < int(command) < 32:
                            temp.append(command)
                        else:
                            print("jour invalide")
                            etape -= 1
                    except:
                        print("un Nombre")
                        etape -= 1
                elif etape == 3:
                    try:
                        if 0 < int(command) < 13:
                            temp.append(command)
                        else:
                            print("mois invalide")
                            etape -= 1
                    except:
                        print("un Nombre")
                        etape -= 1
                elif etape == 4:
                    try:
                        int(command)
                        temp.append(command)
                        print(temp)
                        self._anniv = temp[0] + "/" + temp[1] + "/" + temp[2]
                        run = False
                    except:
                        print("un Nombre")
                        etape -= 1
                etape += 1
        self.save()

    def save(self):
        WriteJson(PathConfFile, {"name": self.name, "city": self.city,
                  "country": self.country, "lang": self.lang, "anniv": self._anniv})





if __name__ == "__main__":
    c = Config()
    print(c)
