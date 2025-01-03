import 'package:flutter/material.dart';
import '../constants/colors.dart';
import '../components/elements/custom_input.dart';
import '../components/elements/custom_button.dart';
import '../components/elements/custom_select.dart';
import '../components/elements/custom_file_picker.dart';
import 'package:file_picker/file_picker.dart';

class ProjectPage extends StatefulWidget {
  const ProjectPage({super.key});

  @override
  State<ProjectPage> createState() => _ProjectPageState();
}

class _ProjectPageState extends State<ProjectPage> {
  final _nameController = TextEditingController();
  String? selectedOs = 'ubuntu-2410';
  String? selectedPlugin = 'gef';
  String? selectedFileName;
  PlatformFile? selectedFile;

  List<Map<String, dynamic>> projects = [
    {
      'id': 'proj-1',
      'name': 'Project 1',
      'os': 'Ubuntu 24.10',
      'plugin': 'GEF',
      'createdAt': DateTime.now().subtract(const Duration(days: 2)),
    },
    // Add more mock data as needed
  ];

  @override
  void dispose() {
    _nameController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.surface,
      body: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Row(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Left side - New Project Form
            SizedBox(
              width: 300,
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text(
                    'New Project',
                    style: TextStyle(
                      fontSize: 20,
                      fontWeight: FontWeight.bold,
                      color: AppColors.textPrimary,
                    ),
                  ),
                  const SizedBox(height: 24),
                  CustomInput(
                    controller: _nameController,
                    hintText: 'Project Name',
                  ),
                  const SizedBox(height: 16),
                  CustomSelect(
                    value: selectedOs ?? '',
                    label: 'Operating System',
                    items: const [
                      {'value': 'ubuntu-2410', 'label': 'Ubuntu 24.10'},
                    ],
                    onChanged: (value) {
                      setState(() {
                        selectedOs = value;
                      });
                    },
                  ),
                  const SizedBox(height: 16),
                  CustomSelect(
                    value: selectedPlugin ?? '',
                    label: 'Debug Plugin',
                    items: const [
                      {'value': 'gef', 'label': 'GEF'},
                      {'value': 'pwndbg', 'label': 'pwndbg'},
                    ],
                    onChanged: (value) {
                      setState(() {
                        selectedPlugin = value;
                      });
                    },
                  ),
                  const SizedBox(height: 16),
                  CustomFilePicker(
                    selectedFileName: selectedFileName,
                    onFileSelected: (file) {
                      setState(() {
                        selectedFileName = file.name;
                        selectedFile = file;
                      });
                    },
                  ),
                  const SizedBox(height: 24),
                  CustomButton(
                    text: 'Create Project',
                    onPressed: () {
                      // TODO: Implement project creation
                    },
                    width: double.infinity,
                  ),
                ],
              ),
            ),
            const SizedBox(width: 48),
            // Right side - Projects Table
            Expanded(
              child: LayoutBuilder(
                builder: (context, constraints) {
                  return SingleChildScrollView(
                    scrollDirection: Axis.horizontal,
                    child: SizedBox(
                      width: constraints.maxWidth,
                      child: DataTable(
                        columnSpacing: 56.0,
                        horizontalMargin: 16.0,
                        dividerThickness: 1,
                        border: TableBorder.all(
                          color: Colors.grey.shade300,
                          width: 1,
                        ),
                        columns: const [
                          DataColumn(
                            label: Text(
                              'Project Name',
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ),
                          ),
                          DataColumn(
                            label: Text(
                              'OS',
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ),
                          ),
                          DataColumn(
                            label: Text(
                              'Plugin',
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ),
                          ),
                          DataColumn(
                            label: Text(
                              'Created',
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ),
                          ),
                          DataColumn(
                            label: Text(
                              'Actions',
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ),
                          ),
                        ],
                        rows: projects.map((project) {
                          return DataRow(
                            cells: [
                              DataCell(Text(project['name'])),
                              DataCell(Text(project['os'])),
                              DataCell(Text(project['plugin'])),
                              DataCell(Text(_formatDateTime(
                                  project['createdAt'] as DateTime))),
                              DataCell(_buildActionButtons(project)),
                            ],
                          );
                        }).toList(),
                      ),
                    ),
                  );
                },
              ),
            ),
          ],
        ),
      ),
    );
  }

  String _formatDateTime(DateTime dateTime) {
    final difference = DateTime.now().difference(dateTime);
    if (difference.inDays > 0) {
      return '${difference.inDays}d ago';
    } else if (difference.inHours > 0) {
      return '${difference.inHours}h ago';
    } else if (difference.inMinutes > 0) {
      return '${difference.inMinutes}m ago';
    } else {
      return 'Just now';
    }
  }

  Widget _buildActionButtons(Map<String, dynamic> project) {
    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        OutlinedButton(
          style: OutlinedButton.styleFrom(
            foregroundColor: Colors.blue,
            side: const BorderSide(color: Colors.blue),
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(4),
            ),
          ),
          onPressed: () {
            // TODO: Create new session
          },
          child: const Text('New Session'),
        ),
        const SizedBox(width: 8),
        OutlinedButton(
          style: OutlinedButton.styleFrom(
            foregroundColor: Colors.red,
            side: const BorderSide(color: Colors.red),
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(4),
            ),
          ),
          onPressed: () {
            showDialog(
              context: context,
              builder: (context) => AlertDialog(
                title: const Text('Delete Project'),
                content: const Text('Are you sure you want to delete this project?'),
                actions: [
                  TextButton(
                    onPressed: () => Navigator.pop(context),
                    child: const Text('Cancel'),
                  ),
                  TextButton(
                    onPressed: () {
                      // TODO: Delete project
                      Navigator.pop(context);
                    },
                    style: TextButton.styleFrom(foregroundColor: Colors.red),
                    child: const Text('Delete'),
                  ),
                ],
              ),
            );
          },
          child: const Text('Delete'),
        ),
      ],
    );
  }
}
