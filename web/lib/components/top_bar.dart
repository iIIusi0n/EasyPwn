import 'package:flutter/material.dart';
import '../constants/colors.dart';
import 'custom_button.dart';
import 'status_badge.dart';

class TopBar extends StatelessWidget {
  final String instanceName;
  final String status;

  const TopBar({
    super.key,
    required this.instanceName,
    required this.status,
  });

  @override
  Widget build(BuildContext context) {
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
            'Instance: $instanceName',
            style: const TextStyle(
              color: AppColors.textPrimary,
              fontSize: 16,
              fontWeight: FontWeight.w500,
            ),
          ),
          const SizedBox(width: 24),
          StatusBadge(
            label: status,
            color: status == 'Running' ? Colors.green : Colors.red,
          ),
          const Spacer(),
          CustomButton(
            text: status == 'Running' ? 'Stop' : 'Start',
            onPressed: () {},
            width: 80,
            height: 30,
            backgroundColor: status == 'Running' ? Colors.red.shade400 : Colors.green.shade400,
            borderColor: status == 'Running' ? Colors.red.shade400 : Colors.green.shade400,
            textColor: Colors.white,
          ),
          const SizedBox(width: 8),
          CustomButton(
            text: 'Restart',
            onPressed: () {},
            width: 80,
            height: 30,
          ),
        ],
      ),
    );
  }
}
