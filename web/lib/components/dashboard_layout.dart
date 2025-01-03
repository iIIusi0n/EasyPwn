import 'package:flutter/material.dart';
import 'top_bar.dart';
import 'side_bar.dart';
import '../constants/colors.dart';

class DashboardLayout extends StatelessWidget {
  final String path;
  final int selectedIndex;
  final Widget child;

  const DashboardLayout({
    super.key,
    required this.path,
    required this.selectedIndex,
    required this.child,
  });

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.surface,
      body: Column(
        children: [
          TopBar(path: path),
          Expanded(
            child: Row(
              children: [
                SideBar(selectedIndex: selectedIndex),
                Expanded(child: child),
              ],
            ),
          ),
        ],
      ),
    );
  }
} 