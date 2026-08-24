import { useState } from 'react'
import { useNavigate, Link } from 'react-router-dom'
import axios from 'axios'

function Login({ onLoginSuccess }) {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')
  const [error, setError] = useState('')
  const navigate = useNavigate()

  const handleSubmit = async (e) => {
    e.preventDefault()
    setError('')
    try {
      const res = await axios.post('/api/login', { username, password })
      localStorage.setItem('user', JSON.stringify(res.data.user))
      onLoginSuccess(res.data.user)
      navigate('/dashboard')
    } catch (err) {
      setError(err.response?.data?.error || 'Login gagal')
    }
  }

  return (
    <div className="min-h-screen bg-canvas text-ink flex items-center justify-center px-6">
      <div className="w-full max-w-sm">
        <Link to="/" className="font-display text-sm text-muted hover:text-ink mb-8 inline-block">
          ← Kembali
        </Link>

        <div className="bg-surface border border-border rounded-2xl p-8">
          <h1 className="font-display text-2xl font-semibold mb-1">Masuk</h1>
          <p className="text-muted text-sm mb-6">Gunakan akun yang sudah terdaftar.</p>

          {error && (
            <div className="bg-accent/10 border border-accent text-accent text-sm rounded-lg px-3 py-2 mb-4">
              {error}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-4">
            <div>
              <label className="block text-xs font-mono text-muted mb-1.5">USERNAME</label>
              <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
                className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
              />
            </div>
            <div>
              <label className="block text-xs font-mono text-muted mb-1.5">PASSWORD</label>
              <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                required
                className="w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors"
              />
            </div>
            <button
              type="submit"
              className="w-full bg-accent hover:bg-accent-soft transition-colors rounded-lg py-2.5 font-medium mt-2"
            >
              Masuk
            </button>
          </form>
        </div>
      </div>
    </div>
  )
}

export default Login