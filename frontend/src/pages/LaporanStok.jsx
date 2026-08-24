import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'
import { FilterBar, Field, TabelKosong, inputClass, rupiah } from '../components/FilterBar'

function LaporanStok({ user, onLogout }) {
  const [lokasiList, setLokasiList] = useState([])
  const [kategoriList, setKategoriList] = useState([])
  const [filter, setFilter] = useState({ lokasi_id: '', kategori_id: '' })
  const [rows, setRows] = useState([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    axios.get('/api/lokasi').then((r) => setLokasiList(r.data || []))
    axios.get('/api/kategori').then((r) => setKategoriList(r.data || []))
  }, [])

  useEffect(() => {
    setLoading(true)
    const q = new URLSearchParams({
      lokasi_id: filter.lokasi_id || 0,
      kategori_id: filter.kategori_id || 0,
    })
    axios
      .get(`/api/laporan/stok?${q}`)
      .then((r) => setRows(r.data || []))
      .finally(() => setLoading(false))
  }, [filter])

  const totalQty = rows.reduce((t, r) => t + r.quantity, 0)
  const totalNilai = rows.reduce((t, r) => t + r.nilai, 0)

  return (
    <Layout title="Laporan Stok" user={user} onLogout={onLogout}>
      <FilterBar onCetak={() => window.print()}>
        <Field label="GUDANG">
          <select
            className={inputClass}
            value={filter.lokasi_id}
            onChange={(e) => setFilter({ ...filter, lokasi_id: e.target.value })}
          >
            <option value="">Semua gudang</option>
            {lokasiList.map((l) => (
              <option key={l.id} value={l.id}>{l.nama_lokasi}</option>
            ))}
          </select>
        </Field>
        <Field label="KATEGORI">
          <select
            className={inputClass}
            value={filter.kategori_id}
            onChange={(e) => setFilter({ ...filter, kategori_id: e.target.value })}
          >
            <option value="">Semua kategori</option>
            {kategoriList.map((k) => (
              <option key={k.id} value={k.id}>{k.nama_kategori}</option>
            ))}
          </select>
        </Field>
      </FilterBar>

      <div className="grid sm:grid-cols-2 gap-4 mb-6">
        <div className="bg-surface border border-border rounded-2xl p-4 sm:p-5">
          <p className="text-xs font-mono text-muted mb-1">TOTAL UNIT</p>
          <p className="font-display text-2xl">{totalQty.toLocaleString('id-ID')}</p>
        </div>
        <div className="bg-surface border border-border rounded-2xl p-4 sm:p-5">
          <p className="text-xs font-mono text-muted mb-1">NILAI PERSEDIAAN</p>
          <p className="font-display text-2xl text-accent">{rupiah(totalNilai)}</p>
        </div>
      </div>

      <div className="bg-surface border border-border rounded-2xl overflow-x-auto">
        <table className="w-full text-sm min-w-[720px]">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">Barang</th>
              <th className="px-5 py-3">Kategori</th>
              <th className="px-5 py-3">Gudang</th>
              <th className="px-5 py-3 text-right">Stok</th>
              <th className="px-5 py-3 text-right">Harga</th>
              <th className="px-5 py-3 text-right">Nilai</th>
            </tr>
          </thead>
          <tbody>
            {loading ? (
              <TabelKosong colSpan={6} pesan="Memuat…" />
            ) : rows.length === 0 ? (
              <TabelKosong colSpan={6} />
            ) : (
              rows.map((r) => (
                <tr key={`${r.lokasi_id}-${r.barang_id}`} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                  <td className="px-5 py-3 font-medium">{r.nama_barang}</td>
                  <td className="px-5 py-3 text-muted">{r.nama_kategori || '-'}</td>
                  <td className="px-5 py-3 text-muted">{r.nama_lokasi}</td>
                  <td className="px-5 py-3 text-right font-mono text-success">
                    {r.quantity} <span className="text-muted">{r.nama_satuan || ''}</span>
                  </td>
                  <td className="px-5 py-3 text-right font-mono text-muted">{rupiah(r.harga)}</td>
                  <td className="px-5 py-3 text-right font-mono text-accent">{rupiah(r.nilai)}</td>
                </tr>
              ))
            )}
          </tbody>
        </table>
      </div>
    </Layout>
  )
}

export default LaporanStok
