package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/username/belajar_go/backend/config"
	"github.com/username/belajar_go/backend/internal/handler"
	"github.com/username/belajar_go/backend/internal/repository"
	"github.com/username/belajar_go/backend/internal/service"
)

func main() {
	config.LoadEnv()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}
	defer db.Close()

	// ---------- Master data ----------
	barangRepo := repository.NewBarangRepository(db)
	barangService := service.NewBarangService(barangRepo)
	barangHandler := handler.NewBarangHandler(barangService)

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	kategoriRepo := repository.NewKategoriRepository(db)
	kategoriService := service.NewKategoriService(kategoriRepo)
	kategoriHandler := handler.NewKategoriHandler(kategoriService)

	satuanRepo := repository.NewSatuanRepository(db)
	satuanService := service.NewSatuanService(satuanRepo)
	satuanHandler := handler.NewSatuanHandler(satuanService)

	supplierRepo := repository.NewSupplierRepository(db)
	supplierService := service.NewSupplierService(supplierRepo)
	supplierHandler := handler.NewSupplierHandler(supplierService)

	lokasiRepo := repository.NewLokasiRepository(db)
	lokasiService := service.NewLokasiService(lokasiRepo)
	lokasiHandler := handler.NewLokasiHandler(lokasiService)

	// ---------- BARU: Transaksi ----------
	stockInRepo := repository.NewStockInRepository(db)
	stockInService := service.NewStockInService(stockInRepo)
	stockInHandler := handler.NewStockInHandler(stockInService)

	stockOutRepo := repository.NewStockOutRepository(db)
	stockOutService := service.NewStockOutService(stockOutRepo)
	stockOutHandler := handler.NewStockOutHandler(stockOutService)

	adjustmentRepo := repository.NewStockAdjustmentRepository(db)
	adjustmentService := service.NewStockAdjustmentService(adjustmentRepo)
	adjustmentHandler := handler.NewStockAdjustmentHandler(adjustmentService)

	transferRepo := repository.NewStockTransferRepository(db)
	transferService := service.NewStockTransferService(transferRepo)
	transferHandler := handler.NewStockTransferHandler(transferService)

	warehouseStockRepo := repository.NewWarehouseStockRepository(db)
	warehouseStockService := service.NewWarehouseStockService(warehouseStockRepo)
	warehouseStockHandler := handler.NewWarehouseStockHandler(warehouseStockService)

	movementRepo := repository.NewStockMovementRepository(db)
	movementService := service.NewStockMovementService(movementRepo)
	movementHandler := handler.NewStockMovementHandler(movementService)

	// ---------- BARU: Laporan ----------
	laporanRepo := repository.NewLaporanRepository(db)
	laporanService := service.NewLaporanService(laporanRepo)
	laporanHandler := handler.NewLaporanHandler(laporanService)
	// ---------- END BARU ----------

	http.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"status":"ok"}`)
	})

	http.HandleFunc("/api/login", userHandler.HandleLogin)

	// Master data
	http.HandleFunc("/api/barang", barangHandler.HandleBarang)
	http.HandleFunc("/api/barang/", barangHandler.HandleBarangByID)

	http.HandleFunc("/api/kategori", kategoriHandler.HandleKategori)
	http.HandleFunc("/api/kategori/", kategoriHandler.HandleKategoriByID)

	http.HandleFunc("/api/satuan", satuanHandler.HandleSatuan)
	http.HandleFunc("/api/satuan/", satuanHandler.HandleSatuanByID)

	http.HandleFunc("/api/supplier", supplierHandler.HandleSupplier)
	http.HandleFunc("/api/supplier/", supplierHandler.HandleSupplierByID)

	http.HandleFunc("/api/lokasi", lokasiHandler.HandleLokasi)
	http.HandleFunc("/api/lokasi/", lokasiHandler.HandleLokasiByID)

	// ---------- BARU: route transaksi ----------
	http.HandleFunc("/api/stock-in", stockInHandler.HandleStockIn)
	http.HandleFunc("/api/stock-in/", stockInHandler.HandleStockInByID)

	http.HandleFunc("/api/stock-out", stockOutHandler.HandleStockOut)
	http.HandleFunc("/api/stock-out/", stockOutHandler.HandleStockOutByID)

	http.HandleFunc("/api/stock-adjustment", adjustmentHandler.HandleAdjustment)
	http.HandleFunc("/api/stock-adjustment/", adjustmentHandler.HandleAdjustmentByID)

	http.HandleFunc("/api/stock-transfer", transferHandler.HandleTransfer)
	http.HandleFunc("/api/stock-transfer/", transferHandler.HandleTransferByID)

	http.HandleFunc("/api/warehouse-stocks", warehouseStockHandler.HandleWarehouseStock)
	http.HandleFunc("/api/stock-movements", movementHandler.HandleStockMovement)

	// Laporan
	http.HandleFunc("/api/laporan/stok", laporanHandler.HandleLaporanStok)
	http.HandleFunc("/api/laporan/kartu-stok", laporanHandler.HandleKartuStok)
	http.HandleFunc("/api/laporan/barang-masuk", laporanHandler.HandleLaporanMasuk)
	http.HandleFunc("/api/laporan/barang-keluar", laporanHandler.HandleLaporanKeluar)
	http.HandleFunc("/api/laporan/pergerakan", laporanHandler.HandleLaporanPergerakan)
	// ---------- END BARU ----------


	http.HandleFunc("/api/profile", userHandler.HandleProfile)
	http.HandleFunc("/api/change-password", userHandler.HandleChangePassword)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server jalan di http://localhost:" + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Error:", err)
	}
}