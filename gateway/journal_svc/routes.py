from journal_svc.pb import journal_svc_pb2 as pb2
from journal_svc.pb import journal_svc_pb2_grpc as pb2_grpc

def register_entry(journal_channel, data):
    username = data.get('username', '')
    title = data.get('title', '')
    content = data.get('content', '')

    stub = pb2_grpc.JournalServiceStub(journal_channel)
    entry = pb2.Entry(username=username, title=title, content=content)
    req = pb2.RegisterEntryRequest(entry=entry)
    
    return stub.RegisterEntry(req)

def get_user_entries(journal_channel, user):
    stub = pb2_grpc.JournalServiceStub(journal_channel)
    req = pb2.GetUserEntriesRequest(username=user)
    
    return stub.GetUserEntries(req)

