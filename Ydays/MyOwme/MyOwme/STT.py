#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# Used for online STT
from re import L
import speech_recognition as sr

# Used to check if we have access to the Internet or not
from connected import isConnect

# Used to play sound to inform the user
from Audio import Audio

# Offline STT library
from vosk import Model, KaldiRecognizer

# Used for live mic STT
import pyaudio


class SpeechToText:

    def __init__(self, offline = False, lang="fr-FR"):
        self.controler = Audio(lang)
        self.lang=lang
        if offline == True:
            # TODO set offline STT to default STT source
            pass

        self.key = "ok maison"

    def ChangeKey(self, newkey):
        self.key = newkey

# TODO Change KeyDetector to support default key
    def KeyDetector(self, text):
        WIP = text.split(" ")

        if WIP[len(WIP)-2].lower() == "ok" and WIP[len(WIP)-1].lower() == "maison":
            self.controler.SystemPlay("siri.wav")
            return self.STT(False, True)
            
    def STT(self, reply, command = False):
        if isConnect():
            return self.OnlineSTT(reply, command)
        else:
            return self.OfflineSTT(reply, command)

    def OnlineSTT(self, reply, command):
        r  = sr.Recognizer()
        with sr.Microphone() as source:
            print("Dites quelque chose")
            audio = r.listen(source)
        try:
            text = r.recognize_google(audio, language=self.lang)
            print("Vous avez dit : " + text)

            if command == True:
                print("La commande est " + text)
                return text
            else:
                if reply == False:
                     return self.KeyDetector(text)
                else:
                    pass
        except sr.UnknownValueError:
            print("L'audio n'as pas été compris")
        except sr.RequestError as e:
            if isConnect():
                print("Connexion Internet OK")
            else:
                print("Connexion Internet pas OK")
            print("Le service Google Speech API ne fonctionne plus" + format(e))

    def OfflineSTT(self, reply, command):
        rec = KaldiRecognizer(Model("./Audio/model/"+self.lang), 16000)

        p = pyaudio.PyAudio()
        stream = p.open(format=pyaudio.paInt16, channels=1, rate=16000, input=True, frames_per_buffer=8000)
        stream.start_stream()

        while True:
            data = stream.read(4000)
            if len(data) == 0:
                break
            if rec.AcceptWaveform(data):
                # print(rec.Result())
                break
            else:
                # pass
                print(rec.PartialResult())
                # print("Message en cours de réception")

        text = (rec.FinalResult().split("\"text\" : \"")[1].split("\"\n"))[0]
        print("Vous avez dit : " + text)

        if command == True:
            print("La commande est " + text)
            return text
        else:
            if reply == False:
                self.KeyDetector(text)
        # print(rec.FinalResult())

