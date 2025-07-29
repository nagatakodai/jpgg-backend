作成日時: 2025-07-29 04:41:15 UTC - 本ファイルはAIによって翻訳されました。

# `/internal`

外部アプリケーションやライブラリからインポートしてほしくないプライベートなコードを置きます。このレイアウトは Go コンパイラによって強制されます。詳しくは Go 1.4 の [リリースノート](https://golang.org/doc/go1.4#internalpackages) を参照してください。`internal` ディレクトリはトップレベルに限らず、プロジェクトツリーの任意の階層に複数作成できます。

共有コードと共有しないコードを分けるために、`internal` 配下にさらに構造を追加することもできます。必須ではありませんが、意図を示すのに役立ちます。アプリケーションコードは `/internal/app` (例: `/internal/app/myapp`) に、複数のアプリで共有するコードは `/internal/pkg` (例: `/internal/pkg/myprivlib`) に置くとよいでしょう。

例:

* https://github.com/hashicorp/terraform/tree/main/internal
* https://github.com/influxdata/influxdb/tree/master/internal
* https://github.com/perkeep/perkeep/tree/master/internal
* https://github.com/jaegertracing/jaeger/tree/main/internal
* https://github.com/moby/moby/tree/master/internal
* https://github.com/satellity/satellity/tree/main/internal
* https://github.com/minio/minio/tree/master/internal

## `/internal/pkg`

例:

* https://github.com/hashicorp/waypoint/tree/main/internal/pkg
