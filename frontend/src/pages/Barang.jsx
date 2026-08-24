import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'

function Barang({ user, onLogout }) {
  const [barangList, setBarangList] = useState([])
  const [form, setForm] = useState({ nama: '', harga: '', stok: '', kategori_id: '', satuan_id: '' })
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
      stok: parseInt(form.stok),
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
      setForm({ nama: '', harga: '', stok: '', kategori_id: '', satuan_id: '' })
      fetchBarang()
    } catch (err) {
      alert('Gagal simpan data: ' + (err.response?.data?.error || err.message))
    }
  }

  const handleEdit = (b) => {
    setForm({
      nama: b.nama,
      harga: b.harga,
      stok: b.stok,
      kategori_id: b.kategori_id || '',
      satuan_id: b.satuan_id || '',
    })
    setEditId(b.id)
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
    setForm({ nama: '', harga: '', stok: '', kategori_id: '', satuan_id: '' })
  }

  return (
    <Layout title="Data Barang" user={user} onLogout={onLogout}>
      <form onSubmit={handleSubmit} className="bg-surface border border-border rounded-2xl p-6 mb-8">
        <h3 className="font-display font-semibold mb-4">
          {editId ? 'Edit Barang' : 'Tambah Barang Baru'}
        </h3>
        <div className="grid sm:grid-cols-4 gap-4 mb-4">
          <div>
            <label className="block text-xs font-mono text-muted mb-1.5">NAMA</label>
            <input
              type="text" name="nama" value={form.nama} onChange={handleChange} required
              className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
            />
          </div>
          <div>
            <label className="block text-xs font-mono text-muted mb-1.5">HARGA</label>
            <input
              type="number" name="harga" value={form.harga} onChange={handleChange} required
              className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
            />
          </div>
          <div>
            <label className="block text-xs font-mono text-muted mb-1.5">STOK</label>
            <input
              type="number" name="stok" value={form.stok} onChange={handleChange} required
              className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
            />
          </div>
          <div>
            <label className="block text-xs font-mono text-muted mb-1.5">SATUAN</label>
            <select
              name="satuan_id"
              value={form.satuan_id}
              onChange={handleChange}
              required
              className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
            >
              <option value="">Pilih satuan</option>
              {satuanList.map((s) => (
                <option key={s.id} value={s.id}>{s.nama_satuan}</option>
              ))}
            </select>
          </div>
        </div>
        <div>
          <label className="block text-xs font-mono text-muted mb-1.5">KATEGORI</label>
          <select
            name="kategori_id"
            value={form.kategori_id}
            onChange={handleChange}
            className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
          >
            <option value="">Tanpa kategori</option>
            {kategoriList.map((k) => (
              <option key={k.id} value={k.id}>{k.nama_kategori}</option>
            ))}
          </select>
        </div>
        <div className="flex gap-3 mt-4">
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

      <div className="bg-surface border border-border rounded-2xl overflow-hidden">
        <table className="w-full text-sm">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">ID</th>
              <th className="px-5 py-3">Nama</th>
              <th className="px-5 py-3">Harga</th>
              <th className="px-5 py-3">Stok</th>
              <th className="px-5 py-3">Satuan</th>
              <th className="px-5 py-3">Kategori</th>
              <th className="px-5 py-3">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {barangList.length === 0 ? (
              <tr><td colSpan="7" className="text-center py-8 text-muted">Belum ada data</td></tr>
            ) : (
              barangList.map((b) => (
                <tr key={b.id} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                  <td className="px-5 py-3 font-mono text-muted">{b.id}</td>
                  <td className="px-5 py-3 font-medium">{b.nama}</td>
                  <td className="px-5 py-3 font-mono text-accent">Rp {b.harga.toLocaleString('id-ID')}</td>
                  <td className="px-5 py-3 font-mono text-success">{b.stok}</td>
                  <td className="px-5 py-3 text-muted">{b.nama_satuan || '-'}</td>
                  <td className="px-5 py-3 text-muted">{b.nama_kategori || '-'}</td>
                  <td className="px-5 py-3">
                    <button onClick={() => handleEdit(b)} className="text-xs text-muted hover:text-ink mr-3">Edit</button>
                    <button onClick={() => handleDelete(b.id)} className="text-xs text-accent hover:text-accent-soft">Hapus</button>
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

export default Barang