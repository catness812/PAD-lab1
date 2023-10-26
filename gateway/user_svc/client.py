import grpc
from flask import request, jsonify
from user_svc.routes import register_user, delete_user
from utils.service_discovery import GRPCServerDiscovery
from utils.circuit_breaker import circuit_breaker
import pybreaker
import time

def start_user_client():
    user_server_discovery = GRPCServerDiscovery("user-grpc-svc")
    grpc_server_address = user_server_discovery.get_next_server()
    if grpc_server_address:
        return grpc.insecure_channel(grpc_server_address)
    else:
        raise RuntimeError("User gRPC server not found")
    
time.sleep(10)
user_channel = start_user_client()

@circuit_breaker
def register_user_handler():
    try: 
        data = request.json
        res = register_user(user_channel, data)
        return jsonify({
            "message": res.message
        })
    except RuntimeError:
        pass
    except pybreaker.CircuitBreakerError:
        return jsonify({"error": "Service is currently unavailable."}), 503

@circuit_breaker
def delete_user_handler():
    try: 
        data = request.json
        res = delete_user(user_channel, data)
        return jsonify({
            "message": res.message
        })
    except RuntimeError:
        pass
    except pybreaker.CircuitBreakerError:
        return jsonify({"error": "Service is currently unavailable."}), 503