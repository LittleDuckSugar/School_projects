#!/usr/bin/env python3
# -*- coding: utf-8 -*-

def day(t):
    for i in t:
        if i in ["aujourd'hui",'demain','lundi','mardi','jeudi','vendredi','samedi','dimanche']:
            return i
    return "aujourd'hui"
