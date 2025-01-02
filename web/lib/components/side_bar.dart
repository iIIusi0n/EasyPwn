import 'package:flutter/material.dart';
import '../constants/colors.dart';
import 'sidebar_item.dart';

class SideBar extends StatelessWidget {
  final int selectedIndex;

  const SideBar({super.key, required this.selectedIndex});

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
            icon: Icons.folder,
            label: 'Project',
            onTap: () {},
            isSelected: selectedIndex == 0,
          ),
          SidebarItem(
            icon: Icons.terminal,
            label: 'Instance',
            onTap: () {},
            isSelected: selectedIndex == 1,
          ),
          SidebarItem(
            icon: Icons.settings,
            label: 'Setting',
            onTap: () {},
            isSelected: selectedIndex == 2,
          ),
          SidebarItem(
            icon: Icons.logout,
            label: 'Logout',
            onTap: () {},
          ),
        ],
      ),
    );
  }
} 