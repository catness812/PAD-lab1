import requests
from consul import Consul

consul_url = "http://localhost:8500/v1/agent/service/register"

gateway = {
    "ID": "gateway",
    "Name": "gateway",
    "Address": "localhost",
    "Port": 5000,
}

svc1 = {
    "ID": "user-grpc-service",
    "Name": "user-grpc-service",
    "Address": "localhost",
    "Port": 50051,
}

svc2 = {
    "ID": "journal-grpc-service",
    "Name": "journal-grpc-service",
    "Address": "localhost",
    "Port": 50052,
}

def register_services():
    requests.put(consul_url, json=gateway)
    requests.put(consul_url, json=svc1)
    requests.put(consul_url, json=svc2)

consul = Consul()

def discover_grpc_server(service_name):
    services = consul.agent.services()
    grpc_service = services.get(service_name)
    if grpc_service:
        grpc_address = f"{grpc_service['Address']}:{grpc_service['Port']}"
        return grpc_address
    return None

def check(svc):
    response = requests.get(f"http://localhost:8500/v1/health/service/{svc}")
    return response.json()
