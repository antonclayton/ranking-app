# Ranking App

A full-stack application with a Go backend and React/TypeScript frontend.

## Prerequisites

General:
- nvm --version -> 0.40.3 (nvm install 20, nvm use 20)
- go version -> go1.24.4 darwin/arm64
- node --version -> v18.18.2
- npm --version -> 6.14.18


### Frontend
- Node.js 18+ (LTS recommended)
- npm 9+ or yarn 1.22+

### Backend
- Go 1.24.4

## Project Structure

```
.
├── ranking-app-backend/    # Go backend
└── ranking-app-frontend/   # React frontend
```

## Setup Instructions

### Backend Setup

1. Navigate to the backend directory:
   ```bash
   cd ranking-app-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Run the server:
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080` by default.

### Frontend Setup

1. Navigate to the frontend directory:
   ```bash
   cd ranking-app-frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   # or
   yarn install
   ```

3. Start the development server:
   ```bash
   npm run dev
   # or
   yarn dev
   ```
   The app will be available at `http://localhost:5173`

## Key Dependencies

### Backend
- [chi](https://github.com/go-chi/chi) - Lightweight, composable router
- [go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite3 driver for Go

### Frontend
- React 19
- TypeScript
- Vite
- Material-UI (MUI)
- React Router

## Environment Variables

Create a `.env` file in the root of both frontend and backend directories as needed.

### Backend (`.env`)
```
PORT=8080
# Add other environment variables here
```

## Development

- **Linting**: `npm run lint` (frontend)
- **Build**: `npm run build` (frontend)

## Database

This project uses SQLite for development. The database file is automatically created when you first run the application.

## License

[MIT](LICENSE)