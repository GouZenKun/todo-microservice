## 概要
golangでTodoマイクロサービスを実装するプロジェクト

## セットアップ手順
1. このリポジトリをダウンロードする
   1. 上右のダウンロードボタンでダウンロードしてZIPを解凍する
   2. `git clone`でダウンロード
2. Goをインストールする
   1. [Go](https://go.dev/doc/install)
3. Mysqlをインストールする
   1. [mysql](https://www.mysql.com/)
5. Dockerをインストールする
   1. [Docker](https://www.docker.com/)
6. Makeコマンド実行するためにGNU Makeをインストールする
   1. [Make](https://www.gnu.org/software/make/)
7. 全部を揃えば、ターミナルでリポジトリを開いて下記のコマンドを実行する
   1. `go mod tidy`

## 利用方法の説明
1. mysqlとapiサーバーを実行する
   1. `make docker_up`を実行する
   2. 実行完了まで待つ
2. CLIでAPI叩く
   1. `make build_cli`でバイナリをビルドしてそれを使って簡単にAPI叩ける
   2. `make cli h`を試す、画面に使えるコマンドを確認できる
   3. コマンドを実行することでAPIを叩く

## アーキテクチャ
- Service-Repository-Controller Pattern
### 選定理由：
- マイクロサービスの機能は決められてClean Architectureまでやる必要ない
- コードジェネレーターでほぼ構造分けられるのでちょうど疎結合がいい
- 業界の人材は既に馴染んでいるのでメンテナンス面は悪くない

## 完成度
### バグ
- sqlmodeがno_zero_in_date使われてtodosテーブルのcreate_at, update_atがバグる
### 未完成
- SSL
- Database migrations
- CLI
- など
