// ---------------------------------------------------------------------
// Konfigurasi axios global + penanganan token.
//
// Cukup diimpor SEKALI (sudah dilakukan di App.jsx). Semua halaman yang
// memakai `import axios from 'axios'` otomatis ikut aturan di sini,
// jadi tidak ada halaman lama yang perlu diubah.
// ---------------------------------------------------------------------
import axios from 'axios'

// Access token disimpan di variabel memori, BUKAN localStorage.
// Konsekuensinya: token hilang saat halaman di-refresh, lalu diambil
// ulang lewat silent refresh di App.jsx.
let accessToken = null

export const setAccessToken = (t) => { accessToken = t }
export const getAccessToken = () => accessToken

// Wajib true supaya cookie refresh token ikut terkirim.
axios.defaults.withCredentials = true

const jalurAuth = '/api/v1/auth'

axios.interceptors.request.use((config) => {
  if (accessToken && !String(config.url || '').startsWith(jalurAuth)) {
    config.headers.Authorization = `Bearer ${accessToken}`
  }
  return config
})

// Satu proses refresh dipakai bersama, supaya 5 request yang barengan
// kena 401 tidak memicu 5 kali rotasi token.
let prosesRefresh = null

export const refreshSession = () => {
  if (!prosesRefresh) {
    prosesRefresh = axios
      .post(`${jalurAuth}/refresh`)
      .then((res) => {
        setAccessToken(res.data.access_token)
        return res.data
      })
      .finally(() => { prosesRefresh = null })
  }
  return prosesRefresh
}

axios.interceptors.response.use(
  (res) => res,
  async (error) => {
    const asli = error.config || {}
    const status = error.response?.status

    const bolehCoba =
      status === 401 &&
      !asli._sudahDicoba &&
      !String(asli.url || '').startsWith(jalurAuth)

    if (!bolehCoba) return Promise.reject(error)

    asli._sudahDicoba = true
    try {
      const data = await refreshSession()
      asli.headers = { ...asli.headers, Authorization: `Bearer ${data.access_token}` }
      return axios(asli)
    } catch (e) {
      setAccessToken(null)
      // Beri tahu App.jsx supaya menendang user ke halaman login.
      window.dispatchEvent(new Event('auth:logout'))
      return Promise.reject(e)
    }
  }
)

export const loginRequest = async (username, password) => {
  const res = await axios.post(`${jalurAuth}/login`, { username, password })
  setAccessToken(res.data.access_token)
  return res.data.user
}

export const logoutRequest = async () => {
  try {
    await axios.post(`${jalurAuth}/logout`)
  } finally {
    setAccessToken(null)
  }
}

export default axios
