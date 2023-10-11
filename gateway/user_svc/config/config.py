import yaml

with open('./user_svc/config.yml', 'r') as config_file:
    config = yaml.safe_load(config_file)

host = config['host']
grpc_port = config['grpc_port']
http_port = config['http_port']
