from consul import Consul
import requests

consul_url = "http://localhost:8500/v1/agent/service/register"

gateway = {
    "ID": "gateway",
    "Name": "gateway",
    "Address": "localhost",
    "Port": 5000,
}

def register_services():
    requests.put(consul_url, json=gateway)

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