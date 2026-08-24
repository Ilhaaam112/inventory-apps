import { useState, useEffect, Fragment } from 'react'
import axios from 'axios'
import { Plus, Trash2, ChevronDown, ArrowRight } from 'lucide-react'
import Layout from '../components/Layout'

const hariIni = () => new Date().toISOString().slice(0, 10)
const barisKosong = { barang_id: '', quantity: '' }

const input =
  'w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors'
const label = 'block text-xs font-mono text-muted mb-1.5'
const labelBaris = 'sm:hidden block text-[11px] font-mono text-muted mb-1'

function TransferGudang({ user, onLogout }) {
  const [barangList, setBarangList] = useState([])
  const [lokasiList, setLokasiList] = useState([])
  const [stokAsal, setStokAsal] = useState({})
  const [riwayat, setRiwayat] = useState([])
  const [detail, setDetail] = useState(null)

  const [header, setHeader] = useState({ from_lokasi_id: '', to_lokasi_id: '', tanggal: hariIni(), catatan: '' })
  const [rows, setRows] = useState([{ ...barisKosong }])
  const [saving, setSaving] = useState(false)

  const fetchRiwayat = async () => {
    const res = await axios.get('/api/stock-transfer')
    setRiwayat(res.data || [])
  }

  useEffect(() => {
    axios.get('/api/barang').then((r) => setBarangList(r.data || []))
    axios.get('/api/lokasi').then((r) => setLokasiList(r.data || []))
    fetchRiwayat()
  }, [])

  useEffect(() => {
    if (!header.from_lokasi_id) return setStokAsal({})
    axios.get(`/api/warehouse-stocks?lokasi_id=${header.from_lokasi_id}`).then((r) => {
      const map = {}
      ;(r.data || []).forEach((w) => { map[w.barang_id] = w.quantity })
      setStokAsal(map)
    })
  }, [header.from_lokasi_id])

  const setRow = (i, field, value) =>
    setRows(rows.map((r, idx) => (idx === i ? { ...r, [field]: value } : r)))
  const addRow = () => setRows([...rows, { ...barisKosong }])
  const removeRow = (i) =>
    setRows(rows.length === 1 ? [{ ...barisKosong }] : rows.filter((_, idx) => idx !== i))

  const tersedia = (barangId) => stokAsal[parseInt(barangId)] ?? 0
  const gudangSama = header.from_lokasi_id && header.from_lokasi_id === header.to_lokasi_id
  const adaYangKurang = rows.some(
    (r) => r.barang_id && r.quantity && parseInt(r.quantity) > tersedia(r.barang_id)
  )

  const handleSubmit = async (e) => {
    e.preventDefault()
    const details = rows
      .filter((r) => r.barang_id && r.quantity)
      .map((r) => ({ barang_id: parseInt(r.barang_id), quantity: parseInt(r.quantity) }))

    if (details.length === 0) {
      alert('Isi minimal satu baris barang dan jumlahnya.')
      return
    }

    setSaving(true)
    try {
      const res = await axios.post('/api/stock-transfer', {
        from_lokasi_id: parseInt(header.from_lokasi_id),
        to_lokasi_id: parseInt(header.to_lokasi_id),
        user_id: user?.id || null,
        tanggal: header.tanggal,
        catatan: header.catatan,
        details,
      })
      alert(`Transfer ${res.data.code} tersimpan. Stok kedua gudang sudah diperbarui.`)
      setHeader({ from_lokasi_id: '', to_lokasi_id: '', tanggal: hariIni(), catatan: '' })
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
    const res = await axios.get(`/api/stock-transfer/${id}`)
    setDetail(res.data)
  }

  const RincianDetail = () => (
    <div className="space-y-1">
      {detail?.catatan && <p className="text-xs text-muted mb-2">{detail.catatan}</p>}
      {detail?.details?.map((d) => (
        <div key={d.id} className="flex justify-between gap-3 text-xs">
          <span className="min-w-0 break-words">{d.nama_barang}</span>
          <span className="font-mono text-muted shrink-0">{d.quantity}</span>
        </div>
      ))}
    </div>
  )

  return (
    <Layout title="Transfer Gudang" user={user} onLogout={onLogout}>
      <form onSubmit={handleSubmit} className="bg-surface border border-border rounded-2xl p-4 sm:p-6 mb-6">
        <h3 className="font-display font-semibold mb-1">Transfer Antar Gudang</h3>
        <p className="text-xs text-muted mb-4">
          Satu transfer tercatat sebagai dua pergerakan: keluar di gudang asal, masuk di gudang tujuan.
        </p>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-4">
          <div>
            <label className={label}>GUDANG ASAL</label>
            <select className={input} required value={header.from_lokasi_id}
              onChange={(e) => setHeader({ ...header, from_lokasi_id: e.target.value })}>
              <option value="">Pilih gudang asal</option>
              {lokasiList.map((l) => <option key={l.id} value={l.id}>{l.nama_lokasi}</option>)}
            </select>
          </div>
          <div>
            <label className={label}>GUDANG TUJUAN</label>
            <select className={`${input} ${gudangSama ? 'border-accent' : ''}`} required value={header.to_lokasi_id}
              onChange={(e) => setHeader({ ...header, to_lokasi_id: e.target.value })}>
              <option value="">Pilih gudang tujuan</option>
              {lokasiList.map((l) => <option key={l.id} value={l.id}>{l.nama_lokasi}</option>)}
            </select>
            {gudangSama && <p className="mt-1 text-[11px] text-accent">Pilih gudang tujuan yang berbeda.</p>}
          </div>
          <div>
            <label className={label}>TANGGAL</label>
            <input type="date" className={input} required value={header.tanggal}
              onChange={(e) => setHeader({ ...header, tanggal: e.target.value })} />
          </div>
        </div>

        <div className="mb-4">
          <label className={label}>CATATAN</label>
          <input type="text" className={input} value={header.catatan}
            onChange={(e) => setHeader({ ...header, catatan: e.target.value })} />
        </div>

        <div className="border-t border-border pt-4">
          <div className="flex items-center justify-between mb-3">
            <p className="text-xs font-mono text-muted tracking-widest uppercase">Detail Barang</p>
            <button type="button" onClick={addRow} className="flex items-center gap-1.5 text-xs text-accent hover:text-accent-soft py-1">
              <Plus size={14} /> Tambah baris
            </button>
          </div>

          <div className="space-y-4 sm:space-y-3">
            {rows.map((row, i) => {
              const stok = tersedia(row.barang_id)
              const kurang = row.barang_id && row.quantity && parseInt(row.quantity) > stok
              return (
                <div key={i}
                  className="rounded-xl border border-border p-3 sm:border-0 sm:p-0 sm:rounded-none grid grid-cols-1 sm:grid-cols-12 gap-3 sm:items-start">
                  <div className="sm:col-span-7">
                    <label className={labelBaris}>BARANG</label>
                    <select className={input} value={row.barang_id} disabled={!header.from_lokasi_id}
                      onChange={(e) => setRow(i, 'barang_id', e.target.value)}>
                      <option value="">{header.from_lokasi_id ? 'Pilih barang' : 'Pilih gudang asal dulu'}</option>
                      {barangList.map((b) => <option key={b.id} value={b.id}>{b.nama}</option>)}
                    </select>
                  </div>
                  <div className="sm:col-span-3">
                    <label className={labelBaris}>JUMLAH</label>
                    <input type="number" min="1" placeholder="Jumlah"
                      className={`${input} ${kurang ? 'border-accent' : ''}`}
                      value={row.quantity} onChange={(e) => setRow(i, 'quantity', e.target.value)} />
                    {row.barang_id && (
                      <p className={`mt-1 text-[11px] font-mono ${kurang ? 'text-accent' : 'text-muted'}`}>
                        {kurang ? `stok kurang, tersedia ${stok}` : `tersedia ${stok} di gudang asal`}
                      </p>
                    )}
                  </div>
                  <div className="sm:col-span-2 flex sm:justify-end">
                    <button type="button" onClick={() => removeRow(i)}
                      className="w-full sm:w-auto flex items-center justify-center gap-2 border border-border sm:border-0 rounded-lg py-2 sm:py-0 sm:p-2.5 text-muted hover:text-accent transition-colors">
                      <Trash2 size={16} />
                      <span className="sm:hidden text-xs">Hapus baris</span>
                    </button>
                  </div>
                </div>
              )
            })}
          </div>
        </div>

        <button type="submit" disabled={saving || gudangSama || adaYangKurang}
          className="w-full sm:w-auto mt-5 bg-accent hover:bg-accent-soft disabled:opacity-50 transition-colors rounded-lg px-5 py-2.5 font-medium text-sm">
          {saving ? 'Menyimpan…' : 'Simpan Transfer'}
        </button>
      </form>

      <h3 className="font-display font-semibold mb-3 md:hidden">Riwayat Transfer</h3>

      {/* Mobile: kartu */}
      <div className="md:hidden space-y-3">
        {riwayat.length === 0 ? (
          <div className="bg-surface border border-border rounded-2xl p-8 text-center text-muted text-sm">Belum ada transfer</div>
        ) : (
          riwayat.map((t) => (
            <div key={t.id} className="bg-surface border border-border rounded-2xl p-4">
              <div className="flex items-start justify-between gap-3">
                <div className="min-w-0">
                  <p className="font-mono text-accent text-sm">{t.code}</p>
                  <p className="text-xs text-muted font-mono mt-0.5">{t.tanggal}</p>
                </div>
                <p className="font-mono shrink-0">{t.total_qty}</p>
              </div>
              <p className="text-xs mt-2 inline-flex items-center gap-1.5 flex-wrap">
                <span className="break-words">{t.nama_from}</span>
                <ArrowRight size={12} className="text-muted shrink-0" />
                <span className="break-words">{t.nama_to}</span>
              </p>
              <button onClick={() => lihatDetail(t.id)}
                className="mt-3 w-full flex items-center justify-center gap-1 text-xs border border-border rounded-lg py-2 text-muted hover:text-ink transition-colors">
                {detail?.id === t.id ? 'Tutup detail' : 'Lihat detail'} <ChevronDown size={13} />
              </button>
              {detail?.id === t.id && (
                <div className="mt-3 pt-3 border-t border-border"><RincianDetail /></div>
              )}
            </div>
          ))
        )}
      </div>

      {/* Tablet & desktop */}
      <div className="hidden md:block bg-surface border border-border rounded-2xl overflow-hidden">
        <div className="px-5 py-4 border-b border-border">
          <h3 className="font-display font-semibold">Riwayat Transfer</h3>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full text-sm min-w-[720px]">
            <thead>
              <tr className="border-b border-border text-left text-xs font-mono text-muted">
                <th className="px-5 py-3">Kode</th>
                <th className="px-5 py-3">Tanggal</th>
                <th className="px-5 py-3">Rute</th>
                <th className="px-5 py-3 text-right">Qty</th>
                <th className="px-5 py-3"></th>
              </tr>
            </thead>
            <tbody>
              {riwayat.length === 0 ? (
                <tr><td colSpan="5" className="text-center py-8 text-muted">Belum ada transfer</td></tr>
              ) : (
                riwayat.map((t) => (
                  <Fragment key={t.id}>
                    <tr className="border-b border-border hover:bg-surface-soft transition-colors">
                      <td className="px-5 py-3 font-mono text-accent">{t.code}</td>
                      <td className="px-5 py-3 font-mono text-muted whitespace-nowrap">{t.tanggal}</td>
                      <td className="px-5 py-3">
                        <span className="inline-flex items-center gap-2">
                          {t.nama_from} <ArrowRight size={13} className="text-muted" /> {t.nama_to}
                        </span>
                      </td>
                      <td className="px-5 py-3 text-right font-mono">{t.total_qty}</td>
                      <td className="px-5 py-3 text-right whitespace-nowrap">
                        <button onClick={() => lihatDetail(t.id)} className="text-xs text-muted hover:text-ink inline-flex items-center gap-1">
                          Detail <ChevronDown size={13} />
                        </button>
                      </td>
                    </tr>
                    {detail?.id === t.id && (
                      <tr className="border-b border-border bg-surface-soft">
                        <td colSpan="5" className="px-5 py-4"><RincianDetail /></td>
                      </tr>
                    )}
                  </Fragment>
                ))
              )}
            </tbody>
          </table>
        </div>
      </div>
    </Layout>
  )
}

export default TransferGudang
