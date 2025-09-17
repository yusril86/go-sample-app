# Go Sample App

## Deskripsi

Go Sample App adalah aplikasi web sederhana yang dibangun menggunakan bahasa pemrograman Go atau native. Proyek ini menggunakan library bawaan Go seperti `net/http` untuk menangani permintaan HTTP dan `database/sql` untuk mengakses database.

## Fitur

- Routing HTTP dasar menggunakan `net/http`.
- Koneksi ke database menggunakan `database/sql`.
- Struktur kode yang sederhana dan mudah dipahami.

## Prasyarat

- Go versi 1.18 atau lebih baru harus terinstal di sistem Anda.
- Database (misalnya MySQL atau PostgreSQL) harus dikonfigurasi jika aplikasi menggunakan database.

## Instalasi dan Menjalankan Aplikasi

1. Clone repositori ini:

   ```bash
   git clone https://github.com/username/go-sample-app.git

   ```

2. Masuk ke direktori proyek:

```bash
  cd go-sample-app
```

3. Konfigurasi file .env (jika diperlukan) untuk pengaturan database dan variabel lingkungan lainnya.
4. Jalankan aplikasi

```bash
   go run main.go
```

5. Akses aplikasi melalui browser di http://localhost:8080.

**Struktur Proyek**:

```bash
go-sample-app/
├── main.go # Entry point aplikasi
├── controller/ # Handler untuk permintaan HTTP dan business logic
├── database/ # Init  database
├── routes/ # routing
└── README.md # Dokumentasi proyek
```
