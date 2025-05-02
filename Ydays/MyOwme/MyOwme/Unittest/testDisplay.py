#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys
sys.path.append('/home/pi/MyOwme/')

from Display import Display
from Weather import Weather
import unittest
# Used in the updateTime methode
from datetime import datetime

class TestDisplay(unittest.TestCase):

    def test_readData(self):
        nextion = Display("/dev/ttyUSB0")
        print("Push music buttun")
        self.assertEqual(nextion.readData(5), "music")

    def test_sendData(self):
        nextion = Display("/dev/ttyUSB0")
        self.assertEqual(nextion.sendData("home.piece1", "txt", "\"Chambre:\""), 29)

    def test_updateTime(self):
        nextion = Display("/dev/ttyUSB0")

        test = len(str(datetime.now().year)) + len("rtc0=ÿÿÿ")
        test += len(str(datetime.now().month)) + len("rtc1=ÿÿÿ")
        test += len(str(datetime.now().day)) + len("rtc2=ÿÿÿ")
        test += len(str(datetime.now().hour)) + len("rtc3=ÿÿÿ")
        test += len(str(datetime.now().minute)) + len("rtc4=ÿÿÿ")
        test += len(str(datetime.now().second)) + len("rtc5=ÿÿÿ")

        self.assertEqual(nextion.updateTime(), test)

    def test___str__(self):
        nextion = Display("/dev/ttyUSB0")

        self.assertEqual(nextion.__str__(), nextion.ser.__str__())

    def test_updateWeather(self):
        city = Weather("Paris", "FR")
        nextion = Display("/dev/ttyUSB0")

        back = len("weather0.updated.txt=\"MAJ en cours\"ÿÿÿ")

        # Fetchs the latest weather informations before updating it
        city.updateData()
        
        now = datetime.now()
        if now.day <= 9:
            day = "0"+str(now.day)
        else:
            day = str(now.day)
        if now.month <= 9:
            month = "0"+str(now.month)
        
        # Hourly forecast (weather1 page)
        for hours in range(0, 8, 1):
            realHour = hours + now.hour +(3 * hours + 1) - hours
            if realHour >= 24:
                realHour -= 24

            back += len("weather1.h"+str(hours+1)+ ".txt=\""+str(realHour)+"h\"ÿÿÿ")
            back += len("weather1.t"+str(hours+1)+ ".txt=\""+str(round(city.one_call.forecast_hourly[hours*3].temperature('celsius')["temp"], 1))+"°\"ÿÿÿ")
            back += len("weather1.pl"+str(hours+1)+ ".txt=\""+str(int(city.one_call.forecast_hourly[hours*3].precipitation_probability*100))+"%\"ÿÿÿ")
            back += len("weather1.wp"+str(hours+1)+".pic="+ nextion.weatherPics[city.one_call.forecast_hourly[hours*3].weather_icon_name]+"ÿÿÿ")
        
        # Current weather info (weather0 page)
        back += len("weather0.temptxt.txt=\""+str(round(city.one_call.current.temperature('celsius')['temp'], 1))+"°C\"ÿÿÿ")
        back += len("weather0.feeltxt.txt=\"Ressentie "+str(round(city.one_call.current.temperature('celsius')['feels_like'], 1))+"°C\"ÿÿÿ")
        back += len("weather0.weatherpic.pic="+ nextion.weatherPics[city.one_call.current.weather_icon_name]+"ÿÿÿ")
        back += len("weather0.wind.val="+ str(city.one_call.current.wnd["deg"])+"ÿÿÿ")
        back += len("weather0.windtxt.txt=\""+str(round(city.one_call.current.wnd["speed"]*3.6, 1))+" km/h\"ÿÿÿ")
        back += len("weather0.statutxt.txt=\""+city.currentWeather().capitalize()+"\"ÿÿÿ")

        # Tells the user which city is updated
        back += len("weather0.citytxt.txt=\""+city.city.name+"\"ÿÿÿ")

        # Tells the user when the latest update happened
        back += len("weather0.updated.txt=\"MAJ "+day+"/"+month+" "+str(now.time())[:5]+"\"ÿÿÿ")



        self.assertEqual(nextion.updateWeather(city), back+10)


if __name__ == '__main__':
    unittest.main()

# # Instance of display
# nextion = Display("COM6")
# # Send data to display
# nextion.sendData("t0", "txt", "\"-12.8°C\"")
# # Read data from dispay
# print(nextion.readData(5))
# # Close serial communication of nextion
# nextion.ser.close()
# # Open communication again
# nextion.openSerial("COM6")
# # Send data to display
# nextion.sendData("t0", "txt", "\"24,6°F\"")
# # Use __str__ override methode
# print(nextion)
# # Close serial communication of nextion
# nextion.ser.close()
