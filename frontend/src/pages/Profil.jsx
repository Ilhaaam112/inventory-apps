import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'

function Profil({ user, onLogout }) {
  const [profile, setProfile] = useState(null)
  const [namaLengkap, setNamaLengkap] = useState('')
  const [passwordForm, setPasswordForm] = useState({ old_password: '', new_password: '', confirm_password: '' })
  const [loading, setLoading] = useState(false)

  const fetchProfile = async () => {
    try {
      const res = await axios.get(`/api/profile?id=${user.id}`)
      setProfile(res.data)
      setNamaLengkap(res.data.nama_lengkap || '')
    } catch (err) {
      console.error('Gagal ambil profil:', err)
    }
  }

  useEffect(() => { fetchProfile() }, [])

  const handleUpdateProfile = async (e) => {
    e.preventDefault()
    setLoading(true)
    try {
      const res = await axios.put('/api/profile', { id: user.id, nama_lengkap: namaLengkap })
      localStorage.setItem('user', JSON.stringify({ ...user, nama_lengkap: res.data.nama_lengkap }))
      alert('Profil berhasil diperbarui')
      window.location.reload()
    } catch (err) {
      alert('Gagal update profil: ' + (err.response?.data?.error || err.message))
    } finally {
      setLoading(false)
    }
  }

  const handleChangePassword = async (e) => {
    e.preventDefault()
    if (passwordForm.new_password !== passwordForm.confirm_password) {
      alert('Konfirmasi password baru tidak cocok')
      return
    }
    setLoading(true)
    try {
      await axios.put('/api/change-password', {
        id: user.id,
        old_password: passwordForm.old_password,
        new_password: passwordForm.new_password,
      })
      alert('Password berhasil diubah')
      setPasswordForm({ old_password: '', new_password: '', confirm_password: '' })
    } catch (err) {
      alert('Gagal ubah password: ' + (err.response?.data?.error || err.message))
    } finally {
      setLoading(false)
    }
  }

  return (
    <Layout title="Profil" user={user} onLogout={onLogout}>
      <div className="grid md:grid-cols-2 gap-6">
        <form onSubmit={handleUpdateProfile} className="bg-surface border border-border rounded-2xl p-6">
          <h3 className="font-display font-semibold mb-4">Info Akun</h3>
          <div className="space-y-4">
            <div>
              <label className="block text-xs font-mono text-muted mb-1.5">USERNAME</label>
              <input type="text" value={profile?.username || ''} disabled
                className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm text-muted cursor-not-allowed" />
            </div>
            <div>
              <label className="block text-xs font-mono text-muted mb-1.5">NAMA LENGKAP</label>
              <input type="text" value={namaLengkap} onChange={(e) => setNamaLengkap(e.target.value)} required
                className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors" />
            </div>
          </div>
          <button type="submit" disabled={loading}
            className="mt-4 bg-accent hover:bg-accent-soft transition-colors rounded-lg px-5 py-2.5 font-medium text-sm disabled:opacity-50">
            Simpan Perubahan
          </button>
        </form>

        <form onSubmit={handleChangePassword} className="bg-surface border border-border rounded-2xl p-6">
          <h3 className="font-display font-semibold mb-4">Ubah Password</h3>
          <div className="space-y-4">
            <div>
              <label className="block text-xs font-mono text-muted mb-1.5">PASSWORD LAMA</label>
              <input type="password" value={passwordForm.old_password}
                onChange={(e) => setPasswordForm({ ...passwordForm, old_password: e.target.value })} required
                className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors" />
            </div>
            <div>
              <label className="block text-xs font-mono text-muted mb-1.5">PASSWORD BARU</label>
              <input type="password" value={passwordForm.new_password}
                onChange={(e) => setPasswordForm({ ...passwordForm, new_password: e.target.value })} required
                className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors" />
            </div>
            <div>
              <label className="block text-xs font-mono text-muted mb-1.5">KONFIRMASI PASSWORD BARU</label>
              <input type="password" value={passwordForm.confirm_password}
                onChange={(e) => setPasswordForm({ ...passwordForm, confirm_password: e.target.value })} required
                className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors" />
            </div>
          </div>
          <button type="submit" disabled={loading}
            className="mt-4 bg-accent hover:bg-accent-soft transition-colors rounded-lg px-5 py-2.5 font-medium text-sm disabled:opacity-50">
            Ubah Password
          </button>
        </form>
      </div>
    </Layout>
  )
}

export default Profil