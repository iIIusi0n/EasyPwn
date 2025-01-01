import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../constants/colors.dart';
import '../components/custom_button.dart';

class HomePage extends StatelessWidget {
  const HomePage({super.key});

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
                'EasyPwn',
                style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.bold,
                  color: AppColors.textPrimary,
                ),
              ),
              const SizedBox(height: 8),
              Text(
                'Click the button below to get started',
                style: TextStyle(
                  fontSize: 16,
                  color: AppColors.greyShade(600),
                ),
                textAlign: TextAlign.center,
              ),
              const SizedBox(height: 32),
              CustomButton(
                text: 'Get to Login',
                onPressed: () => context.go('/login'),
              ),
              const SizedBox(height: 16),
              CustomButton(
                text: 'Go to Register',
                onPressed: () => context.go('/register'),
              ),
              const SizedBox(height: 16),
              CustomButton(
                text: 'Go to Instance',
                onPressed: () => context.go('/i'),
              ),
            ],
          ),
        ),
      ),
    );
  }
} 