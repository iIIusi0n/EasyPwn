import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:xterm/xterm.dart';
import '../constants/colors.dart';
import '../components/chat/chat_message.dart';
import '../components/chat/chat_panel.dart';
import '../components/bottom_bar.dart';
import '../components/elements/custom_button.dart';
import '../services/terminal_service.dart';
import 'dart:convert';
import 'package:google_fonts/google_fonts.dart';
import 'package:go_router/go_router.dart';
import 'dart:async';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import '../services/instance_service.dart';
import '../services/chatbot_service.dart';

class SessionPage extends StatefulWidget {
  final String id;
  const SessionPage({super.key, required this.id});

  @override
  State<SessionPage> createState() => _SessionPageState();
}

class _SessionPageState extends State<SessionPage>
    with SingleTickerProviderStateMixin {
  final FlutterSecureStorage _storage = const FlutterSecureStorage();
  bool isChatExpanded = true;
  bool isDebugConnected = true;
  bool isShellConnected = true;
  final TextEditingController _terminalController = TextEditingController();
  final TextEditingController _chatController = TextEditingController();
  final ScrollController _terminalScrollController = ScrollController();
  final ScrollController _chatScrollController = ScrollController();
  late TabController _tabController;

  final List<ChatMessage> chatMessages = [];
  final FocusNode _chatFocusNode = FocusNode();

  late Terminal debugTerminal;
  late TerminalController debugTerminalController;
  late Terminal shellTerminal;
  late TerminalController shellTerminalController;
  final TerminalService debugTerminalService = TerminalService();
  final TerminalService shellTerminalService = TerminalService();
  late InstanceService _instanceService;
  late ChatbotService _chatbotService;
  final double minChatWidth = 300;
  final double maxChatWidth = 750;
  double chatWidth = 450;

  late StreamSubscription _debugConnectionSubscription;
  late StreamSubscription _shellConnectionSubscription;

  String memoryUsage = '...';

  Future<void> _updateMemoryUsage() async {
    final instance = await _instanceService.getInstance(widget.id);
    if (mounted) {
      setState(() {
        memoryUsage = instance.memory.toString();
      });
    }
  }

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: 2, vsync: this);
    // Add some initial system message
    chatMessages.add(
      ChatMessage(
        sender: 'System',
        message: 'Connected to AI pwnable assistant. Type your messages here.',
        timestamp: DateTime.now(),
        isSystem: true,
      ),
    );
    _initializeTerminal();

    _debugConnectionSubscription =
        debugTerminalService.connectionStatus.listen((status) {
      if (mounted) {
        setState(() {
          isDebugConnected = status;
        });
      }
    });

    _shellConnectionSubscription =
        shellTerminalService.connectionStatus.listen((status) {
      if (mounted) {
        setState(() {
          isShellConnected = status;
        });
      }
    });

    _updateMemoryUsage();
  }

  Future<void> _initializeTerminal() async {
    final token = await _storage.read(key: 'token');
    debugTerminal = Terminal(
      maxLines: 10000,
    );
    shellTerminal = Terminal(
      maxLines: 10000,
    );

    debugTerminalController = TerminalController();
    shellTerminalController = TerminalController();

    _instanceService = InstanceService(token: token!);
    _chatbotService = ChatbotService(token: token);
    debugTerminalService.connect(
        '${Uri.base.scheme == 'https' ? 'wss' : 'ws'}://${Uri.base.host}:${Uri.base.port}/api/stream/session/debugger/${widget.id}',
        token);
    shellTerminalService.connect(
        '${Uri.base.scheme == 'https' ? 'wss' : 'ws'}://${Uri.base.host}:${Uri.base.port}/api/stream/session/shell/${widget.id}',
        token);

    debugTerminalService.onData = (data) {
      if (data is List<int>) {
        debugTerminal.write(const Utf8Decoder().convert(data));
      }
    };

    shellTerminalService.onData = (data) {
      if (data is List<int>) {
        shellTerminal.write(const Utf8Decoder().convert(data));
      }
    };

    debugTerminal.onOutput = (data) {
      final encodedData = const Utf8Encoder().convert(data);
      debugTerminalService.sendCommand(encodedData);
    };

    shellTerminal.onOutput = (data) {
      final encodedData = const Utf8Encoder().convert(data);
      shellTerminalService.sendCommand(encodedData);
    };

    // Handle terminal resize
    debugTerminal.onResize = (w, h, pw, ph) {
      final resizeCommand =
          jsonEncode({'type': 'resize', 'cols': w, 'rows': h});
      debugTerminalService.sendCommand(resizeCommand);
    };

    shellTerminal.onResize = (w, h, pw, ph) {
      final resizeCommand =
          jsonEncode({'type': 'resize', 'cols': w, 'rows': h});
      shellTerminalService.sendCommand(resizeCommand);
    };
  }

  @override
  void dispose() {
    _debugConnectionSubscription.cancel();
    _shellConnectionSubscription.cancel();
    _terminalController.dispose();
    _chatController.dispose();
    _terminalScrollController.dispose();
    _chatScrollController.dispose();
    _tabController.dispose();
    _chatFocusNode.dispose();
    debugTerminalService.dispose();
    shellTerminalService.dispose();
    super.dispose();
  }

  void _handleChatSubmit(String value) async {
    if (value.isEmpty) return;

    // Add user message
    setState(() {
      chatMessages.add(ChatMessage(
        sender: 'You',
        message: value,
        timestamp: DateTime.now(),
      ));
    });
    _chatController.clear();
    _scrollToBottom(_chatScrollController);

    try {
      // Add loading message
      setState(() {
        chatMessages.add(ChatMessage(
          sender: 'Assistant',
          message: '...',
          timestamp: DateTime.now(),
          isLoading: true,
        ));
      });
      _scrollToBottom(_chatScrollController);

      // Get response from API
      final response = await _chatbotService.getResponse(widget.id, value);

      // Replace loading message with actual response
      setState(() {
        chatMessages.removeLast();
        chatMessages.add(ChatMessage(
          sender: 'Assistant',
          message: response,
          timestamp: DateTime.now(),
        ));
      });
    } catch (e) {
      // Replace loading message with error
      setState(() {
        chatMessages.removeLast();
        chatMessages.add(ChatMessage(
          sender: 'System',
          message: 'Failed to get response: $e',
          timestamp: DateTime.now(),
          isSystem: true,
          isError: true,
        ));
      });
    }
    _scrollToBottom(_chatScrollController);
  }

  void _scrollToBottom(ScrollController controller) {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      controller.animateTo(
        controller.position.maxScrollExtent,
        duration: const Duration(milliseconds: 200),
        curve: Curves.easeOut,
      );
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.surface,
      body: Column(
        children: [
          // Main content
          Expanded(
            child: Stack(
              clipBehavior: Clip.none,
              children: [
                Row(
                  children: [
                    // Terminal area
                    Expanded(
                      child: Column(
                        children: [
                          TabBar(
                            controller: _tabController,
                            isScrollable: true,
                            tabAlignment: TabAlignment.start,
                            labelColor: AppColors.textPrimary,
                            unselectedLabelColor: AppColors.greyShade(600),
                            indicatorColor: AppColors.textPrimary,
                            indicatorWeight: 2,
                            padding: const EdgeInsets.symmetric(horizontal: 16),
                            tabs: const [
                              Tab(text: 'Debugger'),
                              Tab(text: 'Shell'),
                            ],
                          ),
                          Expanded(
                            child: TabBarView(
                              controller: _tabController,
                              children: [
                                _buildDebuggerTab(),
                                _buildShellTab(),
                              ],
                            ),
                          ),
                        ],
                      ),
                    ),

                    // Chat panel
                    AnimatedContainer(
                      duration: const Duration(milliseconds: 100),
                      curve: Curves.easeInOut,
                      width: isChatExpanded ? chatWidth : 0,
                      child: SingleChildScrollView(
                        scrollDirection: Axis.horizontal,
                        physics: const NeverScrollableScrollPhysics(),
                        child: SizedBox(
                          width: chatWidth,
                          child: ChatPanel(
                            chatMessages: chatMessages,
                            chatController: _chatController,
                            chatScrollController: _chatScrollController,
                            chatFocusNode: _chatFocusNode,
                            onSubmit: _handleChatSubmit,
                          ),
                        ),
                      ),
                    ),
                  ],
                ),

                // Chat toggle button
                AnimatedPositioned(
                  duration: const Duration(milliseconds: 100),
                  curve: Curves.easeInOut,
                  right: isChatExpanded ? chatWidth : 0,
                  top: MediaQuery.of(context).size.height / 2 - 60,
                  child: GestureDetector(
                    onHorizontalDragUpdate: (details) {
                      if (!isChatExpanded) return;
                      setState(() {
                        chatWidth = (chatWidth - details.delta.dx)
                            .clamp(minChatWidth, maxChatWidth);
                      });
                    },
                    child: MouseRegion(
                      cursor: SystemMouseCursors.resizeLeftRight,
                      child: Material(
                        elevation: 4,
                        color: AppColors.surface,
                        borderRadius: const BorderRadius.only(
                          topLeft: Radius.circular(24),
                          bottomLeft: Radius.circular(24),
                          topRight: Radius.circular(0),
                          bottomRight: Radius.circular(0),
                        ),
                        child: InkWell(
                          borderRadius: const BorderRadius.only(
                            topLeft: Radius.circular(24),
                            bottomLeft: Radius.circular(24),
                            topRight: Radius.circular(0),
                            bottomRight: Radius.circular(0),
                          ),
                          onTap: () {
                            setState(() {
                              isChatExpanded = !isChatExpanded;
                            });
                          },
                          child: Container(
                            width: 22,
                            height: 90,
                            decoration: const ShapeDecoration(
                              color: AppColors.surface,
                              shape: ContinuousRectangleBorder(
                                side: BorderSide(
                                    color: AppColors.border, width: 1.5),
                                borderRadius: BorderRadius.horizontal(
                                  left: Radius.elliptical(24, 48),
                                ),
                              ),
                            ),
                            child: Column(
                              children: [
                                Expanded(
                                  child: Center(
                                    child: Column(
                                      mainAxisSize: MainAxisSize.min,
                                      children: [
                                        Icon(
                                          isChatExpanded
                                              ? Icons.chevron_right
                                              : Icons.chevron_left,
                                          size: 16,
                                          color: AppColors.textSecondary,
                                        ),
                                      ],
                                    ),
                                  ),
                                ),
                              ],
                            ),
                          ),
                        ),
                      ),
                    ),
                  ),
                ),
              ],
            ),
          ),

          // Bottom bar
          BottomBar(
            instanceId: widget.id,
            memoryUsage: "$memoryUsage MB",
          ),
        ],
      ),
    );
  }

  Widget _buildDebuggerTab() {
    return Stack(
      children: [
        Container(
          color: AppColors.surfaceDark,
          child: TerminalView(
            debugTerminal,
            controller: debugTerminalController,
            autofocus: true,
            backgroundOpacity: 0.7,
            textStyle: TerminalStyle(
              fontSize: 14,
              fontFamily: GoogleFonts.robotoMono().fontFamily!,
            ),
            onSecondaryTapDown: (details, offset) async {
              final selection = debugTerminalController.selection;
              if (selection != null) {
                final text = debugTerminal.buffer.getText(selection);
                debugTerminalController.clearSelection();
                await Clipboard.setData(ClipboardData(text: text));
              } else {
                final data = await Clipboard.getData('text/plain');
                final text = data?.text;
                if (text != null) {
                  debugTerminal.paste(text);
                }
              }
            },
          ),
        ),
        if (!isDebugConnected)
          Container(
            color: const Color(0xFF2C1F1F).withOpacity(0.7),
            child: Center(
              child: Container(
                padding:
                    const EdgeInsets.symmetric(horizontal: 20, vertical: 16),
                decoration: BoxDecoration(
                  color: const Color(0xFF1E1E1E).withOpacity(0.95),
                  borderRadius: BorderRadius.circular(8),
                  border: Border.all(
                    color: const Color(0xFF433333),
                    width: 1,
                  ),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.black.withOpacity(0.2),
                      blurRadius: 12,
                      offset: const Offset(0, 4),
                    ),
                  ],
                ),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Container(
                          width: 8,
                          height: 8,
                          decoration: BoxDecoration(
                            color: Colors.red.shade400,
                            borderRadius: BorderRadius.circular(4),
                          ),
                        ),
                        const SizedBox(width: 8),
                        const Text(
                          'Connection lost',
                          style: TextStyle(
                            color: Colors.white70,
                            fontSize: 14,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                      ],
                    ),
                    const SizedBox(height: 12),
                    CustomButton(
                      text: 'Reconnect',
                      width: 100,
                      height: 30,
                      onPressed: () {
                        debugTerminalService.reconnect();
                      },
                      backgroundColor: const Color.fromARGB(255, 196, 73, 73),
                      borderColor: const Color.fromARGB(255, 196, 73, 73),
                      textColor: Colors.white70,
                    ),
                  ],
                ),
              ),
            ),
          ),
        Positioned(
          right: 16,
          bottom: 16,
          child: Container(
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              color: AppColors.surface,
              border: Border.all(color: AppColors.border),
            ),
            child: IconButton(
              icon: const Icon(Icons.logout),
              iconSize: 20,
              color: AppColors.textSecondary,
              onPressed: () {
                context.go('/instances');
              },
            ),
          ),
        ),
      ],
    );
  }

  Widget _buildShellTab() {
    return Stack(
      children: [
        Container(
          color: AppColors.surfaceDark,
          child: TerminalView(
            shellTerminal,
            controller: shellTerminalController,
            autofocus: true,
            backgroundOpacity: 0.7,
            textStyle: TerminalStyle(
              fontSize: 14,
              fontFamily: GoogleFonts.robotoMono().fontFamily!,
            ),
            onSecondaryTapDown: (details, offset) async {
              final selection = shellTerminalController.selection;
              if (selection != null) {
                final text = shellTerminal.buffer.getText(selection);
                shellTerminalController.clearSelection();
                await Clipboard.setData(ClipboardData(text: text));
              } else {
                final data = await Clipboard.getData('text/plain');
                final text = data?.text;
                if (text != null) {
                  shellTerminal.paste(text);
                }
              }
            },
          ),
        ),
        if (!isShellConnected)
          Container(
            color: const Color(0xFF2C1F1F).withOpacity(0.7),
            child: Center(
              child: Container(
                padding:
                    const EdgeInsets.symmetric(horizontal: 20, vertical: 16),
                decoration: BoxDecoration(
                  color: const Color(0xFF1E1E1E).withOpacity(0.95),
                  borderRadius: BorderRadius.circular(8),
                  border: Border.all(
                    color: const Color(0xFF433333),
                    width: 1,
                  ),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.black.withOpacity(0.2),
                      blurRadius: 12,
                      offset: const Offset(0, 4),
                    ),
                  ],
                ),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Row(
                      mainAxisSize: MainAxisSize.min,
                      children: [
                        Container(
                          width: 8,
                          height: 8,
                          decoration: BoxDecoration(
                            color: Colors.red.shade400,
                            borderRadius: BorderRadius.circular(4),
                          ),
                        ),
                        const SizedBox(width: 8),
                        const Text(
                          'Connection lost',
                          style: TextStyle(
                            color: Colors.white70,
                            fontSize: 14,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                      ],
                    ),
                    const SizedBox(height: 12),
                    CustomButton(
                      text: 'Reconnect',
                      width: 100,
                      height: 30,
                      onPressed: () {
                        shellTerminalService.reconnect();
                      },
                      backgroundColor: const Color.fromARGB(255, 196, 73, 73),
                      borderColor: const Color.fromARGB(255, 196, 73, 73),
                      textColor: Colors.white70,
                    ),
                  ],
                ),
              ),
            ),
          ),
        Positioned(
          right: 16,
          bottom: 16,
          child: Container(
            decoration: BoxDecoration(
              shape: BoxShape.circle,
              color: AppColors.surface,
              border: Border.all(color: AppColors.border),
            ),
            child: IconButton(
              icon: const Icon(Icons.logout),
              iconSize: 20,
              color: AppColors.textSecondary,
              onPressed: () {
                context.go('/instances');
              },
            ),
          ),
        ),
      ],
    );
  }
}
