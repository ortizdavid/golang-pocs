#!/bin/bash

BASE_URL="http://localhost:8080"

echo "=== 1. Testing Root Endpoint (GET /) ==="
curl -X GET "$BASE_URL/"
echo -e "\n\n--------------------------------------------------\n"

echo "=== 2. Testing Calculation Simulation (POST /invoices/simulate) ==="
curl -X POST "$BASE_URL/invoices/simulate" \
  -H "Content-Type: application/json" \
  -d '{"unitPrice": 2500.0, "quantity": 2, "taxRate": 0.14}'
echo -e "\n\n--------------------------------------------------\n"

echo "=== 3. Testing Calculate and Save (POST /invoices) ==="
RESPONSE=$(curl -s -X POST "$BASE_URL/invoices" \
  -H "Content-Type: application/json" \
  -d '{"unitPrice": 1500.0, "quantity": 3, "taxRate": 0.14}')
echo "$RESPONSE"
echo -e "\n\n--------------------------------------------------\n"

echo "=== 4. Testing List All invoices (GET /invoices) ==="
curl -X GET "$BASE_URL/invoices"
echo -e "\n\n--------------------------------------------------\n"

echo "=== 5. Testing Get Calculation by ID (GET /invoices/{id}) ==="
# Extracts ID 
INVOICE_ID=$(echo "$RESPONSE" | grep -o '"id":"[^"]*' | head -1 | cut -d'"' -f4)

if [ -n "$INVOICE_ID" ]; then
    echo "Using extracted ID: $INVOICE_ID"
    curl -X GET "$BASE_URL/invoices/$INVOICE_ID"
else
    echo "No ID found from previous step. Run manually:"
    echo "curl -X GET '$BASE_URL/invoices/YOUR_INVOICE_ID'"
fi

echo -e "\n\n=== All quick tests completed! ==="