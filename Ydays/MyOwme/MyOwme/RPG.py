#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os.path
from jsonfile import WriteJson, ReadJson
import random
from STT import SpeechToText
from TTS import synthese
from Audio import Audio
from time import sleep



class RPG:

    def __init__(self):
        self.personnage = Personnage()
        self.run = False
        self.STT = SpeechToText()
        self.TTS = Audio()

    def __SaveExist(self, path: str) -> bool:
        if os.path.isfile(path) == True:
            if os.path.getsize(path) > 0:
                return True
            else:
                return False
        else:
            file = open(path, 'a')
            file.close()
            return False

    def __new(self):
        '''
        Method Private Qui init le Player
        '''
        self.personnage.genre, t = self.__choice("quel est le genre de ton personnage", {
                                                 "A": "féminin", "B": "masculin"})
        self.personnage.race, t = self.__choice("quel est la race de ton personnage", {
                                                "A": "Humain", "B": "Elfe", "C": "Nain"})
        # Peut être prendre le prenom dans CONFIG.json
        self.personnage.name, t = self.__choice("quel est ton nom", {})
        if self.personnage.name != "" and self.personnage.genre != "" and self.personnage.race != "":
            WriteJson("data/RPG/prof.json", self.personnage.GetDict())

    def __loop(self):
        if self.new:
            self.__new()
            self.map = Map()
        else:
            g = ReadJson("data/RPG/prof.json")
            if "start" in g:
                self.map = Map(g["start"])
                if len(self.map.F) <1:
                    self.__new()
                    self.map = Map()
            else:
                self.map = Map()
            self.personnage.SetInfo(g)
        while self.run:
            b, d = self.map.battle
            if b:
                b = Battle(d, self.personnage)
                win = b.fight()
                if win == None : 
                    self.run = False
                    return
                if self.map.rep == {}:
                    print(self.map.hist)
                    self.run = False
                else:
                    if win:
                        self.map.Change("win")
                    else:
                        self.map.Change("loose")
            elif self.map.rep != {}:
                t, rep = self.__choice(self.map.hist, self.map.rep)
                na = self.map.Change(rep)
                g = self.personnage.GetDict()
                g["start"] = self.map.Name
                WriteJson("data/RPG/prof.json", g)
            else:
                print(self.map.hist)
                self.run = False

    def Play(self):
        self.run = True
        self.new = True
        if self.__SaveExist("data/RPG/prof.json"):
            self.new = False
        self.__loop()
        g = self.personnage.GetDict()
        g["start"] = self.map.Name
        WriteJson("data/RPG/prof.json", g)

    def __choice(self, ques: str, reponse: dict):
        if not self.run:
            return
        keys, values = list(reponse.keys()), list(reponse.values())
        ble = False
        #print(ques)
        #tts
        synthese(ques)
        sleep(1)
        self.TTS.SystemPlay("return.wav")
        t = ""
        for i in range(len(reponse.items())):
            print(values)
            t +=  "réponse " + str(keys[i])+ " : "+ str(values[i]) + "\n"
        synthese(t)
        sleep(1)
        self.TTS.SystemPlay("return.wav")
        while ble == False:
            #sst
            command = self.STT.STT(False, True)
            #command = input("mets la réponse").lower()
            if command != None:
                command = command.lower()
                if command == "stop":
                    self.run = False
                    return command, None
                if reponse == {}:
                    if command != '':
                        ble = True
                        return command, None
                for i in range(len(reponse.items())):
                    if command == keys[i].lower() or  command == values[i].lower() :
                        ble = True
                        return values[i], keys[i]


class Personnage:

    def __init__(self):
        self.soin = [1, 10]
        self.genre = -1
        self.race = -1
        self.name = ""
        self.Att = 10
        self.Pv = 100
        self.Lv = 1
        self.PvMax = 100

    def GetDict(self) -> dict:
        return {"Soin": self.soin, "Genre": self.genre, "Race": self.race, "Name": self.name, "Att": self.Att, "Pv": self.Pv, "Lv": self.Lv, "PvMax": self.PvMax}

    def SetInfo(self, data: dict):
        t = data.keys()
        if "Soin" in t and "Genre" in t and "Race" in t and "Name" in t:
            self.genre = data["Genre"]
            self.race = data["Race"]
            self.name = data["Name"]
            self.soin = data["Soin"]  # Set all Data of Player
            if "Pv" in t and "Att" in t and "Lv" in t and "PvMax" in t:
                self.Pv = data["Pv"]
                self.Att = data["Att"]
                self.Lv = data["Lv"]
                self.PvMax = data["PvMax"]
            else:
                print("Error missing Pv | Att | Lv | PvMax")
        else:
            print("Data Incorrect")

    def AddPv(self, m: int):
        self.Pv += m




