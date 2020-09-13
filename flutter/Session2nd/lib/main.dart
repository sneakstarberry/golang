import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;

import 'login.dart';
import 'models.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      home: SessionApp(),
    );
  }
}

class SessionApp extends StatefulWidget {
  @override
  _SessionAppState createState() => _SessionAppState();
}

class _SessionAppState extends State<SessionApp> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Container(
          child: Column(
            children: <Widget>[
              Center(
                child: RaisedButton(
                  onPressed: _incrementCounter,
                  child: Text('Increment Counter'),
                ),
              ),
              Container(
                height: 500,
                child: ListView(
                  children: <Widget>[
                    RaisedButton(
                      child: Text("Login"),
                      onPressed: () { Navigator.push(context,
                          MaterialPageRoute(builder: (context) => LoginPage()));},
                    ),
                    FutureBuilder<List<User>>(
                        future: fetchUsers(http.Client()),
                        builder: (context, snapshot) {
                          return Container(
                            child: ListView.builder(
                                primary: false,
                                shrinkWrap: true,
                                itemCount: snapshot.data.length,
                                itemBuilder: (context, index) {
                                  return Container(
                                    height: 200,
                                    child: Column(
                                      children: <Widget>[
                                        Text(
                                            snapshot.data[index].id.toString()),
                                        Text(snapshot.data[index].username),
                                        Text(
                                            '${snapshot.data[index].updatedAt}'),
                                        Text(
                                            '${snapshot.data[index].createdAt}'),
                                      ],
                                    ),
                                  );
                                }),
                          );
                        })
                  ],
                ),
              )
            ],
          ),
        ),
      );
  }
}

_incrementCounter() async {
  SharedPreferences prefs = await SharedPreferences.getInstance();
  int counter = (prefs.getInt('counter') ?? 0) + 1;
  print('Pressed $counter times.');
  await prefs.setInt('counter', counter);
}
