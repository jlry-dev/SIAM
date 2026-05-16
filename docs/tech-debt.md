# Technical Debt - Bit Battleship

## Definition
Technical debt refers to temporary or inefficient solutions in the codebase that may require future improvement.

---

| ID | Technical Debt | Impact | Priority | Planned Fix |
|----|----------------|--------|----------|-------------|
| TD-001 | Manual .env parsing in main.go | Harder maintenance | High | Replace with configuration utility |
| TD-002 | Limited unit test coverage | Reduced reliability | High | Add more tests |
| TD-003 | Matchmaking logic tightly coupled | Harder scalability | Medium | Separate matchmaking services |
| TD-004 | Minimal error handling in server startup | Difficult debugging | Medium | Improve logging and recovery |
| TD-005 | Metrics ticker runs indefinitely | Possible resource issues | Low | Add proper shutdown handling |

---

# Selected Debt to Fix

## TD-001 - Manual Environment Variable Parsing

### Problem
The application manually parses `.env` files inside `main.go`, making maintenance harder.

### Refactor Goal
Move environment loading into a dedicated helper function/module for cleaner architecture.

