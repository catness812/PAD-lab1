from flask import Flask, request
from utils.timeout import check_timeout
from config.config import host, http_port
from user_svc.client import register_user_handler, delete_user_handler
from journal_svc.client import register_entry_handler, get_user_entries_handler
from utils.service_discovery import register_services, check
from cache.redis import redis_client
from prometheus_flask_exporter import PrometheusMetrics

register_services()

app = Flask(__name__)
metrics = PrometheusMetrics(app)

@app.route('/health/<svc>', methods=['GET'])
def health(svc):
    return check(svc)

@app.route('/users/register', methods=['POST'])
def register_user_route():       
    return check_timeout(register_user_handler)

@app.route('/entries/create', methods=['POST'])
def register_entry_route(): 
    response = check_timeout(register_entry_handler)
        
    user = request.json.get('username')
    if user is not None:
        cache_key = f'entries:{user}'
        redis_client.delete(cache_key)
        
    return response

@app.route('/users/delete', methods=['DELETE'])
def delete_user_route():
    return check_timeout(delete_user_handler)
    
@app.route('/entries/<user>', methods=['GET'])
def get_user_entries_route(user):
    cache_key = f'entries:{user}'
    cached_data = redis_client.get(cache_key)
    if cached_data is not None:
        return cached_data.decode('utf-8')
            
    data = check_timeout(get_user_entries_handler, user)
    redis_client.set(cache_key, str(data))
            
    return data

if __name__ == '__main__':
    app.run(host=f'{host}', port=f'{http_port}', debug=True)