import { useState, useEffect, useCallback } from 'react'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import { refreshSession, logoutRequest, setAccessToken } from './api'

import PublicHome from './pages/PublicHome'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import StokMenipis from './pages/StokMenipis'
import AktivitasTerbaru from './pages/AktivitasTerbaru'
import Barang from './pages/Barang'
import Kategori from './pages/Kategori'
import Satuan from './pages/Satuan'
import Supplier from './pages/Supplier'
import Lokasi from './pages/Lokasi'
import Profil from './pages/Profil'
import BarangMasuk from './pages/BarangMasuk'
import BarangKeluar from './pages/BarangKeluar'
import PenyesuaianStok from './pages/PenyesuaianStok'
import TransferGudang from './pages/TransferGudang'
import LaporanStok from './pages/LaporanStok'
import KartuStok from './pages/KartuStok'
import LaporanBarangMasuk from './pages/LaporanBarangMasuk'
import LaporanBarangKeluar from './pages/LaporanBarangKeluar'
import LaporanPergerakan from './pages/LaporanPergerakan'

function App() {
  const [user, setUser] = useState(null)
  const [memuat, setMemuat] = useState(true)

  // Saat halaman dibuka atau di-refresh, access token di memori hilang.
  // Cookie refresh token masih ada, jadi kita tukar diam-diam jadi
  // access token baru. Kalau gagal, berarti memang belum login.
  useEffect(() => {
    refreshSession()
      .then((data) => setUser(data.user))
      .catch(() => setUser(null))
      .finally(() => setMemuat(false))
  }, [])

  // Interceptor di api.js melempar event ini kalau refresh gagal.
  useEffect(() => {
    const keluarPaksa = () => setUser(null)
    window.addEventListener('auth:logout', keluarPaksa)
    return () => window.removeEventListener('auth:logout', keluarPaksa)
  }, [])

  const handleLogout = useCallback(async () => {
    await logoutRequest()
    setAccessToken(null)
    setUser(null)
  }, [])

  const privateRoute = (Component) =>
    user ? <Component user={user} onLogout={handleLogout} /> : <Navigate to="/login" />

  if (memuat) {
    return (
      <div className="min-h-screen bg-canvas text-muted flex items-center justify-center text-sm">
        Memuat sesi…
      </div>
    )
  }

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<PublicHome />} />
        <Route
          path="/login"
          element={user ? <Navigate to="/dashboard" /> : <Login onLoginSuccess={setUser} />}
        />

        {/* Dashboard */}
        <Route path="/dashboard" element={privateRoute(Dashboard)} />
        <Route path="/dashboard/stok-menipis" element={privateRoute(StokMenipis)} />
        <Route path="/dashboard/aktivitas" element={privateRoute(AktivitasTerbaru)} />

        {/* Master data */}
        <Route path="/barang" element={privateRoute(Barang)} />
        <Route path="/kategori" element={privateRoute(Kategori)} />
        <Route path="/satuan" element={privateRoute(Satuan)} />
        <Route path="/supplier" element={privateRoute(Supplier)} />
        <Route path="/lokasi" element={privateRoute(Lokasi)} />
        <Route path="/profil" element={privateRoute(Profil)} />

        {/* Transaksi */}
        <Route path="/barang-masuk" element={privateRoute(BarangMasuk)} />
        <Route path="/barang-keluar" element={privateRoute(BarangKeluar)} />
        <Route path="/penyesuaian" element={privateRoute(PenyesuaianStok)} />
        <Route path="/transfer" element={privateRoute(TransferGudang)} />

        {/* Laporan */}
        <Route path="/laporan/stok" element={privateRoute(LaporanStok)} />
        <Route path="/laporan/kartu-stok" element={privateRoute(KartuStok)} />
        <Route path="/laporan/barang-masuk" element={privateRoute(LaporanBarangMasuk)} />
        <Route path="/laporan/barang-keluar" element={privateRoute(LaporanBarangKeluar)} />
        <Route path="/laporan/pergerakan" element={privateRoute(LaporanPergerakan)} />
        
        <Route path="/profil" element={privateRoute(Profil)} />

        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
