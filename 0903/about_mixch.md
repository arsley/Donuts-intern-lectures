# 0903

- 担当:長坂

## MixChannel サービス概要

- ミニ版を作り知識を得る
- 映像配信とリアルタイムチャットは別プロトコル?
- 10~30秒の動画を投稿するサービス
- タグ
- ファンになる(=フォローする)
- レベル?

- ライブ映像配信
- 動画の配信HTTPserver


### Service on ...

- AWS
- Cloudfront
- Google Cloud Platform
- firebase?

### API

- インターンではCloudはあまり使わない?

### 非同期タスク

- 処理コストのかかるものは別で行う
- タスクの重複などを防ぐ必要がある
  - SQS?
- queueに入れてtask専用インスタンスにて処理する
  - ユーザ削除が最重

### recommend system

- gRPC
- 一部HTTP(遅い)

### ライブ配信システム

- WebSocketがメイン

## 作るもの

- ブラウザで動くクライアントも作るらしい
