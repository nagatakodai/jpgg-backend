作成日時: 2025-07-29 04:41:15 UTC - 本ファイルはAIによって翻訳されました。

# `/build`

パッケージングと継続的インテグレーション用のディレクトリです。

クラウド (AMI) やコンテナ (Docker)、OS パッケージ (deb, rpm, pkg) の設定やスクリプトは `/build/package` に入れます。

Travis や CircleCI、Drone などの CI 設定やスクリプトは `/build/ci` に入れます。CI ツールによっては設定ファイルの場所に厳しいものもあります。その場合はルートなどツールが期待する場所へのシンボリックリンクを張るなどしてもかまいません。

例:

* https://github.com/cockroachdb/cockroach/tree/master/build
