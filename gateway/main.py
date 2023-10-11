from flask import Flask
from user_svc.client import register_user_handler
from user_svc.util.timeout import check_timeout
from user_svc.config.config import host, http_port

app = Flask(__name__)

@app.route('/users/register', methods=['POST'])
def register_route():
    return check_timeout(register_user_handler)

if __name__ == '__main__':
    app.run(host=f'{host}', port=f'{http_port}', debug=True)
