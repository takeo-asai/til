from pymongo import MongoClient

client = MongoClient()
db = client['scraping-book']
collection = db['items']

for post in collection.find():
    print(post)
