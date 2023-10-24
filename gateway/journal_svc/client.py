import grpc
from flask import request, jsonify
from journal_svc.routes import register_entry, get_user_entries
from utils.service_discovery import discover_grpc_server

def start_client():
    grpc_server_address = discover_grpc_server("journal-grpc-svc")

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

def get_user_entries_handler(user):
    try: 
        res = get_user_entries(channel, user)
        entries_list = []
        for entry in res.entries:
            entry_dict = {
                "title": entry.title,
                "content": entry.content
            }
            entries_list.append(entry_dict)

        return jsonify({
            "message": res.message,
            "entries": entries_list
        })
    except RuntimeError:
        pass
