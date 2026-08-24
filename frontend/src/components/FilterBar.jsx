// Komponen kecil yang dipakai bersama oleh semua halaman laporan:
// baris filter, tombol cetak, dan gaya input yang seragam.

export const inputClass =
  'w-full bg-surface-soft border border-border rounded-lg px-3 py-2.5 text-sm outline-none focus:border-accent transition-colors'
export const labelClass = 'block text-xs font-mono text-muted mb-1.5'

export function Field({ label, children }) {
  return (
    <div>
      <label className={labelClass}>{label}</label>
      {children}
    </div>
  )
}

export function FilterBar({ children, onCetak }) {
  return (
    <div className="bg-surface border border-border rounded-2xl p-6 mb-6 print:hidden">
      <div className="flex items-end justify-between gap-4 flex-wrap">
        <div className="grid sm:grid-cols-2 lg:grid-cols-4 gap-4 flex-1 min-w-[260px]">{children}</div>
        {onCetak && (
          <button
            type="button"
            onClick={onCetak}
            className="border border-border rounded-lg px-4 py-2.5 text-sm text-muted hover:text-ink hover:bg-surface-soft transition-colors"
          >
            Cetak
          </button>
        )}
      </div>
    </div>
  )
}

export function TabelKosong({ colSpan, pesan = 'Tidak ada data pada filter ini' }) {
  return (
    <tr>
      <td colSpan={colSpan} className="text-center py-8 text-muted">{pesan}</td>
    </tr>
  )
}

export const rupiah = (n) => 'Rp ' + Number(n || 0).toLocaleString('id-ID')
export const awalBulan = () => {
  const d = new Date()
  return new Date(d.getFullYear(), d.getMonth(), 1).toISOString().slice(0, 10)
}
export const hariIni = () => new Date().toISOString().slice(0, 10)
