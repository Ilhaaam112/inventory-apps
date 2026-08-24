import { useState, useEffect } from 'react'
import axios from 'axios'
import Layout from '../components/Layout'
import { FilterBar, Field, TabelKosong, inputClass, awalBulan, hariIni } from '../components/FilterBar'

const labelTipe = {
  IN: 'Masuk',
  OUT: 'Keluar',
  ADJUSTMENT: 'Penyesuaian',
  TRANSFER_IN: 'Transfer masuk',
  TRANSFER_OUT: 'Transfer keluar',
}

function KartuStok({ user, onLogout }) {
  const [barangList, setBarangList] = useState([])
  const [lokasiList, setLokasiList] = useState([])
  const [filter, setFilter] = useState({ barang_id: '', lokasi_id: '', start: awalBulan(), end: hariIni() })
  const [kartu, setKartu] = useState(null)
  const [pesan, setPesan] = useState('Pilih barang dan gudang untuk menampilkan kartu stok.')

  useEffect(() => {
    axios.get('/api/barang').then((r) => setBarangList(r.data || []))
    axios.get('/api/lokasi').then((r) => setLokasiList(r.data || []))
  }, [])

  useEffect(() => {
    if (!filter.barang_id || !filter.lokasi_id) {
      setKartu(null)
      setPesan('Pilih barang dan gudang untuk menampilkan kartu stok.')
      return
    }
    const q = new URLSearchParams(filter)
    axios
      .get(`/api/laporan/kartu-stok?${q}`)
      .then((r) => setKartu(r.data))
      .catch((err) => {
        setKartu(null)
        setPesan(err.response?.data?.error || 'Gagal memuat kartu stok')
      })
  }, [filter])

  return (
    <Layout title="Kartu Stok" user={user} onLogout={onLogout}>
      <FilterBar onCetak={() => window.print()}>
        <Field label="BARANG">
          <select
            className={inputClass}
            value={filter.barang_id}
            onChange={(e) => setFilter({ ...filter, barang_id: e.target.value })}
          >
            <option value="">Pilih barang</option>
            {barangList.map((b) => (
              <option key={b.id} value={b.id}>{b.nama}</option>
            ))}
          </select>
        </Field>
        <Field label="GUDANG">
          <select
            className={inputClass}
            value={filter.lokasi_id}
            onChange={(e) => setFilter({ ...filter, lokasi_id: e.target.value })}
          >
            <option value="">Pilih gudang</option>
            {lokasiList.map((l) => (
              <option key={l.id} value={l.id}>{l.nama_lokasi}</option>
            ))}
          </select>
        </Field>
        <Field label="DARI TANGGAL">
          <input
            type="date" className={inputClass} value={filter.start}
            onChange={(e) => setFilter({ ...filter, start: e.target.value })}
          />
        </Field>
        <Field label="SAMPAI TANGGAL">
          <input
            type="date" className={inputClass} value={filter.end}
            onChange={(e) => setFilter({ ...filter, end: e.target.value })}
          />
        </Field>
      </FilterBar>

      {!kartu ? (
        <div className="bg-surface border border-border rounded-2xl p-10 text-center text-muted text-sm">
          {pesan}
        </div>
      ) : (
        <>
          <div className="bg-surface border border-border rounded-2xl p-4 sm:p-5 mb-6 grid grid-cols-2 sm:flex sm:flex-wrap gap-4 sm:gap-8">
            <div>
              <p className="text-xs font-mono text-muted mb-1">BARANG</p>
              <p className="font-display">{kartu.nama_barang} {kartu.nama_satuan ? `(${kartu.nama_satuan})` : ''}</p>
            </div>
            <div>
              <p className="text-xs font-mono text-muted mb-1">GUDANG</p>
              <p className="font-display">{kartu.nama_lokasi}</p>
            </div>
            <div>
              <p className="text-xs font-mono text-muted mb-1">SALDO AWAL</p>
              <p className="font-display">{kartu.saldo_awal}</p>
            </div>
            <div>
              <p className="text-xs font-mono text-muted mb-1">SALDO AKHIR</p>
              <p className="font-display text-accent">{kartu.saldo_akhir}</p>
            </div>
          </div>

          <div className="bg-surface border border-border rounded-2xl overflow-x-auto">
            <table className="w-full text-sm min-w-[720px]">
              <thead>
                <tr className="border-b border-border text-left text-xs font-mono text-muted">
                  <th className="px-5 py-3">Waktu</th>
                  <th className="px-5 py-3">Referensi</th>
                  <th className="px-5 py-3">Jenis</th>
                  <th className="px-5 py-3 text-right">Masuk</th>
                  <th className="px-5 py-3 text-right">Keluar</th>
                  <th className="px-5 py-3 text-right">Saldo</th>
                </tr>
              </thead>
              <tbody>
                <tr className="border-b border-border bg-surface-soft">
                  <td className="px-5 py-3 text-muted" colSpan="5">Saldo awal periode</td>
                  <td className="px-5 py-3 text-right font-mono">{kartu.saldo_awal}</td>
                </tr>
                {kartu.rows?.length === 0 ? (
                  <TabelKosong colSpan={6} pesan="Tidak ada pergerakan pada periode ini" />
                ) : (
                  kartu.rows.map((r, i) => (
                    <tr key={i} className="border-b border-border last:border-0 hover:bg-surface-soft transition-colors">
                      <td className="px-5 py-3 font-mono text-muted">{r.tanggal}</td>
                      <td className="px-5 py-3 font-mono text-accent">{r.referensi}</td>
                      <td className="px-5 py-3 text-muted">{labelTipe[r.type] || r.type}</td>
                      <td className="px-5 py-3 text-right font-mono text-success">{r.masuk || '-'}</td>
                      <td className="px-5 py-3 text-right font-mono text-accent">{r.keluar || '-'}</td>
                      <td className="px-5 py-3 text-right font-mono">{r.saldo_sesudah}</td>
                    </tr>
                  ))
                )}
              </tbody>
            </table>
          </div>
        </>
      )}
    </Layout>
  )
}

export default KartuStok
