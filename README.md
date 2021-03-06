# hack-teamC-2021-summer -再翻訳チャットツール-

システム工学研究会2021夏のハッカソンで制作

## メンバー
- [wind111-lang](https://github.com/wind111-lang):Websocket,Translate API(Go)
- [Minmin](https://github.com/Minmin067):Websocket,JSONの値受け渡し,Webページ作成(Vue)
- [Bayashin](https://github.com/Bayashin):構造体実装,Websocket,Database(Go)
- [SuperTikuwa](https://github.com/SuperTikuwa):チームメンター(全体サポート)

## セットアップ
**1. Dockerのセットアップ**

- まず, mariadbのイメージを持ってくる必要があるので, ```docker pull mariadb```と  
入力して,イメージを取得してください.

- dockerフォルダ内にdocker-compose.ymlがあるので,  
ターミナルでディレクトリを```docker```にいる状態で```docker-compose build```と入力します.
  
- 次に, コンテナを作成して起動するために,```docker-compose up -d```と入力します.  
うまくいかない場合は,```docker-compose up``` 
と入力してください.
  
- コンテナに接続する前に,``` docker ps ```と入力し, 作成されたコンテナを確認します.　確認できたら,  
```docker exec -it <作成された CONTAINER ID> bash```と入力し,接続します.
  
- 接続できたら, ``` mysql -u root -p ```と入力します. すると,パスワードを求められるので,  
 ```admin```と入力して,MySQLにログインしてください.
  
-  ログインが完了したら,```show databases;```と入力して,データベース内に```chat_logs```があるかを確認します.  
確認できたら, ```use chat_logs;```と入力し,```chat_logs```を選択します.そして,```show tables;```と入力し,  
```chat_logs```があることを確認します. 最後に,```show columns from chat_logs;```と入力して,```mysql/db```のディレクトリにある,  
```chat_logs.sql```と内容が一致していれば, セットアップは完了です.

**2. yarnのセットアップ**
- ターミナルで,```frontend/Chat```ディレクトリにいる状態で,```yarn install```と入力して, パッケージを導入します.  
次に, ```yarn build```と入力して,プロジェクトをビルドしてください.

**3. 翻訳APIを呼び出すための設定**
- まず, Azureのポータルサイトにアクセスにアクセスしてください. もしAzureのアカウントを作成していない場合は先にアカウントを作成してください.  
 https://qiita.com/suzukin/items/e50703dd63edbc028dff (学生の場合は,学生プランに申し込みができます.)  
 https://azure.microsoft.com/ja-jp/free/ (学生に該当しない場合はこちら.)  
  
- アカウントが作成できたら,https://docs.microsoft.com/ja-jp/azure/cognitive-services/translator/quickstart-translator?tabs=go の前提条件のタブにある,
Translatorリソースの作成をクリックします. すると, Translatorの作成画面が開かれます.  
- リソースグループのタブに, 新規作成というのがあるので, それをクリックして,
 名前を自由に入力してください.  
 - 入力できたら,下にリソースグループのリージョンが現れますが, 基本的に東日本を選択してください. その下にあるリージョンも同様です.
 次に, 名前を自由に入力して, 価格レベルは, Free F0を選択してください.  
 - 左下の,確認及び作成をクリックすると検証が行われるので,検証が終わったら左下の作成をクリックしてください.
 しばらくするとデプロイが完了するので,完了したらリソースに移動をクリックしてください.  
 - 移動できたら, 左のリソース管理のタブの,キーとエンドポイントをクリックしてください. そして, キーの表示をクリックします.  
 - すると, キーが表示されるので, キー1の部分を```backend/.env```の, ```subscriptionKey```に貼り付けてください. ```location```は,東日本をリージョンとして指定しているので,```japaneast```と入力してください.
 
## 使い方
**4. main.goの実行とWebサイト表示**
- ```backend```ディレクトリ内に,```main.go```があるので,　```go run main.go```と入力します. 
しばらくすると,```Chat server opened```とターミナルに表示されるので,　Webブラウザを開いて, URLに```localhost:8081```と入力してください. そうすると,チャット画面が表示されます.  
他のマシンでチャットアプリに接続する場合は, URLを```main.goを実行しているマシンのIPアドレス:8081```と入力してください.

## 注意:
- ChromeやFirefoxなどの拡張機能に対応しているブラウザで,セキュリティソフト(特にNortonやウイルスバスター辺り)の拡張機能が入っている場合,動作しない可能性があります.
その場合には,拡張機能をオフにしてページをリロードしてください.  
- 現状，　同一LANでないと動作しません.

## その他
- コントリビュートを歓迎しています. 何かありましたらIssuesやPull Requestを投げてくれると幸いです.


