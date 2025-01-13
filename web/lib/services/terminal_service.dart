import 'dart:async';
import 'package:web_socket_channel/web_socket_channel.dart';

class TerminalService {
  late WebSocketChannel channel;
  Function(dynamic)? onData;
  bool isConnected = false;
  String? lastUrl;
  String? token;
  final StreamController<bool> _connectionStatusController = StreamController<bool>.broadcast();

  Stream<bool> get connectionStatus => _connectionStatusController.stream;

  void connect(String url, String token) {
    lastUrl = url;
    this.token = token;
    try {
      final Uri uri = Uri.parse(url);
      final Map<String, String> queryParams = Map.from(uri.queryParameters);
      queryParams['token'] = token;
      final authenticatedUri = uri.replace(queryParameters: queryParams);
      
      channel = WebSocketChannel.connect(authenticatedUri);
      isConnected = true;
      _connectionStatusController.add(true);

      channel.stream.listen(
        (message) {
          if (onData != null) {
            onData!(message);
          }
        },
        onDone: () {
          isConnected = false;
          _connectionStatusController.add(false);
        },
        onError: (error) {
          isConnected = false;
          _connectionStatusController.add(false);
        },
      );
    } catch (e) {
      isConnected = false;
      _connectionStatusController.add(false);
    }
  }

  void reconnect() {
    if (lastUrl != null) {
      connect(lastUrl!, token!);
    }
  }

  void sendCommand(dynamic command) {
    if (isConnected) {
      channel.sink.add(command);
    }
  }

  void dispose() {
    if (isConnected) {
      channel.sink.close();
    }
    _connectionStatusController.close();
  }
} 