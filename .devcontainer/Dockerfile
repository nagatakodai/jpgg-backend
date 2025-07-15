# 1. Go + Debian の公式イメージを土台に
FROM golang:1.24-bullseye

# 2. コンテナ内の作業パス
WORKDIR /workspaces/jpgg-backend

RUN apt-get update -qq && apt-get install -y --no-install-recommends git bash make 

# 3. Go モジュールをインストール
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/home/vscode/.cache/go-build \
go install github.com/air-verse/air@latest && \
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
go install golang.org/x/tools/gopls@latest && \
go install github.com/go-delve/delve/cmd/dlv@latest

# 4. デフォルトは bash で入る
CMD [ "bash" ]