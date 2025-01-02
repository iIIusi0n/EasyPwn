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
      body: Stack(
        children: [
          // Geometric Pattern Background
          CustomPaint(
            painter: GeometricPatternPainter(),
            size: Size.infinite,
          ),
          // Main Content
          SingleChildScrollView(
            child: Center(
              child: Container(
                padding: const EdgeInsets.symmetric(horizontal: 24, vertical: 48),
                constraints: const BoxConstraints(maxWidth: 1200),
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.center,
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    // Header Section
                    const SizedBox(height: 64),
                    const Text(
                      'EasyPwn',
                      style: TextStyle(
                        fontSize: 48,
                        fontWeight: FontWeight.bold,
                        color: AppColors.textPrimary,
                        letterSpacing: -0.5,
                      ),
                    ),
                    const SizedBox(height: 16),
                    Text(
                      'Web Debugger for CTF Pwnable Challenges',
                      style: TextStyle(
                        fontSize: 20,
                        color: AppColors.greyShade(700),
                        letterSpacing: 0.1,
                      ),
                    ),
                    const SizedBox(height: 48),
                    
                    // Features Grid
                    Wrap(
                      spacing: 24,
                      runSpacing: 24,
                      alignment: WrapAlignment.center,
                      children: [
                        _buildFeatureCard(
                          Icons.security,
                          'Advanced Security',
                          'Built-in protection and secure debugging environment',
                        ),
                        _buildFeatureCard(
                          Icons.speed,
                          'Real-time Analysis',
                          'Instant feedback and memory inspection tools',
                        ),
                        _buildFeatureCard(
                          Icons.integration_instructions,
                          'AI Integration',
                          'Smart suggestions and automated vulnerability detection',
                        ),
                      ],
                    ),
                    
                    const SizedBox(height: 64),
                    
                    // CTA Buttons
                    Wrap(
                      spacing: 16,
                      runSpacing: 16,
                      alignment: WrapAlignment.center,
                      children: [
                        CustomButton(
                          text: 'Join now!',
                          onPressed: () => context.go('/register'),
                          backgroundColor: AppColors.textPrimary,
                          textColor: Colors.white,
                          width: 200,
                          height: 48,
                        ),
                        CustomButton(
                          text: 'Go to Dashboard',
                          onPressed: () => context.go('/login'),
                          width: 200,
                          height: 48,
                        ),
                      ],
                    ),
                    const SizedBox(height: 96),
                  ],
                ),
              ),
            ),
          ),
        ],
      ),
    );
  }

  Widget _buildFeatureCard(IconData icon, String title, String description) {
    return Container(
      width: 320,
      padding: const EdgeInsets.all(24),
      decoration: BoxDecoration(
        color: AppColors.surface,
        borderRadius: BorderRadius.circular(8),
        border: Border.all(color: AppColors.border, width: 1.5),
      ),
      child: Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Icon(icon, size: 32, color: AppColors.textSecondary),
          const SizedBox(height: 16),
          Text(
            title,
            style: const TextStyle(
              fontSize: 18,
              fontWeight: FontWeight.bold,
              color: AppColors.textPrimary,
            ),
          ),
          const SizedBox(height: 8),
          Text(
            description,
            style: TextStyle(
              fontSize: 14,
              color: AppColors.greyShade(600),
            ),
            textAlign: TextAlign.center,
          ),
        ],
      ),
    );
  }
}

class GeometricPatternPainter extends CustomPainter {
  @override
  void paint(Canvas canvas, Size size) {
    final paint = Paint()
      ..color = AppColors.greyShade(200).withOpacity(0.5)
      ..strokeWidth = 1
      ..style = PaintingStyle.stroke;

    const spacing = 40.0;
    for (var i = 0.0; i < size.width + size.height; i += spacing) {
      canvas.drawLine(
        Offset(0, i),
        Offset(i, 0),
        paint,
      );
    }
  }

  @override
  bool shouldRepaint(CustomPainter oldDelegate) => false;
}