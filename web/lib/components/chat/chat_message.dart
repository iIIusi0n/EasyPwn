class ChatMessage {
  final String sender;
  final String message;
  final DateTime timestamp;
  final bool isSystem;

  ChatMessage({
    required this.sender,
    required this.message,
    required this.timestamp,
    this.isSystem = false,
  });
} 