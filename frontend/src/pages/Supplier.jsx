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

  useEffect(() => { fetchSupplier() }, [])

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

  const input = 'w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors'
  const label = 'block text-xs font-mono text-muted mb-1.5'
  const btnEdit = 'flex-1 sm:flex-none text-xs border border-border rounded-lg px-3 py-2 text-muted hover:text-ink transition-colors'
  const btnHapus = 'flex-1 sm:flex-none text-xs border border-accent/40 rounded-lg px-3 py-2 text-accent hover:bg-accent/10 transition-colors'

  return (
    <Layout title="Data Supplier" user={user} onLogout={onLogout}>
      <form onSubmit={handleSubmit} className="bg-surface border border-border rounded-2xl p-4 sm:p-6 mb-6">
        <h3 className="font-display font-semibold mb-4">
          {editId ? 'Edit Supplier' : 'Tambah Supplier Baru'}
        </h3>
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4 mb-4">
          <div>
            <label className={label}>NAMA SUPPLIER</label>
            <input type="text" name="nama_supplier" value={form.nama_supplier} onChange={handleChange} required className={input} />
          </div>
          <div>
            <label className={label}>KONTAK</label>
            <input type="text" name="kontak" value={form.kontak} onChange={handleChange} placeholder="No. HP / Email" className={input} />
          </div>
          <div className="sm:col-span-2 lg:col-span-1">
            <label className={label}>ALAMAT</label>
            <input type="text" name="alamat" value={form.alamat} onChange={handleChange} className={input} />
          </div>
        </div>
        <div className="flex flex-col sm:flex-row gap-3">
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

      {/* Mobile: kartu */}
      <div className="md:hidden space-y-3">
        {supplierList.length === 0 ? (
          <div className="bg-surface border border-border rounded-2xl p-8 text-center text-muted text-sm">Belum ada data</div>
        ) : (
          supplierList.map((s) => (
            <div key={s.id} className="bg-surface border border-border rounded-2xl p-4">
              <div className="flex items-start justify-between gap-3">
                <p className="font-medium break-words min-w-0">{s.nama_supplier}</p>
                <span className="text-[11px] font-mono text-muted shrink-0">#{s.id}</span>
              </div>
              <dl className="mt-2 space-y-1 text-xs">
                <div className="flex gap-2">
                  <dt className="text-muted w-16 shrink-0">Kontak</dt>
                  <dd className="break-words min-w-0">{s.kontak || '—'}</dd>
                </div>
                <div className="flex gap-2">
                  <dt className="text-muted w-16 shrink-0">Alamat</dt>
                  <dd className="break-words min-w-0">{s.alamat || '—'}</dd>
                </div>
              </dl>
              <div className="flex gap-2 mt-3">
                <button onClick={() => handleEdit(s)} className={btnEdit}>Edit</button>
                <button onClick={() => handleDelete(s.id)} className={btnHapus}>Hapus</button>
              </div>
            </div>
          ))
        )}
      </div>

      {/* Tablet & desktop */}
      <div className="hidden md:block bg-surface border border-border rounded-2xl overflow-x-auto">
        <table className="w-full text-sm min-w-[720px]">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">ID</th>
              <th className="px-5 py-3">Nama Supplier</th>
              <th className="px-5 py-3">Kontak</th>
              <th className="px-5 py-3">Alamat</th>
              <th className="px-5 py-3 text-right">Aksi</th>
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
                  <td className="px-5 py-3 text-right whitespace-nowrap">
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
