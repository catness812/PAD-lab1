from journal_svc.pb import journal_svc_pb2 as pb2
from journal_svc.pb import journal_svc_pb2_grpc as pb2_grpc

def register_entry(channel, data):
    username = data.get('username', '')
    title = data.get('title', '')
    content = data.get('content', '')

    stub = pb2_grpc.JournalServiceStub(channel)
    entry = pb2.Entry(username=username, title=title, content=content)
    req = pb2.RegisterEntryRequest(entry=entry)
    
    return stub.RegisterEntry(req)

