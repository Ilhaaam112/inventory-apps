package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/username/belajar_go/backend/config"
	"github.com/username/belajar_go/backend/internal/auth"
	"github.com/username/belajar_go/backend/internal/handler"
	mw "github.com/username/belajar_go/backend/internal/middleware"
	"github.com/username/belajar_go/backend/internal/repository"
	"github.com/username/belajar_go/backend/internal/service"
)

func main() {
	config.LoadEnv()

	// Konfigurasi keamanan dibaca lebih dulu. Kalau JWT_SECRET tidak ada
	// atau terlalu pendek, server sengaja menolak jalan.
	authCfg, err := auth.LoadConfig()
	if err != nil {
		log.Fatal("Konfigurasi keamanan tidak valid: ", err)
	}
	jwtManager := auth.NewManager(authCfg)

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Gagal konek database:", err)
	}
	defer db.Close()

	// ================= REPOSITORY =================
	barangRepo := repository.NewBarangRepository(db)
	userRepo := repository.NewUserRepository(db)
	kategoriRepo := repository.NewKategoriRepository(db)
	satuanRepo := repository.NewSatuanRepository(db)
	supplierRepo := repository.NewSupplierRepository(db)
	lokasiRepo := repository.NewLokasiRepository(db)

	stockInRepo := repository.NewStockInRepository(db)
	stockOutRepo := repository.NewStockOutRepository(db)
	adjustmentRepo := repository.NewStockAdjustmentRepository(db)
	transferRepo := repository.NewStockTransferRepository(db)
	warehouseStockRepo := repository.NewWarehouseStockRepository(db)
	movementRepo := repository.NewStockMovementRepository(db)

	laporanRepo := repository.NewLaporanRepository(db)
	dashboardRepo := repository.NewDashboardRepository(db)

	rbacRepo := repository.NewRBACRepository(db)
	refreshRepo := repository.NewRefreshTokenRepository(db)

	// ================= SERVICE =================
	barangService := service.NewBarangService(barangRepo)
	userService := service.NewUserService(userRepo)
	kategoriService := service.NewKategoriService(kategoriRepo)
	satuanService := service.NewSatuanService(satuanRepo)
	supplierService := service.NewSupplierService(supplierRepo)
	lokasiService := service.NewLokasiService(lokasiRepo)

	stockInService := service.NewStockInService(stockInRepo)
	stockOutService := service.NewStockOutService(stockOutRepo)
	adjustmentService := service.NewStockAdjustmentService(adjustmentRepo)
	transferService := service.NewStockTransferService(transferRepo)
	warehouseStockService := service.NewWarehouseStockService(warehouseStockRepo)
	movementService := service.NewStockMovementService(movementRepo)

	laporanService := service.NewLaporanService(laporanRepo)
	dashboardService := service.NewDashboardService(dashboardRepo)
	authService := service.NewAuthService(userRepo, rbacRepo, refreshRepo, jwtManager)

	// ================= HANDLER =================
	barangHandler := handler.NewBarangHandler(barangService)
	userHandler := handler.NewUserHandler(userService)
	kategoriHandler := handler.NewKategoriHandler(kategoriService)
	satuanHandler := handler.NewSatuanHandler(satuanService)
	supplierHandler := handler.NewSupplierHandler(supplierService)
	lokasiHandler := handler.NewLokasiHandler(lokasiService)

	stockInHandler := handler.NewStockInHandler(stockInService)
	stockOutHandler := handler.NewStockOutHandler(stockOutService)
	adjustmentHandler := handler.NewStockAdjustmentHandler(adjustmentService)
	transferHandler := handler.NewStockTransferHandler(transferService)
	warehouseStockHandler := handler.NewWarehouseStockHandler(warehouseStockService)
	movementHandler := handler.NewStockMovementHandler(movementService)

	laporanHandler := handler.NewLaporanHandler(laporanService)
	dashboardHandler := handler.NewDashboardHandler(dashboardService)
	authHandler := handler.NewAuthHandler(authService, authCfg)

	// ================= ROUTING =================
	mux := http.NewServeMux()

	// --- Pembantu supaya pendaftaran route tetap ringkas ---
	//
	//   butuhLogin  = wajib access token valid          -> 401 kalau gagal
	//   izin        = wajib permission sesuai method    -> 403 kalau gagal
	butuhLogin := mw.RequireAuth(jwtManager)

	crud := func(res string) map[string]string {
		return map[string]string{
			http.MethodGet:    res + ".read",
			http.MethodPost:   res + ".create",
			http.MethodPut:    res + ".update",
			http.MethodDelete: res + ".delete",
		}
	}
	bacaTulis := func(res string) map[string]string {
		return map[string]string{
			http.MethodGet:  res + ".read",
			http.MethodPost: res + ".create",
		}
	}
	bacaSaja := func(res string) map[string]string {
		return map[string]string{http.MethodGet: res + ".read"}
	}

	// jaga membungkus handler dengan autentikasi lalu otorisasi.
	jaga := func(h http.HandlerFunc, izin map[string]string) http.Handler {
		return mw.Chain(butuhLogin, mw.RequirePerMethod(izin))(h)
	}

	// ---------- Publik ----------
	mux.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	// Endpoint auth dibatasi lebih ketat: 5 percobaan per menit per IP.
	batasAuth := mw.RateLimit(12*time.Second, 5, authCfg.Production)
	mux.Handle("/api/v1/auth/login", batasAuth(http.HandlerFunc(authHandler.HandleLogin)))
	mux.Handle("/api/v1/auth/refresh", batasAuth(http.HandlerFunc(authHandler.HandleRefresh)))
	mux.Handle("/api/v1/auth/logout", http.HandlerFunc(authHandler.HandleLogout))

	// ---------- Profil ----------
	mux.Handle("/api/profile", jaga(userHandler.HandleProfile, map[string]string{
		http.MethodGet: "profile.read",
		http.MethodPut: "profile.update",
	}))
	mux.Handle("/api/change-password", jaga(userHandler.HandleChangePassword, map[string]string{
		http.MethodPut: "profile.update",
	}))

	// ---------- Master data ----------
	mux.Handle("/api/barang", jaga(barangHandler.HandleBarang, crud("barang")))
	mux.Handle("/api/barang/", jaga(barangHandler.HandleBarangByID, crud("barang")))

	mux.Handle("/api/kategori", jaga(kategoriHandler.HandleKategori, crud("kategori")))
	mux.Handle("/api/kategori/", jaga(kategoriHandler.HandleKategoriByID, crud("kategori")))

	mux.Handle("/api/satuan", jaga(satuanHandler.HandleSatuan, crud("satuan")))
	mux.Handle("/api/satuan/", jaga(satuanHandler.HandleSatuanByID, crud("satuan")))

	mux.Handle("/api/supplier", jaga(supplierHandler.HandleSupplier, crud("supplier")))
	mux.Handle("/api/supplier/", jaga(supplierHandler.HandleSupplierByID, crud("supplier")))

	mux.Handle("/api/lokasi", jaga(lokasiHandler.HandleLokasi, crud("lokasi")))
	mux.Handle("/api/lokasi/", jaga(lokasiHandler.HandleLokasiByID, crud("lokasi")))

	// ---------- Transaksi ----------
	mux.Handle("/api/stock-in", jaga(stockInHandler.HandleStockIn, bacaTulis("stock_in")))
	mux.Handle("/api/stock-in/", jaga(stockInHandler.HandleStockInByID, bacaSaja("stock_in")))

	mux.Handle("/api/stock-out", jaga(stockOutHandler.HandleStockOut, bacaTulis("stock_out")))
	mux.Handle("/api/stock-out/", jaga(stockOutHandler.HandleStockOutByID, bacaSaja("stock_out")))

	mux.Handle("/api/stock-adjustment", jaga(adjustmentHandler.HandleAdjustment, bacaTulis("stock_adjustment")))
	mux.Handle("/api/stock-adjustment/", jaga(adjustmentHandler.HandleAdjustmentByID, bacaSaja("stock_adjustment")))

	mux.Handle("/api/stock-transfer", jaga(transferHandler.HandleTransfer, bacaTulis("stock_transfer")))
	mux.Handle("/api/stock-transfer/", jaga(transferHandler.HandleTransferByID, bacaSaja("stock_transfer")))

	mux.Handle("/api/warehouse-stocks", jaga(warehouseStockHandler.HandleWarehouseStock, bacaSaja("laporan")))
	mux.Handle("/api/stock-movements", jaga(movementHandler.HandleStockMovement, bacaSaja("laporan")))

	// ---------- Laporan ----------
	mux.Handle("/api/laporan/stok", jaga(laporanHandler.HandleLaporanStok, bacaSaja("laporan")))
	mux.Handle("/api/laporan/kartu-stok", jaga(laporanHandler.HandleKartuStok, bacaSaja("laporan")))
	mux.Handle("/api/laporan/barang-masuk", jaga(laporanHandler.HandleLaporanMasuk, bacaSaja("laporan")))
	mux.Handle("/api/laporan/barang-keluar", jaga(laporanHandler.HandleLaporanKeluar, bacaSaja("laporan")))
	mux.Handle("/api/laporan/pergerakan", jaga(laporanHandler.HandleLaporanPergerakan, bacaSaja("laporan")))

	// ---------- Dashboard ----------
	mux.Handle("/api/dashboard/overview", jaga(dashboardHandler.HandleOverview, bacaSaja("dashboard")))
	mux.Handle("/api/dashboard/stok-menipis", jaga(dashboardHandler.HandleStokMenipis, bacaSaja("dashboard")))
	mux.Handle("/api/dashboard/aktivitas", jaga(dashboardHandler.HandleAktivitas, bacaSaja("dashboard")))
	mux.Handle("/api/dashboard/tren", jaga(dashboardHandler.HandleTren, bacaSaja("dashboard")))

	// ================= RANTAI MIDDLEWARE GLOBAL =================
	//
	//   Request -> Recover -> SecurityHeaders -> RateLimit -> CORS -> mux
	//                                                          |
	//                                          (per route) JWT -> Permission -> Handler
	//
	// CORS diletakkan paling dekat dengan mux supaya preflight OPTIONS
	// tetap dibalas walau route-nya butuh login.
	rantai := mw.Chain(
		mw.Recover(authCfg.Production),
		mw.SecurityHeaders(authCfg.Production),
		mw.RateLimit(500*time.Millisecond, 60, authCfg.Production),
		mw.CORS(authCfg.AllowedOrigins),
	)

	// Bersih-bersih refresh token kadaluarsa setiap 12 jam.
	go func() {
		for {
			_ = refreshRepo.DeleteExpired()
			time.Sleep(12 * time.Hour)
		}
	}()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           rantai(mux),
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}

	fmt.Println("Server jalan di http://localhost:" + port)
	if !authCfg.Production {
		fmt.Println("Mode development: cookie Secure=false, HSTS mati.")
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server berhenti: ", err)
	}
}
