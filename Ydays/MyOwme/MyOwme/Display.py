#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# Used for communication
import serial

# Used for listing ports
import serial.tools.list_ports

# Used for updateTime methode
from datetime import datetime

class Display:

    def __str__(self):
        # TODO return also nextion's user's preferences (which UI, lang, bright, color txt, sleep)
        return(str(self.ser))

    def __init__(self, port):
        self.openSerial(port)
        self.musicPics = {
            "play": "22",
            "pause": "21"
        }
        self.weatherPics = {
            "01d": "7",
            "01n": "8",
            "02d": "5",
            "02n": "6",
            "03d": "9",
            "03n": "9",
            "04d": "10",
            "04n": "10",
            "09d": "14",
            "09n": "14",
            "10d": "12",
            "10n": "13",
            "11d": "16",
            "11n": "16",
            "13d": "15",
            "13n": "15",
            "50d": "11",
            "50n": "11"
        }

        self.updateTime()

    def sendData(self, box, param, value):
        # Two ways to send data... Will see with time which one is better
        # data = []

        # endCom = [255, 255, 255]
        # endCom = bytearray(endCom)

        if box == "time":
            toSend = param+"="+value
        else:
            toSend = box+"."+param+"="+value

        # for char in toSend:
        #     data.append(ord(char))

        toSend = toSend.encode('utf-8')
        toSend += bytearray([255, 255, 255])

        # return(self.ser.write(bytearray(data)))
        return(self.ser.write(toSend))

    def readData(self, lenght):
        return self.ser.read(lenght).decode("utf-8")

    def updateTime(self):
        now = datetime.now()
        back = self.sendData("time", "rtc0", str(now.year))
        back += self.sendData("time", "rtc1", str(now.month))
        back += self.sendData("time", "rtc2", str(now.day))
        back += self.sendData("time", "rtc3", str(now.hour))
        back += self.sendData("time", "rtc4", str(now.minute))
        back += self.sendData("time", "rtc5", str(now.second))

        return back

    def openSerial(self, port):
        # self.ser = serial.Serial() # maybe change to this one ?
        
        # TODO autoconnect to nextion display when detected (ortherwise ask for the desired one ?)
        # for p in list(serial.tools.list_ports.comports()):
        #     print(p)
        #     if ("Nextion") in p.description:
        #         print("Connexion au port Nextion")


        self.ser = serial.Serial(port=port, baudrate=9600)

        

    def updateWeather(self, city):
        # Tells the user what we are doing
        back = self.sendData("weather0.updated", "txt", "\"MAJ en cours\"")

        # Fetchs the latest weather informations before updating it
        city.updateData()
        
        now = datetime.now()
        if now.day <= 9:
            day = "0"+str(now.day)
        else:
            day = str(now.day)
        if now.month <= 9:
            month = "0"+str(now.month)
        else:
            month = str(now.month)
        
        # Hourly forecast (weather1 page)
        for hours in range(0, 8, 1):
            realHour = now.hour +(3 * hours + 1)
            if realHour >= 24:
                realHour -= 24

            back += self.sendData("weather1.h"+str(hours+1), "txt", "\""+str(realHour)+"h\"")
            back += self.sendData("weather1.t"+str(hours+1), "txt", "\""+str(round(city.one_call.forecast_hourly[hours*3].temperature('celsius')["temp"], 1))+"°\"")
            back += self.sendData("weather1.pl"+str(hours+1), "txt", "\""+str(int(city.one_call.forecast_hourly[hours*3].precipitation_probability*100))+"%\"")
            back += self.sendData("weather1.wp"+str(hours+1), "pic", self.weatherPics[city.one_call.forecast_hourly[hours*3].weather_icon_name])
        
        # Current weather info (weather0 page)
        back += self.sendData("weather0.temptxt", "txt", "\""+str(round(city.one_call.current.temperature('celsius')['temp'], 1))+"°C\"")
        back += self.sendData("weather0.feeltxt", "txt", "\"Ressentie "+str(round(city.one_call.current.temperature('celsius')['feels_like'], 1))+"°C\"")
        back += self.sendData("weather0.weatherpic", "pic", self.weatherPics[city.one_call.current.weather_icon_name])
        back += self.sendData("weather0.wind", "val", str(city.one_call.current.wnd["deg"]))
        back += self.sendData("weather0.windtxt", "txt", "\""+str(round(city.one_call.current.wnd["speed"]*3.6, 1))+" km/h\"")
        back += self.sendData("weather0.statutxt", "txt", "\""+city.currentWeather().capitalize()+"\"")

        # Tells the user which city is updated
        back += self.sendData("weather0.citytxt", "txt", "\""+city.city.name+"\"")

        # Tells the user when the latest update happened
        back += self.sendData("weather0.updated", "txt", "\"MAJ "+day+"/"+month+" "+str(now.time())[:5]+"\"")
        return back