import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'
import { FilterBar, Field, TabelKosong, inputClass } from '../components/FilterBar'

function StokMenipis({ user, onLogout }) {
  const [lokasiList, setLokasiList] = useState([])
  const [filter, setFilter] = useState({ lokasi_id: '', minimum: '' })
  const [rows, setRows] = useState([])
  const [loading, setLoading] = useState(false)

  useEffect(() => {
    axios.get('/api/lokasi').then((r) => setLokasiList(r.data || []))
  }, [])

  useEffect(() => {
    setLoading(true)
    const q = new URLSearchParams({
      lokasi_id: filter.lokasi_id || 0,
      minimum: filter.minimum || 0,
      limit: 100,
    })
    axios
      .get(`/api/dashboard/stok-menipis?${q}`)
      .then((r) => setRows(r.data || []))
      .finally(() => setLoading(false))
  }, [filter])

  const habis = rows.filter((r) => r.habis).length

  return (
    <Layout title="Stok Menipis" user={user} onLogout={onLogout}>
      <FilterBar onCetak={() => window.print()}>
        <Field label="GUDANG">
          <select className={inputClass} value={filter.lokasi_id}
            onChange={(e) => setFilter({ ...filter, lokasi_id: e.target.value })}>
            <option value="">Semua gudang</option>
            {lokasiList.map((l) => <option key={l.id} value={l.id}>{l.nama_lokasi}</option>)}
          </select>
        </Field>
        <Field label="AMBANG CADANGAN">
          <input
            type="number" min="0" placeholder="0" className={inputClass}
            value={filter.minimum}
            onChange={(e) => setFilter({ ...filter, minimum: e.target.value })}
          />
        </Field>
      </FilterBar>

      <p className="text-xs text-muted mb-4 print:hidden">
        Batas per barang diatur di Data Barang lewat kolom Stok Minimum. Ambang cadangan
        di atas hanya berlaku untuk barang yang batasnya belum diisi.
        {rows.length > 0 && ` Saat ini ${habis} habis dari ${rows.length} baris.`}
      </p>

      <div className="bg-surface border border-border rounded-2xl overflow-hidden">
        <table className="w-full text-sm">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">Barang</th>
              <th className="px-5 py-3">Gudang</th>
              <th className="px-5 py-3 text-right">Stok</th>
              <th className="px-5 py-3 text-right">Minimum</th>
              <th className="px-5 py-3 text-right">Kurang</th>
              <th className="px-5 py-3">Status</th>
            </tr>
          </thead>
          <tbody>
            {loading ? (
              <TabelKosong colSpan={6} pesan="Memuat…" />
            ) : rows.length === 0 ? (
              <TabelKosong colSpan={6} pesan="Semua stok masih di atas batas minimum" />
            ) : (
              rows.map((r) => (
                <tr key={`${r.lokasi_id}-${r.barang_id}`} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                  <td className="px-5 py-3 font-medium">{r.nama_barang}</td>
                  <td className="px-5 py-3 text-muted">{r.nama_lokasi}</td>
                  <td className="px-5 py-3 text-right font-mono text-accent">
                    {r.quantity} <span className="text-muted">{r.nama_satuan || ''}</span>
                  </td>
                  <td className="px-5 py-3 text-right font-mono text-muted">{r.stok_minimum || '-'}</td>
                  <td className="px-5 py-3 text-right font-mono">{r.kekurangan || '-'}</td>
                  <td className="px-5 py-3">
                    <span className={`text-[10px] font-mono border rounded-full px-2 py-0.5 ${
                      r.habis ? 'border-accent/40 text-accent' : 'border-border text-muted'
                    }`}>
                      {r.habis ? 'HABIS' : 'MENIPIS'}
                    </span>
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

export default StokMenipis
