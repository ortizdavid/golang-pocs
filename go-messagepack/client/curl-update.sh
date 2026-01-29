curl -X POST http://localhost:1234/rpc \
  -H "Content-Type: application/json" \
  -d '{
    "method": "ProductService.Update",
    "params": [{
        "Name": "Mouse Gamer Pro",
        "Code": "MOU-99",
        "UnitPrice": 180.00
    }],
    "id": 4
  }'