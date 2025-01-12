import 'package:http/http.dart' as http;

class Instance {
  final String id;
  final String status;
  final int memory;
  final String createdAt;
  final String updatedAt;

  Instance({required this.id, required this.status, required this.memory, required this.createdAt, required this.updatedAt});
}

class InstanceService {
  final String token;

  InstanceService({required this.token});

  Future<List<Instance>> getInstances(String projectId) async {
    final response = await http.get(
      Uri.parse('/api/instance?project_id=$projectId'),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );
    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      if (data == null) return [];
      return (data as List).map((instance) => Instance.fromJson(instance)).toList();
    } else {
      throw Exception('Failed to get instances');
    }
  }

  Future<Instance> createInstance(String projectId) async {
    final response = await http.post(
      Uri.parse('/api/instance?project_id=$projectId'),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );
    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return Instance.fromJson(data);
    } else {
      throw Exception('Failed to create instance');
    }
  }

  Future<void> deleteInstance(String instanceId) async {
    final response = await http.delete(
      Uri.parse('/api/instance/$instanceId'),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );
    if (response.statusCode == 200) {
      return;
    } else {
      throw Exception('Failed to delete instance');
    }
  }
}