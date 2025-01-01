import 'package:flutter/material.dart';
import '../constants/colors.dart';
import 'sidebar_item.dart';

class SideBar extends StatelessWidget {
  const SideBar({super.key});

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 200,
      decoration: const BoxDecoration(
        color: AppColors.surface,
        border: Border(
          right: BorderSide(color: AppColors.border, width: 1.5),
        ),
      ),
      child: Column(
        children: [
          SidebarItem(
            icon: Icons.terminal,
            label: 'Debug',
            onTap: () {},
            isSelected: true,
          ),
          SidebarItem(
            icon: Icons.memory,
            label: 'Memory',
            onTap: () {},
          ),
          SidebarItem(
            icon: Icons.settings,
            label: 'Settings',
            onTap: () {},
          ),
        ],
      ),
    );
  }
} 