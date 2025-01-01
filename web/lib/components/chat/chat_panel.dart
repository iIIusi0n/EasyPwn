import 'package:flutter/material.dart';
import '../../constants/colors.dart';
import 'chat_message.dart';
import 'chat_bubble.dart';

class ChatPanel extends StatelessWidget {
  final List<ChatMessage> chatMessages;
  final TextEditingController chatController;
  final ScrollController chatScrollController;
  final FocusNode chatFocusNode;
  final Function(String) onSubmit;

  const ChatPanel({
    super.key,
    required this.chatMessages,
    required this.chatController,
    required this.chatScrollController,
    required this.chatFocusNode,
    required this.onSubmit,
  });

  @override
  Widget build(BuildContext context) {
    return Container(
      decoration: const BoxDecoration(
        color: AppColors.surface,
        border: Border(
          left: BorderSide(color: AppColors.border, width: 1.5),
        ),
      ),
      child: Column(
        children: [
          // Chat header
          Container(
            height: 40,
            decoration: const BoxDecoration(
              border: Border(
                bottom: BorderSide(color: AppColors.border, width: 1.5),
              ),
            ),
            padding: const EdgeInsets.symmetric(horizontal: 12),
            child: const Row(
              children: [
                Icon(
                  Icons.chat_bubble_outline,
                  size: 18,
                  color: AppColors.textSecondary,
                ),
                SizedBox(width: 8),
                Text(
                  'Chat',
                  style: TextStyle(
                    color: AppColors.textSecondary,
                    fontWeight: FontWeight.w500,
                  ),
                ),
              ],
            ),
          ),

          // Chat messages
          Expanded(
            child: ListView.builder(
              controller: chatScrollController,
              padding: const EdgeInsets.all(16),
              itemCount: chatMessages.length,
              itemBuilder: (context, index) {
                return ChatBubble(
                  message: chatMessages[index],
                  omitSender: index != 0 && chatMessages[index].sender == chatMessages[index - 1].sender,
                );
              },
            ),
          ),

          // Chat input
          Container(
            decoration: const BoxDecoration(
              color: AppColors.surface,
              border: Border(
                top: BorderSide(color: AppColors.border, width: 1.5),
              ),
            ),
            padding: const EdgeInsets.all(8),
            child: Row(
              children: [
                Expanded(
                  child: TextField(
                    controller: chatController,
                    focusNode: chatFocusNode,
                    decoration: InputDecoration(
                      hintText: 'Type a message...',
                      hintStyle: TextStyle(
                        color: AppColors.greyShade(400),
                        fontSize: 14,
                      ),
                      border: InputBorder.none,
                      isDense: true,
                      contentPadding: const EdgeInsets.symmetric(
                        horizontal: 8,
                        vertical: 8,
                      ),
                    ),
                    style: const TextStyle(color: AppColors.textPrimary),
                    onSubmitted: (value) {
                      onSubmit(value);

                      chatFocusNode.requestFocus();
                    },
                  ),
                ),
                IconButton(
                  icon: const Icon(Icons.send),
                  onPressed: () {
                    if (chatController.text.isNotEmpty) {
                      onSubmit(chatController.text);
                    }
                  },
                  color: AppColors.textSecondary,
                  iconSize: 20,
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
} 