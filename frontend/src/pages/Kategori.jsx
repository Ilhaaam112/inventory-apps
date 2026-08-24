import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'

function Kategori({ user, onLogout }) {
  const [list, setList] = useState([])
  const [nama, setNama] = useState('')
  const [editId, setEditId] = useState(null)

  const fetchData = async () => {
    const res = await axios.get('/api/kategori')
    setList(res.data || [])
  }

  useEffect(() => { fetchData() }, [])

  const handleSubmit = async (e) => {
    e.preventDefault()
    try {
      if (editId) {
        await axios.put(`/api/kategori/${editId}`, { nama_kategori: nama })
        setEditId(null)
      } else {
        await axios.post('/api/kategori', { nama_kategori: nama })
      }
      setNama('')
      fetchData()
    } catch (err) {
      alert('Gagal simpan: ' + (err.response?.data?.error || err.message))
    }
  }

  const handleEdit = (k) => {
    setNama(k.nama_kategori)
    setEditId(k.id)
  }

  const handleDelete = async (id) => {
    if (!confirm('Hapus kategori ini? Barang yang memakainya akan jadi tanpa kategori.')) return
    await axios.delete(`/api/kategori/${id}`)
    fetchData()
  }

  const btnEdit = 'flex-1 sm:flex-none text-xs border border-border rounded-lg px-3 py-2 text-muted hover:text-ink transition-colors'
  const btnHapus = 'flex-1 sm:flex-none text-xs border border-accent/40 rounded-lg px-3 py-2 text-accent hover:bg-accent/10 transition-colors'

  return (
    <Layout title="Kategori Barang" user={user} onLogout={onLogout}>
      <form onSubmit={handleSubmit} className="bg-surface border border-border rounded-2xl p-4 sm:p-6 mb-6 flex flex-col sm:flex-row gap-3 sm:items-end">
        <div className="flex-1 min-w-0">
          <label className="block text-xs font-mono text-muted mb-1.5">NAMA KATEGORI</label>
          <input
            type="text"
            value={nama}
            onChange={(e) => setNama(e.target.value)}
            required
            className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
          />
        </div>
        <div className="flex gap-3">
          <button type="submit" className="flex-1 sm:flex-none bg-accent hover:bg-accent-soft transition-colors rounded-lg px-5 py-2.5 font-medium text-sm">
            {editId ? 'Update' : 'Tambah'}
          </button>
          {editId && (
            <button
              type="button"
              onClick={() => { setEditId(null); setNama('') }}
              className="flex-1 sm:flex-none border border-border rounded-lg px-5 py-2.5 text-sm text-muted hover:text-ink transition-colors"
            >
              Batal
            </button>
          )}
        </div>
      </form>

      {/* Mobile: kartu */}
      <div className="md:hidden space-y-3">
        {list.length === 0 ? (
          <div className="bg-surface border border-border rounded-2xl p-8 text-center text-muted text-sm">
            Belum ada kategori
          </div>
        ) : (
          list.map((k) => (
            <div key={k.id} className="bg-surface border border-border rounded-2xl p-4">
              <div className="flex items-start justify-between gap-3 mb-3">
                <p className="font-medium break-words min-w-0">{k.nama_kategori}</p>
                <span className="text-[11px] font-mono text-muted shrink-0">#{k.id}</span>
              </div>
              <div className="flex gap-2">
                <button onClick={() => handleEdit(k)} className={btnEdit}>Edit</button>
                <button onClick={() => handleDelete(k.id)} className={btnHapus}>Hapus</button>
              </div>
            </div>
          ))
        )}
      </div>

      {/* Tablet & desktop: tabel */}
      <div className="hidden md:block bg-surface border border-border rounded-2xl overflow-x-auto">
        <table className="w-full text-sm min-w-[480px]">
          <thead>
            <tr className="border-b border-border text-left text-xs font-mono text-muted">
              <th className="px-5 py-3">ID</th>
              <th className="px-5 py-3">Nama Kategori</th>
              <th className="px-5 py-3 text-right">Aksi</th>
            </tr>
          </thead>
          <tbody>
            {list.length === 0 ? (
              <tr><td colSpan="3" className="text-center py-8 text-muted">Belum ada kategori</td></tr>
            ) : (
              list.map((k) => (
                <tr key={k.id} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                  <td className="px-5 py-3 font-mono text-muted">{k.id}</td>
                  <td className="px-5 py-3 font-medium">{k.nama_kategori}</td>
                  <td className="px-5 py-3 text-right whitespace-nowrap">
                    <button onClick={() => handleEdit(k)} className="text-xs text-muted hover:text-ink mr-3">Edit</button>
                    <button onClick={() => handleDelete(k.id)} className="text-xs text-accent hover:text-accent-soft">Hapus</button>
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

export default Kategori
