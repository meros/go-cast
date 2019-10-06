cast_channel.pb.go : cast_channel.proto
	protoc -I=. --go_out=. ./cast_channel.proto

all : cast_channel.pb.go

test :
	go test
