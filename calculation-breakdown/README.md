# Calculation Breakdown PoC 

Proof of Concept (PoC) built in Go implementing an **audit-ready calculation breakdown** architecture. It handles complex multi-step business rules, isolating pure mathematical functions from database operations, and ensures transactional integrity when persisting calculation steps and final totals.

## Architecture

The project follows a clean, modular structure:
- **`cmd/`**: Application entry point and server initialization.
- **`database/`**: SQLite connection setup and initial schema migration (`tables.sql`).
- **`handlers/`**: HTTP transport layer, routing, and JSON request/response handling.
- **`services/`**: Core business logic containing pure calculation functions and workflow coordination.
- **`repositories/`**: Data persistence layer using Go's `database/sql` with atomic transaction management.
- **`models/`**: Domain entities (`Invoice` as aggregate root and `Step` entities).

## Tech Stack

- **Language**: Go
- **Database**: SQLite (`go-sqlite3`)
- **Routing**: Standard library `net/http` (`net/http.ServeMux`)
- **UUID Generation**: `github.com/google/uuid`

## Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | API overview and available endpoints |
| `POST` | `/calculations` | Execute calculation, record steps, and persist to database |
| `POST` | `/calculations/simulate` | Simulate calculation purely in memory without persisting |
| `GET` | `/calculations` | List all saved calculation invoices |
| `GET` | `/calculations/{id}` | Get calculation details and breakdown steps by ID |

## Getting Started

### Prerequisites

- Go 
- SQLite3 

### Running the Project

Use the provided `Makefile` to manage the project lifecycle:

- **Run the server:**
  ```bash
  make run

```

* **Build the binary:**
```bash
make build

```


* **Clean generated binaries and database files:**
```bash
make clean

```



## Example Usage

### 1. API Discovery (Root)

```bash
curl -X GET http://localhost:8080/

```

### 2. Calculate and Save

Process, persist, and retrieve the full calculation breakdown:

```bash
curl -X POST http://localhost:8080/calculations \
  -H "Content-Type: application/json" \
  -d '{"unitPrice": 1500.0, "quantity": 3, "taxRate": 0.14}'

```

### 3. Simulate Calculation

Simulate the calculation in memory without saving to the database:

```bash
curl -X POST http://localhost:8080/calculations/simulate \
  -H "Content-Type: application/json" \
  -d '{"unitPrice": 2500.0, "quantity": 2, "taxRate": 0.14}'

```

### 4. List All Invoices

Retrieve all saved calculation records:

```bash
curl -X GET http://localhost:8080/calculations

```

### 5. Get Invoice Details by ID

Fetch a specific invoice and its breakdown steps using its UUID:

```bash
curl -X GET "http://localhost:8080/calculations?id=YOUR_INVOICE_UUID_HERE"

```

```