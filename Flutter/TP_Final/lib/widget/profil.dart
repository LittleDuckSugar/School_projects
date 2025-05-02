// ignore_for_file: camel_case_types, must_be_immutable, non_constant_identifier_names

import 'dart:developer';
import 'dart:io';
import 'dart:typed_data';
import 'package:file_picker/file_picker.dart';
import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:tpfinal/contact.dart';
import 'package:tpfinal/functions/firebaseHelper.dart';
import 'package:tpfinal/main.dart';
import 'package:tpfinal/messenger.dart';
import 'package:tpfinal/model/usersfirebase.dart';

class profil extends StatefulWidget {
  const profil({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return myprofil();
  }
}

class myprofil extends State<profil> {
  // ignore: unused_field
  // index of the page
  int _selectedIndex = 2;
  final pages = [const messenger(), const contact(), const profil()];

  UsersFirebase myProfil = UsersFirebase.vide();
  Uint8List? byteFile;
  String? nameFile;
  String? urlFile;
  TextEditingController rename_lastname = new TextEditingController();
  TextEditingController rename_firstname = new TextEditingController();

  //---------------------------------------Methods---------------------------------------
  // getImg allows to find a img into the storage of the phone
  getImg() async {
    FilePickerResult? result = await FilePicker.platform
        .pickFiles(withData: true, type: FileType.media);
    if (result != null) {
      setState(() {
        nameFile = result.files.first.name;
        byteFile = result.files.first.bytes;
        if (byteFile != null) {
          popUpImg();
        }
      });
    }
  }

  // changePage() allows to navigate on the other page
  void changePage(int index) {
    setState(() {
      _selectedIndex = index;
    });
    Navigator.push(context, MaterialPageRoute(builder: (context) {
      return pages[index];
    }));
  }

