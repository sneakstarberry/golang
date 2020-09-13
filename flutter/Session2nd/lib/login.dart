import 'package:flutter/material.dart';

class LoginPage extends StatefulWidget {
  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  @override
  Widget build(BuildContext context) {
    final double screenHeight = MediaQuery.of(context).size.height;
    final double screenWidth = MediaQuery.of(context).size.width;
    return Scaffold(
      appBar: AppBar(
        actions: <Widget>[
          SizedBox(width: 40),
        ],
        centerTitle: true,
        title: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Text("한국어", style: TextStyle(color: Colors.black, fontSize: 15, fontWeight: FontWeight.w600)),
            SizedBox(
              width: 5,
            ),
            Icon(Icons.keyboard_arrow_down, color: Colors.grey, size: 15),
          ],
        ),
        leading: IconButton(
          icon: Icon(
            Icons.menu,
            color: Colors.black,
          ),
          onPressed: () => print("onPressed!!"),
        ),
        backgroundColor: Colors.white,
        elevation: 0,
      ),
      body: Container(
        height: screenHeight - 80,
        color: Colors.white,
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.center,
          children: <Widget>[
            Container(
              height: screenHeight - 150,
              child: ListView(
                padding: EdgeInsets.symmetric(horizontal: 30),
                children: <Widget>[
                  SizedBox(height: 40),
                  Container(
                    height: 80,
                    width: 140,
                    child: Image.asset(
                      "assets/Instagram_logo.PNG",
                      fit: BoxFit.contain,
                    ),
                  ),
                  SizedBox(
                    height: 40,
                  ),
                  Container(
                    height: 40,
                    width: double.infinity,
                    child: FlatButton(
                      color: Color(0xff0195F7),
                      onPressed: () => print("facebook"),
                      child: Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: <Widget>[
                          Icon(
                            Icons.local_activity,
                            color: Colors.white,
                          ),
                          SizedBox(
                            width: 8,
                          ),
                          Text("Facebook으로 계속하기",
                              style: TextStyle(
                                  fontSize: 16,
                                  color: Colors.white,
                                  fontWeight: FontWeight.bold)),
                        ],
                      ),
                    ),
                  ),
                  SizedBox(
                    height: 20,
                  ),
                  Container(
                    height: 30,
                    width: double.infinity,
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: <Widget>[
                        Container(height: 1, width: 130, color: Colors.grey),
                        Text("또는", style: TextStyle(fontSize: 16)),
                        Container(height: 1, width: 130, color: Colors.grey),
                      ],
                    ),
                  ),
                  Form(
                    child: Column(
                      children: <Widget>[
                        SizedBox(
                          height: 20,
                        ),
                        Container(
                          height: 50,
                          child: TextField(
                            decoration: InputDecoration(
                              hintText: "이메일",
                              fillColor: Color(0xfffafafa),
                              filled: true,
                              border: const OutlineInputBorder(),
                            ),
                          ),
                        ),
                        SizedBox(
                          height: 8,
                        ),
                        Container(
                          height: 50,
                          child: TextField(
                            decoration: InputDecoration(
                              hintText: "비밀번호",
                              fillColor: Color(0xfffafafa),
                              filled: true,
                              border: const OutlineInputBorder(),
                            ),
                          ),
                        ),
                      ],
                    ),
                  ),
                  SizedBox(height: 12),
                  InkWell(
                    onTap: () => print("비밀번호를?!"),
                    child: Container(
                      alignment: Alignment.centerRight,
                      child: Text(
                        "비밀번호를 잊으셨나요?",
                        style: TextStyle(
                            color: Color(0xff0195F7),
                            fontWeight: FontWeight.w600),
                      ),
                    ),
                  ),
                  SizedBox(
                    height: 20,
                  ),
                  Container(
                    width: double.infinity,
                    child: FlatButton(
                      color: Color(0xff0195F7),
                      onPressed: () => print("로그인"),
                      child: Text(
                        "로그인",
                        style: TextStyle(
                            color: Colors.white, fontWeight: FontWeight.w800),
                      ),
                    ),
                  ),
                  SizedBox(
                    height: 16,
                  ),
                  Align(
                    alignment: Alignment.center,
                    child: RichText(
                      text: TextSpan(
                        children: <TextSpan>[
                          TextSpan(
                              text: '계정이 없으신가요?',
                              style: TextStyle(
                                  color: Colors.grey,
                                  fontWeight: FontWeight.bold)),
                          TextSpan(
                              text: ' 가입하기',
                              style: TextStyle(
                                  color: Color(0xff0195F7),
                                  fontWeight: FontWeight.w700)),
                        ],
                      ),
                    ),
                  ),
                ],
              ),
            ),
            Container(
              width: double.infinity,
              height: 70,
              decoration: BoxDecoration(
                  color: Color(0xfffafafa),
                  border: Border(
                      top: BorderSide(width: 2, color: Color(0xffe7e7e7)))),
              child: Column(
                mainAxisAlignment: MainAxisAlignment.center,
                crossAxisAlignment: CrossAxisAlignment.center,
                children: <Widget>[
                  Text("from", style: TextStyle(color: Color(0xffb3b3b3), )),
                  SizedBox(height: 6,),
                  Text("FACEBOOK", style: TextStyle(fontWeight: FontWeight.bold, letterSpacing: 3))
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
