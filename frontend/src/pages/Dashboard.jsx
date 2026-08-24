import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'
import Layout from '../components/Layout'

function statusBarang(stok) {
  if (stok === 0) return { label: 'Habis', className: 'bg-red-500/10 text-red-400 border-red-500/30' }
  if (stok <= 10) return { label: 'Menipis', className: 'bg-accent/10 text-accent border-accent/30' }
  return { label: 'Aman', className: 'bg-success/10 text-success border-success/30' }
}

function Dashboard({ user, onLogout }) {
  const navigate = useNavigate()
  const [barangList, setBarangList] = useState([])

  useEffect(() => {
    axios.get('/api/barang').then((res) => setBarangList(res.data || []))
  }, [])

  const totalBarang = barangList.length
  const totalStok = barangList.reduce((sum, b) => sum + b.stok, 0)
  const stokMenipis = barangList.filter((b) => b.stok > 0 && b.stok <= 10).length
  const nilaiInventori = barangList.reduce((sum, b) => sum + b.harga * b.stok, 0)

  const stats = [
    { label: 'Total Barang', value: totalBarang, sub: 'Jenis barang' },
    { label: 'Total Stok', value: totalStok, sub: 'Unit tersedia' },
    { label: 'Stok Menipis', value: stokMenipis, sub: '≤ 10 unit' },
    { label: 'Nilai Inventori', value: `Rp ${nilaiInventori.toLocaleString('id-ID')}`, sub: 'Estimasi total' },
  ]

  return (
    <Layout title="Dashboard" user={user} onLogout={onLogout}>
      <div className="rounded-2xl bg-gradient-to-br from-accent to-accent-soft p-8 mb-6 flex items-center justify-between flex-wrap gap-4">
        <div>
          <p className="font-mono text-xs text-white/80 tracking-widest mb-2">RINGKASAN</p>
          <h2 className="font-display text-2xl font-semibold text-white">
            Halo, {user?.nama_lengkap} 👋
          </h2>
        </div>
        <button
          onClick={() => navigate('/barang')}
          className="bg-white text-accent font-medium rounded-full px-5 py-2.5 text-sm hover:bg-white/90 transition-colors"
        >
          + Tambah Barang
        </button>
      </div>

      <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-8">
        {stats.map((s) => (
          <div key={s.label} className="bg-surface border border-border rounded-2xl p-5">
            <p className="text-xs text-muted mb-3">{s.label}</p>
            <p className="font-display text-2xl font-semibold mb-1">{s.value}</p>
            <p className="text-xs text-muted">{s.sub}</p>
          </div>
        ))}
      </div>

      <div className="bg-surface border border-border rounded-2xl overflow-hidden">
        <div className="flex items-center justify-between px-6 py-4 border-b border-border">
          <h3 className="font-display font-semibold">Barang Terbaru</h3>
          <button onClick={() => navigate('/barang')} className="text-xs text-accent hover:text-accent-soft">
            Lihat semua →
          </button>
        </div>
        <table className="w-full text-sm">
          <thead>
            <tr className="text-left text-xs font-mono text-muted border-b border-border">
              <th className="px-6 py-3">Nama Barang</th>
              <th className="px-6 py-3">Harga</th>
              <th className="px-6 py-3">Stok</th>
              <th className="px-6 py-3">Status</th>
            </tr>
          </thead>
          <tbody>
            {barangList.length === 0 ? (
              <tr><td colSpan="4" className="text-center py-8 text-muted">Belum ada data</td></tr>
            ) : (
              barangList.slice(-5).reverse().map((b) => {
                const status = statusBarang(b.stok)
                return (
                  <tr key={b.id} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                    <td className="px-6 py-3 font-medium">{b.nama}</td>
                    <td className="px-6 py-3 font-mono text-muted">Rp {b.harga.toLocaleString('id-ID')}</td>
                    <td className="px-6 py-3 font-mono">{b.stok}</td>
                    <td className="px-6 py-3">
                      <span className={`text-xs border rounded-full px-2.5 py-1 ${status.className}`}>
                        {status.label}
                      </span>
                    </td>
                  </tr>
                )
              })
            )}
          </tbody>
        </table>
      </div>
    </Layout>
  )
}

export default Dashboard