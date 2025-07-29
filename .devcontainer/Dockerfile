# Go + Debian の公式イメージ
FROM golang:1.24-bullseye

# コンテナ内の作業パス
WORKDIR /workspaces/jpgg-backend

# git, bash, make をインストール
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
        git bash make

# Go モジュールをインストール
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/home/vscode/.cache/go-build \
go install github.com/air-verse/air@latest && \
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest && \
go install golang.org/x/tools/gopls@latest && \
go install github.com/go-delve/delve/cmd/dlv@latest

# デフォルトは bash で入る
CMD [ "bash" ]

# 環境変数
ENV GIN_MODE=debug
