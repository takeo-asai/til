import scrapy

"""
$ scrapy runspider scrapy-yahoo.py -o items.jl 
"""


class YahooNewsSpider(scrapy.Spider):
    name = 'yahoospider'
    allowed_domains = ['www.yahoo.co.jp']
    start_urls = ['https://www.yahoo.co.jp/']

    def parse(self, response):
        for link in response.css('tr > td > a'):
            if link.css('a::attr("href")').extract()[0].count('https://news.yahoo.co.jp/pickup'):
                yield {'title': link.css('a::text').extract()}
