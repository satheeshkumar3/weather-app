

command to create proto file
## protoc --go_out=./generated --go-grpc_out=./generated --proto_path=./proto .\proto\weather.proto
protoc --go_out=./generated --go_opt=paths=source_relative --go-grpc_out=./generated  --go-grpc_opt=paths=source_relative --proto_path=./proto ./proto/weather.proto
