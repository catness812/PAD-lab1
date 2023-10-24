import grpc
from flask import request, jsonify
from user_svc.routes import register_user, delete_user
from utils.service_discovery import discover_grpc_server

def start_client():
    grpc_server_address = discover_grpc_server("user-grpc-svc")

    if grpc_server_address:
        return grpc.insecure_channel(grpc_server_address)
    else:
        raise RuntimeError("gRPC server not found")

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

def delete_user_handler():
    try: 
        data = request.json
        res = delete_user(channel, data)
        return jsonify({
            "message": res.message
        })
    except RuntimeError:
        pass