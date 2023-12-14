import grpc
from flask import request, jsonify
from journal_svc.routes import register_entry, get_user_entries
from utils.service_discovery import GRPCServerDiscovery
from utils.circuit_breaker import get_circuit_breaker

def start_journal_client():
    journal_server_discovery = GRPCServerDiscovery("journal-grpc-svc")
    grpc_server_address = journal_server_discovery.get_next_server()
    if grpc_server_address:
        return grpc.insecure_channel(grpc_server_address)
    else:
        raise RuntimeError("Journal gRPC server not found")

def register_entry_handler():
    try: 
        journal_channel = start_journal_client()
        data = request.json

        if not get_circuit_breaker(journal_channel).allow_request():
            return jsonify({"message": "Circuit breaker open. Service is not available."}), 503

        res = register_entry(journal_channel, data)
        get_circuit_breaker(journal_channel).record_success()
        return jsonify({
            "message": res.message
        })
    
    except RuntimeError:
        pass

    except Exception:
        get_circuit_breaker(journal_channel).record_failure()
        register_entry_handler()

def get_user_entries_handler(user):
    try: 
        journal_channel = start_journal_client()

        if not get_circuit_breaker(journal_channel).allow_request():
            return jsonify({"message": "Circuit breaker open. Service is not available."}), 503

        res = get_user_entries(journal_channel, user)
        entries_list = []
        for entry in res.entries:
            entry_dict = {
                "title": entry.title,
                "content": entry.content
            }
            entries_list.append(entry_dict)

        get_circuit_breaker(journal_channel).record_success()

        return jsonify({
            "message": res.message,
            "entries": entries_list
        })
    
    except RuntimeError:
        pass

    except Exception:
        get_circuit_breaker(journal_channel).record_failure()
        get_user_entries_handler(user)