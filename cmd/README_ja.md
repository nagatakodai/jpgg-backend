作成日時: 2025-07-29 04:41:15 UTC - 本ファイルはAIによって翻訳されました。

# `/cmd`

このプロジェクトのメインアプリケーションを置く場所です。

各アプリケーション用ディレクトリ名は生成したい実行ファイル名と一致させます (例: `/cmd/myapp`)。

アプリケーションディレクトリに多くのコードを置かないようにしましょう。コードが他プロジェクトからインポートして利用できると考えられるなら `/pkg` ディレクトリへ入れてください。再利用される想定がなかったり再利用してほしくない場合は `/internal` ディレクトリに置きます。どんな使われ方をするか分からないので、意図は明確に示しましょう。

一般的には `/internal` や `/pkg` からコードを取り込んで呼び出すだけの小さな `main` 関数のみがここに置かれます。

例:

* https://github.com/vmware-tanzu/velero/tree/main/cmd (非常に小さな `main` 関数のみで残りはパッケージ化されている例)
* https://github.com/moby/moby/tree/master/cmd
* https://github.com/prometheus/prometheus/tree/main/cmd
* https://github.com/influxdata/influxdb/tree/master/cmd
* https://github.com/kubernetes/kubernetes/tree/master/cmd
* https://github.com/dapr/dapr/tree/master/cmd
* https://github.com/ethereum/go-ethereum/tree/master/cmd
