## Setup

### Prerequisites
- Go 1.20 or newer
- [Docker](https://www.docker.com/) (optional, untuk menjalankan database dengan container)
- Git

### 1. Clone Repository
```bash
git clone https://github.com/MekayaStudio/sims.git
cd sims
```

### 2. Install Dependencies
```bash
go mod tidy
```

### 3. Konfigurasi Environment
Salin file `.env.example` menjadi `.env` lalu sesuaikan konfigurasi database dan lainnya:
```bash
cp .env.example .env
```

### 4. Jalankan Database (Opsional, jika pakai Docker)
```bash
docker-compose up -d
```

### 5. Generate APP_KEY
```bash
go run . artisan key:generate
```

### 6. Migrasi & Seeder Database
```bash
go run . artisan migrate
# Jika ingin menambahkan data awal:
go run . artisan db:seed
```

### 7. Jalankan Aplikasi
```bash
go run .
```