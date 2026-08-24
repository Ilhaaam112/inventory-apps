import { useState, useEffect } from 'react'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import PublicHome from './pages/PublicHome'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import Barang from './pages/Barang'
import Kategori from './pages/Kategori'
import Satuan from './pages/Satuan'
import Supplier from './pages/Supplier'
import Lokasi from './pages/Lokasi'
import BarangMasuk from './pages/BarangMasuk'
import BarangKeluar from './pages/BarangKeluar'
import PenyesuaianStok from './pages/PenyesuaianStok'
import TransferGudang from './pages/TransferGudang'
import LaporanStok from './pages/LaporanStok'
import KartuStok from './pages/KartuStok'
import LaporanBarangMasuk from './pages/LaporanBarangMasuk'
import LaporanBarangKeluar from './pages/LaporanBarangKeluar'
import LaporanPergerakan from './pages/LaporanPergerakan'
import Profil from './pages/Profil'

function App() {
  const [user, setUser] = useState(null)

  useEffect(() => {
    const savedUser = localStorage.getItem('user')
    if (savedUser) setUser(JSON.parse(savedUser))
  }, [])

  const handleLogout = () => {
    localStorage.removeItem('user')
    setUser(null)
  }

  // semua halaman dalam dashboard butuh login
  const privateRoute = (Component) =>
    user ? <Component user={user} onLogout={handleLogout} /> : <Navigate to="/login" />

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<PublicHome />} />
        <Route
          path="/login"
          element={user ? <Navigate to="/dashboard" /> : <Login onLoginSuccess={setUser} />}
        />

        <Route path="/dashboard" element={privateRoute(Dashboard)} />
        <Route path="/barang" element={privateRoute(Barang)} />
        <Route path="/kategori" element={privateRoute(Kategori)} />
        <Route path="/satuan" element={privateRoute(Satuan)} />
        <Route path="/supplier" element={privateRoute(Supplier)} />
        <Route path="/lokasi" element={privateRoute(Lokasi)} />

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