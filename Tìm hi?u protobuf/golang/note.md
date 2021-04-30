- Gen code từ file proto sang golang thì như này
```
protoc -I src/ --go_out=src/ src/simple/simple.proto
```

Với 
    + `protoc` là compiler
    + `-I src/` định nghĩa cái root để tìm file
    + `--go_out=src/` đặt đầu ra của file ở đâu
    + `src/simple/simple.proto` full path của file
    