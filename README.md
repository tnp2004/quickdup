# QuickDup

A simple app to create and share notes quickly. Write a note, get a short code, share it with anyone. The note disappears after a while.

## What It Does

- Create notes fast
- Get a 6-character code (like `a3b9c2`)
- Share that code with others
- Others type the code to see your note
- Notes go away after they expire

## Tech Stack

| Part | Technology |
|------|-----------|
| Backend | Go + Echo framework |
| Frontend | Next.js + Tailwind CSS |
| Database | PostgreSQL |
| Auth | JWT |

## Project Structure

```
quickdup/
├── app/              # Main app entry
├── configs/          # Config files
├── modules/
│   ├── auth/         # Login and signup
│   ├── notes/        # Create and find notes
│   ├── users/        # User accounts
│   └── servers/      # API routes
├── pkg/
│   ├── databases/    # Database setup
│   └── utils/        # Helper functions
└── web/              # Next.js frontend
```

## How It Works

1. User writes a note in the web app
2. Backend saves it to PostgreSQL
3. Backend creates a short random code
4. User shares the code
5. Someone enters the code
6. Backend finds the note and shows it
7. After expiry time, the code stops working

## Running the App

### Backend

```bash
# 1. Create a config file at configs/config.yaml
# 2. Run the server
go run app/main.go
```

### Frontend

```bash
cd web
npm install
npm run dev
```

## Config File Example

Create `configs/config.yaml`:

```yaml
server:
  host: "localhost"
  port: "3000"
  allowOrigins:
    - "*"
  bodyLimit: "10M"
  timeout: 30

database:
  host: "localhost"
  port: "5432"
  name: "quickdup"
  username: "postgres"
  password: "yourpassword"
  sslmode: "disable"

auth:
  jwt:
    accessSecretKey: "your-access-secret"
    refreshSecretKey: "your-refresh-secret"
    accessTokenExpireDuration: 900    # 15 minutes
    refreshTokenExpireDuration: 604800 # 7 days
  cloudinary:
    apiKey: "your-cloudinary-key"
    apiSecret: "your-cloudinary-secret"
```