class Battle:

    def __init__(self, dict, player):
        self.STT = SpeechToText()
        self.TTS = Audio()
        self.player = player
        self.stats = dict["en"]
        self.pv = self.stats["pv"]
        j = ReadJson("data/RPG/prof.json")
        if "pvCurrent" in j.keys():
            self.pv = j["pvCurrent"]
        self.name = self.stats["name"]
        self.att = self.stats["att"]
        self.rep = self.stats["replique"]

    def choice(self, ques: str, reponse: dict):
        keys, values = list(reponse.keys()), list(reponse.values())
        ble = False
        synthese(ques)
        sleep(1)
        self.TTS.SystemPlay("return.wav")
        t = ""
        for i in range(len(reponse.items())):
            t +=  "réponse " + str(keys[i]) + " : "+ str(values[i]) + "\n"
        synthese(t)
        sleep(1)
        self.TTS.SystemPlay("return.wav")
        while ble == False:
            command = self.STT.STT(False, True)
            # command = input("mets la réponse").lower()
            if command != None:
                command = command.lower()
                if command == "stop":
                    self.run = False
                    g = ReadJson("data/RPG/prof.json")
                    g2 = self.player.GetDict()
                    g2["start"] = g["start"]
                    g2["pvCurrent"] = self.pv
                    WriteJson("data/RPG/prof.json", g2)
                    return None , None
                if reponse == {}:
                    if command != '':
                        ble = True
                        return command, None
                for i in range(len(reponse.items())):
                    print("command :", command)
                    if command == keys[i].lower() or  command == values[i].lower():
                        ble = True
                        return values[i], keys[i]

    def fight(self):
        self.run = True
        while self.run:
            keys, valu = self.choice("que veut tu faire", {
                                     "A": "attaquer", "B": "Soins", "C": "ma vie"})
            if keys == None and valu == None :
                return
            # calcule des dégâts à nerf
            valdegap = (self.player.Att + random.randint(0,
                        int(self.player.Att*0.3)))*self.player.Lv
            valdegae = self.att + random.randint(0, int(self.att*0.3))
            if valu == "A":
                self.pv -= valdegap
                synthese("tu as infliger "+ str(valdegap))
                sleep(1)
                self.TTS.SystemPlay("return.wav")
                synthese("il lui reste plus que "+ str(self.pv))
                sleep(1)
                self.TTS.SystemPlay("return.wav")
                self.player.AddPv(-valdegae)
                synthese("tu as subi "+ str(valdegae))
                sleep(1)
                self.TTS.SystemPlay("return.wav")
                synthese("il ne te reste plus que "+ str(self.player.Pv))
                sleep(1)
                self.TTS.SystemPlay("return.wav")
            elif valu == "B":
                if self.player.soin[0] > 0:
                    self.player.soin[0] -= 1
                    temp = self.player.Pv
                    if self.player.PvMax-self.player.Pv < self.player.soin[1]:
                        self.player.Pv = self.player.PvMax
                    else:
                        self.player.Pv += self.player.soin[1]
                    synthese("vous passer de "+ str(temp)+ " à "+ str(self.player.Pv))
                    sleep(1)
                    self.TTS.SystemPlay("return.wav")
                else:
                    synthese("vous n'avez pas de soin")
                    sleep(1)
                    self.TTS.SystemPlay("return.wav")
            elif valu == "C":
                synthese("il te reste "+ str(self.player.Pv)+
                      " et "+ str(self.pv)+ " à "+ str(self.name))
            if valdegae >= self.player.Pv <= 0:
                run = False
                return False
            if valdegap > self.pv <= 0:
                run = False
                self.__loot()
                j = ReadJson("data/RPG/prof.json")
                if "pvCurrent" in j.keys():
                    del(j["pvCurrent"])
                    WriteJson("data/RPG/prof.json", j)
                return True

    def __loot(self):
        for k in self.stats["loot"].keys():
            if k == "att":
                synthese("l'attaque + "+ str(self.stats["loot"][k]))
                sleep(1)
                self.TTS.SystemPlay("return.wav")
                self.player.Att += self.stats["loot"][k]
            elif k == "soin":
                synthese("le nombre de soin + "+ str(self.stats["loot"][k]))
                sleep(1)
                self.TTS.SystemPlay("return.wav")
                self.player.soin[0] += self.stats["loot"][k]
            elif k == "effsoin":
                synthese("l'efficacité des soins + "+ str(self.stats["loot"][k]))
                sleep(1)
                self.TTS.SystemPlay("return.wav")
                self.player.soin[1] += self.stats["loot"][k]
            elif k == "pvmax":
                synthese("t pv on fortement augmenté + "+str(self.stats["loot"][k]))
                sleep(1)
                self.TTS.SystemPlay("return.wav")
                self.player.PvMax += self.stats["loot"][k]
                self.player.Pv += self.stats["loot"][k]
            elif k == "lv":
                synthese("+ "+ str (self.stats["loot"][k])+ " niveau")
                self.player.Lv += self.stats["loot"][k]

class Map:
    def __init__(self, start="start"):
        self.dict = ReadJson("data/RPG/RPG.json")
        self.Name = start
        self.F = self.dict[start]["f"]

    def Change(self, rep):
        if rep in self.dict[self.Name]["rep"].keys():
            self.Name = self.F[rep]
            self.F = self.dict[self.Name]["f"]
            return self.Name
        return None

    def __str__(self):
        return self.Name + ", les fils" + str(self.F)

    @property
    def hist(self):
        return self.dict[self.Name]["hist"]

    @property
    def rep(self):
        return self.dict[self.Name]["rep"]

    @property
    def battle(self):
        if "battle" in self.dict[self.Name].keys():
            return True, self.dict[self.Name]["battle"]
        else:
            return False, None

if __name__ == "__main__":
    rpg = RPG()
    rpg.Play()