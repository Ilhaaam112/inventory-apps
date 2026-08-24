import { Link } from 'react-router-dom'

function PublicHome() {
  return (
    <div className="min-h-screen bg-canvas text-ink">
      {/* Navbar */}
      <nav className="flex items-center justify-between px-6 md:px-12 py-6 border-b border-border">
        <span className="font-display text-lg font-semibold tracking-tight">
          belajar<span className="text-accent">Go</span>.
        </span>
        <Link
          to="/login"
          className="rounded-full bg-accent px-5 py-2 text-sm font-medium hover:bg-accent-soft transition-colors"
        >
          Masuk
        </Link>
      </nav>

      {/* Hero */}
      <section className="px-6 md:px-12 py-20 grid md:grid-cols-2 gap-12 items-center max-w-6xl mx-auto">
        <div>
          <p className="font-mono text-xs text-accent tracking-widest mb-4">
            SISTEM STOK · V1.0
          </p>
          <h1 className="font-display text-4xl md:text-5xl font-semibold leading-tight mb-6">
            Setiap barang,<br />tercatat rapi.
          </h1>
          <p className="text-muted text-base md:text-lg mb-8 max-w-md">
            Kelola data barang, pantau stok, dan catat harga dalam satu tempat.
            Sederhana, cepat, dan selalu update.
          </p>
          <Link
            to="/login"
            className="inline-block rounded-full bg-accent px-7 py-3 font-medium hover:bg-accent-soft transition-colors"
          >
            Masuk ke Akun →
          </Link>
        </div>

        {/* Signature: kartu label gudang */}
        <div className="flex justify-center">
          <div className="relative w-72 rounded-2xl bg-surface border border-border shadow-2xl p-6 -rotate-3">
            <div className="absolute -top-3 left-1/2 -translate-x-1/2 w-6 h-6 rounded-full bg-canvas border border-border" />
            <p className="font-mono text-xs text-muted mb-1">SKU-00184</p>
            <h3 className="font-display text-xl font-semibold mb-4">Kabel USB-C</h3>
            <div className="flex justify-between items-end">
              <div>
                <p className="text-xs text-muted mb-1">Harga</p>
                <p className="font-mono text-lg text-accent">Rp 20.000</p>
              </div>
              <div className="text-right">
                <p className="text-xs text-muted mb-1">Stok</p>
                <p className="font-mono text-lg text-success">40 unit</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Fitur */}
      <section className="px-6 md:px-12 py-16 border-t border-border">
        <div className="max-w-6xl mx-auto grid md:grid-cols-3 gap-8">
          {[
            { no: '01', title: 'Catat', desc: 'Tambahkan barang baru lengkap dengan harga dan jumlah stok.' },
            { no: '02', title: 'Pantau', desc: 'Lihat seluruh daftar barang dan stok yang tersedia kapan saja.' },
            { no: '03', title: 'Kelola', desc: 'Ubah atau hapus data barang secara langsung dari dashboard.' },
          ].map((item) => (
            <div key={item.no}>
              <p className="font-mono text-accent text-sm mb-3">{item.no}</p>
              <h3 className="font-display text-lg font-semibold mb-2">{item.title}</h3>
              <p className="text-muted text-sm">{item.desc}</p>
            </div>
          ))}
        </div>
      </section>

      <footer className="px-6 md:px-12 py-8 border-t border-border text-center text-muted text-xs">
        Dibuat untuk belajar Go &amp; React — {new Date().getFullYear()}
      </footer>
    </div>
  )
}

export default PublicHome