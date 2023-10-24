from flask import Flask
from utils.timeout import check_timeout
from config.config import host, http_port
from user_svc.client import register_user_handler, delete_user_handler
from journal_svc.client import register_entry_handler, get_user_entries_handler
from utils.service_discovery import register_services, check

register_services()

app = Flask(__name__)

@app.route('/health/<svc>', methods=['GET'])
def health(svc):
    return check(svc)

@app.route('/users/register', methods=['POST'])
def register_user_route():
    return check_timeout(register_user_handler)

@app.route('/entries/create', methods=['POST'])
def register_entry_route():
    return check_timeout(register_entry_handler)

@app.route('/users/delete', methods=['DELETE'])
def delete_user_route():
    return check_timeout(delete_user_handler)

@app.route('/entries/<user>', methods=['GET'])
def get_user_entries_route(user):
    return check_timeout(get_user_entries_handler, user)

if __name__ == '__main__':
    app.run(host=f'{host}', port=f'{http_port}', debug=True)