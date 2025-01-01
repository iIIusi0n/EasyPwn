import 'package:flutter/material.dart';
import '../../constants/colors.dart';
import 'chat_message.dart';

class ChatBubble extends StatelessWidget {
  final ChatMessage message;
  final bool omitSender;

  const ChatBubble({
    super.key,
    required this.message,
    this.omitSender = false,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(top: omitSender ? 0 : 12, bottom: 4),
      child: Column(
        crossAxisAlignment: message.sender == 'You' ? CrossAxisAlignment.end : CrossAxisAlignment.start,
        children: [
          if (!omitSender)
            Text(
              message.sender,
              style: TextStyle(
                color: AppColors.greyShade(600),
              fontSize: 12,
              fontWeight: FontWeight.w500,
              ),
            ),
          Container(
            margin: const EdgeInsets.only(top: 6),
            padding: const EdgeInsets.symmetric(
              horizontal: 12,
              vertical: 8,
            ),
            decoration: BoxDecoration(
              color: AppColors.surface,
              borderRadius: BorderRadius.circular(4),
              border: Border.all(
                color: AppColors.border,
                width: 1.5,
              ),
            ),
            child: Text(
              message.message,
              style: const TextStyle(color: AppColors.textPrimary),
            ),
          ),
        ],
      ),
    );
  }
} 