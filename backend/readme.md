 curl -X GET "http://api.weatherapi.com/v1/history.json?key=1084236536cc47d486f32720242006&q=37.7749,-122.4194&dt=2024-04-01&end_dt=2024-05-01"





TOKEN="your_jwt_token"

curl -X GET http://localhost:8080/profile \
-H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzE4OTQ0MDgxfQ.gBDCv87ENX1Iv909pgpa66alKM4aZodz1qBc_FDuIJg"


{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyIiwiZXhwIjoxNzE4OTQ0MDgxfQ.gBDCv87ENX1Iv909pgpa66alKM4aZodz1qBc_FDuIJg"}



curl -X GET \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG5fZG9lIiwiZXhwIjoxNzE4OTUzMzk0fQ.v_sWLx9NEo8l5UhjMvi7wZq6SnWfXHqtVfajdRmaJzE" \
     http://localhost:8080/profile
