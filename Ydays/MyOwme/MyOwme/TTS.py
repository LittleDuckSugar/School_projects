#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import os

# String to audio file

# V1 tout

def synthese(textetospeech: str, lang: str="fr-FR"):
    """
    param lang take value of param lang of pico2wave(Français: fr-FR,anglais US: en-US,anglais GB: en-GB, Espagnol: es-ES,Allemand: de-DE,Italien: it-IT)
    """
    os.system(
        "pico2wave -l {} -w Audio/System/{}/return.wav \"{}\"".format(lang,lang, textetospeech))


if __name__ == "__main__":
    synthese("hello world", "en-US")

# import linecache

# Ligne d'un fichier txt en fichier audio
# def synthesealine(line, filepath):
#    synthese(linecache.getline(filepath, line))
