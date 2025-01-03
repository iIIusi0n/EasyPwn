import 'package:go_router/go_router.dart';
import 'pages/login_page.dart';
import 'pages/register_page.dart';
import 'pages/session.dart';
import 'pages/home_page.dart';
import 'pages/instance.dart';
import 'pages/project.dart';
final router = GoRouter(
  initialLocation: '/',
  routes: [
    GoRoute(
      path: '/',
      builder: (context, state) => const HomePage(),
    ),
    GoRoute(
      path: '/login',
      builder: (context, state) => const LoginPage(),
    ),
    GoRoute(
      path: '/register',
      builder: (context, state) => const RegisterPage(),
    ),
    GoRoute(
      path: '/projects',
      builder: (context, state) => const ProjectPage(),
    ),
    GoRoute(
      path: '/instances',
      builder: (context, state) => const InstancePage(),
    ),
    GoRoute(
      path: '/session/:id',
      builder: (context, state) => SessionPage(id: state.pathParameters['id']!),
    ),
  ],
); 