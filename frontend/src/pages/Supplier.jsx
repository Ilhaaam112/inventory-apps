import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'

function Supplier({ user, onLogout }) {
  const [supplierList, setSupplierList] = useState([])
  const [form, setForm] = useState({ nama_supplier: '', kontak: '', alamat: '' })
  const [editId, setEditId] = useState(null)

  const fetchSupplier = async () => {
    try {
      const res = await axios.get('/api/supplier')
      setSupplierList(res.data || [])
    } catch (err) {
      console.error('Gagal ambil data:', err)
    }
  }

  useEffect(() => {
    fetchSupplier()
  }, [])

  const handleChange = (e) => setForm({ ...form, [e.target.name]: e.target.value })

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      if (editId) {
        await axios.put(`/api/supplier/${editId}`, form)
        setEditId(null)
      } else {
        await axios.post('/api/supplier', form)
      }
      setForm({ nama_supplier: '', kontak: '', alamat: '' })
      fetchSupplier()
    } catch (err) {
      alert('Gagal simpan data: ' + (err.response?.data?.error || err.message))
    }
  }

  const handleEdit = (s) => {
    setForm({ nama_supplier: s.nama_supplier, kontak: s.kontak || '', alamat: s.alamat || '' })
    setEditId(s.id)
  }

  const handleDelete = async (id) => {
    if (!confirm('Yakin mau hapus supplier ini?')) return
    try {
      await axios.delete(`/api/supplier/${id}`)
      fetchSupplier()
    } catch {
      alert('Gagal hapus data')
    }
  }

  const handleCancelEdit = () => {
    setEditId(null)
    setForm({ nama_supplier: '', kontak: '', alamat: '' })
  }

  return (
    <Layout title="Data Supplier" user={user} onLogout={onLogout}>
      <form onSubmit={handleSubmit} className="bg-surface border border-border rounded-2xl p-6 mb-8">
        <h3 className="font-display font-semibold mb-4">
          {editId ? 'Edit Supplier' : 'Tambah Supplier Baru'}
        </h3>
        <div className="grid sm:grid-cols-3 gap-4 mb-4">
          <div>
            <label className="block text-xs font-mono text-muted mb-1.5">NAMA SUPPLIER</label>
            <input
              type="text" name="nama_supplier" value={form.nama_supplier} onChange={handleChange} required
              className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
            />
          </div>
          <div>
            <label className="block text-xs font-mono text-muted mb-1.5">KONTAK</label>
            <input
              type="text" name="kontak" value={form.kontak} onChange={handleChange}
              placeholder="No. HP / Email"
              className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
            />
          </div>
          <div>
            <label className="block text-xs font-mono text-muted mb-1.5">ALAMAT</label>
            <input
              type="text" name="alamat" value={form.alamat} onChange={handleChange}
              className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
            />
          </div>
        </div>
        <div className="flex gap-3">
          <button type="submit" className="bg-accent hover:bg-accent-soft transition-colors rounded-lg px-5 py-2.5 font-medium text-sm">
            {editId ? 'Update Supplier' : 'Tambah Supplier'}
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
              <th className="px-5 py-3">Nama Supplier</th>
              <th className="px-5 py-3">Kontak</th>
              <th className="px-5 py-3">Alamat</th>
              <th className="px-5 py-3">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {supplierList.length === 0 ? (
              <tr><td colSpan="5" className="text-center py-8 text-muted">Belum ada data</td></tr>
            ) : (
              supplierList.map((s) => (
                <tr key={s.id} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                  <td className="px-5 py-3 font-mono text-muted">{s.id}</td>
                  <td className="px-5 py-3 font-medium">{s.nama_supplier}</td>
                  <td className="px-5 py-3 text-muted">{s.kontak || '-'}</td>
                  <td className="px-5 py-3 text-muted">{s.alamat || '-'}</td>
                  <td className="px-5 py-3">
                    <button onClick={() => handleEdit(s)} className="text-xs text-muted hover:text-ink mr-3">Edit</button>
                    <button onClick={() => handleDelete(s.id)} className="text-xs text-accent hover:text-accent-soft">Hapus</button>
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

export default Supplier
