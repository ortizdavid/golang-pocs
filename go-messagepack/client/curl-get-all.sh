curl -X POST http://localhost:1234/rpc \
  -H "Content-Type: application/json" \
  -d '{
    "method": "ProductService.GetAll",
    "params": [{}],
    "id": 6
  }'

