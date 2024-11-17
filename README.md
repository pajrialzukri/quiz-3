# Quiz-3 API

## Deskripsi

Quiz-3 API adalah aplikasi backend yang dibangun menggunakan **Go** dan **Gin Framework**. Aplikasi ini menyediakan API untuk mengelola **Users**, **Categories**, dan **Books**. API ini dilengkapi dengan autentikasi berbasis **JWT** dan dokumentasi API menggunakan **Swagger**.

## Fitur Utama

- **User Authentication**: Menggunakan JWT untuk autentikasi user.
- **CRUD Operations**: CRUD untuk **Users**, **Categories**, dan **Books**.
- **Protected Routes**: Beberapa endpoint dilindungi oleh middleware autentikasi JWT.
- **Swagger Documentation**: API dokumentasi otomatis dengan Swagger.

## Teknologi

- **Go** (Golang)
- **Gin Framework**: Untuk routing dan middleware.
- **JWT**: Untuk autentikasi token berbasis JSON.
- **PostgreSQL**: Sebagai database untuk menyimpan data pengguna, kategori, dan buku.
- **Swagger**: Untuk dokumentasi API.
- **Golang's bcrypt**: Untuk hashing password.
