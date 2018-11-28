# go-api-practice
<img width="952" alt="2018-11-29_03h17_43" src="https://user-images.githubusercontent.com/18478417/49185774-24a95d00-f3a6-11e8-9902-5fd418f69026.png">

## Quick Start
### Go serve
※ You need to install [dep](https://github.com/golang/dep)

```bash
make setup
go run main.go
curl localhost:8080/users
```

You can visit `localhost:8080/static/`

### Nuxt serve
To check implimentation for CORS

```bash
make setup
go run main.go
cd public
npm run dev
```

You can visit `localhost:3000/`

Our API server origin is `localhost:8080`, so CORS is required, and our API implement it.


## demo
#### New
<img width="951" alt="2018-11-29_03h17_53" src="https://user-images.githubusercontent.com/18478417/49185801-37239680-f3a6-11e8-9a23-7556576fc534.png">

#### Edit
<img width="957" alt="2018-11-29_03h18_03" src="https://user-images.githubusercontent.com/18478417/49185807-3985f080-f3a6-11e8-967e-f0aebc543b5c.png">
