// ignore_for_file: camel_case_types, prefer_final_fields

import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:flutter/material.dart';
import 'package:tpfinal/messenger.dart';
import 'package:tpfinal/model/usersfirebase.dart';
import 'package:tpfinal/functions/firebaseHelper.dart';
import 'package:flutter/src/widgets/async.dart';
import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:tpfinal/widget/profil.dart';

class contact extends StatefulWidget {
  const contact({Key? key}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return contactState();
  }
}

class contactState extends State<contact> {
  int _selectedIndex = 1;
  final pages = [const messenger(), const contact(), const profil()];

  // changePage() allows to navigate on the other page
  void changePage(int index) {
    setState(() {
      _selectedIndex = index;
    });
    Navigator.push(context, MaterialPageRoute(builder: (context) {
      return pages[index];
    }));
  }

  @override
  late List UserFirstName;
  late List UserLastName;
  String Search = "";

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

  /*
  printList(documents){

  }

   */

  Widget bodyPage() {
    return Container(
        child: Column(children: [
      TextField(
        onChanged: (String value) async {
          setState(() {
            Search = value;
            print("test");
            Text(Search);
          });
        },
      ),

      StreamBuilder<QuerySnapshot>(
          //stream: firebaseHelper().fireUser.where( "NOM",isEqualTo: Search).snapshots(),
          stream: firebaseHelper().fireUser.snapshots(),
          builder: (context, snapshot) {
            print("test");
            List documents = snapshot.data!.docs;
            return ListView.builder(
              itemCount: documents.length,
              itemBuilder: (context, index) {
                var result = documents[index];
                print("${result.firstname}  ${result.lastname}");
                return ListTile(
                  title: Text("${result.firstname}  ${result.lastname}"),
                );
              },
            );
          }),

      //---------demandes d'ami reçues-------------------
      const Text("demandes d'ami reçues"),
      /*
        ListView.builder(
          //List friendsadd = firebaseHelper.getFriendsRequests();
          shrinkWrap: true,
          itemCount: friendsAdd.length,
          itemBuilder: (context, index) {
            UsersFirebase user = UsersFirebase(friendsAdd[index]);
            return Center(
                child: Card(
                    child: Column(
                        mainAxisSize: MainAxisSize.min,
                        children: <Widget>[
                   ListTile(
                      title: Text("${user.firstname}  ${user.lastname}")),
                      //title: Text("test")),
                          Row(
                            mainAxisAlignment: MainAxisAlignment.end,
                            children: <Widget>[
                          TextButton(
                            child: const Text('ACCEPTER LA DEMANDE'),
                            onPressed: () {firebaseHelper.addFriend(user.uid)},
                          ),
                      const SizedBox(width: 8),
                    ],
                  ),
                ])));
          });*/
    ]));
  }
}

      /*
        return Column(
        children:
        [
        TextField(
        onChanged: (value) {
        setState(() {
        Search = value;

        });
        },
        ),
        ]
        );

     */



