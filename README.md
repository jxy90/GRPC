# GRPC
- part1-protobuf 
  - 使用grpc
    - protoc --csharp_out=csharp *.proto
    - protoc --proto_path src/ --go_out=src/ src/first/person.proto 
    - protoc --proto_path src/ --go_out=src/ src/second/enum.proto 
- prat2-grpc-server
  - protoc --proto_path=./protos ./protos/*.proto --go_out=plugins=grpc:./
  - openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj /CN=localhost
  - 