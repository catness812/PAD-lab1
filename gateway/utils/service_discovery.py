import requests

consul_url = "http://consul:8500/v1/agent/service/register"

gateway = {
    "ID": "gateway",
    "Name": "gateway",
    "Address": "localhost",
    "Port": 5000,
}

def register_services():
    requests.put(consul_url, json=gateway)


def check(svc):
    response = requests.get(f"http://consul:8500/v1/health/service/{svc}")
    return response.json()

class GRPCServerDiscovery:
    def __init__(self, service_name):
        self.service_name = service_name
        self.servers = []
        self.index = 0
    
    def discover_grpc_servers(self):
        url = 'http://consul:8500/v1/agent/services'
        response = requests.get(url)
        if response.status_code == 200:
            services = response.json()
            for service in services.values():
                if service['Service'] == self.service_name:
                    address = f"{service['Address']}:{service['Port']}"
                    self.servers.append(address)
            if not self.servers:
                raise RuntimeError(f"No gRPC servers found for service {self.service_name}")
    
    def get_next_server(self):
        if not self.servers:
            self.discover_grpc_servers()
        server = self.servers[self.index]
        self.index = (self.index + 1) % len(self.servers)
        return server