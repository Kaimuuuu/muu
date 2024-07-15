# Project setup

## Running appliation

```sh
# replace the following ${replace_me} with correspond environment variable values
FRONTEND_URL=${frontend_url} MONGO_URL=${mongodb_url} MONGO_DB=${mongodb_db} PORT=${port} JWT_SECRET=${jwt_secret} go run main.go

# example
FRONTEND_URL=localhost:3000 MONGO_URL=mongodb://localhost:27017 MONGO_DB=github.com/Kaimuuuu/muu PORT=3001 JWT_SECRET=secret go run main.go
```

when the application is running use https://jwt.io/ to generate a useable JWT to create initial admin user

replace PAYLOAD with the following json **warning: this is a backdoor token so use this token create initial user only**
```json
{
  "employeeId": "0",
  "role": 0
}
```

replace VERIFY SIGNATURE with your ${jwt_secret} and copy the token

```sh
# create initial user

curl -X 'POST' \
  --json '{"name":${name},"age":${age},"role":0,"email":${email}}' \
  -H 'Content-Type: application/json' \
  -H 'Authorization: Bearer ${jwt_token}' \
  ${backend_url}/employee

```

create /public directory to store image
```sh
mkdir public
```
