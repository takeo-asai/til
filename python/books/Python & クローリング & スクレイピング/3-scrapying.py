import requests
import bs4
from urllib.parse import urljoin
from functools import reduce
import time


def list_books(html, base_url):
    soup = bs4.BeautifulSoup(html, "html.parser")
    # links = soup.find_all('a', {"itemprop": "url"})
    links = soup.select('#listBook a[itemprop="url"]')  # CSS セレクタの場合は select
    hrefs = list(map(lambda x: urljoin(base_url, x.get("href")), links))
    return hrefs


def get_book_detail(html, base_url):
    soup = bs4.BeautifulSoup(html, "html.parser")
    title = soup.find("h1", {"id": "bookTitle"}).get_text()
    price = soup.find("span", {"class": "buy"}).get_text()
    content = reduce(lambda s, x: s + x,
                     [y.get_text() for y in soup.find_all("h3")])
    book = {
        "url": base_url,
        "title": title,
        "price": price,
        "content": content,
    }
    return book


def main():
    session = requests.Session()
    base_url = "https://gihyo.jp/dp"
    response = session.get(base_url)
    urls = list_books(response.content, base_url)
    for url in urls:
        time.sleep(1)
        resp = session.get(url)
        book = get_book_detail(resp.content, url)
        print(book)
        print("\n")


if __name__ == '__main__':
    main()


# TODO: db へ保存する
# response = requests.get(url)
# print(list_books(response.content))
# response = requests.get("https://gihyo.jp/dp/ebook/2019/978-4-297-10507-5")
# print(get_book_detail(response.content))
