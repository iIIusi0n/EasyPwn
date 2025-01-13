import 'package:flutter/material.dart';
import '../constants/colors.dart';
import '../components/elements/custom_input.dart';
import '../components/elements/custom_button.dart';
import '../components/elements/custom_select.dart';
import '../components/elements/custom_file_picker.dart';
import 'package:file_picker/file_picker.dart';
import '../services/project_service.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:go_router/go_router.dart';
import '../services/instance_service.dart';

class ProjectPage extends StatefulWidget {
  const ProjectPage({super.key});

  @override
  State<ProjectPage> createState() => _ProjectPageState();
}

class _ProjectPageState extends State<ProjectPage> {
  final _nameController = TextEditingController();
  String? selectedOs;
  String? selectedPlugin;
  String? selectedFileName;
  PlatformFile? selectedFile;

  final _storage = const FlutterSecureStorage();
  late ProjectService _projectService;
  late InstanceService _instanceService;
  List<Project> projects = [];
  List<Os> osList = [];
  List<Plugin> pluginList = [];
  bool _isLoading = true;

  @override
  void initState() {
    super.initState();
    _initializeData();
  }

  Future<void> _initializeData() async {
    final token = await _storage.read(key: 'token');
    if (token == null) {
      if (mounted) context.go('/login');
      return;
    }

    _projectService = ProjectService(token: token);
    _instanceService = InstanceService(token: token);
    
    try {
      final futures = await Future.wait([
        _projectService.getProjects(),
        _projectService.getOsList(),
        _projectService.getPluginList(),
      ]);

      if (mounted) {
        setState(() {
          projects = futures[0] as List<Project>;
          osList = futures[1] as List<Os>;
          pluginList = futures[2] as List<Plugin>;
          _isLoading = false;
        });

        if (osList.isNotEmpty) {
          selectedOs = osList.first.id;
        }
        if (pluginList.isNotEmpty) {
          selectedPlugin = pluginList.first.id;
        }
      }
    } catch (e) {
      if (mounted) context.go('/login');
    }
  }

  @override
  void dispose() {
    _nameController.dispose();
    super.dispose();
  }

  Future<void> _handleCreateProject() async {
    if (_nameController.text.isEmpty ||
        selectedOs == null ||
        selectedPlugin == null ||
        selectedFile == null ||
        selectedFile!.bytes == null) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Please fill in all fields')),
      );
      return;
    }
    
    try {
      final _ = await _projectService.createProject(
        _nameController.text,
        selectedOs!,
        selectedPlugin!,
        selectedFile!.bytes!,
        selectedFile!.name,
      );
      
      final updatedProjects = await _projectService.getProjects();
      setState(() {
        projects = updatedProjects;
      });

      _nameController.clear();
      setState(() {
        selectedFileName = null;
        selectedFile = null;
      });
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Failed to create project')),
        );
      }
    }
  }

  Future<void> _handleDeleteProject(String projectId) async {
    try {
      await _projectService.deleteProject(projectId);
      final updatedProjects = await _projectService.getProjects();
      setState(() {
        projects = updatedProjects;
      });
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Failed to delete project')),
        );
      }
    }
  }

  String getOsLabel(String label) {
    switch (label) {
      case 'ubuntu-2410':
        return 'Ubuntu 24.10';
      default:
        return label;
    }
  }

  String getPluginLabel(String label) {
    switch (label) {
      case 'gef':
        return 'GEF';
      case 'pwndbg':
        return 'pwndbg';
      default:
        return label;
    }
  }

  @override
  Widget build(BuildContext context) {
    if (_isLoading) {
      return const Scaffold(
        body: Center(child: CircularProgressIndicator()),
      );
    }
    
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
                    items: osList.map((os) => {
                      'value': os.id,
                      'label': getOsLabel(os.name),
                    }).toList(),
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
                    items: pluginList.map((plugin) => {
                      'value': plugin.id,
                      'label': getPluginLabel(plugin.name),
                    }).toList(),
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
                    onPressed: _handleCreateProject,
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
                          borderRadius: BorderRadius.circular(4),
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
                              DataCell(Text(project.name)),
                              DataCell(Text(project.osName)),
                              DataCell(Text(project.pluginName)),
                              DataCell(Text(_formatDateTime(
                                  DateTime.parse(project.createdAt)))),
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

  Widget _buildActionButtons(Project project) {
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
          onPressed: () async {
            try {
              final _ = await _instanceService.createInstance(project.id);
              if (mounted) {
                context.go('/instances', extra: project.id);
              }
            } catch (e) {
              if (mounted) {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Failed to create instance')),
                );
              }
            }
          },
          child: const Text('New Instance'),
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
                      _handleDeleteProject(project.id);
                      Navigator.pop(context);
                    },
                    style: TextButton.styleFrom(foregroundColor: Colors.red),
                    child: const Text('Delete'),
                  ),
                ],
                shape: RoundedRectangleBorder(
                  borderRadius: BorderRadius.circular(4),
                ),
              ),
            );
          },
          child: const Text('Delete'),
        ),
      ],
    );
  }
}
