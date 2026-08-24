# Inventory Apps

Sistem informasi inventaris multi-gudang. Setiap perubahan stok — masuk, keluar,
penyesuaian, maupun transfer antar gudang — tercatat sebagai jejak yang bisa
ditelusuri, bukan sekadar angka yang berubah diam-diam.

Backend REST API dengan **Go** (standard library, tanpa framework) dan frontend
**React + Vite + Tailwind CSS**.

---

## Daftar Isi

- [Fitur](#fitur)
- [Teknologi](#teknologi)
- [Arsitektur](#arsitektur)
- [Konsep Inti](#konsep-inti)
- [Instalasi](#instalasi)
- [Struktur Folder](#struktur-folder)
- [Endpoint API](#endpoint-api)
- [Keamanan](#keamanan)
- [Batasan yang Diketahui](#batasan-yang-diketahui)
- [Rencana Pengembangan](#rencana-pengembangan)

---

## Fitur

### Dashboard
- **Overview** — jumlah barang, total unit, nilai stok, dan grafik pergerakan
  stok dengan rentang tanggal yang bisa dipilih
- **Stok Menipis** — barang yang menyentuh atau di bawah batas minimum, per gudang
- **Aktivitas Terbaru** — riwayat pergerakan stok terakhir beserta pelakunya

### Master Data
Barang, Kategori, Satuan, Supplier, dan Lokasi/Gudang.

### Transaksi
| Transaksi | Efek pada stok |
|---|---|
| Barang Masuk | menambah stok gudang tujuan, mencatat supplier dan harga beli |
| Barang Keluar | mengurangi stok, ditolak kalau stok tidak mencukupi |
| Penyesuaian Stok | menyamakan stok sistem dengan hasil hitung fisik |
| Transfer Gudang | memindahkan stok antar gudang, tercatat sebagai dua pergerakan |

### Laporan
Laporan Stok, Kartu Stok, Laporan Barang Masuk, Laporan Barang Keluar, dan
Laporan Pergerakan Stok (saldo awal → mutasi → saldo akhir). Semuanya bisa
difilter per periode dan per gudang, serta bisa dicetak.

---

## Teknologi

**Backend**
- Go 1.21+ dengan `net/http` (tanpa framework)
- MySQL 8
- `golang-jwt/jwt/v5` untuk JWT
- `golang.org/x/crypto/bcrypt` untuk hashing password
- `golang.org/x/time/rate` untuk rate limiting

**Frontend**
- React 18 + Vite
- React Router
- Tailwind CSS
- Axios
- Lucide React (ikon)

---

## Arsitektur

Backend memakai layered architecture. Setiap modul punya file sendiri di tiap
lapisan, jadi menambah fitur baru tidak menyentuh kode modul lain.

```
Request
   ↓
Recover  →  Security Headers  →  Rate Limit  →  CORS
   ↓
JWT Authentication  (401 kalau token tidak valid)
   ↓
Permission / RBAC   (403 kalau tidak berhak)
   ↓
Handler      parsing request & format response
   ↓
Service      validasi & aturan bisnis
   ↓
Repository   query database
   ↓
MySQL
```

Frontend tidak pernah menghitung atau mengubah stok. React hanya mengirim niat
transaksi; Go yang memvalidasi dan memutuskan.

---

## Konsep Inti

### Stok milik gudang, bukan milik barang

Stok disimpan per pasangan (gudang, barang) di tabel `warehouse_stocks`, bukan
sebagai satu angka di tabel `barang`. Tanpa ini, sistem tidak bisa menjawab
"ada berapa Keyboard di Gudang Cabang?".

### Satu pintu untuk mengubah stok

Seluruh transaksi bermuara ke satu fungsi, `applyMovement()`, yang dijalankan di
dalam database transaction. Fungsi itu yang mengunci baris stok, mengecek
kecukupan, memperbarui saldo, dan mencatat pergerakan. Kalau satu langkah gagal,
seluruh transaksi rollback — tidak ada stok yang tersimpan setengah jalan.

### Semua pergerakan tercatat

Tabel `stock_movements` menyimpan setiap perubahan lengkap dengan `stock_before`
dan `stock_after`. Dari situ Kartu Stok bisa dibuat tanpa menghitung ulang
riwayat, dan setiap angka stok bisa dipertanggungjawabkan asal-usulnya.

Nilai `quantity` disimpan bertanda: positif menambah, negatif mengurangi.
Konsekuensinya `SUM(quantity)` selalu sama dengan stok berjalan.

```
                     PRODUK
                        │
                 WAREHOUSE_STOCKS
                        │
      ┌─────────────────┼─────────────────┐
      ▼                 ▼                 ▼
  BARANG MASUK    BARANG KELUAR     PENYESUAIAN
      │                 │                 │
      └─────────────────┼─────────────────┘
                        ▼
                 STOCK_MOVEMENTS
                        ▲
                        │
                    TRANSFER
                   ↙        ↘
            Gudang A        Gudang B
          TRANSFER_OUT    TRANSFER_IN
```

---

## Instalasi

### Prasyarat
- Go 1.21 atau lebih baru
- Node.js 18 atau lebih baru
- MySQL 8

### 1. Clone

```bash
git clone https://github.com/<username>/belajar_go.git
cd belajar_go
```

> Ganti juga module path di `backend/go.mod` dan seluruh import
> `github.com/username/belajar_go` sesuai username GitHub kamu.

### 2. Database

Buat database, lalu jalankan migrasi secara berurutan:

```sql
CREATE DATABASE belajar_go CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
```

```
database/migration_transaksi.sql     -- warehouse_stocks, stock_movements, 4 transaksi
database/migration_dashboard.sql     -- kolom stok_minimum
database/migration_auth.sql          -- roles, permissions, refresh_tokens
```

Migrasi auth otomatis membuat role `admin` dengan seluruh permission dan
menjadikan user lama sebagai admin.

### 3. Backend

```bash
cd backend
cp .env.example .env
```

Isi `JWT_SECRET` minimal 32 karakter — server sengaja menolak jalan tanpanya:

```bash
openssl rand -base64 48
```

Lalu:

```bash
go mod tidy
go run ./cmd/api
```

Server berjalan di `http://localhost:8080`.

### 4. Frontend

```bash
cd frontend
npm install
npm run dev
```

Aplikasi berjalan di `http://localhost:5173`.

Pastikan `vite.config.js` mem-proxy `/api` ke backend:

```js
server: {
  proxy: { '/api': 'http://localhost:8080' }
}
```

---

## Struktur Folder

```
backend/
├── cmd/api/main.go              wiring & routing
├── config/                      env & koneksi database
└── internal/
    ├── auth/                    JWT, bcrypt, token acak
    ├── middleware/              CORS, rate limit, JWT, RBAC, security headers
    ├── model/                   struct & tag JSON
    ├── repository/              query database
    ├── service/                 validasi & aturan bisnis
    └── handler/                 HTTP handler

frontend/src/
├── api.js                       axios interceptor & penyimpanan token
├── App.jsx                      routing & state autentikasi
├── components/                  Layout, Sidebar, FilterBar
└── pages/                       seluruh halaman

database/                        file migrasi SQL
```

---

## Endpoint API

### Autentikasi

| Method | Endpoint | Keterangan |
|---|---|---|
| POST | `/api/v1/auth/login` | menukar kredensial jadi access token + cookie refresh |
| POST | `/api/v1/auth/refresh` | rotasi refresh token, terbitkan access token baru |
| POST | `/api/v1/auth/logout` | mencabut seluruh sesi |

### Master Data

`/api/barang`, `/api/kategori`, `/api/satuan`, `/api/supplier`, `/api/lokasi`
— masing-masing mendukung `GET`, `POST`, dan `GET`/`PUT`/`DELETE` dengan `/{id}`.

### Transaksi

| Method | Endpoint |
|---|---|
| GET, POST | `/api/stock-in`, `/api/stock-out`, `/api/stock-adjustment`, `/api/stock-transfer` |
| GET | endpoint yang sama dengan `/{id}` untuk detail |
| GET | `/api/warehouse-stocks?lokasi_id=&barang_id=` |
| GET | `/api/stock-movements?lokasi_id=&barang_id=&limit=` |

### Laporan & Dashboard

| Method | Endpoint |
|---|---|
| GET | `/api/laporan/stok?lokasi_id=&kategori_id=` |
| GET | `/api/laporan/kartu-stok?barang_id=&lokasi_id=&start=&end=` |
| GET | `/api/laporan/barang-masuk?start=&end=&lokasi_id=&supplier_id=` |
| GET | `/api/laporan/barang-keluar?start=&end=&lokasi_id=` |
| GET | `/api/laporan/pergerakan?start=&end=&lokasi_id=` |
| GET | `/api/dashboard/overview?start=&end=` |
| GET | `/api/dashboard/stok-menipis?lokasi_id=&minimum=&limit=` |
| GET | `/api/dashboard/aktivitas?limit=` |

Semua endpoint selain `/api/health` dan `/api/v1/auth/*` membutuhkan header
`Authorization: Bearer <access_token>`.

---

## Keamanan

- **Access token** berumur 15 menit, disimpan di memori JavaScript — bukan
  `localStorage`
- **Refresh token** berumur 7 hari, dikirim sebagai cookie `HttpOnly` + `SameSite`,
  dan hanya hash SHA-256-nya yang disimpan di database
- **Refresh token rotation dengan reuse detection** — token yang dipakai dua kali
  dianggap dicuri, dan seluruh rantai sesinya langsung dicabut
- **RBAC** — autentikasi (401) dan otorisasi (403) ditangani middleware terpisah,
  dengan permission berpola `resource.action`
- **Password** di-hash bcrypt cost 12
- **Rate limiting** per IP, lebih ketat pada endpoint login dan refresh
- **CORS** hanya menerima origin dari daftar putih, tidak pernah `*`
- **Validasi input** dilakukan di backend, tidak bergantung pada React
- Seluruh query memakai prepared statement
- Error database dan stack trace tidak pernah dikirim ke client

Secret tidak pernah ditulis di source code — semuanya lewat environment variable.

---

## Batasan yang Diketahui

Beberapa hal sengaja belum dikerjakan, dan sebaiknya diketahui sebelum dipakai
di lingkungan nyata:

- **Belum ada pembatalan transaksi.** Kolom `status` selalu berisi `POSTED`.
  Kesalahan input dikoreksi lewat Penyesuaian Stok, bukan dihapus.
- **Belum ada harga pokok (HPP).** Nilai persediaan dihitung dari `barang.harga`.
  Untuk akuntansi persediaan yang benar, perlu rata-rata bergerak atau FIFO.
- **Rate limiter disimpan di memori** — hilang saat restart dan tidak berlaku
  lintas instance. Untuk produksi berskala, pindahkan ke Redis.
- **Access token tidak bisa dicabut** sebelum kedaluwarsa. Logout mencabut
  refresh token; access token yang sudah terbit tetap sah maksimal 15 menit.
- **Batas stok minimum berlaku per barang**, belum bisa berbeda per gudang.
- **Belum ada HTTPS.** Wajib dijalankan di belakang reverse proxy ber-TLS saat
  produksi, karena cookie `Secure` tidak akan terkirim lewat HTTP.

---

## Rencana Pengembangan

- [ ] Alur draft → posting → void untuk transaksi
- [ ] Harga pokok rata-rata bergerak
- [ ] Role tambahan (staff gudang, viewer)
- [ ] Ekspor laporan ke Excel dan PDF
- [ ] Batas stok minimum per gudang
- [ ] Audit log perubahan master data
- [ ] Barcode / QR untuk pencarian barang

---

## Lisensi

MIT
