# CI/CD Pipeline Diagram

```text
Developer pushes code to main
          │
          ▼
    GitHub Actions CI
          │
          ▼
  Build and Run Tests
          │
          ▼
     Start Server
          │
          ▼
      Smoke Test
          │
          ▼
  Auto Deploy to Production
```
