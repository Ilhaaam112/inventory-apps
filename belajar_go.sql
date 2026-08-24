-- phpMyAdmin SQL Dump
-- version 5.2.1deb3
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Aug 24, 2026 at 11:28 PM
-- Server version: 8.0.46-0ubuntu0.24.04.3
-- PHP Version: 8.3.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `belajar_go`
--

-- --------------------------------------------------------

--
-- Table structure for table `barang`
--

CREATE TABLE `barang` (
  `id` int NOT NULL,
  `nama` varchar(100) NOT NULL,
  `harga` decimal(12,2) NOT NULL DEFAULT '0.00',
  `stok` int NOT NULL DEFAULT '0',
  `stok_minimum` int NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `kategori_id` int DEFAULT NULL,
  `satuan_id` int DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `barang`
--

INSERT INTO `barang` (`id`, `nama`, `harga`, `stok`, `stok_minimum`, `created_at`, `kategori_id`, `satuan_id`) VALUES
(1, 'Kabel USB-C', 20000.00, 500, 0, '2026-08-23 23:04:15', 1, 1),
(2, 'Mouse Wireless', 85000.00, 0, 0, '2026-08-23 23:04:15', 1, 1),
(3, 'Keyboard Mechanical', 350000.00, 500, 0, '2026-08-23 23:04:15', 1, 1),
(5, 'Connector RJ45', 2500.00, 500, 500, '2026-08-23 23:59:49', 4, 1),
(6, 'Mouse', 65.00, 0, 10, '2026-08-24 05:12:04', 3, 1);

-- --------------------------------------------------------

--
-- Table structure for table `kategori`
--

CREATE TABLE `kategori` (
  `id` int NOT NULL,
  `nama_kategori` varchar(100) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `kategori`
--

INSERT INTO `kategori` (`id`, `nama_kategori`, `created_at`) VALUES
(1, 'Aksesoris Komputer', '2026-08-23 23:45:51'),
(2, 'Peralatan Kantor', '2026-08-23 23:45:51'),
(3, 'Elektronik', '2026-08-23 23:45:51'),
(4, 'Peralatan Network', '2026-08-23 23:58:48'),
(5, 'Pelaratan Listrik', '2026-08-24 22:40:20');

-- --------------------------------------------------------

--
-- Table structure for table `lokasi`
--

CREATE TABLE `lokasi` (
  `id` int NOT NULL,
  `nama_lokasi` varchar(150) NOT NULL,
  `keterangan` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `lokasi`
--

INSERT INTO `lokasi` (`id`, `nama_lokasi`, `keterangan`, `created_at`) VALUES
(1, 'Gudang Utama', 'Gudang pusat penyimpanan', '2026-08-24 00:08:30'),
(2, 'Gudang Cabang Kota', 'Gudang cabang kota', '2026-08-24 00:08:30'),
(3, 'Gudang Cabang Desa', 'Gudang Cabang Desa', '2026-08-24 00:14:50');

-- --------------------------------------------------------

--
-- Table structure for table `permissions`
--

CREATE TABLE `permissions` (
  `id` int NOT NULL,
  `code` varchar(60) NOT NULL,
  `keterangan` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `permissions`
--

INSERT INTO `permissions` (`id`, `code`, `keterangan`) VALUES
(1, 'barang.read', 'Lihat data barang'),
(2, 'barang.create', 'Tambah barang'),
(3, 'barang.update', 'Ubah barang'),
(4, 'barang.delete', 'Hapus barang'),
(5, 'kategori.read', 'Lihat kategori'),
(6, 'kategori.create', 'Tambah kategori'),
(7, 'kategori.update', 'Ubah kategori'),
(8, 'kategori.delete', 'Hapus kategori'),
(9, 'satuan.read', 'Lihat satuan'),
(10, 'satuan.create', 'Tambah satuan'),
(11, 'satuan.update', 'Ubah satuan'),
(12, 'satuan.delete', 'Hapus satuan'),
(13, 'supplier.read', 'Lihat supplier'),
(14, 'supplier.create', 'Tambah supplier'),
(15, 'supplier.update', 'Ubah supplier'),
(16, 'supplier.delete', 'Hapus supplier'),
(17, 'lokasi.read', 'Lihat gudang'),
(18, 'lokasi.create', 'Tambah gudang'),
(19, 'lokasi.update', 'Ubah gudang'),
(20, 'lokasi.delete', 'Hapus gudang'),
(21, 'stock_in.read', 'Lihat barang masuk'),
(22, 'stock_in.create', 'Simpan barang masuk'),
(23, 'stock_out.read', 'Lihat barang keluar'),
(24, 'stock_out.create', 'Simpan barang keluar'),
(25, 'stock_adjustment.read', 'Lihat penyesuaian stok'),
(26, 'stock_adjustment.create', 'Simpan penyesuaian stok'),
(27, 'stock_transfer.read', 'Lihat transfer gudang'),
(28, 'stock_transfer.create', 'Simpan transfer gudang'),
(29, 'laporan.read', 'Lihat seluruh laporan'),
(30, 'dashboard.read', 'Lihat dashboard'),
(31, 'profile.read', 'Lihat profil sendiri'),
(32, 'profile.update', 'Ubah profil sendiri');

-- --------------------------------------------------------

--
-- Table structure for table `refresh_tokens`
--

CREATE TABLE `refresh_tokens` (
  `id` bigint NOT NULL,
  `user_id` int NOT NULL,
  `token_hash` char(64) NOT NULL,
  `family_id` char(32) NOT NULL,
  `user_agent` varchar(255) NOT NULL DEFAULT '',
  `ip` varchar(45) NOT NULL DEFAULT '',
  `expires_at` datetime NOT NULL,
  `used_at` datetime DEFAULT NULL,
  `revoked_at` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `refresh_tokens`
--

INSERT INTO `refresh_tokens` (`id`, `user_id`, `token_hash`, `family_id`, `user_agent`, `ip`, `expires_at`, `used_at`, `revoked_at`, `created_at`) VALUES
(1, 1, '6a4663a421e7666521e731c45c1d15d7aa89fa1322c86be2c792baa06ad8c6e4', 'b977387d5071a53cccb73cb206f7c0b8', 'PostmanRuntime/7.49.1', '::1', '2026-08-31 22:35:18', NULL, NULL, '2026-08-25 05:35:17'),
(2, 1, '44de8bfb780b87a6e4a0c71df2ce2f13f95f20f2858ddfdb0cd65b70de14c003', '8a021803fd3709a764d78e9087178688', 'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:153.0) Gecko/20100101 Firefox/153.0', '127.0.0.1', '2026-08-31 22:38:22', '2026-08-25 05:55:11', NULL, '2026-08-25 05:38:21'),
(3, 1, '3053f42750f131f11b885d4722259bd1bf8a2fe2dcc00459c0b304245c829348', '8a021803fd3709a764d78e9087178688', 'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:153.0) Gecko/20100101 Firefox/153.0', '127.0.0.1', '2026-08-31 22:55:11', '2026-08-25 05:57:26', NULL, '2026-08-25 05:55:11'),
(4, 1, '3dad1987f30e266f80d9b95226a9e1feb2203a84d670355060d72c74e5cb3c6a', '8a021803fd3709a764d78e9087178688', 'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:153.0) Gecko/20100101 Firefox/153.0', '127.0.0.1', '2026-08-31 22:57:27', '2026-08-25 06:25:31', NULL, '2026-08-25 05:57:26'),
(5, 1, 'bf0d6753641a5b0d233d3e40a77a4937021739f66d2e7bdd8263a573d26121d5', '8a021803fd3709a764d78e9087178688', 'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:153.0) Gecko/20100101 Firefox/153.0', '127.0.0.1', '2026-08-31 23:25:32', '2026-08-25 06:26:28', NULL, '2026-08-25 06:25:31'),
(6, 1, '8b612a79ebe3f405041a0c27c1df82e50afd78d593fb00a5ef34a4ce4c7a4861', '8a021803fd3709a764d78e9087178688', 'Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:153.0) Gecko/20100101 Firefox/153.0', '127.0.0.1', '2026-08-31 23:26:29', NULL, NULL, '2026-08-25 06:26:28');

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` int NOT NULL,
  `name` varchar(50) NOT NULL,
  `keterangan` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `name`, `keterangan`) VALUES
(1, 'admin', 'Akses penuh ke seluruh modul');

-- --------------------------------------------------------

--
-- Table structure for table `role_permissions`
--

CREATE TABLE `role_permissions` (
  `role_id` int NOT NULL,
  `permission_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `role_permissions`
--

INSERT INTO `role_permissions` (`role_id`, `permission_id`) VALUES
(1, 1),
(1, 2),
(1, 3),
(1, 4),
(1, 5),
(1, 6),
(1, 7),
(1, 8),
(1, 9),
(1, 10),
(1, 11),
(1, 12),
(1, 13),
(1, 14),
(1, 15),
(1, 16),
(1, 17),
(1, 18),
(1, 19),
(1, 20),
(1, 21),
(1, 22),
(1, 23),
(1, 24),
(1, 25),
(1, 26),
(1, 27),
(1, 28),
(1, 29),
(1, 30),
(1, 31),
(1, 32);

-- --------------------------------------------------------

--
-- Table structure for table `satuan`
--

CREATE TABLE `satuan` (
  `id` int NOT NULL,
  `nama_satuan` varchar(50) NOT NULL,
  `keterangan` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `satuan`
--

INSERT INTO `satuan` (`id`, `nama_satuan`, `keterangan`, `created_at`) VALUES
(1, 'Pcs', 'Per buah', '2026-08-24 00:08:30'),
(2, 'Box', 'Per kardus', '2026-08-24 00:08:30'),
(3, 'Kg', 'Kilogram', '2026-08-24 00:08:30'),
(4, 'Dozen/Lusin', 'Lusin', '2026-08-24 00:14:11');

-- --------------------------------------------------------

--
-- Table structure for table `stock_adjustments`
--

CREATE TABLE `stock_adjustments` (
  `id` int NOT NULL,
  `code` varchar(30) NOT NULL,
  `lokasi_id` int NOT NULL,
  `user_id` int DEFAULT NULL,
  `tanggal` date NOT NULL,
  `alasan` varchar(255) NOT NULL DEFAULT '',
  `status` varchar(20) NOT NULL DEFAULT 'POSTED',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `stock_adjustment_details`
--

CREATE TABLE `stock_adjustment_details` (
  `id` int NOT NULL,
  `stock_adjustment_id` int NOT NULL,
  `barang_id` int NOT NULL,
  `system_stock` int NOT NULL,
  `actual_stock` int NOT NULL,
  `difference` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `stock_ins`
--

CREATE TABLE `stock_ins` (
  `id` int NOT NULL,
  `code` varchar(30) NOT NULL,
  `supplier_id` int DEFAULT NULL,
  `lokasi_id` int NOT NULL,
  `user_id` int DEFAULT NULL,
  `tanggal` date NOT NULL,
  `catatan` varchar(255) NOT NULL DEFAULT '',
  `status` varchar(20) NOT NULL DEFAULT 'POSTED',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `stock_ins`
--

INSERT INTO `stock_ins` (`id`, `code`, `supplier_id`, `lokasi_id`, `user_id`, `tanggal`, `catatan`, `status`, `created_at`) VALUES
(1, 'IN-0001', 3, 3, 1, '2026-08-24', 'rutin bulanan', 'POSTED', '2026-08-24 08:14:26'),
(2, 'IN-0002', 2, 2, 1, '2026-08-24', 'pembelian bulanan', 'POSTED', '2026-08-24 11:44:27'),
(3, 'IN-0003', 3, 3, 1, '2026-08-24', 'pembelian rutin', 'POSTED', '2026-08-25 06:26:47');

-- --------------------------------------------------------

--
-- Table structure for table `stock_in_details`
--

CREATE TABLE `stock_in_details` (
  `id` int NOT NULL,
  `stock_in_id` int NOT NULL,
  `barang_id` int NOT NULL,
  `quantity` int NOT NULL,
  `harga_beli` decimal(15,2) NOT NULL DEFAULT '0.00'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `stock_in_details`
--

INSERT INTO `stock_in_details` (`id`, `stock_in_id`, `barang_id`, `quantity`, `harga_beli`) VALUES
(1, 1, 5, 1000, 2500.00),
(2, 2, 3, 500, 350000.00),
(3, 3, 1, 500, 20000.00);

-- --------------------------------------------------------

--
-- Table structure for table `stock_movements`
--

CREATE TABLE `stock_movements` (
  `id` int NOT NULL,
  `barang_id` int NOT NULL,
  `lokasi_id` int NOT NULL,
  `type` enum('IN','OUT','ADJUSTMENT','TRANSFER_IN','TRANSFER_OUT') NOT NULL,
  `reference_type` varchar(30) NOT NULL,
  `reference_id` int NOT NULL,
  `quantity` int NOT NULL,
  `stock_before` int NOT NULL,
  `stock_after` int NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `stock_movements`
--

INSERT INTO `stock_movements` (`id`, `barang_id`, `lokasi_id`, `type`, `reference_type`, `reference_id`, `quantity`, `stock_before`, `stock_after`, `created_at`) VALUES
(1, 5, 3, 'IN', 'stock_in', 1, 1000, 0, 1000, '2026-08-24 08:14:26'),
(2, 5, 3, 'OUT', 'stock_out', 1, -500, 1000, 500, '2026-08-24 08:14:45'),
(3, 5, 3, 'TRANSFER_OUT', 'stock_transfer', 1, -100, 500, 400, '2026-08-24 08:15:09'),
(4, 5, 2, 'TRANSFER_IN', 'stock_transfer', 1, 100, 0, 100, '2026-08-24 08:15:09'),
(5, 3, 2, 'IN', 'stock_in', 2, 500, 0, 500, '2026-08-24 11:44:27'),
(6, 1, 3, 'IN', 'stock_in', 3, 500, 0, 500, '2026-08-25 06:26:47');

-- --------------------------------------------------------

--
-- Table structure for table `stock_outs`
--

CREATE TABLE `stock_outs` (
  `id` int NOT NULL,
  `code` varchar(30) NOT NULL,
  `lokasi_id` int NOT NULL,
  `user_id` int DEFAULT NULL,
  `tanggal` date NOT NULL,
  `tujuan` varchar(150) NOT NULL DEFAULT '',
  `catatan` varchar(255) NOT NULL DEFAULT '',
  `status` varchar(20) NOT NULL DEFAULT 'POSTED',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `stock_outs`
--

INSERT INTO `stock_outs` (`id`, `code`, `lokasi_id`, `user_id`, `tanggal`, `tujuan`, `catatan`, `status`, `created_at`) VALUES
(1, 'OUT-0001', 3, 1, '2026-08-24', 'penjualan', 'penjualan bulanan', 'POSTED', '2026-08-24 08:14:45');

-- --------------------------------------------------------

--
-- Table structure for table `stock_out_details`
--

CREATE TABLE `stock_out_details` (
  `id` int NOT NULL,
  `stock_out_id` int NOT NULL,
  `barang_id` int NOT NULL,
  `quantity` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `stock_out_details`
--

INSERT INTO `stock_out_details` (`id`, `stock_out_id`, `barang_id`, `quantity`) VALUES
(1, 1, 5, 500);

-- --------------------------------------------------------

--
-- Table structure for table `stock_transfers`
--

CREATE TABLE `stock_transfers` (
  `id` int NOT NULL,
  `code` varchar(30) NOT NULL,
  `from_lokasi_id` int NOT NULL,
  `to_lokasi_id` int NOT NULL,
  `user_id` int DEFAULT NULL,
  `tanggal` date NOT NULL,
  `catatan` varchar(255) NOT NULL DEFAULT '',
  `status` varchar(20) NOT NULL DEFAULT 'POSTED',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `stock_transfers`
--

INSERT INTO `stock_transfers` (`id`, `code`, `from_lokasi_id`, `to_lokasi_id`, `user_id`, `tanggal`, `catatan`, `status`, `created_at`) VALUES
(1, 'TRF-0001', 3, 2, 1, '2026-08-24', 'supply rj 45', 'POSTED', '2026-08-24 08:15:09');

-- --------------------------------------------------------

--
-- Table structure for table `stock_transfer_details`
--

CREATE TABLE `stock_transfer_details` (
  `id` int NOT NULL,
  `stock_transfer_id` int NOT NULL,
  `barang_id` int NOT NULL,
  `quantity` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `stock_transfer_details`
--

INSERT INTO `stock_transfer_details` (`id`, `stock_transfer_id`, `barang_id`, `quantity`) VALUES
(1, 1, 5, 100);

-- --------------------------------------------------------

--
-- Table structure for table `supplier`
--

CREATE TABLE `supplier` (
  `id` int NOT NULL,
  `nama_supplier` varchar(150) NOT NULL,
  `kontak` varchar(100) DEFAULT NULL,
  `alamat` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `supplier`
--

INSERT INTO `supplier` (`id`, `nama_supplier`, `kontak`, `alamat`, `created_at`) VALUES
(1, 'CV Sumber Makmur', '0812xxxxxxx', 'Jl. Raya No. 1', '2026-08-24 00:08:30'),
(2, 'PT Jaya Abadi', '0813xxxxxxx', 'Jl. Industri No. 2', '2026-08-24 00:08:30'),
(3, 'PT. ABC', '0267367324672', 'SERANG', '2026-08-24 00:14:24');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `username` varchar(50) NOT NULL,
  `password` varchar(255) NOT NULL,
  `nama_lengkap` varchar(100) DEFAULT NULL,
  `role_id` int DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `username`, `password`, `nama_lengkap`, `role_id`, `created_at`) VALUES
(1, 'admin', '$2a$10$C5P8Vu6vqCdtpakOuv5S1eiJd5lNeE3ckWgOHnnOreGZXJ0SJwQia', 'Ilham', 1, '2026-08-23 23:18:55');

-- --------------------------------------------------------

--
-- Table structure for table `warehouse_stocks`
--

CREATE TABLE `warehouse_stocks` (
  `id` int NOT NULL,
  `lokasi_id` int NOT NULL,
  `barang_id` int NOT NULL,
  `quantity` int NOT NULL DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `warehouse_stocks`
--

INSERT INTO `warehouse_stocks` (`id`, `lokasi_id`, `barang_id`, `quantity`) VALUES
(1, 3, 5, 400),
(2, 2, 5, 100),
(3, 2, 3, 500),
(4, 3, 1, 500);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `barang`
--
ALTER TABLE `barang`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_barang_kategori` (`kategori_id`),
  ADD KEY `fk_barang_satuan` (`satuan_id`);

--
-- Indexes for table `kategori`
--
ALTER TABLE `kategori`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `lokasi`
--
ALTER TABLE `lokasi`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `permissions`
--
ALTER TABLE `permissions`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Indexes for table `refresh_tokens`
--
ALTER TABLE `refresh_tokens`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `token_hash` (`token_hash`),
  ADD KEY `idx_rt_user` (`user_id`),
  ADD KEY `idx_rt_family` (`family_id`),
  ADD KEY `idx_rt_expires` (`expires_at`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- Indexes for table `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD PRIMARY KEY (`role_id`,`permission_id`),
  ADD KEY `fk_rp_perm` (`permission_id`);

--
-- Indexes for table `satuan`
--
ALTER TABLE `satuan`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `stock_adjustments`
--
ALTER TABLE `stock_adjustments`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`),
  ADD KEY `fk_adj_lokasi` (`lokasi_id`);

--
-- Indexes for table `stock_adjustment_details`
--
ALTER TABLE `stock_adjustment_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_adjd_header` (`stock_adjustment_id`),
  ADD KEY `fk_adjd_barang` (`barang_id`);

--
-- Indexes for table `stock_ins`
--
ALTER TABLE `stock_ins`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`),
  ADD KEY `fk_in_supplier` (`supplier_id`),
  ADD KEY `fk_in_lokasi` (`lokasi_id`);

--
-- Indexes for table `stock_in_details`
--
ALTER TABLE `stock_in_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_ind_header` (`stock_in_id`),
  ADD KEY `fk_ind_barang` (`barang_id`);

--
-- Indexes for table `stock_movements`
--
ALTER TABLE `stock_movements`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_mv_barang` (`barang_id`),
  ADD KEY `idx_mv_lokasi` (`lokasi_id`),
  ADD KEY `idx_mv_ref` (`reference_type`,`reference_id`),
  ADD KEY `idx_mv_created` (`created_at`);

--
-- Indexes for table `stock_outs`
--
ALTER TABLE `stock_outs`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`),
  ADD KEY `fk_out_lokasi` (`lokasi_id`);

--
-- Indexes for table `stock_out_details`
--
ALTER TABLE `stock_out_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_outd_header` (`stock_out_id`),
  ADD KEY `fk_outd_barang` (`barang_id`);

--
-- Indexes for table `stock_transfers`
--
ALTER TABLE `stock_transfers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`),
  ADD KEY `fk_trf_from` (`from_lokasi_id`),
  ADD KEY `fk_trf_to` (`to_lokasi_id`);

--
-- Indexes for table `stock_transfer_details`
--
ALTER TABLE `stock_transfer_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_trfd_header` (`stock_transfer_id`),
  ADD KEY `fk_trfd_barang` (`barang_id`);

--
-- Indexes for table `supplier`
--
ALTER TABLE `supplier`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `username` (`username`),
  ADD KEY `fk_users_role` (`role_id`);

--
-- Indexes for table `warehouse_stocks`
--
ALTER TABLE `warehouse_stocks`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uq_lokasi_barang` (`lokasi_id`,`barang_id`),
  ADD KEY `fk_ws_barang` (`barang_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `barang`
--
ALTER TABLE `barang`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `kategori`
--
ALTER TABLE `kategori`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `lokasi`
--
ALTER TABLE `lokasi`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=33;

--
-- AUTO_INCREMENT for table `refresh_tokens`
--
ALTER TABLE `refresh_tokens`
  MODIFY `id` bigint NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `satuan`
--
ALTER TABLE `satuan`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `stock_adjustments`
--
ALTER TABLE `stock_adjustments`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `stock_adjustment_details`
--
ALTER TABLE `stock_adjustment_details`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `stock_ins`
--
ALTER TABLE `stock_ins`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `stock_in_details`
--
ALTER TABLE `stock_in_details`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `stock_movements`
--
ALTER TABLE `stock_movements`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `stock_outs`
--
ALTER TABLE `stock_outs`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `stock_out_details`
--
ALTER TABLE `stock_out_details`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `stock_transfers`
--
ALTER TABLE `stock_transfers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `stock_transfer_details`
--
ALTER TABLE `stock_transfer_details`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `supplier`
--
ALTER TABLE `supplier`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `warehouse_stocks`
--
ALTER TABLE `warehouse_stocks`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `barang`
--
ALTER TABLE `barang`
  ADD CONSTRAINT `fk_barang_kategori` FOREIGN KEY (`kategori_id`) REFERENCES `kategori` (`id`) ON DELETE SET NULL,
  ADD CONSTRAINT `fk_barang_satuan` FOREIGN KEY (`satuan_id`) REFERENCES `satuan` (`id`) ON DELETE SET NULL;

--
-- Constraints for table `refresh_tokens`
--
ALTER TABLE `refresh_tokens`
  ADD CONSTRAINT `fk_rt_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD CONSTRAINT `fk_rp_perm` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `fk_rp_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `stock_adjustments`
--
ALTER TABLE `stock_adjustments`
  ADD CONSTRAINT `fk_adj_lokasi` FOREIGN KEY (`lokasi_id`) REFERENCES `lokasi` (`id`);

--
-- Constraints for table `stock_adjustment_details`
--
ALTER TABLE `stock_adjustment_details`
  ADD CONSTRAINT `fk_adjd_barang` FOREIGN KEY (`barang_id`) REFERENCES `barang` (`id`),
  ADD CONSTRAINT `fk_adjd_header` FOREIGN KEY (`stock_adjustment_id`) REFERENCES `stock_adjustments` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `stock_ins`
--
ALTER TABLE `stock_ins`
  ADD CONSTRAINT `fk_in_lokasi` FOREIGN KEY (`lokasi_id`) REFERENCES `lokasi` (`id`),
  ADD CONSTRAINT `fk_in_supplier` FOREIGN KEY (`supplier_id`) REFERENCES `supplier` (`id`);

--
-- Constraints for table `stock_in_details`
--
ALTER TABLE `stock_in_details`
  ADD CONSTRAINT `fk_ind_barang` FOREIGN KEY (`barang_id`) REFERENCES `barang` (`id`),
  ADD CONSTRAINT `fk_ind_header` FOREIGN KEY (`stock_in_id`) REFERENCES `stock_ins` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `stock_movements`
--
ALTER TABLE `stock_movements`
  ADD CONSTRAINT `fk_mv_barang` FOREIGN KEY (`barang_id`) REFERENCES `barang` (`id`),
  ADD CONSTRAINT `fk_mv_lokasi` FOREIGN KEY (`lokasi_id`) REFERENCES `lokasi` (`id`);

--
-- Constraints for table `stock_outs`
--
ALTER TABLE `stock_outs`
  ADD CONSTRAINT `fk_out_lokasi` FOREIGN KEY (`lokasi_id`) REFERENCES `lokasi` (`id`);

--
-- Constraints for table `stock_out_details`
--
ALTER TABLE `stock_out_details`
  ADD CONSTRAINT `fk_outd_barang` FOREIGN KEY (`barang_id`) REFERENCES `barang` (`id`),
  ADD CONSTRAINT `fk_outd_header` FOREIGN KEY (`stock_out_id`) REFERENCES `stock_outs` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `stock_transfers`
--
ALTER TABLE `stock_transfers`
  ADD CONSTRAINT `fk_trf_from` FOREIGN KEY (`from_lokasi_id`) REFERENCES `lokasi` (`id`),
  ADD CONSTRAINT `fk_trf_to` FOREIGN KEY (`to_lokasi_id`) REFERENCES `lokasi` (`id`);

--
-- Constraints for table `stock_transfer_details`
--
ALTER TABLE `stock_transfer_details`
  ADD CONSTRAINT `fk_trfd_barang` FOREIGN KEY (`barang_id`) REFERENCES `barang` (`id`),
  ADD CONSTRAINT `fk_trfd_header` FOREIGN KEY (`stock_transfer_id`) REFERENCES `stock_transfers` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

--
-- Constraints for table `warehouse_stocks`
--
ALTER TABLE `warehouse_stocks`
  ADD CONSTRAINT `fk_ws_barang` FOREIGN KEY (`barang_id`) REFERENCES `barang` (`id`),
  ADD CONSTRAINT `fk_ws_lokasi` FOREIGN KEY (`lokasi_id`) REFERENCES `lokasi` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
