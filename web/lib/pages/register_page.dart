import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../components/elements/custom_input.dart';
import '../components/elements/custom_button.dart';
import '../constants/colors.dart';
import 'dart:async';

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

  @override
  void dispose() {
    _emailController.dispose();
    _passwordController.dispose();
    _confirmPasswordController.dispose();
    _confirmationCodeController.dispose();
    super.dispose();
  }

  void _startEmailTimer() {
    setState(() {
      _canSendEmail = false;
      _isConfirmationEnabled = true;
      _remainingSeconds = 180; // 3 minutes
    });

    Timer.periodic(const Duration(seconds: 1), (timer) {
      setState(() {
        if (_remainingSeconds > 0) {
          _remainingSeconds--;
        } else {
          _canSendEmail = true;
          timer.cancel();
        }
      });
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
                      onPressed: () {
                        if (_canSendEmail && _emailController.text.isNotEmpty) {
                          _startEmailTimer();
                        }
                      },
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
                text: 'Sign Up',
                onPressed: () {
                  // Add register logic
                },
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