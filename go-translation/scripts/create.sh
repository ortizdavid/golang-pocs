curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -H "Accept-Language: pt" \
  -d '{
    "name": "Teclado Mecânico RGB",
    "code": "KEY-RGB-01",
    "unit_price": 350.50
  }'


curl -X POST http://localhost:8080/products \
  -H "Content-Type: application/json" \
  -H "Accept-Language: en" \
  -d '{
    "name": "Phone",
    "code": "PH-AC-02",
    "unit_price": 950.50
  }'

