# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html

from pymongo import MongoClient


class SixScrapyPipeline(object):
    def process_item(self, item, spider):
        return item


class MongoPipeline(object):

    def open_spider(self, spider):
        self.client = MongoClient()
        self.db = self.client['scraping-book']
        self.collection = self.db['items']

    def close_spider(self, spider):
        self.client.close()

    def process_item(self, item, spider):
        # insert_one の引数は変更されるのでコピーを渡せとのこと
        self.collection.insert_one(dict(item))
        return item
