from email.mime import image
from guizero import *
from guizero import App, Combo, Text, MenuBar, TextBox, PushButton
import random
choix = False
french = False

# ---------------------------------------------Def-----------------------------------------------------------------------

def calcul(selected_value) :
    Raf = random.randint(1, 2)
    if selected_value == "Easy":
        if Raf == 1:
            a = random.randint(1, 10)
            b = random.randint(1, 10)
            Ores = str(a)  + "*" + str(b)
        if Raf == 2:
            a = random.randint(1, 10)
            b = random.randint(2, 3)
            Ores = str(a) + "%" + str(b)
    if selected_value == "Medium":
        if Raf == 1:
            a = random.randint(1, 9)
            b = random.randint(11, 99)
            Ores = str(a) + "*" + str(b)
        if Raf == 2:
            a = random.randint(11, 99)
            b = random.randint(2, 3 )
            Ores = str(a) + "%" + str(b)
    if selected_value == "Hard":
        if Raf == 1:
            a = random.randint(11, 99)
            b = random.randint(11, 99)
            Ores = str(a) + "*" + str(b)
        if Raf == 2:
            a = random.randint(11, 99)
            b = random.randint(11, 99)
            Ores = str(a) + "%" + str(b)
    if selected_value == "Master":
        if Raf == 2:
            a = random.randint(101, 999)
            b = random.randint(2, 4)
            Ores = str(a) + "%" + str(b)
        if Raf == 1:
            a = random.randint(101, 999)
            b = random.randint(101, 999)
            Ores = str(a) + "*" + str(b)
    print("calcul is : ", Ores)
    return Ores

    

def update_bg():
    app.bg = bg_combo.value

def update_text():
    app.text_color = text_combo.value
     
def button_pressed():
    emp = 0
    cal = calcul("Easy")
    test = app.question("Calcul", cal)
    if test is not None:
        
        for m in cal:
            if m.isdigit():
                emp = emp + int(m)
        print("result is ", emp)
        if test == emp:
            info("result", "good job !")
        else:
            info("result","sorry, bad answer")

def change_text_size(slider_value):
    welcome_message.size = slider_value


def msg(mm, res, pp):
    langue = app.yesno("Edit", mm)
    if langue == True:
        app.info("Edit", res)
        if pp == "reset":
            app.bg = "black"
            app.text_color = "white"
        if pp == "fr":
            french = True


def edit_function1():
    msg("Switch to french ?", "the language have been switch to french", "fr")

def edit_function2():

    msg("reset the colors ?", "the colors'll be refreshed", "reset")
    



def file_function():
    window = Window(app, title="Calul mental")
    window.display()
def quitF():
    app.destroy()

# ------------------------------------------English version-----------------------------------------------------------

if french == False:
    


    app = App(title="Mental calculation", width=800, height=480)
    colors = ["black", "white", "red", "green", "blue"]

    app.bg = "black"
    app.text_color = "white"

    title1 = Text(app, text="Background color")
    bg_combo = Combo(app, options=colors, selected=app.bg, command=update_bg)

    title2 = Text(app, text="Text color")
    text_combo = Combo(app, options=colors, selected=app.text_color, command=update_text)


    menubar = MenuBar(app,
                    toplevel=["File", "Settings"],
                    options=[
                        [ ["New", file_function] , ["Quit", quitF]],
                        [ ["langage", edit_function1], ["reset", edit_function2] ]
                    ])

    welcome_message = Text(app, text="choose a level",font="Banshrift", size="20", color="gray")
    choice = Combo(app, options=["Easy", "Medium", "Hard", "Master"], grid=[1,0], command=calcul)
    update_text = PushButton(app, command=button_pressed, text="Start")
    text_size = Slider(app, command=change_text_size, start=20, end=40)




# ------------------------------------------French version-----------------------------------------------------------------

# elif french == True:

#     app = App(title="Calcul mental", width=800, height=480)
#     colors = ["noire" , "blanc", "rouge", "vert", "bleu"]

#     app.bg = "black"
#     app.text_color = "white"

#     title1 = Text(app, text="Couleur de fond")
#     bg_combo = Combo(app, options=colors, selected=app.bg, command=update_bg)

#     title2 = Text(app, text="Couleur de texte")
#     text_combo = Combo(app, options=colors, selected=app.text_color, command=update_text)

#     menubar = MenuBar(app,
#                     toplevel=["File", "Edit"],
#                     options=[
#                         [ ["Nouveau", file_function] ],
#                         [ ["language", edit_function1], ["rafraichir les couleurs", edit_function2] ]
#                     ])


#     welcome_message = Text(app, text="choississez un niveau",font="Banshrift", size="20", color="gray")
#     choice = Combo(app, options=["Facile", "Moyen", "Difficile", "Expert"], grid=[1,0])

#     update_text = PushButton(app, command=button_pressed, text="Commencer")
#     text_size = Slider(app, command=change_text_size, start=20, end=40)
    app.update()
app.display()