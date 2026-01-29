curl -X POST http://localhost:1234/rpc \
  -H "Content-Type: application/json" \
  -d '{
    "method": "ProductService.GetByID",
    "params": [1],
    "id": 2
  }'