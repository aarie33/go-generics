# Go Generics
- versi go 1.18
- go get github.com/stretchr/testify

- Biasanya tipe karakter menggunakan dengan 1 huruf saja. ex: T
- Wajib memiliki `constraint`

### Comparable
- Tipe data yang bisa dibandingkan

### Type Parameter Inheritance
- Golang sebenarnya tidak mendukung inheritance, tapi bisa menggunakan interface

### Type Set
- adalah sebuah interface

### Type Aproximation
- use ~ to approximate a type

### Type Inference
- Go Generics akan menginfer tipe data dari parameter yang diberikan

### Generic Type
- Tipe data yang bisa digunakan untuk tipe data lain

### Generic Struct
- Struct yang bisa digunakan untuk tipe data lain

### Generic Interface
- Generic juga dapat digunakan untuk interface

### In Line Type Constraint
- jika mau menggantikan `any` bisa menggunakan `int | string` dsb

### Experimental Package
- beberapa package yang memungkinkan untuk menggunakan generics
- `go get golang.org/x/exp/...`
- isinya belum stabil
- isinya termasuk :
  - constraint package
  - maps & slices package