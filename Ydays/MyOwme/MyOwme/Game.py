from Audio import Audio
from math import sqrt
from time import sleep
import random
from STT import SpeechToText
from TTS import synthese

class Game : #jeu calcul mental

    def __init__(self) :
        self.play = False
        self.controller = Audio()
        self.stop = False
        self.res = int
        self.calcul = int
        self.STT = SpeechToText()
        self.Otext = list(range(2))
        self.langage = 'fr'


    def Start(self) :
        self.play = True
        while (self.play) :
            command = self.STT.STT(False,True)
            sleep(3)
        
            if command != None:
                if command == "stop":
                    self.play = False

                if command == "facile" :
                    Raf = random.randint(1, 2)
                    if Raf == 1:
                        a = random.randint(1, 10)
                        b = random.randint(1, 10)
                        Ores = a * b
                    if Raf == 2:
                        a = random.randint(1, 10)
                        b = random.randint(2, 3)
                        Ores = a % b
                if command == "moyen":
                    Raf = random.randint(1, 3)
                    if Raf == 1:
                        a = random.randint(1, 9)
                        b = random.randint(11, 99)
                        Ores = a * b
                    if Raf == 2:
                        a = random.randint(11, 99)
                        b = random.randint(2, 3 )
                        Ores = a % b
                    if Raf == 3:
                        a = random.randint(1, 9)
                        self.Otext.append('carré de ')
                        self.Otext.append(a)
                        Ores = sqrt(a)
                if command == "difficile":
                    Raf = random.randint(1, 2)
                    if Raf == 1:
                        a = random.randint(11, 99)
                        b = random.randint(11, 99)
                        Ores = a * b
                    if Raf == 2:
                        a = random.randint(11, 99)
                        b = random.randint(11, 99)
                        Ores = a % b
                if command == "extreme":
                    Raf = random.randint(1, 3)
                    if Raf == 3:
                        a = random.randint(16, 99)
                        self.Otext.append('carré de ')
                        self.Otext.append(a)
                        Ores = sqrt(a)
                    if Raf == 2:
                        a = random.randint(101, 999)
                        b = random.randint(2, 4)
                        Ores = a % b
                    if Raf == 1:
                        a = random.randint(101, 999)
                        b = random.randint(101, 999)
                        Ores = a * b
            res = self.STT.STT(False,True)
            if res != None:
                if res == Ores:
                    # self.controller.SystemPlay("Audio/System/mixkit-quick-win-video-game-notification-269.wav")
                    print("good")
                else :
                    # self.controller.SystemPlay("Audio/System/mixkit-losing-piano-2024.wav")
                    print("false")
            self.play = False
