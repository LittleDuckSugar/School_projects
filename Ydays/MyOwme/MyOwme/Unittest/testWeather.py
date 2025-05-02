#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
sys.path.append('/home/pi/MyOwme/')

from Weather import Weather
import unittest

class TestWeather(unittest.TestCase):

    def test___str__(self):
        city = Weather("Paris", "FR")
        
        self.assertEqual(city.__str__(), city.one_call.__str__())

if __name__ == '__main__':
    unittest.main()
