#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import time
import vlc
import os

class Audio:
    
    def __init__(self,lang="fr-FR"):
        self.vlc_instance = vlc.Instance()
        self.lang = lang
        self.player = self.vlc_instance.media_player_new()
        
        self.volume = 50


    # SystemPlay plays system's sounds
    #ajouter le fichier de la langue
    def SystemPlay(self, audioSystemName):

        if self.player.is_playing():
            self.player.pause()
            
        media = self.vlc_instance.media_new("Audio/System/"+self.lang+"/"+audioSystemName)
        self.player.set_media(media)
        self.player.play()
        time.sleep(1)
        duration = self.player.get_length() / 1000
        time.sleep(duration)
            
    # MusicPlay plays user's musics/playlists
    def MusicPlay(self, musicName, repeat):
        # self.vlc_instance.vlm_set_loop(repeat, True)
        while repeat != 0:
            media = self.vlc_instance.media_new("Audio/System/"+musicName)
            self.player.set_media(media)
            self.player.play()
            time.sleep(1)
            duration = self.player.get_length() / 1000
            time.sleep(duration)
            repeat-=1
        

    def MusicPandP(self):
        if self.player.is_playing():
            self.player.pause()
        else:
            self.player.play()
            self.pause = False
        
    # VolumeUp increases the volume of a specific channel
    def VolumeUp(self):
        if self.volume >= 99:
            self.volume = 100
        else :
            self.volume+=5
            
    def VolumeDown(self):
        if self.volume <= 1:
            self.volume = 0
        else:
            self.volume -= 5
        self.player.audio_set_volume(self.volume)
            



    # VolumeGetter return the volume of a specific channel
    def VolumeGetter(self):
        return self.volume

    # VolumeSetter apply the volume to a specific channel
    def VolumeSetter(self, volume):
        self.volume = volume
        self.player.audio_set_volume(self.volume)


    def addPlaylist(self, playlistName):
        # Add a playlist on a new empty media list
        self.mediaList = self.player.media_list_new()
        path = r"Audio/System/"+playlistName
        songs = os.listdir(path)
        for s in songs:
            self.mediaList.add_media(self.Player.media_new(os.path.join(path,s)))
        self.listPlayer = self.Player.media_list_player_new()
        self.listPlayer.set_media_list(self.mediaList)

    def rmPlaylist(self):
        # To remove the current playlist
        self.player.media_list_remove_index()
        
    def pPlay(self):
        self.listPlayer.play()

    def pNext(self):
        self.listPlayer.next()

    def pPause(self):
        self.listPlayer.pause()

    def pPrevious(self):
        self.listPlayer.previous()
        
    def pStop(self):
        self.listPlayer.stop()

    def pPlayback(self):
        self.listPlayer.set_playback_mode()
