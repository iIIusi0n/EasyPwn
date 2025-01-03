import 'package:flutter/material.dart';
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
      pageBuilder: (context, state) => const NoTransitionPage(
        child: const HomePage(),
      ),
    ),
    GoRoute(
      path: '/login',
      pageBuilder: (context, state) => const NoTransitionPage(
        child: LoginPage(),
      ),
    ),
    GoRoute(
      path: '/register',
      pageBuilder: (context, state) => const NoTransitionPage(
        child: RegisterPage(),
      ),
    ),
    GoRoute(
      path: '/projects',
      pageBuilder: (context, state) => const NoTransitionPage(
        child: ProjectPage(),
      ),
    ),
    GoRoute(
      path: '/instances',
      pageBuilder: (context, state) => const NoTransitionPage(
        child: InstancePage(),
      ),
    ),
    GoRoute(
      path: '/session/:id',
      pageBuilder: (context, state) => NoTransitionPage(
        child: SessionPage(id: state.pathParameters['id']!),
      ),
    ),
  ],
); 