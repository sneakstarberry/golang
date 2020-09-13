import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import 'dart:async';
import 'dart:convert';
User userFromJson(String str) => User.fromJson(json.decode(str));

String userToJson(User data) => json.encode(data.toJson());


List<User> parseUsers(String responseBody) {
  final parsed = json.decode(responseBody);
  final result = parsed['response'];
  var test = result.map<User>((json) => User.fromJson(json)).toList();
  print(test[0]);
  return result.map<User>((json) => User.fromJson(json)).toList();
}

Future<List<User>> fetchUsers(http.Client client) async {
  final response = await client.get('http://10.0.2.2:8888/api/v1/users');
  return compute(parseUsers, response.body);
}

class User {
    int id;
    String username;
    String email;
    String password;
    String avatarPath;
    DateTime createdAt;
    DateTime updatedAt;

    User({
        this.id,
        this.username,
        this.email,
        this.password,
        this.avatarPath,
        this.createdAt,
        this.updatedAt,
    });

    factory User.fromJson(Map<String, dynamic> json) => User(
        id: json["id"],
        username: json["username"],
        email: json["email"],
        password: json["password"],
        avatarPath: json["avatar_path"],
        createdAt: DateTime.parse(json["created_at"]),
        updatedAt: DateTime.parse(json["updated_at"]),
    );

    Map<String, dynamic> toJson() => {
        "id": id,
        "username": username,
        "email": email,
        "password": password,
        "avatar_path": avatarPath,
        "created_at": createdAt.toIso8601String(),
        "updated_at": updatedAt.toIso8601String(),
    };
}
