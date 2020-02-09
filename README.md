# Lash Semi Chat
プロジェクト用のあれ

## DBの初期化について
気が向いたら自動化します。
事前に、`docker-compose up`でコンテナを起動させておく。

すでに起動してる人は、`docker-compose down --rmi local`ってしておくと色々吹っ飛ぶ

1. `docker exec -it lschat-mysql bash`でコンテナの中に入る
2. `mysql -u root -p`でmysqlに接続する
3. `source /docker-entrypoint-initdb.d/init.d/ddl.sql`でdbとテーブルを定義しておく
4. `source /docker-entrypoint-initdb.d/init.d/dml.sql`でテストデータを仕込む


## DBのファイル構成について
ファイル構成について色々
- 基本はmysqlディレクトリ以下にファイルを追加する
- ddl.sqlでDB，テーブルの定義を行う
- dml.sqlでテストデータの登録なり，crudを行う

```
.
├── README.md
├── mysql
    └── init.d
        ├── ddl.sql // db作ってテーブル作る
        └── dml.sql // テストデータのセット
```
