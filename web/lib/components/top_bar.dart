import 'package:flutter/material.dart';
import '../constants/colors.dart';

class TopBar extends StatelessWidget {
  final String path;

  const TopBar({
    super.key,
    required this.path,
  });

  @override
  Widget build(BuildContext context) {
    final pathSegments = path.split('/');
    final lastSegment = pathSegments.isNotEmpty ? pathSegments.last : '';
    final parentPath = pathSegments.length > 1 
        ? '${pathSegments.sublist(0, pathSegments.length - 1).join('/')}/' 
        : '';

    return Container(
      height: 60,
      decoration: const BoxDecoration(
        color: AppColors.surface,
        border: Border(
          bottom: BorderSide(color: AppColors.border, width: 1.5),
        ),
      ),
      padding: const EdgeInsets.symmetric(horizontal: 16),
      child: Row(
        children: [
          const Icon(Icons.memory, color: AppColors.textSecondary),
          const SizedBox(width: 12),
          Text(
            parentPath,
            style: const TextStyle(
              color: AppColors.textSecondary,
              fontSize: 16,
              fontWeight: FontWeight.w500,
            ),
          ),
          Text(
            lastSegment,
            style: const TextStyle(
              color: AppColors.textSecondary,
              fontSize: 16,
              fontWeight: FontWeight.w700,
            ),
          ),
        ],
      ),
    );
  }
}
