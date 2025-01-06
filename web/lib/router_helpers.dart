import 'package:flutter/material.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:go_router/go_router.dart';
import 'services/auth_service.dart';

const _storage = FlutterSecureStorage();
final _authService = AuthService();

Future<String?> redirectIfAuthenticated(BuildContext context, GoRouterState state) async {
  final token = await _storage.read(key: 'token');
  if (token == null) return null;

  try {
    await _authService.validateToken(token);
    return '/projects';
  } catch (e) {
    await _storage.delete(key: 'token');
    return null;
  }
} 