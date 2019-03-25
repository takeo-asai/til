# wget によるスクレイピング

```bash
$ wget -r --no-parent -w 1 -l 1 --restrict-file-names=nocontrol https://gihyo.jp/dp/
```

- Microdata(itemprop) を使用することでページのデザインの変更に強くする
- sed の @ は / の代わりに使える

```bash
$ cat gihyo.jp/dp/index.html | grep 'itemprop="name"' | sed -E 's@<br/>@@' | sed -E 's@<[^>]*>@@g' | sed 's/^ *//'
```

minify や prettier されている html は grep しづらいためこの方法は不向き。python などの一般的なプログラミング言語を使ったスクレイピングの方が汎用的。
