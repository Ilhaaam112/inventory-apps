import { useState, useEffect, Fragment } from 'react'
import axios from 'axios'
import { Plus, Trash2, ChevronDown } from 'lucide-react'
import Layout from '../components/Layout'

const hariIni = () => new Date().toISOString().slice(0, 10)
const barisKosong = { barang_id: '', actual_stock: '' }

const input =
  'w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors'
const label = 'block text-xs font-mono text-muted mb-1.5'

function PenyesuaianStok({ user, onLogout }) {
  const [barangList, setBarangList] = useState([])
  const [lokasiList, setLokasiList] = useState([])
  const [stokGudang, setStokGudang] = useState({})
  const [riwayat, setRiwayat] = useState([])
  const [detail, setDetail] = useState(null)

  const [header, setHeader] = useState({ lokasi_id: '', tanggal: hariIni(), alasan: '' })
  const [rows, setRows] = useState([{ ...barisKosong }])
  const [saving, setSaving] = useState(false)

  const fetchRiwayat = async () => {
    const res = await axios.get('/api/stock-adjustment')
    setRiwayat(res.data || [])
  }

  useEffect(() => {
    axios.get('/api/barang').then((r) => setBarangList(r.data || []))
    axios.get('/api/lokasi').then((r) => setLokasiList(r.data || []))
    fetchRiwayat()
  }, [])

  useEffect(() => {
    if (!header.lokasi_id) return setStokGudang({})
    axios.get(`/api/warehouse-stocks?lokasi_id=${header.lokasi_id}`).then((r) => {
      const map = {}
      ;(r.data || []).forEach((w) => { map[w.barang_id] = w.quantity })
      setStokGudang(map)
    })
  }, [header.lokasi_id])

  const setRow = (i, field, value) =>
    setRows(rows.map((r, idx) => (idx === i ? { ...r, [field]: value } : r)))
  const addRow = () => setRows([...rows, { ...barisKosong }])
  const removeRow = (i) =>
    setRows(rows.length === 1 ? [{ ...barisKosong }] : rows.filter((_, idx) => idx !== i))

  const stokSistem = (barangId) => stokGudang[parseInt(barangId)] ?? 0
  const selisih = (row) =>
    row.barang_id && row.actual_stock !== '' ? parseInt(row.actual_stock) - stokSistem(row.barang_id) : null

  const handleSubmit = async (e) => {
    e.preventDefault()
    const details = rows
      .filter((r) => r.barang_id && r.actual_stock !== '')
      .map((r) => ({ barang_id: parseInt(r.barang_id), actual_stock: parseInt(r.actual_stock) }))

    if (details.length === 0) {
      alert('Isi minimal satu baris barang dan stok fisiknya.')
      return
    }

    setSaving(true)
    try {
      const res = await axios.post('/api/stock-adjustment', {
        lokasi_id: parseInt(header.lokasi_id),
        user_id: user?.id || null,
        tanggal: header.tanggal,
        alasan: header.alasan,
        details,
      })
      alert(`Penyesuaian ${res.data.code} tersimpan. Stok sistem sudah disamakan dengan stok fisik.`)
      setHeader({ lokasi_id: '', tanggal: hariIni(), alasan: '' })
      setRows([{ ...barisKosong }])
      fetchRiwayat()
    } catch (err) {
      alert('Gagal simpan: ' + (err.response?.data?.error || err.message))
    } finally {
      setSaving(false)
    }
  }

  const lihatDetail = async (id) => {
    if (detail?.id === id) return setDetail(null)
    const res = await axios.get(`/api/stock-adjustment/${id}`)
    setDetail(res.data)
  }

  return (
    <Layout title="Penyesuaian Stok" user={user} onLogout={onLogout}>
      <form onSubmit={handleSubmit} className="bg-surface border border-border rounded-2xl p-6 mb-8">
        <h3 className="font-display font-semibold mb-1">Penyesuaian Stok</h3>
        <p className="text-xs text-muted mb-4">
          Masukkan hasil hitung fisik. Sistem menghitung sendiri selisihnya terhadap stok tercatat.
        </p>

        <div className="grid sm:grid-cols-3 gap-4 mb-4">
          <div>
            <label className={label}>GUDANG</label>
            <select
              className={input} required
              value={header.lokasi_id}
              onChange={(e) => setHeader({ ...header, lokasi_id: e.target.value })}
            >
              <option value="">Pilih gudang</option>
              {lokasiList.map((l) => (
                <option key={l.id} value={l.id}>{l.nama_lokasi}</option>
              ))}
            </select>
          </div>
          <div>
            <label className={label}>TANGGAL</label>
            <input
              type="date" className={input} required
              value={header.tanggal}
              onChange={(e) => setHeader({ ...header, tanggal: e.target.value })}
            />
          </div>
          <div>
            <label className={label}>ALASAN</label>
            <input
              type="text" className={input} required placeholder="Barang rusak / hilang / salah catat"
              value={header.alasan}
              onChange={(e) => setHeader({ ...header, alasan: e.target.value })}
            />
          </div>
        </div>

        <div className="border-t border-border pt-4">
          <div className="flex items-center justify-between mb-3">
            <p className="text-xs font-mono text-muted tracking-widest uppercase">Detail Barang</p>
            <button type="button" onClick={addRow} className="flex items-center gap-1.5 text-xs text-accent hover:text-accent-soft">
              <Plus size={14} /> Tambah baris
            </button>
          </div>

          <div className="space-y-3">
            {rows.map((row, i) => {
              const diff = selisih(row)
              return (
                <div key={i} className="grid grid-cols-12 gap-3 items-end">
                  <div className="col-span-12 sm:col-span-5">
                    <select
                      className={input}
                      value={row.barang_id}
                      onChange={(e) => setRow(i, 'barang_id', e.target.value)}
                      disabled={!header.lokasi_id}
                    >
                      <option value="">{header.lokasi_id ? 'Pilih barang' : 'Pilih gudang dulu'}</option>
                      {barangList.map((b) => (
                        <option key={b.id} value={b.id}>{b.nama}</option>
                      ))}
                    </select>
                  </div>
                  <div className="col-span-4 sm:col-span-2">
                    <input
                      type="text" readOnly className={`${input} text-muted`}
                      value={row.barang_id ? `sistem ${stokSistem(row.barang_id)}` : 'sistem –'}
                    />
                  </div>
                  <div className="col-span-4 sm:col-span-2">
                    <input
                      type="number" min="0" placeholder="Stok fisik" className={input}
                      value={row.actual_stock}
                      onChange={(e) => setRow(i, 'actual_stock', e.target.value)}
                    />
                  </div>
                  <div className="col-span-3 sm:col-span-2 pb-2.5">
                    <span className={`text-sm font-mono ${diff === null ? 'text-muted' : diff < 0 ? 'text-accent' : diff > 0 ? 'text-success' : 'text-muted'}`}>
                      {diff === null ? '—' : diff > 0 ? `+${diff}` : diff}
                    </span>
                  </div>
                  <div className="col-span-1 flex justify-end">
                    <button type="button" onClick={() => removeRow(i)} className="p-2.5 text-muted hover:text-accent transition-colors">
                      <Trash2 size={16} />
                    </button>
                  </div>
                </div>
              )
            })}
          </div>
        </div>

        <button
          type="submit"
          disabled={saving}
          className="mt-5 bg-accent hover:bg-accent-soft disabled:opacity-50 transition-colors rounded-lg px-5 py-2.5 font-medium text-sm"
        >
          {saving ? 'Menyimpan…' : 'Simpan Penyesuaian'}
        </button>
      </form>

      <div className="bg-surface border border-border rounded-2xl overflow-hidden">
        <div className="px-5 py-4 border-b border-border">
          <h3 className="font-display font-semibold">Riwayat Penyesuaian</h3>
        </div>
        <table className="w-full text-sm">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">Kode</th>
              <th className="px-5 py-3">Tanggal</th>
              <th className="px-5 py-3">Gudang</th>
              <th className="px-5 py-3">Alasan</th>
              <th className="px-5 py-3">Item</th>
              <th className="px-5 py-3"></th>
            </tr>
          </thead>
          <tbody>
            {riwayat.length === 0 ? (
              <tr><td colSpan="6" className="text-center py-8 text-muted">Belum ada penyesuaian</td></tr>
            ) : (
              riwayat.map((t) => (
                <Fragment key={t.id}>
                  <tr className="border-b border-border hover:bg-surface-soft transition-colors">
                    <td className="px-5 py-3 font-mono text-accent">{t.code}</td>
                    <td className="px-5 py-3 font-mono text-muted">{t.tanggal}</td>
                    <td className="px-5 py-3">{t.nama_lokasi || '-'}</td>
                    <td className="px-5 py-3 text-muted">{t.alasan}</td>
                    <td className="px-5 py-3 font-mono">{t.total_item}</td>
                    <td className="px-5 py-3 text-right">
                      <button onClick={() => lihatDetail(t.id)} className="text-xs text-muted hover:text-ink inline-flex items-center gap-1">
                        Detail <ChevronDown size={13} />
                      </button>
                    </td>
                  </tr>
                  {detail?.id === t.id && (
                    <tr className="border-b border-border bg-surface-soft">
                      <td colSpan="6" className="px-5 py-4">
                        <div className="space-y-1">
                          {detail.details?.map((d) => (
                            <div key={d.id} className="flex justify-between text-xs">
                              <span>{d.nama_barang}</span>
                              <span className="font-mono text-muted">
                                {d.system_stock} → {d.actual_stock}{' '}
                                <span className={d.difference < 0 ? 'text-accent' : 'text-success'}>
                                  ({d.difference > 0 ? `+${d.difference}` : d.difference})
                                </span>
                              </span>
                            </div>
                          ))}
                        </div>
                      </td>
                    </tr>
                  )}
                </Fragment>
              ))
            )}
          </tbody>
        </table>
      </div>
    </Layout>
  )
}

export default PenyesuaianStok
