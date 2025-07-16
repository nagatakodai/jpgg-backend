#!/usr/bin/env bash
set -euo pipefail      # コケたら即終了

echo "▶️  DevContainer 初期セットアップ開始..."

# Go モジュールをダウンロード
echo "👍  Go モジュールをダウンロード中..."
go mod download

echo "🎉  DevContainer セットアップ完了！"