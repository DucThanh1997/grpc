protoc -I src/ --go_out=src/ src/simple/simple.proto
protoc -I complex/ --go_out=complex/ complex/complex.proto