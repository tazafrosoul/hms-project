grpc:
	rm -rf grpc/users/pb/*
	cd grpc/users && protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative usermgt.proto


# build commands
buildserver: grpc
	go build -o bin/server server/cmd/main.go

buildusermgt: grpc
	go build -o bin/usermgt usermgt/cmd/main.go

build: buildserver buildusermgt


# run commands
runserver: buildserver
	alacritty -e ./bin/server &

runusermgt: buildusermgt
	alacritty -e ./bin/usermgt

run: runserver runusermgt

# clean commands
clean:
	rm -rf bin/*