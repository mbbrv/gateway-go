## Prerequisites

Make sure you have the following tools installed on your system:

- Docker
- Docker Compose

## Getting Started

Follow these steps to run the Golang server along with PostgreSQL and Redis using Docker Compose.

### Step 1: Clone the Repository

```bash
git clone https://github.com/your-username/your-repo.git
cd your-repo
```

### Step 2: Run the Docker Compose

```bash
docker-compose up
```

### Step 3: Check the Services

- Golang Server: [http://localhost:8000](http://localhost:8000)
- PostgreSQL: localhost:5432
- Redis: localhost:6379

### Run Tests

```bash
    go test ./...
```
