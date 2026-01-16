curl -X PUT http://localhost:8080/products/1 \
  -H "Content-Type: application/json" \
  -H "Accept-Language: en" \
  -d '{
    "name": "Wireless Keyboard",
    "unit_price": 400.00
  }'