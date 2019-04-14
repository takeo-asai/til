# Python & クローリング & スクレイピング

Python 3 系

## 1 章

### Shell Script によるスクレイピング

```bash
$ wget "url" | grep "keyword" | sed "replace" > output.txt
```

## 2 章

### 標準 Python によるスクレイピング

1 章の内容を python を使って rewrite

```python
html = fetch("url")
items = scrape(html)
save("output.db", items)
```

## 3 章

### ライブラリを使用した Python によるスクレイピング

```bash
docker run -p 27017:27017 --name dev-mongo mongo
```

## 4 章

### エラー処理、繰り返し処理、robot が従うべき挙動など

## 5 章

### API, 自然言語処理, 可視化

```bash
$ pip install selenium
$ brew cask install phantomjs
```

## 6 章

### Scrapy と画像処理

- web ページからのリンク抽出
- robots.txt, xml sitemap 認識
- インターバル、並列処理、重複防止、リトライ
- デーモン化、ジョブ管理

```bash
$ scrapy startproject six_scrapy
$ scrapy genspider news news.yahoo.co.jp
$ scrapy shell https://news.yahoo.co.jp/pickup/6320347
$ scrapy crawl news -o news.jl
```

readability, extracontent を使用することで、不特定多数の web サイトから本文が取り出せる（可能性が高い）

- css セレクタなしでもそれなりには取れるのでは？

## 7 章

### AWS へクローラーを deploy する

クローリングとスクレイピングの分離
