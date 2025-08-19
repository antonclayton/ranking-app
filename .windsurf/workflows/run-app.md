---
description: Running the frontend and backend dev environment 
---

This workflow starts both the Go backend and the React frontend in separate terminals, allowing you to run the full application stack with a single command.

// turbo-all

1. Start the Backend Server
This step navigates to the `ranking-app-backend` directory and starts the Go API server.

```bash
cd /Users/aclayton/dev/go-ranking-app/ranking-app-backend
go run cmd/main.go

2. Start the Frontend Server This step navigates to the ranking-app-frontend directory and starts the Vite development server.
```bash
nvm use 20
cd /Users/aclayton/dev/go-ranking-app/ranking-app-frontend
npm run dev