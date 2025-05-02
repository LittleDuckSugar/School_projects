#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# Import pygame's audio handler
from pygame import mixer

class Audio:
    
    def __init__(self,lang="fr-FR"):
        # Audio setup
        mixer.init()
        self.Sound = mixer.Sound
        self.Music = mixer.music
        self.lang = lang
        # TODO get the default audio volume from the config.txt file
        self.volumeSys = 0.7
        self.volumeMus = 0.7

    # SystemPlay plays system's sounds
    #ajouter le fichier de la langue
    def SystemPlay(self, audioSystemName):
        self.Sound.play(mixer.Sound("Audio/System/"+self.lang+"/"+audioSystemName))

    # MusicPlay plays user's musics/playlists
    def MusicPlay(self, musicName, repeat = 0):
        test = mixer.Sound("Audio/UserSounds/"+musicName)  
        if self.Music.get_busy():
            self.Music.queue("Audio/UserSounds/"+musicName)
        else:
            self.Music.load("Audio/UserSounds/"+musicName)
        self.Music.play(repeat) # -1 means it will repeat forever
        self.Music.set_volume(self.volumeMus)
        return test

    def MusicPandP(self):
        
        if self.Music.get_busy():
            self.Music.pause()
            self.pause = True
        else:
            self.Music.unpause()
            self.pause = False
        
    # VolumeUp increases the volume of a specific channel
    def VolumeUp(self, channelNum):
        if self.VolumeGetter(channelNum) >= 0.9:
            self.VolumeSetter(channelNum, 1.0)
        else:
            self.VolumeSetter(channelNum, self.VolumeGetter(channelNum)+0.1)

    # VolumeDown decreases the volume of a specific channel
    def VolumeDown(self, channelNum):
        if self.VolumeGetter(channelNum) <= 0.1:
            self.VolumeSetter(channelNum, 0)
        else:
            self.VolumeSetter(channelNum, self.VolumeGetter(channelNum)-0.1)


    # VolumeGetter return the volume of a specific channel
    def VolumeGetter(self, channelNum):
        return self.mixer.channelNum.get_volume()

    # VolumeSetter apply the volume to a specific channel
    def VolumeSetter(self, channelNum, volume):
        self.mixer.channelNum.set_volume(volume)


#http://www.pygame.org/docs/ref/mixer.html

# from audio import MusicPlay, VolumeUp, VolumeSetter, VolumeGetter
# from time import sleep

# MusicPlay("test.wav")

# VolumeSetter(1, 0.5)

# while True:
#     sleep(2)
#     VolumeUp(1)
#     if VolumeGetter(1) == 1:
#         VolumeSetter(1,0)


# Starting the mixer
# mixer.init()
#
# Loading the song
# mixer.music.load("Audio/test.wav")
#
# # Setting the volume
# mixer.music.set_volume(0.7)
#
# # Start playing the song
# mixer.music.play()
#
# # infinite loop
# while True:
#
#     print("Press 'p' to pause, 'r' to resume")
#     print("Press 'e' to exit the program")
#     query = input("  ")
#
#     if query == 'p':
#
#         # Pausing the music
#         mixer.music.pause()
#     elif query == 'r':
#
#         # Resuming the music
#         mixer.music.unpause()
#     elif query == 'e':
#
#         # Stop the mixer
#         mixer.music.stop()
#         break
