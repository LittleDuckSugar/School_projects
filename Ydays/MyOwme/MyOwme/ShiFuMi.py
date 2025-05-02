#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from STT import SpeechToText
from random import choice
from Audio import Audio
from time import sleep

# from jsonfile import ReadJson

class ShiFuMi:
    """
    Jeu de ShiFuMi
    """
    def __init__(self,lang ="fr-FR"):
        self.run = False
        self.controller = Audio(lang)
        self.STT = SpeechToText()
        self.Pos = list(range(3))  # 0 : Pierre ; 1 : Feuille ; 2 : Cisaux # représente les combinaisons

    def Play(self):
        #TODO READ FILE json for get good response in all lang
        # d = ReadJson("Path") #Si on veut charger les reponse depuis un json
        d = {"pierre":0,"cisaux":2,"feuille":1,"papier":1,"rock":0} #représente les possibilités 
        self.run = True
        while self.run:
            sleep(2)
            self.controller.SystemPlay("/ShiFuMi/instructions.wav")
            command = self.STT.STT(False, True)
            playerCh = None
            if command != None:
                command = command.lower()
                print(command)
                if command == "stop":
                    self.run = False
                else:
                    try : 
                        playerCh = d[command]
                    except : 
                        playerCh = None
                if playerCh != None:
                    rbCh = choice(self.Pos)
                    print(rbCh)
                    if (rbCh == 2 and playerCh == 0) or (rbCh == 0 and playerCh == 2):
                        #verification des possibilité
                        if rbCh < playerCh:
                            sleep(2)
                            self.controller.SystemPlay("/ShiFuMi/loose.wav")
                            print("rb win")  # print rb win
                        else:
                            sleep(2)
                            self.controller.SystemPlay("/ShiFuMi/loose.wav")
                            print("player win")  # print player win
                    else:
                        if rbCh == playerCh:
                            sleep(2)
                            self.controller.SystemPlay("/ShiFuMi/equ.wav")
                            print("egality")  # egality
                        elif rbCh < playerCh:
                            sleep(2)
                            self.controller.SystemPlay("/ShiFuMi/win.wav")
                            print("Player Win")  # Player win
                        else:
                            sleep(2)
                            self.controller.SystemPlay("/ShiFuMi/loose.wav")
                            print("Rb Win")  # Rb win
                    self.run = False
