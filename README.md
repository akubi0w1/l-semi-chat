# Lash Semi Chat
プロジェクト用のあれ

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
