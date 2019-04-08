import requests
import bs4
from urllib.parse import urljoin
from functools import reduce

url = "https://gihyo.jp/dp"


def list_books(html):
    soup = bs4.BeautifulSoup(html, "html.parser")
    links = soup.find_all('a', {"itemprop": "url"})
    hrefs = list(map(lambda x: urljoin(url, x.get("href")), links))
    return hrefs


def get_book_detail(html):
    soup = bs4.BeautifulSoup(html, "html.parser")
    title = soup.find("h1", {"id": "bookTitle"}).get_text()
    price = soup.find("span", {"class": "buy"}).get_text()
    content = reduce(lambda s, x: s + x,
                     [y.get_text() for y in soup.find_all("h3")])
    book = {
        "title": title,
        "price": price,
        "content": content,
    }
    return book


# TODO: db へ保存する
# response = requests.get(url)
# print(list_books(response.content))
response = requests.get("https://gihyo.jp/dp/ebook/2019/978-4-297-10507-5")
print(get_book_detail(response.content))
