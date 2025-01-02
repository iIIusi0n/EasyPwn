import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:xterm/xterm.dart';
import '../constants/colors.dart';
import '../components/chat/chat_message.dart';
import '../components/chat/chat_panel.dart';
import '../components/top_bar.dart';
import '../components/side_bar.dart';
import '../components/bottom_bar.dart';
import '../components/custom_button.dart';
import '../services/terminal_service.dart';
import 'dart:convert';
import 'package:google_fonts/google_fonts.dart';

class SessionPage extends StatefulWidget {
  final String id;
  const SessionPage({super.key, required this.id});

  @override
  State<SessionPage> createState() => _SessionPageState();
}

class _SessionPageState extends State<SessionPage> with SingleTickerProviderStateMixin {
  bool isChatExpanded = true;
  bool isConnected = true;
  final TextEditingController _terminalController = TextEditingController();
  final TextEditingController _chatController = TextEditingController();
  final ScrollController _terminalScrollController = ScrollController();
  final ScrollController _chatScrollController = ScrollController();
  late TabController _tabController;

  final List<ChatMessage> chatMessages = [];
  final FocusNode _chatFocusNode = FocusNode();

  late Terminal terminal;
  late TerminalController terminalController;
  final TerminalService terminalService = TerminalService();

  final double minChatWidth = 300;
  final double maxChatWidth = 750;
  double chatWidth = 450;

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: 2, vsync: this);
    // Add some initial system message
    chatMessages.add(
      ChatMessage(
        sender: 'System',
        message: 'Connected to debug session. Type your messages here.',
        timestamp: DateTime.now(),
        isSystem: true,
      ),
    );
    _initializeTerminal();

    terminalService.connectionStatus.listen((status) {
      setState(() {
        isConnected = status;
      });
    });
  }

  void _initializeTerminal() {
    terminal = Terminal(
      maxLines: 10000,
    );
    
    terminalController = TerminalController();

    terminalService.connect('${Uri.base.scheme == 'https' ? 'wss' : 'ws'}://${Uri.base.host}:${Uri.base.port}/ws');
    
    terminalService.onData = (data) {
      if (data is List<int>) {
        terminal.write(const Utf8Decoder().convert(data));
      }
    };

    terminal.onOutput = (data) {
      final encodedData = const Utf8Encoder().convert(data);
      terminalService.sendCommand(encodedData);
    };

    // Handle terminal resize
    terminal.onResize = (w, h, pw, ph) {
      final resizeCommand = jsonEncode({
        'type': 'resize',
        'cols': w,
        'rows': h
      });
      terminalService.sendCommand(resizeCommand);
    };
  }

  @override
  void dispose() {
    _terminalController.dispose();
    _chatController.dispose();
    _terminalScrollController.dispose();
    _chatScrollController.dispose();
    _tabController.dispose();
    _chatFocusNode.dispose();
    terminalService.dispose();
    super.dispose();
  }

  void _handleChatSubmit(String value) {
    if (value.isEmpty) return;

    setState(() {
      chatMessages.add(ChatMessage(
        sender: 'You',
        message: value,
        timestamp: DateTime.now(),
      ));
    });
    _chatController.clear();
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
          // Top bar
          const TopBar(path: 'bifrost/uuid-23nif2u-34kb'),

          // Main content
          Expanded(
            child: Row(
              children: [
                // Sidebar
                const SideBar(selectedIndex: 1),

                // Terminal area with chat panel
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
                                      _buildDebuggerTab(),
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
                              chatWidth = (chatWidth - details.delta.dx).clamp(minChatWidth, maxChatWidth);
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
                                      side: BorderSide(color: AppColors.border, width: 1.5),
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
              ],
            ),
          ),

          // Bottom bar
          const BottomBar(
            instanceAddress: 'localhost:8080',
            memoryUsage: '124MB',
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
            terminal,
            controller: terminalController,
            autofocus: true,
            backgroundOpacity: 0.7,
            textStyle: TerminalStyle(
              fontSize: 14,
              fontFamily: GoogleFonts.robotoMono().fontFamily!,
            ),
            onSecondaryTapDown: (details, offset) async {
              final selection = terminalController.selection;
              if (selection != null) {
                final text = terminal.buffer.getText(selection);
                terminalController.clearSelection();
                await Clipboard.setData(ClipboardData(text: text));
              } else {
                final data = await Clipboard.getData('text/plain');
                final text = data?.text;
                if (text != null) {
                  terminal.paste(text);
                }
              }
            },
          ),
        ),
        if (!isConnected)
          Container(
            color: const Color(0xFF2C1F1F).withOpacity(0.7),
            child: Center(
              child: Container(
                padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 16),
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
                        terminalService.reconnect();
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
      ],
    );
  }
}
