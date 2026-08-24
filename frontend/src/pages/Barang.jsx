import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'

const kosong = { nama: '', harga: '', stok_minimum: '', kategori_id: '', satuan_id: '' }

function Barang({ user, onLogout }) {
  const [barangList, setBarangList] = useState([])
  const [form, setForm] = useState(kosong)
  const [editId, setEditId] = useState(null)
  const [kategoriList, setKategoriList] = useState([])
  const [satuanList, setSatuanList] = useState([])

  const fetchBarang = async () => {
    try {
      const res = await axios.get('/api/barang')
      setBarangList(res.data || [])
    } catch (err) {
      console.error('Gagal ambil data:', err)
    }
  }

  useEffect(() => {
    fetchBarang()
    axios.get('/api/kategori').then((res) => setKategoriList(res.data || []))
    axios.get('/api/satuan').then((res) => setSatuanList(res.data || []))
  }, [])

  const handleChange = (e) => setForm({ ...form, [e.target.name]: e.target.value })

  const handleSubmit = async (e) => {
    e.preventDefault()
    const payload = {
      nama: form.nama,
      harga: parseFloat(form.harga),
      stok_minimum: form.stok_minimum ? parseInt(form.stok_minimum) : 0,
      kategori_id: form.kategori_id ? parseInt(form.kategori_id) : null,
      satuan_id: form.satuan_id ? parseInt(form.satuan_id) : null,
    }
    try {
      if (editId) {
        await axios.put(`/api/barang/${editId}`, payload)
        setEditId(null)
      } else {
        await axios.post('/api/barang', payload)
      }
      setForm(kosong)
      fetchBarang()
    } catch (err) {
      alert('Gagal simpan data: ' + (err.response?.data?.error || err.message))
    }
  }

  const handleEdit = (b) => {
    setForm({
      nama: b.nama,
      harga: b.harga,
      stok_minimum: b.stok_minimum ?? '',
      kategori_id: b.kategori_id || '',
      satuan_id: b.satuan_id || '',
    })
    setEditId(b.id)
    window.scrollTo({ top: 0, behavior: 'smooth' })
  }

  const handleDelete = async (id) => {
    if (!confirm('Yakin mau hapus barang ini?')) return
    try {
      await axios.delete(`/api/barang/${id}`)
      fetchBarang()
    } catch {
      alert('Gagal hapus data')
    }
  }

  const handleCancelEdit = () => {
    setEditId(null)
    setForm(kosong)
  }

  const input = 'w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors'
  const label = 'block text-xs font-mono text-muted mb-1.5'
  const btnEdit = 'flex-1 sm:flex-none text-xs border border-border rounded-lg px-3 py-2 text-muted hover:text-ink transition-colors'
  const btnHapus = 'flex-1 sm:flex-none text-xs border border-accent/40 rounded-lg px-3 py-2 text-accent hover:bg-accent/10 transition-colors'

  return (
    <Layout title="Data Barang" user={user} onLogout={onLogout}>
      <form onSubmit={handleSubmit} className="bg-surface border border-border rounded-2xl p-4 sm:p-6 mb-6">
        <h3 className="font-display font-semibold mb-1">
          {editId ? 'Edit Barang' : 'Tambah Barang Baru'}
        </h3>
        <p className="text-xs text-muted mb-4">
          Stok tidak diisi di sini. Jumlah stok hanya berubah lewat menu Transaksi.
        </p>

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
          <div>
            <label className={label}>NAMA</label>
            <input type="text" name="nama" value={form.nama} onChange={handleChange} required className={input} />
          </div>
          <div>
            <label className={label}>HARGA</label>
            <input type="number" name="harga" value={form.harga} onChange={handleChange} required className={input} />
          </div>
          <div>
            <label className={label}>STOK MINIMUM</label>
            <input type="number" min="0" name="stok_minimum" placeholder="0"
              value={form.stok_minimum} onChange={handleChange} className={input} />
          </div>
          <div>
            <label className={label}>SATUAN</label>
            <select name="satuan_id" value={form.satuan_id} onChange={handleChange} required className={input}>
              <option value="">Pilih satuan</option>
              {satuanList.map((s) => (
                <option key={s.id} value={s.id}>{s.nama_satuan}</option>
              ))}
            </select>
          </div>
        </div>

        <div>
          <label className={label}>KATEGORI</label>
          <select name="kategori_id" value={form.kategori_id} onChange={handleChange} className={input}>
            <option value="">Tanpa kategori</option>
            {kategoriList.map((k) => (
              <option key={k.id} value={k.id}>{k.nama_kategori}</option>
            ))}
          </select>
        </div>

        <div className="flex flex-col sm:flex-row gap-3 mt-4">
          <button type="submit" className="bg-accent hover:bg-accent-soft transition-colors rounded-lg px-5 py-2.5 font-medium text-sm">
            {editId ? 'Update Barang' : 'Tambah Barang'}
          </button>
          {editId && (
            <button type="button" onClick={handleCancelEdit} className="border border-border rounded-lg px-5 py-2.5 text-sm text-muted hover:text-ink transition-colors">
              Batal
            </button>
          )}
        </div>
      </form>

      {/* Mobile: kartu */}
      <div className="md:hidden space-y-3">
        {barangList.length === 0 ? (
          <div className="bg-surface border border-border rounded-2xl p-8 text-center text-muted text-sm">Belum ada data</div>
        ) : (
          barangList.map((b) => {
            const menipis = b.stok_minimum > 0 && b.stok <= b.stok_minimum
            return (
              <div key={b.id} className="bg-surface border border-border rounded-2xl p-4">
                <div className="flex items-start justify-between gap-3">
                  <div className="min-w-0">
                    <p className="font-medium break-words">{b.nama}</p>
                    <p className="text-xs font-mono text-accent mt-0.5">
                      Rp {b.harga.toLocaleString('id-ID')}
                    </p>
                  </div>
                  <div className="text-right shrink-0">
                    <p className={`font-mono text-lg ${menipis ? 'text-accent' : 'text-success'}`}>{b.stok}</p>
                    <p className="text-[11px] text-muted">{b.nama_satuan || 'unit'}</p>
                  </div>
                </div>

                <div className="flex flex-wrap gap-x-4 gap-y-1 mt-3 text-xs text-muted">
                  <span>Min: <span className="font-mono">{b.stok_minimum || '—'}</span></span>
                  <span>Kategori: {b.nama_kategori || '—'}</span>
                  <span className="font-mono">#{b.id}</span>
                </div>

                <div className="flex gap-2 mt-3">
                  <button onClick={() => handleEdit(b)} className={btnEdit}>Edit</button>
                  <button onClick={() => handleDelete(b.id)} className={btnHapus}>Hapus</button>
                </div>
              </div>
            )
          })
        )}
      </div>

      {/* Tablet & desktop */}
      <div className="hidden md:block bg-surface border border-border rounded-2xl overflow-x-auto">
        <table className="w-full text-sm min-w-[860px]">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">ID</th>
              <th className="px-5 py-3">Nama</th>
              <th className="px-5 py-3 text-right">Harga</th>
              <th className="px-5 py-3 text-right">Stok</th>
              <th className="px-5 py-3 text-right">Min</th>
              <th className="px-5 py-3">Satuan</th>
              <th className="px-5 py-3">Kategori</th>
              <th className="px-5 py-3 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {barangList.length === 0 ? (
              <tr><td colSpan="8" className="text-center py-8 text-muted">Belum ada data</td></tr>
            ) : (
              barangList.map((b) => {
                const menipis = b.stok_minimum > 0 && b.stok <= b.stok_minimum
                return (
                  <tr key={b.id} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                    <td className="px-5 py-3 font-mono text-muted">{b.id}</td>
                    <td className="px-5 py-3 font-medium">{b.nama}</td>
                    <td className="px-5 py-3 text-right font-mono text-accent whitespace-nowrap">Rp {b.harga.toLocaleString('id-ID')}</td>
                    <td className={`px-5 py-3 text-right font-mono ${menipis ? 'text-accent' : 'text-success'}`}>{b.stok}</td>
                    <td className="px-5 py-3 text-right font-mono text-muted">{b.stok_minimum || '-'}</td>
                    <td className="px-5 py-3 text-muted">{b.nama_satuan || '-'}</td>
                    <td className="px-5 py-3 text-muted">{b.nama_kategori || '-'}</td>
                    <td className="px-5 py-3 text-right whitespace-nowrap">
                      <button onClick={() => handleEdit(b)} className="text-xs text-muted hover:text-ink mr-3">Edit</button>
                      <button onClick={() => handleDelete(b.id)} className="text-xs text-accent hover:text-accent-soft">Hapus</button>
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

export default Barang
