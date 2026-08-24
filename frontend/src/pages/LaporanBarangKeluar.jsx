import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'
import { FilterBar, Field, TabelKosong, inputClass, awalBulan, hariIni } from '../components/FilterBar'

function LaporanBarangKeluar({ user, onLogout }) {
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
      .get(`/api/laporan/barang-keluar?${q}`)
      .then((r) => setRows(r.data || []))
      .finally(() => setLoading(false))
  }, [filter])

  const totalQty = rows.reduce((t, r) => t + r.quantity, 0)

  return (
    <Layout title="Laporan Barang Keluar" user={user} onLogout={onLogout}>
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

      <div className="bg-surface border border-border rounded-2xl p-4 sm:p-5 mb-6">
        <p className="text-xs font-mono text-muted mb-1">TOTAL UNIT KELUAR</p>
        <p className="font-display text-2xl text-accent">{totalQty.toLocaleString('id-ID')}</p>
      </div>

      <div className="bg-surface border border-border rounded-2xl overflow-x-auto">
        <table className="w-full text-sm min-w-[760px]">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">Tanggal</th>
              <th className="px-5 py-3">Kode</th>
              <th className="px-5 py-3">Gudang</th>
              <th className="px-5 py-3">Tujuan</th>
              <th className="px-5 py-3">Barang</th>
              <th className="px-5 py-3 text-right">Qty</th>
            </tr>
          </thead>
          <tbody>
            {loading ? (
              <TabelKosong colSpan={6} pesan="Memuat…" />
            ) : rows.length === 0 ? (
              <TabelKosong colSpan={6} />
            ) : (
              rows.map((r, i) => (
                <tr key={i} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                  <td className="px-5 py-3 font-mono text-muted">{r.tanggal}</td>
                  <td className="px-5 py-3 font-mono text-accent">{r.code}</td>
                  <td className="px-5 py-3 text-muted">{r.nama_lokasi}</td>
                  <td className="px-5 py-3 text-muted">{r.tujuan || '-'}</td>
                  <td className="px-5 py-3 font-medium">{r.nama_barang}</td>
                  <td className="px-5 py-3 text-right font-mono text-accent">
                    {r.quantity} <span className="text-muted">{r.nama_satuan || ''}</span>
                  </td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </Layout>
  )
}

export default LaporanBarangKeluar
