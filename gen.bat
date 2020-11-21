cd models/protos

protoc --micro_out=../ --go_out=../ products.proto

cd .. && cd ..