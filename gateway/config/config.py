import yaml

with open('./config.yml', 'r') as config_file:
    config = yaml.safe_load(config_file)

host = config['host']
http_port = config['http_port']
