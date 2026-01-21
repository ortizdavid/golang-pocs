curl -X POST http://localhost:1234/rpc \
  -H "Content-Type: application/json" \
  -d '{
    "method": "ProductService.CalculateTotal",
    "params": [{
        "productId": 1,
        "quantity": 5
    }],
    "id": 100
  }'