# -*- coding: utf-8 -*-
from scrapy.spiders import CrawlSpider, Rule
from scrapy.linkextractors import LinkExtractor


class NewsCrawlSpider(CrawlSpider):
    name = 'news_crawl'
    allowed_domains = ['news.yahoo.co.jp']
    start_urls = ['http://news.yahoo.co.jp/']
    # 複数が定義でき、上から順に適用される
    rules = (
        Rule(LinkExtractor(allow=r'pickup/\d+$'), callback='parse_topics'),
    )

    def parse_topics(self, response):
        title = response.css('title::text').extract_first()
        url = response.url
        yield {'title': title, 'url': url}
