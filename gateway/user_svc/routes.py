from user_svc.pb import user_svc_pb2 as pb2
from user_svc.pb import user_svc_pb2_grpc as pb2_grpc

def register_user(channel, data):
    username = data.get('username', '')
    password = data.get('password', '')

    stub = pb2_grpc.UserServiceStub(channel)
    req = pb2.RegisterRequest(username=username, password=password)
    
    return stub.RegisterUser(req)

