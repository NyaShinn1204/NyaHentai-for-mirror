<p align="center">
    <h3 align="center">NyaHentai for Mirror</h3>
    <p align="center">
        GoとPHPを使ってNyaHentaiのミラーを作成します<br />
    </p>
</p>

## インストール

**※ Goのバージョン1.23以上が必要です**

> [!TIP]
> これは開發バージョンです！

まず:8190のポートを:80/api/に流す必要があります。

その次にApacheが実行されているフォルダにdoujinshiフォルダを突っ込んでください (デフォルトだと/var/www/html/)

現状閲覧しか作っていないので検索などマイリストなどは作っていません。

```bash
git clone https://github.com/NyaShinn1204/NyaHentai-for-mirror

cd NyaHentai-for-mirror/api-go

go run main.go
```


## 免責事項

急に使えなくなっても知りません。一日暇だったから作っただけです02