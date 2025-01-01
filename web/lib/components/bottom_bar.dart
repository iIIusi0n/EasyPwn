import 'package:flutter/material.dart';
import '../constants/colors.dart';

class BottomBar extends StatelessWidget {
  final String instanceAddress;
  final String memoryUsage;

  const BottomBar({
    super.key,
    required this.instanceAddress,
    required this.memoryUsage,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 24,
      decoration: const BoxDecoration(
        color: AppColors.surface,
        border: Border(
          top: BorderSide(color: AppColors.border, width: 1.5),
        ),
      ),
      padding: const EdgeInsets.symmetric(horizontal: 16),
      child: Row(
        children: [
          Text(
            'Connected to: $instanceAddress',
            style: TextStyle(
              color: AppColors.greyShade(600),
              fontSize: 12,
            ),
          ),
          const Spacer(),
          Text(
            'Memory: $memoryUsage',
            style: TextStyle(
              color: AppColors.greyShade(600),
              fontSize: 12,
            ),
          ),
        ],
      ),
    );
  }
} 