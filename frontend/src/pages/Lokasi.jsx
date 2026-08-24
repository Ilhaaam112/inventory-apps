import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'

function Lokasi({ user, onLogout }) {
  const [lokasiList, setLokasiList] = useState([])
  const [form, setForm] = useState({ nama_lokasi: '', keterangan: '' })
  const [editId, setEditId] = useState(null)

  const fetchLokasi = async () => {
    try {
      const res = await axios.get('/api/lokasi')
      setLokasiList(res.data || [])
    } catch (err) {
      console.error('Gagal ambil data:', err)
    }
  }

  useEffect(() => { fetchLokasi() }, [])

  const handleChange = (e) => setForm({ ...form, [e.target.name]: e.target.value })

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      if (editId) {
        await axios.put(`/api/lokasi/${editId}`, form)
        setEditId(null)
      } else {
        await axios.post('/api/lokasi', form)
      }
      setForm({ nama_lokasi: '', keterangan: '' })
      fetchLokasi()
    } catch (err) {
      alert('Gagal simpan data: ' + (err.response?.data?.error || err.message))
    }
  }

  const handleEdit = (l) => {
    setForm({ nama_lokasi: l.nama_lokasi, keterangan: l.keterangan || '' })
    setEditId(l.id)
  }

  const handleDelete = async (id) => {
    if (!confirm('Yakin mau hapus lokasi ini?')) return
    try {
      await axios.delete(`/api/lokasi/${id}`)
      fetchLokasi()
    } catch {
      alert('Gagal hapus data')
    }
  }

  const handleCancelEdit = () => {
    setEditId(null)
    setForm({ nama_lokasi: '', keterangan: '' })
  }

  const input = 'w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors'
  const label = 'block text-xs font-mono text-muted mb-1.5'
  const btnEdit = 'flex-1 sm:flex-none text-xs border border-border rounded-lg px-3 py-2 text-muted hover:text-ink transition-colors'
  const btnHapus = 'flex-1 sm:flex-none text-xs border border-accent/40 rounded-lg px-3 py-2 text-accent hover:bg-accent/10 transition-colors'

  return (
    <Layout title="Data Lokasi / Gudang" user={user} onLogout={onLogout}>
      <form onSubmit={handleSubmit} className="bg-surface border border-border rounded-2xl p-4 sm:p-6 mb-6">
        <h3 className="font-display font-semibold mb-4">
          {editId ? 'Edit Lokasi' : 'Tambah Lokasi Baru'}
        </h3>
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 mb-4">
          <div>
            <label className={label}>NAMA LOKASI / GUDANG</label>
            <input type="text" name="nama_lokasi" value={form.nama_lokasi} onChange={handleChange} required
              placeholder="Gudang Utama, Gudang Cabang, dll" className={input} />
          </div>
          <div>
            <label className={label}>KETERANGAN</label>
            <input type="text" name="keterangan" value={form.keterangan} onChange={handleChange}
              placeholder="Opsional" className={input} />
          </div>
        </div>
        <div className="flex flex-col sm:flex-row gap-3">
          <button type="submit" className="bg-accent hover:bg-accent-soft transition-colors rounded-lg px-5 py-2.5 font-medium text-sm">
            {editId ? 'Update Lokasi' : 'Tambah Lokasi'}
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
        {lokasiList.length === 0 ? (
          <div className="bg-surface border border-border rounded-2xl p-8 text-center text-muted text-sm">Belum ada data</div>
        ) : (
          lokasiList.map((l) => (
            <div key={l.id} className="bg-surface border border-border rounded-2xl p-4">
              <div className="flex items-start justify-between gap-3">
                <p className="font-medium break-words min-w-0">{l.nama_lokasi}</p>
                <span className="text-[11px] font-mono text-muted shrink-0">#{l.id}</span>
              </div>
              <p className="text-xs text-muted mt-1 break-words">{l.keterangan || '—'}</p>
              <div className="flex gap-2 mt-3">
                <button onClick={() => handleEdit(l)} className={btnEdit}>Edit</button>
                <button onClick={() => handleDelete(l.id)} className={btnHapus}>Hapus</button>
              </div>
            </div>
          ))
        )}
      </div>

      {/* Tablet & desktop */}
      <div className="hidden md:block bg-surface border border-border rounded-2xl overflow-x-auto">
        <table className="w-full text-sm min-w-[600px]">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">ID</th>
              <th className="px-5 py-3">Nama Lokasi</th>
              <th className="px-5 py-3">Keterangan</th>
              <th className="px-5 py-3 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {lokasiList.length === 0 ? (
              <tr><td colSpan="4" className="text-center py-8 text-muted">Belum ada data</td></tr>
            ) : (
              lokasiList.map((l) => (
                <tr key={l.id} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                  <td className="px-5 py-3 font-mono text-muted">{l.id}</td>
                  <td className="px-5 py-3 font-medium">{l.nama_lokasi}</td>
                  <td className="px-5 py-3 text-muted">{l.keterangan || '-'}</td>
                  <td className="px-5 py-3 text-right whitespace-nowrap">
                    <button onClick={() => handleEdit(l)} className="text-xs text-muted hover:text-ink mr-3">Edit</button>
                    <button onClick={() => handleDelete(l.id)} className="text-xs text-accent hover:text-accent-soft">Hapus</button>
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

export default Lokasi
