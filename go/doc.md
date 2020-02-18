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

1/8
- archive, tag のロギング？
- そういえばアーカイブの認証どうすんの？
- archive, pathっているの？thread idでよくね？



!!!!
めっちゃ挟み込んだので。。。ちょっと長文になるけど、なんとなく思ったアドバイスですー参考までに。

## ログの仕様について
ログの仕様がちょっと変わって、引数が複数取れるようになったので、今まで
```
logger.Error(fmt.Sprintf("xxx handler: %s", err))
```
みたいにやってたところを、
```
logger.Error("xxx handler: ", err)
```
みたいにかけるようになりましたー。
以降こんな感じで書いてくれて大丈夫です！
Sprintfのほうが好きって場合はそれで全然大丈夫です！

## ログの排出タイミングとか
もしログの排出タイミング迷ってたら、めちゃ自己流だけどアドバイスです。
より内側にある関数を呼び出す時は、外側でlog出さなくてもokです。基本的にログは内側でエラーが起こった瞬間に出すようにしてるので...
具体的には、handlerからInteractorの関数を呼び出す場合、Interactorでログを排出しているので、handlerのエラーハンドリングではlogを出さない！って具合。interactor->repositoryも同様です。
多分自分も徹底できてないのでなんとも言えないけどな！！！ごめんなさい！！！！

## ログの種類
loggerの使い分けだけど、Errorはサーバエラー(5xx系)、Warnはクライアントエラー(4xx系)ってイメージで使ってます！
Infoは、サーバの稼働状況を吐いたり。Debugはその名の通りdebugで使ってくださいー！

## エラーレスポンス作成時のラップについて
レスポンス時、response.HttpError(w, err)で、errをBadRequest(err)みたいなラップするかの判断も、ログと同様で、内側を呼び出しているか？で決めてもらって大丈夫。
内側を呼んでる場合は内側の関数でエラーをラップするので、最終的に戻ってくるエラーはラップされた状態で戻ってきてくれます！多分！
~まあ、ラップされてなくても、デフォルトで500がつくので、多少ミスっても大丈夫です。~

## インタフェース定義について
インタフェース定義していく時に、関数の引数を型だけじゃなくて一緒に名前まで定義しておくと呼び出す時に使う側が楽になると思う。
例えばだけど、
```
StoreArchive(string, string, string, string, int) error
```
より
```
StoreArchive(archiveID, path, password, threadID string, isPublic int) error
```
のが初めて見る人にも優しいよねって具合に。メンテする時に死ぬかもしれないので、後者にしてもらえると嬉しいー！




まあ、ラップされてなくても、デフォルトで500がつくので、多少ミスっても大丈夫です。

TODO:スレッドの人数上限気にしている？
