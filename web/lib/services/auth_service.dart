import 'package:http/http.dart' as http;
import 'dart:convert';

class AuthService {
  Future<String> login(String email, String password) async {
    final response = await http.post(
      Uri.parse('/api/auth/login'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'email': email,
        'password': password,
      }),
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return data['token'];
    } else {
      throw Exception('Failed to login');
    }
  }

  Future<String> register(String email, String password, String confirmationCode) async {
    final response = await http.post(
      Uri.parse('/api/auth/register'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'email': email,
        'password': password,
        'code': confirmationCode,
      }),
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return data['token'];
    } else {
      throw Exception('Failed to register');
    }
  }

  Future<void> sendConfirmationEmail(String email) async {
    final response = await http.post(
      Uri.parse('/api/auth/confirm'),
      headers: {'Content-Type': 'application/json'},
      body: jsonEncode({
        'email': email,
      }),
    );

    if (response.statusCode != 200) {
      throw Exception('Failed to send confirmation email');
    }
  }
}