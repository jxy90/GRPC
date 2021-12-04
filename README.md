# GRPC
- part1-protobuf 
  - 使用grpc
    - protoc --csharp_out=csharp *.proto
    - protoc --proto_path src/ --go_out=src/ src/first/person.proto 
    - protoc --proto_path src/ --go_out=src/ src/second/enum.proto 
- prat2-grpc-server
  - protoc --proto_path=./protos ./protos/*.proto --go_out=plugins=grpc:./
  - x509
    - openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj /CN=localhost


# 解决:transport: authentication handshake failed: x509: certificate relies on lega
## CA证书
### 生成.key  私钥文件
- $ openssl genrsa -out ca.key 2048
### 生成.csr 证书签名请求文件
- $ openssl req -new -key ca.key -out ca.csr  -subj "/C=GB/L=China/O=lixd/CN=www.lixueduan.com"
### 自签名生成.crt 证书文件
- $ openssl req -new -x509 -days 3650 -key ca.key -out ca.crt  -subj "/C=GB/L=China/O=lixd/CN=www.lixueduan.com"
## 服务端证书
### 生成.key  私钥文件
$ openssl genrsa -out server.key 2048

### 生成.csr 证书签名请求文件
$ openssl req -new -key server.key -out server.csr \
-subj "/C=GB/L=China/O=lixd/CN=www.lixueduan.com" \
-reqexts SAN \
-config <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.lixueduan.com,DNS:*.refersmoon.com"))

### 签名生成.crt 证书文件
$ openssl x509 -req -days 3650 \
-in server.csr -out server.crt \
-CA ca.crt -CAkey ca.key -CAcreateserial \
-extensions SAN \
-extfile <(cat /etc/ssl/openssl.cnf <(printf "\n[SAN]\nsubjectAltName=DNS:*.lixueduan.com,DNS:*.refersmoon.com"))
