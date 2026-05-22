# System Architecture - Bit Battleship

## Overview

Bit Battleship is a real-time multiplayer backend system built using Go, WebSockets, PostgreSQL, and deployed on Koyeb.

---

## Architecture Diagram (Text-Based)

Client (Browser / Game UI)
        ↓ WebSocket / HTTP
Go Backend Server (Koyeb)
        ↓
Matchmaking Service
        ↓
Game Room Manager (In-memory state)
        ↓
PostgreSQL Database

---

## Core Components

### 1. Server Layer
Handles HTTP/WebSocket communication with players.

### 2. Matchmaking System
Pairs players into game rooms using worker pools.

### 3. Game Engine
Manages:
- ship placement
- attacks
- turns
- win detection

### 4. Store Layer
Manages active game rooms in memory.

### 5. Database Layer
Stores completed games and results.

---

## Deployment

- Hosted on Koyeb
- Auto-deploy from GitHub main branch
- Environment variables used for configuration

---

## Design Decisions

- In-memory rooms for speed
- Goroutines for concurrency
- Channels for event handling
- PostgreSQL for persistence
