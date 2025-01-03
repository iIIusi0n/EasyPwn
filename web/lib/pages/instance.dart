import 'package:flutter/material.dart';
import '../constants/colors.dart';
import '../components/top_bar.dart';
import '../components/side_bar.dart';

class InstancePage extends StatefulWidget {
  const InstancePage({super.key});

  @override
  State<InstancePage> createState() => _InstancePageState();
}

class _InstancePageState extends State<InstancePage> {
  String? selectedProjectId;
  List<Map<String, dynamic>> instances = [
    {
      'id': 'inst-1',
      'projectId': 'proj-1',
      'createdAt': DateTime.now().subtract(const Duration(hours: 2)),
      'status': 'Running',
      'memoryUsage': '124MB',
      'os': 'ubuntu-2410',
      'plugin': 'gef',
    },
    // Add more mock data as needed
  ];

  @override
  void initState() {
    super.initState();
    // TODO: Fetch projects and instances from backend
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.surface,
      body: Column(
        children: [
          const TopBar(path: 'user/instances'),

          // Main content
          Expanded(
            child: Row(
              children: [
                const SideBar(selectedIndex: 1),
                
                // Instance content
                Expanded(
                  child: Padding(
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
                            items: const [
                              DropdownMenuItem(
                                value: 'proj-1',
                                child: Text('Project 1'),
                              ),
                              // Add more projects
                            ],
                            onChanged: (value) {
                              setState(() {
                                selectedProjectId = value;
                              });
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
                                    dividerThickness: 1,
                                    border: TableBorder.all(
                                      color: Colors.grey.shade300,
                                      width: 1,
                                    ),
                                    columns: const [
                                      DataColumn(
                                        label: Text(
                                          'Created',
                                          style: TextStyle(fontWeight: FontWeight.bold),
                                        ),
                                      ),
                                      DataColumn(
                                        label: Text(
                                          'Status',
                                          style: TextStyle(fontWeight: FontWeight.bold),
                                        ),
                                      ),
                                      DataColumn(
                                        label: Text(
                                          'Memory',
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
                                          'Actions',
                                          style: TextStyle(fontWeight: FontWeight.bold),
                                        ),
                                      ),
                                    ],
                                    rows: instances
                                        .where((instance) =>
                                            selectedProjectId == null ||
                                            instance['projectId'] == selectedProjectId)
                                        .map((instance) {
                                      return DataRow(
                                        cells: [
                                          DataCell(Text(_formatDateTime(
                                              instance['createdAt'] as DateTime))),
                                          DataCell(_buildStatusCell(instance['status'])),
                                          DataCell(Text(instance['memoryUsage'])),
                                          DataCell(Text(instance['os'])),
                                          DataCell(Text(instance['plugin'])),
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
                ),
              ],
            ),
          ),
        ],
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
      case 'pending':
        statusColor = Colors.orange;
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

  Widget _buildActionButtons(Map<String, dynamic> instance) {
    final bool isRunning = instance['status'].toString().toLowerCase() == 'running';
    
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
              // TODO: Open instance
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
          onPressed: () {
            // TODO: Start/Restart instance
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
            onPressed: () {
              // TODO: Stop instance
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
                    onPressed: () {
                      // TODO: Delete instance
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
