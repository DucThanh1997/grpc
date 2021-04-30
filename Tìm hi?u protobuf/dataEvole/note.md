- Không thay đổi tag
- add thêm field mới, và cái file proto cũ sẽ ko đọc nó (ignore đó)
- Xóa field nhưng tag number sẽ không được dùng nữa
- Không nên thay đổi kiểu dữ liệu

- Khi bạn xóa field
    + Nếu thằng code cũ không tìm thấy field đó, default value sẽ được sử dụng
    + Nếu đọc data cũ thì thằng field bị xóa sẽ được drop
    + reserved số tag và tên của field nếu xóa nó còn không rất dễ bị conflict