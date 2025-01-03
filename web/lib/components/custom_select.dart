import 'package:flutter/material.dart';
import '../constants/colors.dart';

class CustomSelect extends StatelessWidget {
  final String value;
  final String label;
  final List<Map<String, String>> items;
  final Function(String?) onChanged;

  const CustomSelect({
    super.key,
    required this.value,
    required this.label,
    required this.items,
    required this.onChanged,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 400,
      decoration: BoxDecoration(
        color: AppColors.surface,
        borderRadius: BorderRadius.circular(4),
        border: Border.all(color: AppColors.border, width: 1.5),
      ),
      child: DropdownButtonHideUnderline(
        child: ButtonTheme(
          alignedDropdown: true,
          child: DropdownButton<String>(
            value: value,
            isExpanded: true,
            icon: const Icon(Icons.keyboard_arrow_down),
            iconSize: 24,
            elevation: 2,
            style: const TextStyle(color: AppColors.textPrimary),
            dropdownColor: AppColors.surface,
            padding: const EdgeInsets.symmetric(horizontal: 16),
            items: items.map<DropdownMenuItem<String>>((Map<String, String> item) {
              return DropdownMenuItem<String>(
                value: item['value'],
                child: Text(item['label'] ?? ''),
              );
            }).toList(),
            onChanged: onChanged,
          ),
        ),
      ),
    );
  }
} 