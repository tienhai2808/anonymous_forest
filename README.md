# Anonymous Forest (Chốn An Yên)

Anonymous Forest (Chốn An Yên) is a full-stack web application for anonymous sharing and reading of thoughts, confessions, and stories. It is designed to provide a safe, tranquil space for users to express themselves and connect with others without revealing their identity.

---

## Table of Contents
- [Features](#features)
- [Architecture](#architecture)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Usage](#usage)
- [Docker & Deployment](#docker--deployment)
- [Contributing](#contributing)
- [License](#license)

---

## Features
- **Anonymous Posting:** Share your thoughts without registration.
- **Empathy & Protest:** React to posts with empathy or protest.
- **Commenting:** Leave comments on posts.
- **Rate Limiting:** Limits on daily posts and reads to encourage mindful sharing.
- **Temporary Links:** Shareable post links expire after 7 days.
- **Moderation:** Posts with 3 protests are automatically deleted.
- **Responsive UI:** Modern, mobile-friendly interface.

## Architecture

```
┌────────────┐      ┌────────────┐      ┌────────────┐
│  Frontend  │ <--> │  Backend   │ <--> │  Database  │
│ (Next.js)  │      │  (Go API)  │      │MongoDB/Redis│
└────────────┘      └────────────┘      └────────────┘
```

## Tech Stack
- **Frontend:** Next.js, React, Tailwind CSS, TypeScript
- **Backend:** Go (Fiber), MongoDB, Redis
- **Containerization:** Docker, Docker Compose

## Project Structure

```
anonymous_forest/
├── backend/      # Go API server
│   ├── cmd/api/  # Main entrypoint
│   ├── configs/  # Config files
│   ├── internal/ # Business logic, models, handlers
│   └── ...
├── frontend/     # Next.js app
│   ├── src/      # App source code
│   └── ...
├── docker-compose.local.yml
├── docker-compose.prod.yml
└── README.md
```

## Getting Started

### Prerequisites
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

### Local Development

1. **Clone the repository:**
	```sh
	git clone https://github.com/tienhai2808/anonymous_forest.git
	cd anonymous_forest
	```

2. **Start all services:**
	```sh
	docker-compose -f docker-compose.local.yml up --build -d
	```

3. **Access the app:**
	- Frontend: [http://localhost:3000](http://localhost:3000)
	- Backend API: [http://localhost:5000/backend/api](http://localhost:5000/backend/api)

### Manual Setup (Advanced)

#### Backend
```sh
cd backend
go mod download
make run
```

#### Frontend
```sh
cd frontend
npm install
npm run dev
```

## Configuration

- **Backend:** See `backend/configs/config.yml` for server, database, and cache settings.
- **Frontend:** Environment variables are set via Docker or `.env.local` for API URLs.
- **Docker Compose:** Edit `docker-compose.local.yml` for service configuration.

## Usage

### Posting & Reading
- Each user can post up to 5 times and read up to 10 posts per day.
- Posts are anonymous and can be shared via a temporary link (valid for 7 days).
- Posts with 3 protests are automatically deleted.

### API Endpoints
- The backend exposes RESTful endpoints under `/backend/api` (see code for details).

## Docker & Deployment

### Build & Run (Production)
```sh
docker-compose -f docker-compose.prod.yml up --build -d
```

### Services
- **frontend:** Next.js app (port 3000)
- **backend:** Go API (port 5000)
- **mongodb:** Database (port 27017)
- **redis:** Cache (port 6379)

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements or bug fixes.

## License

This project is licensed under the MIT License.
