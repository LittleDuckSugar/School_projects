import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';

class UsersFirebase {
  //Values
  late String uid;
  late String mail;
  String? firstname;
  String? lastname;
  String? avatar;
  List? contacts = [];
  List? friendRequest = [];

  //Constructor
  UsersFirebase.vide();

  UsersFirebase(DocumentSnapshot snapshot) {
    Map<String, dynamic> map = snapshot.data() as Map<String, dynamic>;
    uid = snapshot.id;
    lastname = map["NOM"];
    firstname = map["PRENOM"];
    mail = map["MAIL"];
    avatar = map["AVATAR"];
    contacts = map["CONTACT"];
    friendRequest = map["Demande d'ami"];
  }
}
