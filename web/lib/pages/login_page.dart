import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import '../components/elements/custom_input.dart';
import '../components/elements/custom_button.dart';
import '../constants/colors.dart';
import '../services/auth_service.dart';

const _storage = FlutterSecureStorage();

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();
  final _authService = AuthService();
  bool _isLoading = false;
  String? _errorMessage;

  @override
  void dispose() {
    _emailController.dispose();
    _passwordController.dispose();
    super.dispose();
  }

  VoidCallback get _handleLogin => () async {
    setState(() {
      _isLoading = true;
      _errorMessage = null;
    });

    try {
      final token = await _authService.login(
        _emailController.text,
        _passwordController.text,
      );
      
      await _storage.write(key: 'token', value: token);
      
      if (mounted) {
        context.go('/projects');
      }
    } catch (e) {
      setState(() {
        _errorMessage = 'Invalid email or password';
      });
    } finally {
      if (mounted) {
        setState(() {
          _isLoading = false;
        });
      }
    }
  };

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.background,
      body: Center(
        child: Container(
          padding: const EdgeInsets.all(48),
          constraints: const BoxConstraints(maxWidth: 400),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              const Text(
                'Welcome back',
                style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.bold,
                  color: AppColors.textSecondary,
                ),
              ),
              const SizedBox(height: 8),
              Text(
                'Please enter your details to sign in',
                style: TextStyle(
                  fontSize: 16,
                  color: AppColors.greyShade(600),
                ),
              ),
              const SizedBox(height: 32),
              if (_errorMessage != null)
                Padding(
                  padding: const EdgeInsets.only(bottom: 16),
                  child: Text(
                    _errorMessage!,
                    style: const TextStyle(
                      color: Colors.red,
                      fontSize: 14,
                    ),
                  ),
                ),
              CustomInput(
                controller: _emailController,
                hintText: 'Email',
                enabled: !_isLoading,
              ),
              const SizedBox(height: 16),
              CustomInput(
                controller: _passwordController,
                hintText: 'Password',
                isPassword: true,
                enabled: !_isLoading,
              ),
              const SizedBox(height: 24),
              CustomButton(
                text: _isLoading ? 'Signing in...' : 'Sign In',
                onPressed: _isLoading ? (() {}) : _handleLogin,
              ),
              const SizedBox(height: 24),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    'Don\'t have an account? ',
                    style: TextStyle(color: AppColors.greyShade(600)),
                  ),
                  MouseRegion(
                    cursor: SystemMouseCursors.click,
                    child: GestureDetector(
                      onTap: () {
                        context.go('/register');
                      },
                      child: const Text(
                        'Sign up',
                        style: TextStyle(
                          color: AppColors.textSecondary,
                          fontWeight: FontWeight.w500,
                        ),
                      ),
                    ),
                  ),
                ],
              ),
            ],
          ),
        ),
      ),
    );
  }
} 