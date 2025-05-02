#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
sys.path.append('/home/pi/MyOwme/')

from STT import SpeechToText
import unittest

class TestSTT(unittest.TestCase):

    def test_STT(self):
        STT = SpeechToText()
        
        self.assertEqual(STT.STT(False), "coucou")

if __name__ == '__main__':
    unittest.main()
