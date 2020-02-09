# 住み分けの話
- サーバの接続、ルーティング、sqlの接続(main)
- リクエストを受け取る、レスポンスを作る(handler)
- 処理(service)
    - ロジック(interactor)
    - DBをぶっ叩く(repository)
- モデル定義(domain)

```
go
├── cmd
│   └── main.go // 一番いじってはいけない処理。いじらない
├── conf        // server, dbのコンフィグ置く場所。基本的にはいじらない
├── go.mod
├── go.sum
└── pkg
    ├── domain  // モデルをおいておく場所。外のことは何も知らない。
    │   ├── response.go  // errorレスポンスをどうにか定義してある
    │   └── user.go
    ├── interface // 外部のルールを使う。service, domainを知ってる
    │   ├── auth    // 認証とかパスワードの暗号化とかのユーティリティ。いじらない
    │   │   ├── hash.go
    │   │   └── jwt.go
    │   ├── database    // dbアクセスの実装。いじらない。
    │   │   └── sql_handler.go
    │   ├── dcontext    // contextの書き込みと読み出し。あんまいじらん
    │   │   └── dcontext.go
    │   ├── handler     // メソッドによって分離させたり、レスポンス作ったり。
    │   │   ├── account_handler.go
    │   │   ├── app_handler.go
    │   │   └── auth_handler.go
    │   └── server      // serverが提供する機能みたいな。
    │       ├── middleware  // 認証が必要な時に噛ませるのが入ってる。必要なら作る。いじらない
    │       │   └── auth.go
    │       ├── response    // レスポンス定義。この辺エラーハンドリングの関係で書き直すかも。いじらない。
    │       │   └── response.go
    │       ├── router      // urlのマッピング。path1種類ごとに1つのハンドラで定義します。
    │       │   └── router.go
    │       └── server.go   // serverの起動とか。いじらない
    └── service  // 処理。domainのみを知ってる
        ├── interactor  // メイン処理。
        │   ├── account_interactor.go
        │   ├── auth.go
        │   └── auth_interactor.go
        └── repository  // db関連の処理はこっち。
            ├── account_repository.go
            ├── auth_repository.go
            └── sql_handler.go
```

12/25
- 今のとこ、パスワードのハッシュとか、handlerでやってるので、ここをserviceに移行させるかも知らん。考え中です。
    - パスワードハッシュはおそらく移行できたはず。serviceへの以降も行ったはず。