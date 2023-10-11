import grpc
from flask import request, jsonify
from journal_svc.routes import register_entry
from service_discovery import discover_grpc_server

def start_client():
    grpc_server_address = discover_grpc_server("journal-grpc-service")

    if grpc_server_address:
        return grpc.insecure_channel(grpc_server_address)
    else:
        raise RuntimeError("gRPC server not found")

channel = start_client()

def register_entry_handler():
    try: 
        data = request.json
        res = register_entry(channel, data)
        return jsonify({
            "message": res.message
        })
    except RuntimeError:
        pass