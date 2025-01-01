import 'package:flutter/material.dart';
import '../constants/colors.dart';

class CustomButton extends StatelessWidget {
  final String text;
  final VoidCallback onPressed;
  final double? width;
  final double? height;
  final Color? backgroundColor;
  final Color? borderColor;
  final Color? textColor;

  const CustomButton({
    super.key,
    required this.text,
    required this.onPressed,
    this.width = 400,
    this.height = 40,
    this.backgroundColor = AppColors.surface,
    this.borderColor = AppColors.borderDark,
    this.textColor = AppColors.textPrimary,
  });

  @override
  Widget build(BuildContext context) {
    // Calculate relative values based on height
    final double verticalPadding = (height ?? 40) * 0.2; // 20% of height
    final double fontSize = (height ?? 40) * 0.375; // 37.5% of height

    return SizedBox(
      width: width,
      height: height,
      child: TextButton(
        onPressed: onPressed,
        style: TextButton.styleFrom(
          backgroundColor: backgroundColor,
          padding: EdgeInsets.symmetric(vertical: verticalPadding),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(4),
            side: BorderSide(color: borderColor!, width: 1.5),
          ),
        ),
        child: Text(
          text,
          style: TextStyle(
            color: textColor,
            fontSize: fontSize,
            fontWeight: FontWeight.w600,
            letterSpacing: 0.3,
          ),
        ),
      ),
    );
  }
} 