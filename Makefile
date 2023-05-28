userproto:
	cd usermgt/internal/transport/grpc && protoc --go_out=usersgrpc --go_opt=paths=source_relative --go-grpc_out=usersgrpc --go-grpc_opt=paths=source_relative usermgt.proto
	# protoc --go_out=usersgrpc --go-grpc_out=usersgrpc usermgt.proto