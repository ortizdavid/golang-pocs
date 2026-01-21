curl -X POST http://localhost:1234/rpc \
  -H "Content-Type: application/json" \
  -d '{
    "method": "ProductService.Create",
    "params": [{
        "Name": "PC Lenvovo",
        "Code": "PC-001",
        "UnitPrice": 1150.00
    }],
    "id": 1
  }'