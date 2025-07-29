# jpgg-backend

# 🏗️ Backend 開発環境セットアップ手順（初回クローン時）

> **対象**: このリポジトリを _はじめて_ クローンする開発者  
> **前提**: WSL 2／Docker Desktop／VS Code + Dev Containers 拡張がインストール済み 🐳

---

## 1. リポジトリを取得

```bash
# WSL2 側 (例: Ubuntu) に配置
cd ~/jpgg/jpgg-backend # 任意の作業ルート
git clone github.com/nagatakodai/jpgg-backend
cd jpgg-backend
```

> Windows ドライブ（`C:\` など）に置くとファイル I/O が遅くなるので注意。

---

## 2. Dev Container を起動

1. VS Code で **「Open Folder…」 → `backend` フォルダ** を選択  
2. 右下ポップアップで **「Reopen in Container」** をクリック  
3. 初回ビルドには数分かかります  
   *（`air`, `gopls`, `dlv`, `golangci‑lint` などを自動インストール）*

---

## 3. モジュール初期化（初回のみ）

```bash
# コンテナ内ターミナルで実行
go mod init github.com/nagatakodai/jpgg-backend      # ← 既にある場合は不要
go get github.com/gin-gonic/gin@latest      # Gin を取得
go mod tidy                                 # 依存整理
```

---

## 4. 開発用サーバーを起動

```bash
air   # ホットリロード付きで :8080 が立ち上がる
```

ブラウザで `http://localhost:8080/ping` → `{"message":"pong"}` が返れば OK 🎉

---

## 5. よく使うコマンド

| 用途                 | コマンド |
| -------------------- | -------- |
| 依存追加             | `go get <module>@latest && go mod tidy` |
| 静的解析 (lint)      | `golangci-lint run ./...` |
| 単体テスト           | `go test ./...` |
| デバッガ（CLI）      | `dlv debug` |
| コンテナ再ビルド      | VS Code → Command Palette → **Dev Containers: Rebuild** |

---

## 6. トラブルシューティング

| 症状                                   | 解決策 |
| -------------------------------------- | ------ |
| Dev Container ビルドに失敗             | Docker Desktop が起動済みか確認し、VS Code 左下 **Rebuild Container** |
| `go: no modules found`                 | `go mod init` → `go mod tidy` を忘れていないか確認 |
| API が 404/タイムアウト                | `air` が起動しているか確認／`GIN_MODE=debug` で再起動 |

---

## 7. 参考リンク

* **開発環境の詳細** → `.devcontainer/` ディレクトリ  
* **CI/CD・インフラ構築** → Confluence 「AWS環境構築ガイド」
* **ディレクトリ構成** → [リンク](README_ja.md)