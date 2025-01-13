class ChatMessage {
  final String sender;
  final String message;
  final DateTime timestamp;
  final bool isSystem;
  final bool isLoading;
  final bool isError;

  ChatMessage({
    required this.sender,
    required this.message,
    required this.timestamp,
    this.isSystem = false,
    this.isLoading = false,
    this.isError = false,
  });
} 