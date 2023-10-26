import grpc
from flask import request, jsonify
from journal_svc.routes import register_entry, get_user_entries
from utils.service_discovery import GRPCServerDiscovery
from utils.circuit_breaker import circuit_breaker
import pybreaker
import time

def start_journal_client():
    journal_server_discovery = GRPCServerDiscovery("journal-grpc-svc")
    grpc_server_address = journal_server_discovery.get_next_server()
    if grpc_server_address:
        return grpc.insecure_channel(grpc_server_address)
    else:
        raise RuntimeError("Journal gRPC server not found")
    
time.sleep(10)
journal_channel = start_journal_client()

@circuit_breaker
def register_entry_handler():
    try: 
        data = request.json
        res = register_entry(journal_channel, data)
        return jsonify({
            "message": res.message
        })
    except RuntimeError:
        pass
    except pybreaker.CircuitBreakerError:
        return jsonify({"error": "Service is currently unavailable."}), 503

@circuit_breaker
def get_user_entries_handler(user):
    try: 
        res = get_user_entries(journal_channel, user)
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
    except pybreaker.CircuitBreakerError:
        return jsonify({"error": "Service is currently unavailable."}), 503
