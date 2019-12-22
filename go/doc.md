# 住み分けの話
サーバの接続、ルーティング、sqlの接続(main)
リクエストを受け取って、レスポンスを作る(handler)
ビジネスロジック(service)
DBをぶっ叩く(repository)

```
go/
    cmd/
        main.go
    config/
        server_conf.go
        db_conf.go
    interface/
        handler/
            各種handler.go
        server/
            server.go
            response.go
            router.go
        database/
            sql_handler.go
    service/
        repository/
            user_repository.go
        interactor/
            user_interactor.go
    domain/
        user.go
```

domain: モデル定義
service: ビジネスロジック
interface: リクエスト、レスポンス関係