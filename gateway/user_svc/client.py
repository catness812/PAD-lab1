import grpc
from user_svc.config.config import host, grpc_port
from flask import request, jsonify
from user_svc.routes import register_user

def start_client():
    return grpc.insecure_channel(f'{host}:{grpc_port}')

channel = start_client()

def register_user_handler():
    try: 
        data = request.json
        res = register_user(channel, data)
        return jsonify({
            "message": res.message
        })
    except RuntimeError:
        pass