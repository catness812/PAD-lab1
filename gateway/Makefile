proto_user:
	python -m grpc_tools.protoc -I ./user_svc/proto --python_out=./user_svc/pb --grpc_python_out=./user_svc/pb ./user_svc/proto/user_svc.proto
	
proto_journal:
	python -m grpc_tools.protoc -I ./journal_svc/proto --python_out=./journal_svc/pb --grpc_python_out=./journal_svc/pb ./journal_svc/proto/journal_svc.proto