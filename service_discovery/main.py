import subprocess
import os

consul_process = None
consul_started = False

def start_consul_dev():
    global consul_process, consul_started
    if not consul_started:
        try:
            consul_process = subprocess.Popen(['consul', 'agent', '-dev'])
            consul_started = True
            print("Consul agent started in development mode.")
        except FileNotFoundError:
            print("Consul executable not found. Make sure Consul is installed and in your system's PATH.")

def stop_consul():
    global consul_process
    if consul_started and consul_process is not None:
        try:
            subprocess.run(['consul', 'leave'], check=True)
            print("Consul agent stopped.")
        except FileNotFoundError:
            print("Consul executable not found. Make sure Consul is installed and in your system's PATH.")
        except subprocess.CalledProcessError:
            os.exit(0)
    else:
        print("Consul process was not started.")

if __name__ == "__main__":
    while True:
        try:
            start_consul_dev()
        except KeyboardInterrupt:
            stop_consul()
