from redis import Redis

redis_client = Redis(host='journaling-app-redis', port=6379, db=0)