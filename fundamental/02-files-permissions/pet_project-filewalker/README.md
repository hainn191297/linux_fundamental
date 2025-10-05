# Pet Project: filewalker

Goal: duyệt một thư mục (mặc định: `.`), in ra đường dẫn, inode, quyền và kích thước.

## Chạy thử

```bash
cd fundamental/02-files-permissions/pet_project-filewalker
go run .

# Chỉ định thư mục khác
go run . /tmp
```

## Gợi ý mở rộng

- Bộ lọc theo mẫu tên và phần mở rộng.
- Đếm tổng số file, tổng size, và top-N file lớn nhất.
- Ghi kết quả ra JSON Line (`.jsonl`).

