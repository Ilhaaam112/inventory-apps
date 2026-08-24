import { useState, useEffect } from 'react'
import axios from 'axios'
import {
  ArrowDownToLine, ArrowUpFromLine, SlidersHorizontal, ArrowLeftRight,
} from 'lucide-react'
import Layout from '../components/Layout'
import { FilterBar, Field, inputClass } from '../components/FilterBar'

const gaya = {
  IN: { label: 'Barang masuk', icon: ArrowDownToLine, warna: 'text-success' },
  OUT: { label: 'Barang keluar', icon: ArrowUpFromLine, warna: 'text-accent' },
  ADJUSTMENT: { label: 'Penyesuaian', icon: SlidersHorizontal, warna: 'text-muted' },
  TRANSFER_IN: { label: 'Transfer masuk', icon: ArrowLeftRight, warna: 'text-success' },
  TRANSFER_OUT: { label: 'Transfer keluar', icon: ArrowLeftRight, warna: 'text-accent' },
}

function AktivitasTerbaru({ user, onLogout }) {
  const [limit, setLimit] = useState(30)
  const [rows, setRows] = useState([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    axios
      .get(`/api/dashboard/aktivitas?limit=${limit}`)
      .then((r) => setRows(r.data || []))
      .finally(() => setLoading(false))
  }, [limit])

  return (
    <Layout title="Aktivitas Terbaru" user={user} onLogout={onLogout}>
      <FilterBar onCetak={() => window.print()}>
        <Field label="JUMLAH BARIS">
          <select className={inputClass} value={limit} onChange={(e) => setLimit(Number(e.target.value))}>
            <option value={20}>20 terakhir</option>
            <option value={30}>30 terakhir</option>
            <option value={50}>50 terakhir</option>
            <option value={100}>100 terakhir</option>
          </select>
        </Field>
      </FilterBar>

      <div className="bg-surface border border-border rounded-2xl overflow-hidden">
        {loading ? (
          <p className="text-center py-10 text-muted text-sm">Memuat…</p>
        ) : rows.length === 0 ? (
          <p className="text-center py-10 text-muted text-sm">Belum ada pergerakan stok</p>
        ) : (
          <div className="divide-y divide-border">
            {rows.map((r, i) => {
              const g = gaya[r.type] || { label: r.type, icon: SlidersHorizontal, warna: 'text-muted' }
              const Icon = g.icon
              return (
                <div key={i} className="flex items-start gap-4 px-5 py-4 hover:bg-surface-soft transition-colors">
                  <div className={`mt-0.5 ${g.warna}`}>
                    <Icon size={17} />
                  </div>
                  <div className="flex-1 min-w-0">
                    <p className="text-sm">
                      <span className="font-medium">{r.nama_barang}</span>
                      <span className="text-muted"> · {g.label} di {r.nama_lokasi}</span>
                    </p>
                    <p className="text-xs text-muted mt-0.5 font-mono">
                      {r.waktu} · {r.referensi}
                      {r.nama_user ? ` · ${r.nama_user}` : ''}
                    </p>
                  </div>
                  <div className="text-right shrink-0">
                    <p className={`font-mono text-sm ${r.quantity >= 0 ? 'text-success' : 'text-accent'}`}>
                      {r.quantity > 0 ? `+${r.quantity}` : r.quantity}
                    </p>
                    <p className="text-xs text-muted font-mono">sisa {r.stock_after}</p>
                  </div>
                </div>
              )
            })}
          </div>
        )}
      </div>
    </Layout>
  )
}

export default AktivitasTerbaru
