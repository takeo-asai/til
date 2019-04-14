# -*- coding: utf-8 -*-
import scrapy


class NewsSpider(scrapy.Spider):
    name = 'news'
    allowed_domains = ['news.yahoo.co.jp']
    start_urls = ['http://news.yahoo.co.jp/']

    def parse(self, response):
        for url in response.css('.topics a::attr("href")').re(r'pickup/\d+$'):
            yield scrapy.Request(response.urljoin(url), self.parse_topics)

    def parse_topics(self, response):
        title = response.css('title::text').extract_first()
        url = response.url
        yield {'title': title, 'url': url}
