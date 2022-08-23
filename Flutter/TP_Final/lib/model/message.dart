import 'package:cloud_firestore/cloud_firestore.dart';

class Message {
  String id = "";
  String who = "";
  String receiver = "";
  String content = "";
  DateTime dateMessage = DateTime.now();

  // constructor
  Message.vide();

  Message(DocumentSnapshot snapshot) {
    id = snapshot.id;
    Map<String, dynamic> mapMessage = snapshot.data() as Map<String, dynamic>;
    who = mapMessage["De"];
    receiver = mapMessage["Destinataire"];
    content = mapMessage["Texte"];
    Timestamp time = mapMessage["Date"];
    dateMessage = time.toDate();
  }
}
