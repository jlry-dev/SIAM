# Product Backlog - Bit Battleship

## Project Description
Bit Battleship is a multiplayer online Battleship game where two players can create or join rooms, arrange ships, and battle each other in real time.

---

# User Stories

## 1. Create Game Room
- Priority: High
- Story Points: 3

### User Story
As a player, I want to create a game room so that I can invite another player.

### Acceptance Criteria
- Player can create a room
- System generates a room code
- Room waits for another player

---

## 2. Join Game Room
- Priority: High
- Story Points: 3

### User Story
As a player, I want to join a room using a room code so that I can play with friends.

### Acceptance Criteria
- Player can enter room code
- Invalid room codes are rejected
- Player joins successfully

---

## 3. Arrange Ships
- Priority: High
- Story Points: 5

### User Story
As a player, I want to place ships on the board before the game starts.

### Acceptance Criteria
- Ships can be placed on the grid
- Ships cannot overlap
- Ships cannot go outside the board

---

## 4. Start Match
- Priority: High
- Story Points: 2

### User Story
As a player, I want the match to begin once both players are ready.

### Acceptance Criteria
- Match starts automatically
- Both players receive game state updates

---

## 5. Turn-Based Gameplay
- Priority: High
- Story Points: 3

### User Story
As a player, I want turns to alternate properly so that gameplay is fair.

### Acceptance Criteria
- Only one player attacks per turn
- Turn changes after attack

---

## 6. Attack Enemy Grid
- Priority: High
- Story Points: 5

### User Story
As a player, I want to attack enemy coordinates so that I can destroy ships.

### Acceptance Criteria
- Player can select coordinates
- Hit or miss is displayed
- Coordinates cannot be reused

---

## 7. Display Hit and Miss Effects
- Priority: Medium
- Story Points: 2

### User Story
As a player, I want visual feedback for hits and misses so that I understand the game state.

### Acceptance Criteria
- Hits display a hit marker
- Misses display a miss marker

---

## 8. Win Detection
- Priority: Medium
- Story Points: 3

### User Story
As a player, I want the game to detect a winner automatically.

### Acceptance Criteria
- System detects when all ships are destroyed
- Winner message is displayed

---

## 9. Restart Game
- Priority: Low
- Story Points: 2

### User Story
As a player, I want to restart the game after finishing a match.

### Acceptance Criteria
- Players can start a new match
- Boards reset properly

---

## 10. Leave Room
- Priority: Low
- Story Points: 1

### User Story
As a player, I want to leave a room safely.

### Acceptance Criteria
- Player can exit room
- Opponent is notified
