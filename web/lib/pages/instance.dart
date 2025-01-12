import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';
import '../constants/colors.dart';
import '../services/instance_service.dart';
import '../services/project_service.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class InstancePage extends StatefulWidget {
  const InstancePage({super.key});

  @override
  State<InstancePage> createState() => _InstancePageState();
}

class _InstancePageState extends State<InstancePage> {
  String? selectedProjectId;
  bool _sortAscending = true;
  int _sortColumnIndex = 0;

  List<Instance> instances = [];
  List<Project> projects = [];

  final _storage = const FlutterSecureStorage();
  late InstanceService _instanceService;
  late ProjectService _projectService;
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
      ]);

      if (mounted) {
        setState(() {
          projects = futures[0] as List<Project>;
          _isLoading = false;
        });
      }
    } catch (e) {
      if (mounted) context.go('/login');
    }
  }

  Future<void> _refreshInstances() async {
    if (selectedProjectId == null) return;
    
    try {
      final updatedInstances = await _instanceService.getInstances(selectedProjectId!);
      setState(() {
        instances = updatedInstances;
      });
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Failed to fetch instances')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.surface,
      body: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            // Project selector
            SizedBox(
              width: 300,
              child: DropdownButtonFormField<String>(
                value: selectedProjectId,
                decoration: const InputDecoration(
                  labelText: 'Select Project',
                  border: OutlineInputBorder(),
                ),
                items: projects.map((project) => DropdownMenuItem(
                  value: project.id,
                  child: Text(project.name),
                )).toList(),
                onChanged: (value) {
                  setState(() {
                    selectedProjectId = value;
                  });
                  _refreshInstances();
                },
              ),
            ),
            
            const SizedBox(height: 12),
            
            // Instances table
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
                        sortAscending: _sortAscending,
                        sortColumnIndex: _sortColumnIndex,
                        dividerThickness: 1,
                        border: TableBorder.all(
                          color: Colors.grey.shade300,
                          width: 1,
                          borderRadius: BorderRadius.circular(4),
                        ),
                        columns: [
                          DataColumn(
                            label: const Padding(
                              padding: EdgeInsets.only(right: 8.0),
                              child: Text(
                                'Created',
                                style: TextStyle(fontWeight: FontWeight.bold),
                              ),
                            ),
                            onSort: (columnIndex, ascending) {
                              setState(() {
                                _sortColumnIndex = columnIndex;
                                _sortAscending = ascending;
                                instances.sort((a, b) {
                                  final DateTime aDate = DateTime.parse(a.createdAt);
                                  final DateTime bDate = DateTime.parse(b.createdAt);
                                  return ascending
                                      ? aDate.compareTo(bDate)
                                      : bDate.compareTo(aDate);
                                });
                              });
                            },
                          ),
                          DataColumn(
                            label: const Padding(
                              padding: EdgeInsets.only(right: 8.0),
                              child: Text(
                                'Status',
                                style: TextStyle(fontWeight: FontWeight.bold),
                              ),
                            ),
                            onSort: (columnIndex, ascending) {
                              setState(() {
                                _sortColumnIndex = columnIndex;
                                _sortAscending = ascending;
                                instances.sort((a, b) {
                                  final String aStatus = a.status;
                                  final String bStatus = b.status;
                                  return ascending
                                      ? aStatus.compareTo(bStatus)
                                      : bStatus.compareTo(aStatus);
                                });
                              });
                            },
                          ),
                          const DataColumn(
                            label: Text(
                              'Memory',
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ),
                          ),
                          const DataColumn(
                            label: Text(
                              'Actions',
                              style: TextStyle(fontWeight: FontWeight.bold),
                            ),
                          ),
                        ],
                        rows: instances.map((instance) {
                          return DataRow(
                            cells: [
                              DataCell(Text(_formatDateTime(DateTime.parse(instance.createdAt)))),
                              DataCell(_buildStatusCell(instance.status)),
                              DataCell(Text('${instance.memory} MB')),
                              DataCell(_buildActionButtons(instance)),
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

  Widget _buildStatusCell(String status) {
    Color statusColor;
    switch (status.toLowerCase()) {
      case 'running':
        statusColor = Colors.green;
        break;
      case 'stopped':
        statusColor = Colors.red;
        break;
      default:
        statusColor = Colors.grey;
    }

    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        Container(
          width: 8,
          height: 8,
          decoration: BoxDecoration(
            color: statusColor,
            shape: BoxShape.circle,
          ),
        ),
        const SizedBox(width: 8),
        Text(status),
      ],
    );
  }

  Widget _buildActionButtons(Instance instance) {
    final bool isRunning = instance.status.toString().toLowerCase() == 'running';
    
    return Row(
      mainAxisSize: MainAxisSize.min,
      children: [
        if (isRunning) ...[
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
              context.go('/session/${instance.id}');
            },
            child: const Text('Open'),
          ),
          const SizedBox(width: 8),
        ],
        OutlinedButton(
          style: OutlinedButton.styleFrom(
            foregroundColor: Colors.green,
            side: const BorderSide(color: Colors.green),
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
            shape: RoundedRectangleBorder(
              borderRadius: BorderRadius.circular(4),
            ),
          ),
          onPressed: () async {
            try {
              await _instanceService.startInstance(instance.id);
              // Refresh instances after starting
              final updatedInstances = await _instanceService.getInstances(selectedProjectId!);
              setState(() {
                instances = updatedInstances;
              });
            } catch (e) {
              if (mounted) {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Failed to start instance')),
                );
              }
            }
          },
          child: Text(isRunning ? 'Restart' : 'Start'),
        ),
        if (isRunning) ...[
          const SizedBox(width: 8),
          OutlinedButton(
            style: OutlinedButton.styleFrom(
              foregroundColor: Colors.orange,
              side: const BorderSide(color: Colors.orange),
              padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(4),
              ),
            ),
            onPressed: () async {
              try {
                await _instanceService.stopInstance(instance.id);
                // Refresh instances after stopping
                final updatedInstances = await _instanceService.getInstances(selectedProjectId!);
                setState(() {
                  instances = updatedInstances;
                });
              } catch (e) {
                if (mounted) {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('Failed to stop instance')),
                  );
                }
              }
            },
            child: const Text('Stop'),
          ),
        ],
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
                title: const Text('Delete Instance'),
                content: const Text('Are you sure you want to delete this instance?'),
                actions: [
                  TextButton(
                    onPressed: () => Navigator.pop(context),
                    child: const Text('Cancel'),
                  ),
                  TextButton(
                    onPressed: () async {
                      Navigator.pop(context);
                      try {
                        await _instanceService.deleteInstance(instance.id);
                        // Refresh instances after deletion
                        final updatedInstances = await _instanceService.getInstances(selectedProjectId!);
                        setState(() {
                          instances = updatedInstances;
                        });
                      } catch (e) {
                        if (mounted) {
                          ScaffoldMessenger.of(context).showSnackBar(
                            const SnackBar(content: Text('Failed to delete instance')),
                          );
                        }
                      }
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
