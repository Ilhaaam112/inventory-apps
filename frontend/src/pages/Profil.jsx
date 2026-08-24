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
      // Tanpa ?id= : server selalu memakai id dari access token.
      const res = await axios.get('/api/profile')
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
      await axios.put('/api/profile', { nama_lengkap: namaLengkap })
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
    if (passwordForm.new_password.length < 8) {
      alert('Password baru minimal 8 karakter')
      return
    }
    setLoading(true)
    try {
      await axios.put('/api/change-password', {
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

  const input = 'w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors'
  const inputMati = 'w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm text-muted cursor-not-allowed'
  const label = 'block text-xs font-mono text-muted mb-1.5'
  const tombol = 'w-full sm:w-auto mt-4 bg-accent hover:bg-accent-soft transition-colors rounded-lg px-5 py-2.5 font-medium text-sm disabled:opacity-50'

  return (
    <Layout title="Profil" user={user} onLogout={onLogout}>
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4 sm:gap-6">
        <form onSubmit={handleUpdateProfile} className="bg-surface border border-border rounded-2xl p-4 sm:p-6">
          <h3 className="font-display font-semibold mb-4">Info Akun</h3>
          <div className="space-y-4">
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div>
                <label className={label}>USERNAME</label>
                <input type="text" value={profile?.username || ''} disabled className={inputMati} />
              </div>
              <div>
                <label className={label}>ROLE</label>
                <input type="text" value={profile?.role || '-'} disabled className={inputMati} />
              </div>
            </div>
            <div>
              <label className={label}>NAMA LENGKAP</label>
              <input type="text" value={namaLengkap} maxLength={100}
                onChange={(e) => setNamaLengkap(e.target.value)} required className={input} />
            </div>
          </div>
          <button type="submit" disabled={loading} className={tombol}>Simpan Perubahan</button>
        </form>

        <form onSubmit={handleChangePassword} className="bg-surface border border-border rounded-2xl p-4 sm:p-6">
          <h3 className="font-display font-semibold mb-1">Ubah Password</h3>
          <p className="text-xs text-muted mb-4">Minimal 8 karakter.</p>
          <div className="space-y-4">
            <div>
              <label className={label}>PASSWORD LAMA</label>
              <input type="password" autoComplete="current-password" value={passwordForm.old_password}
                onChange={(e) => setPasswordForm({ ...passwordForm, old_password: e.target.value })} required className={input} />
            </div>
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              <div>
                <label className={label}>PASSWORD BARU</label>
                <input type="password" autoComplete="new-password" minLength={8} value={passwordForm.new_password}
                  onChange={(e) => setPasswordForm({ ...passwordForm, new_password: e.target.value })} required className={input} />
              </div>
              <div>
                <label className={label}>KONFIRMASI</label>
                <input type="password" autoComplete="new-password" minLength={8} value={passwordForm.confirm_password}
                  onChange={(e) => setPasswordForm({ ...passwordForm, confirm_password: e.target.value })} required className={input} />
              </div>
            </div>
          </div>
          <button type="submit" disabled={loading} className={tombol}>Ubah Password</button>
        </form>
      </div>
    </Layout>
  )
}

export default Profil
