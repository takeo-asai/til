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

## 7 章

### AWS へクローラーを deploy する
