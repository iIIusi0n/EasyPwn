import 'package:go_router/go_router.dart';
import 'pages/login_page.dart';
import 'pages/register_page.dart';
import 'pages/instance.dart';
import 'pages/instances_page.dart';
import 'pages/home_page.dart';

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
      path: '/i',
      builder: (context, state) => const InstancesPage(),
    ),
    GoRoute(
      path: '/i/:id',
      builder: (context, state) => InstancePage(id: state.pathParameters['id']!),
    ),
  ],
); 