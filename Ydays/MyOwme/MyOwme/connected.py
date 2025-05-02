#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# File is used to check if we have access to the Internet or not
import urllib.request

def isConnect():
    try:
        urllib.request.urlopen('http://google.com')
        return True
    except:
        return False
