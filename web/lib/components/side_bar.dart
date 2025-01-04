import 'package:flutter/material.dart';
import '../constants/colors.dart';
import 'sidebar_item.dart';
import 'package:go_router/go_router.dart';

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
            label: 'Projects',
            onTap: () => context.go('/projects'),
            isSelected: selectedIndex == 0,
          ),
          SidebarItem(
            icon: Icons.terminal,
            label: 'Instances',
            onTap: () => context.go('/instances'),
            isSelected: selectedIndex == 1,
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