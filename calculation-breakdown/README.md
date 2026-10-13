# Calculation Breakdown PoC

A lightweight Go proof-of-concept demonstrating an **audit-ready calculation architecture**, separating pure business logic from persistence.

## Architecture

- **`cmd/`**: Application entry point
- **`database/`**: SQLite schema and connection
- **`handlers/`**: HTTP transport & routing (`net/http`)
- **`services/`**: Core calculation logic & rules
- **`repositories/`**: Database persistence (Transactions)
- **`models/`**: Domain entities

## Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `GET` | `/` | API discovery |
| `POST` | `/invoices` | Calculate, save steps, and persist |
| `POST` | `/invoices/simulate` | Simulate calculation in-memory |
| `GET` | `/invoices` | List all saved invoices |
| `GET` | `/invoices/{id}` | Get invoice breakdown by ID |

## Quick Start

```bash
# Run the application
make run

# Build binary
make build

```

## Example Usage

**Calculate & Save:**

```bash
curl -X POST http://localhost:8080/invoices \
  -H "Content-Type: application/json" \
  -d '{"unitPrice": 1500.0, "quantity": 3, "taxRate": 0.14}'

```

**Simulate:**

```bash
curl -X POST http://localhost:8080/invoices/simulate \
  -H "Content-Type: application/json" \
  -d '{"unitPrice": 2500.0, "quantity": 2, "taxRate": 0.14}'

```

**Get by ID:**

```bash
curl http://localhost:8080/invoices/YOUR_UUID_HERE

```