  // popUpImg() displays to the user if he wants to record the img
  popUpImg() {
    showDialog(
        barrierDismissible: false,
        context: context,
        builder: (context) {
          if (Platform.isIOS) {
            return CupertinoAlertDialog(
              title: const Text("Souhaitez-vous enregistrer cette image ?"),
              content: Image.memory(byteFile!),
              actions: [
                ElevatedButton(
                  onPressed: () {
                    Navigator.pop(context);
                  },
                  child: const Text("Annuler"),
                ),
                ElevatedButton(
                    onPressed: () {
                      firebaseHelper()
                          .storageImg(nameFile!, byteFile!)
                          .then((value) {
                        setState(() {
                          urlFile = value;
                          Map<String, dynamic> map = {
                            "AVATAR": urlFile,
                          };
                          myProfil.avatar = urlFile;
                          firebaseHelper().updateUser(myProfil.uid, map);
                          Navigator.pop(context);
                        });
                      });
                    },
                    child: const Text("Valider")),
              ],
            );
          } else {
            return AlertDialog(
              title: const Text("Voulez-vous enrgistrer cette image ?"),
              content: Image.memory(byteFile!),
              actions: [
                ElevatedButton(
                  onPressed: () {
                    //Annuler
                    Navigator.pop(context);
                  },
                  child: const Text("Annuler"),
                ),
                ElevatedButton(
                    onPressed: () {
                      //enregsiter notre image dans la base de donnée
                      firebaseHelper()
                          .storageImg(nameFile!, byteFile!)
                          .then((String urlImage) {
                        setState(() {
                          urlFile = urlImage;
                          Map<String, dynamic> map = {
                            "AVATAR": urlFile,
                          };
                          firebaseHelper().updateUser(myProfil.uid, map);
                        });
                      });

                      Navigator.pop(context);
                    },
                    child: const Text("Valider"))
              ],
            );
          }
        });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Container(
          padding: const EdgeInsets.all(20),
          child: bodyPage(),
        ),
        bottomNavigationBar: BottomNavigationBar(
          selectedFontSize: 15,
          selectedIconTheme: const IconThemeData(
              color: Color.fromARGB(255, 98, 23, 189), size: 30),
          selectedItemColor: const Color.fromARGB(255, 98, 23, 189),
          selectedLabelStyle: const TextStyle(fontWeight: FontWeight.bold),
          type: BottomNavigationBarType.fixed,
          items: const <BottomNavigationBarItem>[
            BottomNavigationBarItem(
              icon: Icon(Icons.chat),
              label: 'Messages',
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.book),
              label: 'Contacts',
            ),
            BottomNavigationBarItem(
              icon: Icon(Icons.person),
              label: 'Profil',
            ),
          ],
          currentIndex: _selectedIndex,
          onTap: changePage,
        ));
  }

  Widget bodyPage() {
    firebaseHelper().getID().then((String id) {
      setState(() {
        String userID = id;
        firebaseHelper().getUser(userID).then((UsersFirebase myuser) {
          setState(() {
            myProfil = myuser;
          });
        });
      });
    });

    return Container(
        padding: const EdgeInsets.all(20),
        alignment: Alignment.center,
        child: Column(
          children: [
            const SizedBox(
              height: 25,
            ),
            InkWell(
              child: Container(
                height: 160,
                width: 160,
                decoration: BoxDecoration(
                    shape: BoxShape.circle,
                    image: DecorationImage(
                        fit: BoxFit.fill,
                        image: (myProfil.avatar != null)
                            ? NetworkImage(myProfil.avatar!)
                            : const NetworkImage(
                                "https://firebasestorage.googleapis.com/v0/b/cours-1---dev-mobile.appspot.com/o/image%2Fuser.png?alt=media&token=fb5511e3-0aad-4cc3-9c15-cfd3c054cf52"))),
              ),
            ),
            const SizedBox(
              height: 25,
            ),

            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                IconButton(
                  onPressed: () {
                    log("Pressed - add picture");
                    getImg();
                  },
                  icon: const Icon(Icons.add_a_photo),
                  color: const Color.fromARGB(213, 71, 71, 71),
                  iconSize: 25,
                ),
                IconButton(
                  onPressed: () {
                    log("Pressed - change picture");
                    getImg();
                  },
                  icon: const Icon(Icons.add_photo_alternate),
                  color: const Color.fromARGB(213, 71, 71, 71),
                  iconSize: 25,
                ),
              ],
            ),

            const SizedBox(
              height: 30,
            ),
            // ---------Change lastname-------------
            TextField(
              decoration: InputDecoration(
                  border: const OutlineInputBorder(),
                  hintText: "Modifier votre nom",
                  labelText: "${myProfil.lastname}"),
              controller: rename_lastname,
            ),
            const SizedBox(
              height: 30,
            ),
            // ---------Change firstname-------------
            TextField(
              decoration: InputDecoration(
                  border: const OutlineInputBorder(),
                  hintText: "Modifier votre prenom",
                  labelText: "${myProfil.firstname}"),
              controller: rename_firstname,
            ),
            const SizedBox(
              height: 30,
            ),
            // ------------Save data------------------
            ElevatedButton.icon(
              onPressed: () {
                log("Pressed");
                myProfil.firstname = rename_firstname.text;
                myProfil.lastname = rename_lastname.text;
                Map<String, dynamic> changemap = {
                  "NOM": myProfil.lastname,
                  "PRENOM": myProfil.firstname
                };

                firebaseHelper().updateUser(myProfil.uid, changemap);
              },
              icon: const Icon(
                Icons.save,
                color: Colors.white,
                size: 25,
              ),
              label: const Text("Sauvegarder vos modifications"),
              style: ElevatedButton.styleFrom(
                primary: const Color.fromARGB(186, 74, 37, 160),
              ),
            ),

            const SizedBox(
              height: 30,
            ),

            // ------------logout User------------------
            InkWell(
              onTap: () {
                firebaseHelper().logoutFirebase().then((value) {
                  log("Déconnexion de l'utilisateur");
                  Navigator.push(context, MaterialPageRoute(builder: (context) {
                    return const MyHomePage(
                      title: "",
                    );
                  }));
                });
              },
              child: const Text(
                "Déconnexion",
                style: TextStyle(color: Colors.red, fontSize: 18),
              ),
            )
          ],
        ));
  }
}
