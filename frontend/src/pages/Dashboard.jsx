import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import axios from 'axios'
import { Package, Warehouse, Boxes, AlertTriangle, ArrowRight } from 'lucide-react'
import Layout from '../components/Layout'
import { rupiah } from '../components/FilterBar'

function Kartu({ icon: Icon, label, nilai, sub, warna = 'text-ink' }) {
  return (
    <div className="bg-surface border border-border rounded-2xl p-5">
      <div className="flex items-center gap-2 mb-2 text-muted">
        <Icon size={15} />
        <p className="text-xs font-mono tracking-widest uppercase">{label}</p>
      </div>
      <p className={`font-display text-2xl ${warna}`}>{nilai}</p>
      {sub && <p className="text-xs text-muted mt-1">{sub}</p>}
    </div>
  )
}

function Dashboard({ user, onLogout }) {
  const [data, setData] = useState(null)
  const [gagal, setGagal] = useState(false)

  useEffect(() => {
    axios
      .get('/api/dashboard/overview')
      .then((r) => setData(r.data))
      .catch(() => setGagal(true))
  }, [])

  if (gagal) {
    return (
      <Layout title="Overview" user={user} onLogout={onLogout}>
        <div className="bg-surface border border-border rounded-2xl p-10 text-center text-muted text-sm">
          Gagal memuat ringkasan. Pastikan server Go berjalan dan migrasi dashboard sudah dijalankan.
        </div>
      </Layout>
    )
  }

  if (!data) {
    return (
      <Layout title="Overview" user={user} onLogout={onLogout}>
        <div className="bg-surface border border-border rounded-2xl p-10 text-center text-muted text-sm">Memuat…</div>
      </Layout>
    )
  }

  const puncak = Math.max(1, ...data.tren.map((t) => Math.max(t.masuk, t.keluar)))
  const perluPerhatian = data.stok_menipis + data.stok_habis

  return (
    <Layout title="Overview" user={user} onLogout={onLogout}>
      <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-6">
        <Kartu icon={Package} label="Jenis Barang" nilai={data.total_barang}
          sub={`${data.total_gudang} gudang · ${data.total_supplier} supplier`} />
        <Kartu icon={Boxes} label="Total Unit" nilai={data.total_unit.toLocaleString('id-ID')}
          sub="tersebar di seluruh gudang" />
        <Kartu icon={Warehouse} label="Nilai Stok" nilai={rupiah(data.nilai_persediaan)}
          sub="unit × harga barang" warna="text-accent" />
        <Kartu icon={AlertTriangle} label="Perlu Perhatian" nilai={perluPerhatian}
          sub={`${data.stok_habis} habis · ${data.stok_menipis} menipis`}
          warna={perluPerhatian > 0 ? 'text-accent' : 'text-success'} />
      </div>

      <div className="grid lg:grid-cols-3 gap-4 mb-6">
        <div className="bg-surface border border-border rounded-2xl p-5">
          <p className="text-xs font-mono text-muted tracking-widest uppercase mb-3">Hari Ini</p>
          <div className="flex gap-8">
            <div>
              <p className="font-display text-2xl text-success">+{data.masuk_hari_ini}</p>
              <p className="text-xs text-muted">unit masuk</p>
            </div>
            <div>
              <p className="font-display text-2xl text-accent">−{data.keluar_hari_ini}</p>
              <p className="text-xs text-muted">unit keluar</p>
            </div>
          </div>
          <p className="text-xs text-muted mt-4 pt-4 border-t border-border">
            {data.transaksi_bulan_ini} transaksi bulan ini
          </p>
        </div>

        <div className="bg-surface border border-border rounded-2xl p-5 lg:col-span-2">
          <p className="text-xs font-mono text-muted tracking-widest uppercase mb-4">7 Hari Terakhir</p>
          <div className="flex items-end justify-between gap-2 h-32">
            {data.tren.map((t) => (
              <div key={t.tanggal} className="flex-1 flex flex-col items-center gap-1">
                <div className="w-full flex items-end justify-center gap-1 h-24">
                  <div
                    className="w-1/2 bg-success/70 rounded-t"
                    style={{ height: `${(t.masuk / puncak) * 100}%` }}
                    title={`Masuk ${t.masuk}`}
                  />
                  <div
                    className="w-1/2 bg-accent/70 rounded-t"
                    style={{ height: `${(t.keluar / puncak) * 100}%` }}
                    title={`Keluar ${t.keluar}`}
                  />
                </div>
                <span className="text-[10px] font-mono text-muted">{t.tanggal.slice(8)}</span>
              </div>
            ))}
          </div>
          <div className="flex gap-4 mt-3 text-xs text-muted">
            <span className="flex items-center gap-1.5"><span className="w-2.5 h-2.5 rounded-sm bg-success/70" /> Masuk</span>
            <span className="flex items-center gap-1.5"><span className="w-2.5 h-2.5 rounded-sm bg-accent/70" /> Keluar</span>
          </div>
        </div>
      </div>

      <div className="grid sm:grid-cols-2 gap-4">
        <Link to="/dashboard/stok-menipis" className="bg-surface border border-border rounded-2xl p-5 hover:border-accent/40 transition-colors flex items-center justify-between">
          <div>
            <p className="font-display font-semibold mb-1">Stok Menipis</p>
            <p className="text-xs text-muted">Lihat barang yang perlu dipesan ulang</p>
          </div>
          <ArrowRight size={18} className="text-muted" />
        </Link>
        <Link to="/dashboard/aktivitas" className="bg-surface border border-border rounded-2xl p-5 hover:border-accent/40 transition-colors flex items-center justify-between">
          <div>
            <p className="font-display font-semibold mb-1">Aktivitas Terbaru</p>
            <p className="text-xs text-muted">Riwayat pergerakan stok terakhir</p>
          </div>
          <ArrowRight size={18} className="text-muted" />
        </Link>
      </div>
    </Layout>
  )
}

export default Dashboard
