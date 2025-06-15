#! /bin/bash
for i in {1..10000}
do
    curl -X POST http://localhost:8080/api \
        -H "Content-Type: application/json" \
        -d '{"id": 1, "name": "John Doe", "email": "john.doe@example.com"}'
done