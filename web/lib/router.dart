import 'package:go_router/go_router.dart';
import 'pages/login_page.dart';
import 'pages/register_page.dart';
import 'pages/session.dart';
import 'pages/home_page.dart';
import 'pages/instance.dart';
import 'pages/project.dart';
import 'components/dashboard_layout.dart';
import 'router_helpers.dart';

final router = GoRouter(
  initialLocation: '/',
  routes: [
    // Auth routes (no shell)
    GoRoute(
      path: '/login',
      redirect: redirectIfAuthenticated,
      pageBuilder: (context, state) => const NoTransitionPage(
        child: LoginPage(),
      ),
    ),
    GoRoute(
      path: '/register',
      redirect: redirectIfAuthenticated,
      pageBuilder: (context, state) => const NoTransitionPage(
        child: RegisterPage(),
      ),
    ),
    // Home route (no shell)
    GoRoute(
      path: '/',
      builder: (context, state) => const HomePage(),
    ),
    
    // Main app routes (with shell)
    ShellRoute(
      builder: (context, state, child) {
        final path = state.uri.path;
        final selectedIndex = switch (path) {
          '/projects' => 0,
          '/instances' => 1,
          String s when s.startsWith('/session/') => 1,
          _ => -1,
        };
        
        return DashboardLayout(
          path: path,
          selectedIndex: selectedIndex,
          child: child,
        );
      },
      routes: [
        GoRoute(
          path: '/projects',
          builder: (context, state) => const ProjectPage(),
        ),
        GoRoute(
          path: '/instances',
          builder: (context, state) => InstancePage(
            initialProjectId: state.extra as String?,
          ),
        ),
        GoRoute(
          path: '/session/:id',
          builder: (context, state) => SessionPage(
            id: state.pathParameters['id'] ?? '',
          ),
        ),
      ],
    ),
  ],
); 