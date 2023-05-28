userspb:
	rm -rf grpc/users/pb/*
	cd grpc/users && protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative usermgt.proto