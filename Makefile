grpc: grpc/users/usermgt.proto
	rm -rf grpc/users/pb/*
	cd grpc/users && protoc --go_out=pb --go_opt=paths=source_relative --go-grpc_out=pb --go-grpc_opt=paths=source_relative usermgt.proto


# build commands
copymod: go.mod
	cp -R common/ grpc/ server/mod
	cp go.mod server/mod/
	cp -R common/ grpc/ usermgt/mod
	cp go.mod usermgt/mod/


dockserver: grpc copymod
	cd ./server && \
	docker build -t apiserver:latest .

dockusermgt: grpc
	go build -o bin/usermgt usermgt/cmd/main.go

build: buildserver buildusermgt


# run commands
runserver: dockserver
	./bin/server

runusermgt: dockusermgt
	./bin/usermgt

run: runserver runusermgt

# clean commands
clean:
	rm -rf server/mod/*
	rm -rf usermgt/mod/*