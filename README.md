# sample-golang-structure
```azure
Directory
├── cmd
    ├── config.yaml
    ├── main.go
├── config
    ├── db.go
    ├── log.go
    ├── server.go
    ├── system.go
├── global
    ├── global.go
├── initialize
    ├── log-adaptor
        ├── xorm.go
    ├── init_db.go
    ├── initialize.go
    ├── router.go
    ├── viper.go
├── model
    ├── request
        ├── CreatedArticleRequest.go
        ├── DeleteArticleRequest.go
        ├── ListArticleRequest.go
        ├── UpdateArticleRequest.go
    ├── Article.go
├── repository
    ├── article.go
├── response
    ├── common.go
    ├── response.go
├── router
    ├── Router.go
├── service
    ├── Article.go
├── until
    ├── ratelog.go
main
```


```azure
Run App
cd ./cmd/
go run .
```

```azure
Build App
cd ./cmd/
go build main.go
```






<h5>Power by Voeun SO</h5>