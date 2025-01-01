import 'package:flutter/material.dart';

class AppColors {
  static const Color primary = Color(0xFF2C2C2C);
  static const Color secondary = Color(0xFF404040);

  static const Color surface = Color(0xFFFAFAFA);
  static const Color surfaceDark = Color(0xFF1E1E1E);
  static const Color background = Color(0xFFFFFFFF);

  static const Color textPrimary = Color(0xFF1A1A1A);
  static const Color textSecondary = Color(0xFF404040);
  static const Color textConsoleLight = Color(0xFFF0F0F0);

  static const Color border = Color(0xFFE0E0E0);
  static const Color borderDark = Color(0xFF2C2C2C);


  static Color greyShade(int weight) {
    return Colors.grey[weight]!;
  }
} 