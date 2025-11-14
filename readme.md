# BE Sharing Vision

A backend API for managing articles built with Go and MySQL.

## Quick Start

### Step 1: Copy and Configure Environment Variables

```bash
cp .env.example .env
```

Edit `.env` and update the values if needed (default values should work for local development):

```env
DB_ROOT_PWD=admin
DB_NAME=sv_article
DB_USER=admin
DB_PASSWORD=admin
DB_HOST=localhost
DB_PORT=3306
APP_PORT=8080
```

### Step 2: Install Go Dependencies

```bash
go mod download
go mod tidy
```

### Step 3: Choose Your Setup

#### Option A: Using Docker (Easiest)

Start MySQL with Docker Compose:

```bash
docker-compose up -d
```

Then run the app:

```bash
go run main.go
```

#### Option B: Without Docker

Install MySQL manually, create the database and user, then run:

```bash
go run main.go
```

### Step 4: Access the API

The app will be running at `http://localhost:8080`
note: app port depend on APP_PORT in .env, make sure to hit the correct port

## Stopping

**With Docker:**
```bash
docker-compose down
then stop the Go process (Ctrl+C)
```

**Without Docker:**
Just stop the Go process (Ctrl+C)

##Postman Collections
https://web.postman.co/workspace/22e19e24-b8e9-46c1-8b45-97f0f043ca9e/collection/39463387-48170647-20da-448b-a5ac-dc9f5440fff6?action=share&source=copy-link&creator=39463387