import 'dart:developer';
import 'dart:io';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tpfinal/functions/firebaseHelper.dart';
import 'package:tpfinal/main.dart';

// ignore: camel_case_types
class register extends StatefulWidget {
  const register({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return registerState();
  }
}

// ignore: camel_case_types
class registerState extends State<register> {
  //values
  late String mail;
  late String password;
  late String lastname;
  late String firstname;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text("Inscription"),
        centerTitle: true,
      ),
      body: Container(
        padding: const EdgeInsets.all(20),
        child: bodyPage(),
      ),
    );
  }

  //show popUp when register is done
  popUpRegister() {
    showDialog(
        barrierDismissible: false,
        context: context,
        builder: (context) {
          if (Platform.isIOS) {
            return CupertinoAlertDialog(
              title: const Text(
                  "L'inscription est réussi ! Connectez-vous, en cliquant sur connexion."),
              actions: [
                ElevatedButton(
                    onPressed: () {
                      Navigator.pop(context);
                    },
                    child: const Text("OK"))
              ],
            );
          } else {
            return AlertDialog(
              title: const Text(
                  "L'inscription est réussi ! Connectez-vous,  en cliquant sur connexion."),
              actions: [
                ElevatedButton(
                    onPressed: () {
                      Navigator.pop(context);
                    },
                    child: const Text("OK"))
              ],
            );
          }
        });
  }

  Widget bodyPage() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.spaceAround,
      children: [
        //------------mail-----------------
        TextField(
          onChanged: (value) {
            setState(() {
              mail = value;
            });
          },
          decoration: InputDecoration(
              filled: true,
              fillColor: Colors.white,
              hintText: "Entrez votre mail",
              hintStyle:
                  const TextStyle(color: Color.fromARGB(255, 98, 23, 189)),
              border:
                  OutlineInputBorder(borderRadius: BorderRadius.circular(20))),
        ),

        //------------password-----------------
        TextField(
          obscureText: true,
          onChanged: (value) {
            setState(() {
              password = value;
            });
          },
          decoration: InputDecoration(
              filled: true,
              fillColor: Colors.white,
              hintText: "Entrez votre mot de passe",
              hintStyle:
                  const TextStyle(color: Color.fromARGB(255, 98, 23, 189)),
              border:
                  OutlineInputBorder(borderRadius: BorderRadius.circular(20))),
        ),

        //------------lastname-----------------
        TextField(
          onChanged: (value) {
            setState(() {
              lastname = value;
            });
          },
          decoration: InputDecoration(
              filled: true,
              fillColor: Colors.white,
              hintText: "Entrez votre nom",
              hintStyle:
                  const TextStyle(color: Color.fromARGB(255, 98, 23, 189)),
              border:
                  OutlineInputBorder(borderRadius: BorderRadius.circular(20))),
        ),
        //------------firstname-----------------
        TextField(
          onChanged: (value) {
            setState(() {
              firstname = value;
            });
          },
          decoration: InputDecoration(
              filled: true,
              fillColor: Colors.white,
              hintText: "Entrez votre prénom",
              hintStyle:
                  const TextStyle(color: Color.fromARGB(255, 98, 23, 189)),
              border:
                  OutlineInputBorder(borderRadius: BorderRadius.circular(20))),
        ),

        //------------button-----------------
        ElevatedButton(
            style: ElevatedButton.styleFrom(
                textStyle: const TextStyle(fontSize: 20),
                primary: const Color.fromARGB(255, 98, 23, 189),
                shape: RoundedRectangleBorder(
                    borderRadius: BorderRadius.circular(20))),
            onPressed: () {
              log("Inscription réussi.");
              popUpRegister();
              firebaseHelper()
                  .registerFirebase(lastname, firstname, mail, password);
            },
            child: const Text("Inscription")),

        //------------linkLogin-----------------
        InkWell(
          onTap: () {
            Navigator.push(context,
                MaterialPageRoute(builder: (BuildContext context) {
              return const MyHomePage(
                title: '',
              );
            }));
          },
          child: const Text(
            "Connexion",
            style: TextStyle(color: Colors.blue),
          ),
        ),
      ],
    );
  }
}
