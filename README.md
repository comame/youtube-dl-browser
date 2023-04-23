## 依存関係

- go
- ffmpeg
- yt-dlp

## セットアップ

> ℹ️ Windows では WSL で動かすとよさそう

### Chrome 拡張機能のインストール

chrome://extensions で `chrome-extension/` をインストールする

### サーバを起動する

`http://localhost:5906` でアクセスできるようにしておくこと

```
$ go install github.com/comame/youtube-dl-browser@latest
$ youtube-dl-browser
```

### 保存ディレクトリを設定する

拡張機能のアイコンを右クリックし、「オプション」
