#!/bin/bash

Base URL of the running server

BASE_URL="http://localhost:8080"

echo "=== 1. Testing Root Endpoint (GET /) ==="
curl -X GET "$BASE_URL/"
echo -e "\n\n--------------------------------------------------\n"

echo "=== 2. Testing Calculation Simulation (POST /calculations/simulate) ==="
curl -X POST "$BASE_URL/calculations/simulate" 

-H "Content-Type: application/json" 

-d '{"unitPrice": 1500.0, "quantity": 3, "taxRate": 0.14}'
echo -e "\n\n--------------------------------------------------\n"

echo "=== 3. Testing Calculate and Save (POST /calculations) ==="

Capturing the response to extract an ID if needed, but printing directly for now

RESPONSE=$(curl -s -X POST "$BASE_URL/calculations" 

-H "Content-Type: application/json" 

-d '{"unitPrice": 2500.0, "quantity": 5, "taxRate": 0.14}')
echo "$RESPONSE"
echo -e "\n--------------------------------------------------\n"

echo "=== 4. Testing List All Calculations (GET /calculations) ==="
curl -X GET "$BASE_URL/calculations"
echo -e "\n\n--------------------------------------------------\n"

To test GET /calculations/{id}, we can fetch the list or use a placeholder ID

(Replace INVOICE_ID with a valid ID from the list above if you want to test a specific one)

echo "=== 5. Testing Get Calculation by ID (GET /calculations/{id}) ==="

Attempting to fetch a sample ID or showing how to query it:

echo "Run this manually with a real ID: curl -X GET '$BASE_URL/calculations/YOUR_INVOICE_ID'"

echo -e "\n=== All quick tests completed! ==="