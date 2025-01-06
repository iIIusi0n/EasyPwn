import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../components/elements/custom_input.dart';
import '../components/elements/custom_button.dart';
import '../constants/colors.dart';
import 'dart:async';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import '../services/auth_service.dart';

const _storage = FlutterSecureStorage();

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  State<RegisterPage> createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final _emailController = TextEditingController();
  final _passwordController = TextEditingController();
  final _confirmPasswordController = TextEditingController();
  final _confirmationCodeController = TextEditingController();
  bool _isConfirmationEnabled = false;
  bool _canSendEmail = true;
  int _remainingSeconds = 0;
  Timer? _timer;
  final _authService = AuthService();
  bool _isLoading = false;
  String? _errorMessage;
  String? _successMessage;

  @override
  void dispose() {
    _timer?.cancel();
    _emailController.dispose();
    _passwordController.dispose();
    _confirmPasswordController.dispose();
    _confirmationCodeController.dispose();
    super.dispose();
  }

  VoidCallback get _handleRegister => () async {
    if (_passwordController.text != _confirmPasswordController.text) {
      setState(() {
        _errorMessage = 'Passwords do not match';
        _successMessage = null;
      });
      return;
    }

    setState(() {
      _isLoading = true;
      _errorMessage = null;
      _successMessage = null;
    });

    try {
      final token = await _authService.register(
        _emailController.text,
        _passwordController.text,
        _confirmationCodeController.text,
      );

      _storage.write(key: 'token', value: token);

      if (mounted) {
        context.go('/projects');
      }
    } catch (e) {
      setState(() {
        _errorMessage = 'Failed to register. Please try again.';
        _successMessage = null;
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  };

  VoidCallback get _handleConfirmation => () async {
    if (_emailController.text.isEmpty) {
      setState(() {
        _errorMessage = 'Please enter your email';
      });
      return;
    }

    if (!_canSendEmail) {
      setState(() {
        _errorMessage = 'Please wait for the email to be sent';
      });
      return;
    }

    setState(() {
      _errorMessage = null;
      _successMessage = 'Sending confirmation email...';
    });

    try {
      await _authService.sendConfirmationEmail(_emailController.text);
      setState(() {
        _successMessage = 'Confirmation email sent successfully';
        _errorMessage = null;
      });
      _startEmailTimer();
    } catch (e) {
      setState(() {
        _successMessage = null;
        _errorMessage = 'Failed to send confirmation email';
      });
    } 
  };

  void _startEmailTimer() {
    setState(() {
      _canSendEmail = false;
      _isConfirmationEnabled = true;
      _remainingSeconds = 180; // 3 minutes
    });

    _timer?.cancel();
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      if (mounted) {
        setState(() {
          if (_remainingSeconds > 0) {
            _remainingSeconds--;
          } else {
            _canSendEmail = true;
            timer.cancel();
          }
        });
      }
    });
  }

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
                'Create account',
                style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.bold,
                  color: AppColors.textSecondary,
                ),
              ),
              const SizedBox(height: 8),
              Text(
                'Please enter your details to sign up',
                style: TextStyle(
                  fontSize: 16,
                  color: AppColors.greyShade(600),
                ),
              ),
              const SizedBox(height: 32),
              if (_successMessage != null)
                Padding(
                  padding: const EdgeInsets.only(bottom: 16),
                  child: Text(
                    _successMessage!,
                    style: const TextStyle(
                      color: Colors.green,
                      fontSize: 14,
                    ),
                  ),
                ),
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
              ),
              const SizedBox(height: 16),
              Row(
                children: [
                  Expanded(
                    child: CustomInput(
                      controller: _confirmationCodeController,
                      hintText: 'Confirmation Code',
                      enabled: _isConfirmationEnabled,
                    ),
                  ),
                  const SizedBox(width: 8),
                  SizedBox(
                    width: 60,
                    child: CustomButton(
                      text: _canSendEmail 
                          ? 'Send' 
                          : '${(_remainingSeconds / 60).floor()}:${(_remainingSeconds % 60).toString().padLeft(2, '0')}',
                      onPressed: _handleConfirmation,
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 16),
              CustomInput(
                controller: _passwordController,
                hintText: 'Password',
                isPassword: true,
              ),
              const SizedBox(height: 16),
              CustomInput(
                controller: _confirmPasswordController,
                hintText: 'Confirm Password',
                isPassword: true,
              ),
              const SizedBox(height: 24),
              CustomButton(
                text: _isLoading ? 'Signing up...' : 'Sign Up',
                onPressed: _isLoading ? (() {}) : _handleRegister,
              ),
              const SizedBox(height: 24),
              Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    'Already have an account? ',
                    style: TextStyle(color: AppColors.greyShade(600)),
                  ),
                  MouseRegion(
                    cursor: SystemMouseCursors.click,
                    child: GestureDetector(
                      onTap: () {
                        context.go('/login');
                      },
                      child: const Text(
                        'Sign in',
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