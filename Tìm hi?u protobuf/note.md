## Định nghĩa kiểu message
- Dưới đây là 1 ví dụ đơn giản, bạn muốn định nghĩa 1 search request message. Trong file .proto bạn sẽ viết như này

```
syntax = "proto3" // chỉ định bạn sẽ viết

message SearchRequest {
    string query = 1;
    int32 page_number = 2;
    int32 result_per_page = 3;
}
// Search Request định nghĩa 3 trường (tên và value)

// Mỗi 1 trường có 1 con số unique. Các con số này dùng để identify các trường của bạn trong `message binary format`, và không nên thay đổi khi file của bạn đã đưa vào sử dụng
```

## Các luật của field

Các trường của message có thể là
- Singular: chỉ có 0 hoặc 1 trường
- Repeated: các trường này được lặp lại nhiều lần kiểu array

## Reversed Fields
Nếu bạn update message bằng cách xóa 1 filed hoặc comment nó lại, user có thể dùng lại các số đó cho những trường mà họ muốn update sau này. Tuy nhiên việc này có thể gây ra 1 đống rắc rối khi bạn load 1 phiên bản cũ của .proto (conflict). Để khắc phục bạn nên điền những field number là reserved. protocol buffer sẽ complain nếu có ai đó sử dụng số này trong tương lai

```
message Foo {
    reserved 2, 15, 9 to 11;
    reserved "foo", "bar";
}
```

## Enumerations
