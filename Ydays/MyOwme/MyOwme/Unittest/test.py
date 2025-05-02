import speech_recognition as s_r
print("Your speech_recognition version is: "+s_r.__version__)
r = s_r.Recognizer()
my_mic_device = s_r.Microphone(device_index=1)
with my_mic_device as source:
    print("what do you need from the weather ?")
    r.adjust_for_ambient_noise(source)
    audio = r.listen(source)
my_string=r.recognize_google(audio)
print(my_string)    
