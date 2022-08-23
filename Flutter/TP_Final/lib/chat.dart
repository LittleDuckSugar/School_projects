// ignore_for_file: camel_case_types

import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:tpfinal/functions/firebaseHelper.dart';
import 'package:tpfinal/model/usersfirebase.dart';

class chat extends StatefulWidget {
  const chat({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return chatState();
  }
}

class chatState extends State<chat> {
  late String msgtext;
  UsersFirebase user = UsersFirebase.vide();
  UsersFirebase friends = UsersFirebase.vide();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        elevation: 0,
        automaticallyImplyLeading: false,
        backgroundColor: Colors.white,
        flexibleSpace: SafeArea(
          child: Container(
            padding: const EdgeInsets.only(right: 16),
            child: Row(
              children: <Widget>[
                IconButton(
                  onPressed: () {
                    Navigator.pop(context);
                  },
                  icon: const Icon(
                    Icons.arrow_back,
                    color: Colors.black,
                  ),
                ),
                const SizedBox(
                  width: 2,
                ),
                const CircleAvatar(
                  backgroundImage: NetworkImage(
                      "https://firebasestorage.googleapis.com/v0/b/cours-1---dev-mobile.appspot.com/o/image%2Fuser.png?alt=media&token=fb5511e3-0aad-4cc3-9c15-cfd3c054cf52"),
                  maxRadius: 20,
                ),
                const SizedBox(
                  width: 12,
                ),
                Expanded(
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.start,
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: const [
                      Text(
                        "Damien DD",
                        style: TextStyle(
                            fontSize: 16, fontWeight: FontWeight.w600),
                      ),
                      SizedBox(
                        height: 6,
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
      body: bodyPage(),
    );
  }

  Widget bodyPage() {
    firebaseHelper().getID().then((String id) {
      setState(() {
        String userID = id;
        firebaseHelper().getUser(userID).then((UsersFirebase usr) {
          setState(() {
            user = usr;
          });
        });
      });
    });
    String otheruserID = "xHoEOiQPv9aLGbC5eoVBlMT8vX42";
    firebaseHelper().getUser(otheruserID).then((UsersFirebase oth) {
      setState(() {
        friends = oth;
      });
    });
    return Stack(
      children: [
        Container(
          child: Text("coucou"),
        ),
        Align(
          alignment: Alignment.bottomLeft,
          child: Container(
            padding: const EdgeInsets.only(left: 10, bottom: 10, top: 10),
            height: 60,
            width: double.infinity,
            color: Colors.white,
            child: Row(
              children: [
                const SizedBox(
                  width: 15,
                ),
                Expanded(
                  child: TextField(
                    onChanged: (value) {
                      setState(() {
                        msgtext = value;
                      });
                    },
                    decoration: const InputDecoration(
                        hintText: "Ecrire un message...",
                        hintStyle: TextStyle(color: Colors.black54),
                        border: InputBorder.none),
                  ),
                ),
                const SizedBox(
                  width: 15,
                ),
                FloatingActionButton(
                  onPressed: () {
                    log("Création du message");
                    firebaseHelper().sendMsg(msgtext, user, friends);
                  },
                  child: const Icon(
                    Icons.send,
                    color: Colors.white,
                    size: 18,
                  ),
                  backgroundColor: Colors.blue,
                  elevation: 0,
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }
}
