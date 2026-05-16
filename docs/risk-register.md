# Risk Register - Bit Battleship

## Risk Scoring Guide
- Likelihood: 1 (Low) to 5 (High)
- Impact: 1 (Low) to 5 (High)
- Risk Score = Likelihood × Impact

---

| ID | Risk Description | Likelihood | Impact | Risk Score | Mitigation Plan | Risk Owner |
|----|------------------|------------|--------|------------|-----------------|------------|
| R1 | Database connection failure | 3 | 5 | 15 | Add reconnection handling and validate DATABASE_URL | Kirk Espina |
| R2 | Matchmaking queue bugs | 4 | 4 | 16 | Test matchmaking logic with multiple users | Backend Developer |
| R3 | Players disconnect during match | 4 | 4 | 16 | Add disconnect handling and room cleanup | Backend Developer |
| R4 | WebSocket synchronization issues | 3 | 5 | 15 | Implement state validation and testing | Backend Developer |
| R5 | Ships overlapping incorrectly | 2 | 4 | 8 | Add placement validation checks | Frontend Developer |
| R6 | Merge conflicts in GitHub | 3 | 3 | 9 | Use feature branches and review PRs carefully | Team Lead |
| R7 | Delays due to unfinished tasks | 4 | 3 | 12 | Assign tasks clearly and track sprint progress | Scrum Master |
| R8 | UI bugs affecting gameplay | 3 | 3 | 9 | Perform regular testing before merges | QA Tester |
