import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'
import axios from 'axios'
import { Package, Warehouse, Boxes, AlertTriangle, ArrowRight } from 'lucide-react'
import Layout from '../components/Layout'
import { rupiah } from '../components/FilterBar'

const iso = (d) => d.toISOString().slice(0, 10)
const hariIni = () => iso(new Date())
const mundur = (hari) => {
  const d = new Date()
  d.setDate(d.getDate() - hari)
  return iso(d)
}

function Kartu({ icon: Icon, label, nilai, sub, warna = 'text-ink' }) {
  return (
    <div className="bg-surface border border-border rounded-2xl p-4 sm:p-5">
      <div className="flex items-center gap-2 mb-2 text-muted">
        <Icon size={15} />
        <p className="text-xs font-mono tracking-widest uppercase">{label}</p>
      </div>
      <p className={`font-display text-xl sm:text-2xl break-words ${warna}`}>{nilai}</p>
      {sub && <p className="text-xs text-muted mt-1">{sub}</p>}
    </div>
  )
}

function Dashboard({ user, onLogout }) {
  const [data, setData] = useState(null)
  const [gagal, setGagal] = useState(false)
  const [rentang, setRentang] = useState({ start: mundur(6), end: hariIni() })

  useEffect(() => {
    const q = new URLSearchParams(rentang)
    axios
      .get(`/api/dashboard/overview?${q}`)
      .then((r) => { setData(r.data); setGagal(false) })
      .catch(() => setGagal(true))
  }, [rentang])

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

  const tren = data.tren || []
  const puncak = Math.max(1, ...tren.map((t) => Math.max(t.masuk, t.keluar)))
  const adaPergerakan = tren.some((t) => t.masuk || t.keluar)
  const perluPerhatian = (data.stok_menipis || 0) + (data.stok_habis || 0)
  const inputKecil =
    'bg-surface-soft border border-border rounded-lg px-2.5 py-1.5 text-xs outline-none focus:border-accent transition-colors'

  return (
    <Layout title="Overview" user={user} onLogout={onLogout}>
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-3 sm:gap-4 mb-4 sm:mb-6">
        <Kartu icon={Package} label="Jenis Barang" nilai={data.total_barang}
          sub={`${data.total_gudang} gudang · ${data.total_supplier} supplier`} />
        <Kartu icon={Boxes} label="Total Unit" nilai={(data.total_unit || 0).toLocaleString('id-ID')}
          sub="tersebar di seluruh gudang" />
        <Kartu icon={Warehouse} label="Nilai Stok" nilai={rupiah(data.nilai_persediaan)}
          sub="unit × harga barang" warna="text-accent" />
        <Kartu icon={AlertTriangle} label="Perlu Perhatian" nilai={perluPerhatian}
          sub={`${data.stok_habis || 0} habis · ${data.stok_menipis || 0} menipis`}
          warna={perluPerhatian > 0 ? 'text-accent' : 'text-success'} />
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-3 sm:gap-4 mb-4 sm:mb-6">
        <div className="bg-surface border border-border rounded-2xl p-4 sm:p-5">
          <p className="text-xs font-mono text-muted tracking-widest uppercase mb-3">Hari Ini</p>
          <div className="flex gap-8">
            <div>
              <p className="font-display text-2xl text-success">+{data.masuk_hari_ini || 0}</p>
              <p className="text-xs text-muted">unit masuk</p>
            </div>
            <div>
              <p className="font-display text-2xl text-accent">−{data.keluar_hari_ini || 0}</p>
              <p className="text-xs text-muted">unit keluar</p>
            </div>
          </div>
          <p className="text-xs text-muted mt-4 pt-4 border-t border-border">
            {data.transaksi_bulan_ini || 0} transaksi bulan ini
          </p>
        </div>

        <div className="bg-surface border border-border rounded-2xl p-4 sm:p-5 lg:col-span-2">
          <div className="flex items-center justify-between gap-3 mb-4 flex-wrap">
            <p className="text-xs font-mono text-muted tracking-widest uppercase">Pergerakan Stok</p>
            <div className="flex items-center gap-2">
              <input type="date" className={inputKecil} value={rentang.start}
                onChange={(e) => setRentang({ ...rentang, start: e.target.value })} />
              <span className="text-muted text-xs">–</span>
              <input type="date" className={inputKecil} value={rentang.end}
                onChange={(e) => setRentang({ ...rentang, end: e.target.value })} />
            </div>
          </div>

          <div className="flex gap-1.5 mb-4">
            {[
              { label: '7 hari', hari: 6 },
              { label: '30 hari', hari: 29 },
              { label: '90 hari', hari: 89 },
            ].map((p) => (
              <button
                key={p.label}
                onClick={() => setRentang({ start: mundur(p.hari), end: hariIni() })}
                className="text-[11px] font-mono border border-border rounded-full px-2.5 py-1 text-muted hover:text-ink hover:border-accent/40 transition-colors"
              >
                {p.label}
              </button>
            ))}
          </div>

          {tren.length === 0 ? (
            <p className="text-center text-muted text-sm py-10">Tidak ada data pada rentang ini</p>
          ) : (
            <>
              <div className="flex items-end justify-between gap-1 h-32 overflow-x-auto">
                {tren.map((t) => (
                  <div key={t.tanggal} className="flex-1 min-w-[14px] flex flex-col items-center gap-1">
                    <div className="w-full flex items-end justify-center gap-[2px] h-24">
                      {/* min-height 2px supaya batang bernilai 0 tetap terlihat */}
                      <div
                        className="w-1/2 bg-success/70 rounded-t"
                        style={{ height: `${Math.max((t.masuk / puncak) * 100, 2)}%` }}
                        title={`${t.tanggal} · masuk ${t.masuk}`}
                      />
                      <div
                        className="w-1/2 bg-accent/70 rounded-t"
                        style={{ height: `${Math.max((t.keluar / puncak) * 100, 2)}%` }}
                        title={`${t.tanggal} · keluar ${t.keluar}`}
                      />
                    </div>
                    {tren.length <= 31 && (
                      <span className="text-[10px] font-mono text-muted">{t.tanggal.slice(8)}</span>
                    )}
                  </div>
                ))}
              </div>

              <div className="flex gap-4 mt-3 text-xs text-muted">
                <span className="flex items-center gap-1.5"><span className="w-2.5 h-2.5 rounded-sm bg-success/70" /> Masuk</span>
                <span className="flex items-center gap-1.5"><span className="w-2.5 h-2.5 rounded-sm bg-accent/70" /> Keluar</span>
                {!adaPergerakan && (
                  <span className="ml-auto">Belum ada transaksi pada rentang ini</span>
                )}
              </div>
            </>
          )}
        </div>
      </div>

      <div className="grid grid-cols-1 sm:grid-cols-2 gap-3 sm:gap-4">
        <Link to="/dashboard/stok-menipis" className="bg-surface border border-border rounded-2xl p-4 sm:p-5 hover:border-accent/40 transition-colors flex items-center justify-between">
          <div>
            <p className="font-display font-semibold mb-1">Stok Menipis</p>
            <p className="text-xs text-muted">Lihat barang yang perlu dipesan ulang</p>
          </div>
          <ArrowRight size={18} className="text-muted" />
        </Link>
        <Link to="/dashboard/aktivitas" className="bg-surface border border-border rounded-2xl p-4 sm:p-5 hover:border-accent/40 transition-colors flex items-center justify-between">
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
