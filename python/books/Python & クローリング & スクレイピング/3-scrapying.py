import requests
import bs4
from urllib.parse import urljoin
from functools import reduce
import time
from pymongo import MongoClient


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


def store_book(collection, book):
    collection.insert_one(book)


def find_book(collection, url):
    return collection.find_one({"url": url})


def main():
    client = MongoClient()  # localhost:27017
    collection = client.scrapying.books
    collection.create_index("url", unique=True)
    base_url = "https://gihyo.jp/dp"
    session = requests.Session()
    response = session.get(base_url)
    urls = list_books(response.content, base_url)
    for url in urls:
        if not find_book(collection, url):
            time.sleep(1)
            resp = session.get(url)
            book = get_book_detail(resp.content, url)
            store_book(collection, book)
            print(book)
            print("\n")


if __name__ == '__main__':
    main()
