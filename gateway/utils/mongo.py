import pymongo
from pymongo import MongoClient

def connect_to_mongo(database_url):
    try:
        client = MongoClient(database_url)
        return client
    except Exception as e:
        print(f"Unable to connect to MongoDB. Error: {e}")
        return None

def ping_mongo(client):
    try:
        client.admin.command('ping')
        return True
    except Exception as e:
        print(f"Error pinging MongoDB: {e}")
    return False

mongo_url = "mongodb://root:root@journaling-app-mongo:27017"
mongo_client = connect_to_mongo(mongo_url)

if mongo_client:
    if ping_mongo(mongo_client):
        print("MongoDB is accessible.")
    else:
        print("Failed to ping MongoDB.")
    mongo_client.close()
else:
    print("Unable to connect to MongoDB.")
