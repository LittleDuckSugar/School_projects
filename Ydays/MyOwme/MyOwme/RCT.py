#!/usr/bin/env python3
# -*- coding: utf-8 -*-
from jsonfile import ReadJson

out = {}


class RCT:
    def __init__(self, lang="fr-FR"):
        self.d = ReadJson("data/RCT/"+lang+".json")
        self.keys = list(self.d.keys())

    def __str__(self) -> str:
        return self.d.__str__()

    def __exec(self, out, t, word):
        '''
        execute la fonction récuperér et return ça sorti
        '''
        for i in list(out.keys()):
            if i != "__builtins__" and i != "require":
                return word, out[i](t)
        return word, None

    def Rec(self, t: str):
        '''
        execute la fonction lié au keyword et return ça sorti
        '''
        t = t.lower().split(" ")
        for i in self.keys:
            for g in range(len(t)):
                if t[g] in self.d[i]["keyswords"]:
                    out = {"require": require}  #injecte require dans l'environement d'execution
                    exec(self.d[i]["func"], out) #recupère la fonction via le keyword et avec son environement d'execution
                    return self.__exec(out, t[g+1:], i)
        return None, None

# Fonction envoyé au script dans fr.json ou en.json


def require(path, params, nameOffEntryFunc="main"):
    '''
    permet d'executer un fichier python avec comme entrée une fonction main (modifiable avec nameOffEntryFunc)
    '''
    with open(path, "r") as f:
        t = f.read()
        temp = {}
        exec(t, temp)
        if nameOffEntryFunc in temp:
            return temp[nameOffEntryFunc](params)
        return None



if __name__ == "__main__":
    t = input(">>")
    g = RCT()
    print(g.Rec(t))
