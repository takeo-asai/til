from pymongo import MongoClient

client = MongoClient()
db = client.test
collection = db.spots

collection.insert_one({'name': '東京スカイツリー', 'prefecture': '東京'})
collection.insert_many([{'name': '東京ディズニーランド', 'prefecture': '千葉'}, {
                       'name': '東京ドーム', 'prefecture': '東京'}])

for spot in collection.find():
    print(spot)

for spot in collection.find({"prefecture": "東京"}):
    print(spot)
