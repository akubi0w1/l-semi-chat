# client

## mockupの作成について
npmがインストールされている必要があります。

sassのコンパイルを自動化するためにgulpを導入してます。

sass: cssが書きやすくなるメタ言語
gulp: タスクランナー。コンパイル、トランスパイルなどの自動化が可能。

`gulp`コマンドを使うため、以下で、gulpを動かせるようにしてください。

```
// gulpコマンドを実行できるように、グローバルインストールする
$ npm install -g gulp
```

`client`(gulpfile.jsがある)ディレクトリに移動し、以下コマンドで、ファイルの監視と自動コンパイルが始まります。
`client/src/sass/*.scss`のファイルを監視してくれる。

```
$ gulp sass-watch

// 単発でコンパイルするなら
$ gulp sass
```

## いまのとこの、html構造

こんな感じ。。。

```
----------------------------
| nav |     container      |
|     | sidebar | clontent |
|     |         |          |
```