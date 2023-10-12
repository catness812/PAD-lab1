from flask import Flask
from util.timeout import check_timeout
from config.config import host, http_port
from user_svc.client import register_user_handler
from journal_svc.client import register_entry_handler
from service_discovery import register_services, check

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

if __name__ == '__main__':
    app.run(host=f'{host}', port=f'{http_port}', debug=True)