import 'dart:async';
import 'package:web_socket_channel/web_socket_channel.dart';

class TerminalService {
  late WebSocketChannel channel;
  Function(dynamic)? onData;
  bool isConnected = false;
  String? lastUrl;
  final StreamController<bool> _connectionStatusController = StreamController<bool>.broadcast();

  Stream<bool> get connectionStatus => _connectionStatusController.stream;

  void connect(String url) {
    lastUrl = url;
    try {
      channel = WebSocketChannel.connect(Uri.parse(url));
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
      connect(lastUrl!);
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