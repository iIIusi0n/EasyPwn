import 'package:http/http.dart' as http;
import 'dart:convert';

class Project {
  final String id;
  final String name;
  final String userId;
  final String filePath;
  final String fileName;
  final String osName;
  final String pluginName;
  final String createdAt;

  Project({
    required this.id,
    required this.name,
    required this.userId,
    required this.filePath,
    required this.fileName,
    required this.osName,
    required this.pluginName,
    required this.createdAt,
  });

  factory Project.fromJson(Map<String, dynamic> json) {
    return Project(
      id: json['project_id'],
      name: json['name'],
      userId: json['user_id'],
      filePath: json['file_path'],
      fileName: json['file_name'],
      osName: json['os_name'],
      pluginName: json['plugin_name'],
      createdAt: json['created_at'],
    );
  }
}

class Os {
  final String id;
  final String name;

  Os({required this.id, required this.name});

  factory Os.fromJson(Map<String, dynamic> json) {
    return Os(
      id: json['id'],
      name: json['name'],
    );
  }
}

class Plugin {
  final String id;
  final String name;

  Plugin({required this.id, required this.name});

  factory Plugin.fromJson(Map<String, dynamic> json) {
    return Plugin(
      id: json['id'],
      name: json['name'],
    );
  }
}

class ProjectService {
  final String token;

  ProjectService({required this.token});

  Future<List<Os>> getOsList() async {
    final response = await http.get(
      Uri.parse('/api/project/os'),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );
    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      if (data == null) return [];
      return (data as List).map((os) => Os.fromJson(os)).toList();
    } else {
      throw Exception('Failed to get OS list');
    }
  }

  Future<List<Plugin>> getPluginList() async {
    final response = await http.get(
      Uri.parse('/api/project/plugin'),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );
    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      if (data == null) return [];
      return (data as List).map((plugin) => Plugin.fromJson(plugin)).toList();
    } else {
      throw Exception('Failed to get plugin list');
    }
  }

  Future<List<Project>> getProjects() async {
    final response = await http.get(
      Uri.parse('/api/project'),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );
    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      if (data == null) return [];
      return (data as List).map((project) => Project.fromJson(project)).toList();
    } else {
      throw Exception('Failed to get projects');
    }
  }

  Future<String> createProject(String projectName, String osId, String pluginId, List<int> fileBytes, String filename) async {
    var request = http.MultipartRequest('POST', Uri.parse('/api/project'));
    
    request.headers['Authorization'] = 'Bearer $token';
    
    request.fields['project_name'] = projectName;
    request.fields['os_id'] = osId;
    request.fields['plugin_id'] = pluginId;
    
    request.files.add(
      http.MultipartFile.fromBytes(
        'file',
        fileBytes,
        filename: filename
      )
    );

    final streamedResponse = await request.send();
    final response = await http.Response.fromStream(streamedResponse);

    if (response.statusCode != 200) {
      throw Exception('Failed to create project');
    }
    final data = jsonDecode(response.body);
    return data['project_id'];
  }

  Future<void> deleteProject(String projectId) async {
    final response = await http.delete(
      Uri.parse('/api/project/$projectId'),
      headers: {
        'Authorization': 'Bearer $token',
      },
    );
    if (response.statusCode != 200) {
      throw Exception('Failed to delete project');
    }
  }
}
