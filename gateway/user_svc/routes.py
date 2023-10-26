from user_svc.pb import user_svc_pb2 as pb2
from user_svc.pb import user_svc_pb2_grpc as pb2_grpc

def register_user(user_channel, data):
    username = data.get('username', '')
    password = data.get('password', '')

    stub = pb2_grpc.UserServiceStub(user_channel)
    user = pb2.User(username=username, password=password)
    req = pb2.RegisterUserRequest(user=user)
    
    return stub.RegisterUser(req)

def delete_user(user_channel, data):
    username = data.get('username', '')
    password = data.get('password', '')

    stub = pb2_grpc.UserServiceStub(user_channel)
    user = pb2.User(username=username, password=password)
    
    return stub.DeleteUser(user)

