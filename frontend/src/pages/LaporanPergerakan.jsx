import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'
import { FilterBar, Field, TabelKosong, inputClass, awalBulan, hariIni } from '../components/FilterBar'

function LaporanPergerakan({ user, onLogout }) {
  const [lokasiList, setLokasiList] = useState([])
  const [filter, setFilter] = useState({ start: awalBulan(), end: hariIni(), lokasi_id: '' })
  const [rows, setRows] = useState([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    axios.get('/api/lokasi').then((r) => setLokasiList(r.data || []))
  }, [])

  useEffect(() => {
    setLoading(true)
    const q = new URLSearchParams({
      start: filter.start,
      end: filter.end,
      lokasi_id: filter.lokasi_id || 0,
    })
    axios
      .get(`/api/laporan/pergerakan?${q}`)
      .then((r) => setRows(r.data || []))
      .finally(() => setLoading(false))
  }, [filter])

  const bergerak = rows.filter((r) => r.masuk || r.keluar || r.penyesuaian)

  return (
    <Layout title="Laporan Pergerakan Stok" user={user} onLogout={onLogout}>
      <FilterBar onCetak={() => window.print()}>
        <Field label="DARI TANGGAL">
          <input type="date" className={inputClass} value={filter.start}
            onChange={(e) => setFilter({ ...filter, start: e.target.value })} />
        </Field>
        <Field label="SAMPAI TANGGAL">
          <input type="date" className={inputClass} value={filter.end}
            onChange={(e) => setFilter({ ...filter, end: e.target.value })} />
        </Field>
        <Field label="GUDANG">
          <select className={inputClass} value={filter.lokasi_id}
            onChange={(e) => setFilter({ ...filter, lokasi_id: e.target.value })}>
            <option value="">Semua gudang</option>
            {lokasiList.map((l) => <option key={l.id} value={l.id}>{l.nama_lokasi}</option>)}
          </select>
        </Field>
      </FilterBar>

      <p className="text-xs text-muted mb-4 print:hidden">
        Saldo awal + masuk − keluar + penyesuaian = saldo akhir. Baris tanpa mutasi tetap
        ditampilkan supaya posisi stok terbaca utuh ({bergerak.length} dari {rows.length} baris bergerak).
      </p>

      <div className="bg-surface border border-border rounded-2xl overflow-x-auto">
        <table className="w-full text-sm min-w-[820px]">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">Barang</th>
              <th className="px-5 py-3">Gudang</th>
              <th className="px-5 py-3 text-right">Saldo Awal</th>
              <th className="px-5 py-3 text-right">Masuk</th>
              <th className="px-5 py-3 text-right">Keluar</th>
              <th className="px-5 py-3 text-right">Penyesuaian</th>
              <th className="px-5 py-3 text-right">Saldo Akhir</th>
            </tr>
          </thead>
          <tbody>
            {loading ? (
              <TabelKosong colSpan={7} pesan="Memuat…" />
            ) : rows.length === 0 ? (
              <TabelKosong colSpan={7} />
            ) : (
              rows.map((r) => (
                <tr key={`${r.lokasi_id}-${r.barang_id}`} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                  <td className="px-5 py-3 font-medium">{r.nama_barang}</td>
                  <td className="px-5 py-3 text-muted">{r.nama_lokasi}</td>
                  <td className="px-5 py-3 text-right font-mono text-muted">{r.saldo_awal}</td>
                  <td className="px-5 py-3 text-right font-mono text-success">{r.masuk || '-'}</td>
                  <td className="px-5 py-3 text-right font-mono text-accent">{r.keluar || '-'}</td>
                  <td className="px-5 py-3 text-right font-mono">
                    {r.penyesuaian ? (r.penyesuaian > 0 ? `+${r.penyesuaian}` : r.penyesuaian) : '-'}
                  </td>
                  <td className="px-5 py-3 text-right font-mono font-medium">{r.saldo_akhir}</td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </Layout>
  )
}

export default LaporanPergerakan
