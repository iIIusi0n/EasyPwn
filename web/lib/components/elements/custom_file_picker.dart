import 'package:flutter/material.dart';
import 'package:file_picker/file_picker.dart';
import '../../constants/colors.dart';

class CustomFilePicker extends StatelessWidget {
  final String? selectedFileName;
  final Function(PlatformFile) onFileSelected;

  const CustomFilePicker({
    super.key,
    required this.selectedFileName,
    required this.onFileSelected,
  });

  Future<void> _pickFile() async {
    try {
      FilePickerResult? result = await FilePicker.platform.pickFiles(
        type: FileType.any,
        allowMultiple: false,
      );

      if (result != null && result.files.isNotEmpty) {
        onFileSelected(result.files.first);
      }
    } catch (e) {
      debugPrint('Error picking file: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 400,
      decoration: BoxDecoration(
        color: AppColors.surface,
        borderRadius: BorderRadius.circular(4),
        border: Border.all(color: AppColors.border, width: 1.5),
      ),
      child: Material(
        color: Colors.transparent,
        child: InkWell(
          onTap: _pickFile,
          borderRadius: BorderRadius.circular(4),
          child: Padding(
            padding: const EdgeInsets.symmetric(
              horizontal: 16,
              vertical: 18,
            ),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.start,
              children: [
                const Icon(Icons.upload_file, 
                  color: AppColors.textSecondary,
                  size: 18,
                ),
                const SizedBox(width: 12),
                Expanded(
                  child: Text(
                    selectedFileName ?? 'Choose File',
                    style: TextStyle(
                      color: selectedFileName != null 
                          ? AppColors.textPrimary 
                          : AppColors.greyShade(400),
                      fontSize: 14,
                    ),
                    overflow: TextOverflow.ellipsis,
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
} 