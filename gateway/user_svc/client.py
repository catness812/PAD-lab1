import grpc
from flask import request, jsonify
from user_svc.routes import register_user, delete_user
from utils.service_discovery import GRPCServerDiscovery
from utils.circuit_breaker import get_circuit_breaker
from utils.mongo import ping_mongo, connect_to_mongo
from utils.postgres import ping_postgres, connect_to_postgres

def start_user_client():
    user_server_discovery = GRPCServerDiscovery("user-grpc-svc")
    grpc_server_address = user_server_discovery.get_next_server()
    if grpc_server_address:
        return grpc.insecure_channel(grpc_server_address)
    else:
        raise RuntimeError("User gRPC server not found")

def register_user_handler():
    try:
        user_channel = start_user_client()
        data = request.json

        if not get_circuit_breaker(user_channel).allow_request():
            return jsonify({"message": "Circuit breaker open. Service is not available."}), 503

        res = register_user(user_channel, data)
        get_circuit_breaker(user_channel).record_success()
        return jsonify({"message": res.message})

    except RuntimeError:
        pass

    except Exception:
        get_circuit_breaker(user_channel).record_failure()
        register_user_handler()
    
def delete_user_handler():
    try:
        postgres_url = "postgres://postgres:pass@journaling-app-postgres:5432/journaling-app-db"
        postgres_conn = connect_to_postgres(postgres_url)

        if postgres_conn:
            if not ping_postgres(postgres_conn):
                return jsonify({"message": "PostgreSQL is not accessible."}), 503
        else:
            return jsonify({"message": "Unable to connect to PostgreSQL."}), 503

        mongo_url = "mongodb://root:root@journaling-app-mongo:27017"
        mongo_client = connect_to_mongo(mongo_url)

        if mongo_client:
            if not ping_mongo(mongo_client):
                return jsonify({"message": "MongoDB is not accessible."}), 503
        else:
            return jsonify({"message": "Unable to connect to MongoDB."}), 503

        user_channel = start_user_client()
        data = request.json

        if not get_circuit_breaker(user_channel).allow_request():
            return jsonify({"message": "Circuit breaker open. Service is not available."}), 503

        res = delete_user(user_channel, data)
        get_circuit_breaker(user_channel).record_success()
        return jsonify({"message": res.message})
    
    except RuntimeError:
        pass

    except Exception as e:
        print(f"Unhandled exception: {e}")
        get_circuit_breaker(user_channel).record_failure()
        delete_user_handler()
    