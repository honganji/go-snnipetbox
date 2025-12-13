# Snippetbox

テキストスニペットを共有するためのWebアプリケーションです。

## プロジェクト構成

```
.
├── cmd/
│   └── web/          # アプリケーションのエントリーポイントとHTTPハンドラー
├── ui/
│   ├── html/         # HTMLテンプレート
│   │   ├── pages/    # ページテンプレート
│   │   └── partials/ # 再利用可能なテンプレートパーツ
│   └── static/       # 静的アセット
│       ├── css/
│       ├── js/
│       └── img/
```

## 環境変数

プロジェクトルートに `.env` ファイルを作成し、以下の変数を設定してください：

```env
ADDR=:4000
DSN=web:pass@/snippetbox
```

### 変数一覧

| 変数 | 説明 | デフォルト値 |
|------|------|-------------|
| `ADDR` | HTTPサーバーのアドレスとポート | `:4000` |
| `DSN` | MySQLデータソース名 | `root:123@/snippetbox` |

### DSNフォーマット

MySQLのDSNは以下の形式に従います：

```
[ユーザー名]:[パスワード]@[プロトコル]([ホスト]:[ポート])/[データベース名]
```

例：
- `web:pass@/snippetbox` - デフォルトソケットを使用したローカル接続
- `web:pass@tcp(localhost:3306)/snippetbox` - localhostへのTCP接続
