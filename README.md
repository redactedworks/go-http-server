# go-backend-playground
Simple go server with Docker build file

Project structure follows [goland-standards](https://github.com/golang-standards/project-layout) recommendations.

```bash
grpcurl -d '{"user_id":"test"}' -plaintext localhost:9099 playground.UserService.GetUser
```

### For Mongo

`docker-compose up -d`

`mongosh "mongodb://root:example@localhost:27017"`