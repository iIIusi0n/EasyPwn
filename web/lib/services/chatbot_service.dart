import 'dart:convert';

import 'package:http/http.dart' as http;

class ChatbotResponse {
  final String response;

  ChatbotResponse({required this.response});

  factory ChatbotResponse.fromJson(Map<String, dynamic> json) {
    return ChatbotResponse(response: json['response'].replaceAll('\\n', '\n'));
  }
}

class ChatbotService {
  final String token;

  ChatbotService({required this.token});

  Future<String> getResponse(String instanceId, String message) async {
    final response = await http.post(
      Uri.parse('/api/instance/$instanceId/chat'),
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode({
        'message': message,
      }),
    );
    if (response.statusCode == 200) {
      return ChatbotResponse.fromJson(jsonDecode(response.body)).response;
    } else {
      throw Exception('Failed to get response');
    }
  }
}
