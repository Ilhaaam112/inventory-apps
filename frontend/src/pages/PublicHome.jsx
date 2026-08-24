import { useState } from 'react'
import { Link } from 'react-router-dom'
import {
  ArrowRight,
  Package,
  Tags,
  Ruler,
  Truck,
  Warehouse,
  ArrowDownToLine,
  ArrowUpFromLine,
  RefreshCw,
  ArrowLeftRight,
  History,
  ShieldCheck,
  BarChart3,
  ScrollText,
  Activity,
  Target,
  Compass,
  Mail,
  MapPin,
  Link2,
  ChevronDown,
} from 'lucide-react'

const masterData = [
  { name: 'Data Barang', icon: Package, desc: 'Nama, harga, stok total, kategori & satuan.' },
  { name: 'Kategori', icon: Tags, desc: 'Pengelompokan barang.' },
  { name: 'Satuan', icon: Ruler, desc: 'Pcs, box, kg — konsisten di semua transaksi.' },
  { name: 'Supplier', icon: Truck, desc: 'Sumber barang masuk.' },
  { name: 'Lokasi / Gudang', icon: Warehouse, desc: 'Titik penyimpanan stok.' },
]

const transaksi = [
  { name: 'Barang Masuk', icon: ArrowDownToLine, desc: 'Stok gudang bertambah, tercatat dari supplier mana.' },
  { name: 'Barang Keluar', icon: ArrowUpFromLine, desc: 'Stok berkurang, sistem cek kecukupan dulu.' },
  { name: 'Penyesuaian Stok', icon: RefreshCw, desc: 'Selisih stok sistem vs fisik, otomatis dihitung.' },
  { name: 'Transfer Gudang', icon: ArrowLeftRight, desc: 'Pindah stok antar gudang, dua sisi tercatat.' },
]

const laporan = [
  { name: 'Laporan Stok', icon: BarChart3, desc: 'Posisi stok terkini di semua gudang.' },
  { name: 'Kartu Stok', icon: ScrollText, desc: 'Riwayat keluar-masuk per barang, satu kartu satu SKU.' },
  { name: 'Laporan Barang Masuk', icon: ArrowDownToLine, desc: 'Rekap penerimaan per supplier & periode.' },
  { name: 'Laporan Barang Keluar', icon: ArrowUpFromLine, desc: 'Rekap pengeluaran per gudang & periode.' },
  { name: 'Laporan Pergerakan', icon: Activity, desc: 'Semua tipe movement digabung dalam satu linimasa.' },
]

const misi = [
  'Setiap perubahan stok tervalidasi di server, bukan ditebak dari frontend.',
  'Satu database transaction untuk satu transaksi — tidak ada stok yang setengah tersimpan.',
  'Semua pergerakan, sekecil apa pun, tercatat di stock_movements dan bisa ditelusuri.',
  'Struktur data (satuan, kategori, gudang) konsisten, bukan teks bebas yang gampang berantakan.',
]

const faqList = [
  {
    q: 'Stok yang ditampilkan itu total atau per gudang?',
    a: 'Dua-duanya. Kolom stok di Data Barang adalah total dari semua gudang, tapi sistem tetap melacak stok per gudang lewat tabel stok_gudang — jadi kamu bisa lihat rincian per lokasi kapan saja.',
  },
  {
    q: 'Apa yang terjadi kalau transaksi gagal di tengah jalan?',
    a: 'Karena setiap transaksi dibungkus database transaction, kalau ada satu langkah yang gagal (misalnya stok kurang), semua perubahan pada transaksi itu dibatalkan (rollback). Tidak ada stok yang berubah setengah-setengah.',
  },
  {
    q: 'Kenapa transaksi tidak bisa dihapus?',
    a: 'Barang Masuk, Barang Keluar, Penyesuaian, dan Transfer Gudang adalah catatan historis. Kalau ada kesalahan, cara yang benar adalah membuat transaksi baru untuk mengoreksinya, supaya riwayat stock_movements tetap lengkap dan bisa diaudit.',
  },
  {
    q: 'Apakah saya bisa tahu siapa yang melakukan suatu transaksi?',
    a: 'Bisa. Setiap baris di stock_movements menyimpan user_id yang melakukan transaksi, jadi kartu stok dan laporan pergerakan bisa menunjukkan siapa mengubah apa.',
  },
  {
    q: 'Kenapa Satuan wajib diisi tapi Kategori boleh kosong?',
    a: 'Satuan menentukan makna angka stok (pcs, box, kg), jadi wajib supaya datanya tidak ambigu. Kategori sifatnya cuma pengelompokan untuk memudahkan pencarian, jadi boleh dikosongkan.',
  },
]

const tickerEntries = [
  { tipe: 'IN', label: 'Keyboard Mekanik', detail: '+50 · Gudang Utama', tone: 'text-success' },
  { tipe: 'OUT', label: 'Mouse Wireless', detail: '−12 · Gudang Cabang', tone: 'text-accent' },
  { tipe: 'TRANSFER', label: 'Kabel USB-C', detail: '30 · Utama → Cabang', tone: 'text-ink' },
  { tipe: 'ADJUSTMENT', label: 'Headset Gaming', detail: '−2 · stok fisik', tone: 'text-muted' },
  { tipe: 'IN', label: 'Monitor 24"', detail: '+8 · Gudang Utama', tone: 'text-success' },
  { tipe: 'OUT', label: 'Webcam HD', detail: '−5 · Gudang Utama', tone: 'text-accent' },
]

const alurKerja = [
  { no: '01', title: 'Input transaksi', desc: 'Pilih gudang, barang, dan jumlah dari form React — tidak ada angka stok yang diketik manual.' },
  { no: '02', title: 'Validasi di server', desc: 'Go mengecek data barang, gudang, dan kecukupan stok sebelum apa pun disimpan.' },
  { no: '03', title: 'Stok & jejak tercatat', desc: 'Dalam satu database transaction: stok gudang diperbarui dan baris baru masuk ke stock_movements.' },
  { no: '04', title: 'Bisa ditelusuri kapan saja', desc: 'Setiap kenaikan atau penurunan stok punya riwayat lengkap — dari mana, ke mana, oleh siapa.' },
]

function TickerCard({ entry }) {
  const Icon = { IN: ArrowDownToLine, OUT: ArrowUpFromLine, TRANSFER: ArrowLeftRight, ADJUSTMENT: RefreshCw }[entry.tipe]
  return (
    <div className="flex items-center gap-3 shrink-0 bg-surface border border-border rounded-xl px-4 py-3 mx-2 w-64">
      <div className={`shrink-0 w-8 h-8 rounded-lg bg-surface-soft flex items-center justify-center ${entry.tone}`}>
        <Icon size={15} />
      </div>
      <div className="min-w-0">
        <p className="text-xs font-mono text-muted tracking-wider">{entry.tipe}</p>
        <p className="text-sm font-medium truncate">{entry.label}</p>
        <p className={`text-xs font-mono ${entry.tone}`}>{entry.detail}</p>
      </div>
    </div>
  )
}

function FaqItem({ item, isOpen, onToggle }) {
  return (
    <div className="border border-border rounded-xl bg-surface overflow-hidden">
      <button
        onClick={onToggle}
        className="w-full flex items-center justify-between gap-4 px-5 py-4 text-left"
      >
        <span className="font-medium text-sm">{item.q}</span>
        <ChevronDown
          size={16}
          className={`shrink-0 text-muted transition-transform duration-200 ${isOpen ? 'rotate-180' : ''}`}
        />
      </button>
      <div
        className={`grid transition-all duration-200 ease-in-out ${isOpen ? 'grid-rows-[1fr]' : 'grid-rows-[0fr]'}`}
      >
        <div className="overflow-hidden">
          <p className="px-5 pb-4 text-sm text-muted">{item.a}</p>
        </div>
      </div>
    </div>
  )
}

function PublicHome() {
  const doubledTicker = [...tickerEntries, ...tickerEntries]
  const [openFaq, setOpenFaq] = useState(0)

  return (
    <div className="min-h-screen bg-canvas text-ink">
      <style>{`
        @keyframes Inventory AppsTicker {
          from { transform: translateX(0); }
          to { transform: translateX(-50%); }
        }
        .Inventory Apps-ticker-track {
          animation: Inventory AppsTicker 28s linear infinite;
        }
        .Inventory Apps-ticker-mask {
          -webkit-mask-image: linear-gradient(to right, transparent, black 8%, black 92%, transparent);
          mask-image: linear-gradient(to right, transparent, black 8%, black 92%, transparent);
        }
        @media (prefers-reduced-motion: reduce) {
          .Inventory Apps-ticker-track {
            animation: none;
          }
        }
      `}</style>

      {/* Navbar */}
      <nav className="flex items-center justify-between px-6 md:px-12 py-6 border-b border-border">
        <span className="font-display text-lg font-semibold tracking-tight">
          Inventory<span className="text-accent">Apps</span>.
        </span>
        <div className="hidden md:flex items-center gap-7 text-sm text-muted">
          <a href="#tentang" className="hover:text-ink transition-colors">Tentang</a>
          <a href="#visi-misi" className="hover:text-ink transition-colors">Visi &amp; Misi</a>
          <a href="#modul" className="hover:text-ink transition-colors">Modul</a>
          <a href="#alur" className="hover:text-ink transition-colors">Alur Kerja</a>
          <a href="#faq" className="hover:text-ink transition-colors">FAQ</a>
          <a href="#kontak" className="hover:text-ink transition-colors">Kontak</a>
        </div>
        <Link
          to="/login"
          className="rounded-full bg-accent px-5 py-2 text-sm font-medium hover:bg-accent-soft transition-colors"
        >
          Masuk
        </Link>
      </nav>

      {/* Hero */}
      <section className="px-6 md:px-12 pt-20 pb-16 grid md:grid-cols-2 gap-12 items-center max-w-6xl mx-auto">
        <div>
          <p className="font-mono text-xs text-accent tracking-widest mb-4">
            MULTI-GUDANG · JEJAK STOK LENGKAP
          </p>
          <h1 className="font-display text-4xl md:text-5xl font-semibold leading-tight mb-6">
            Stok bergerak,<br />kamu selalu tahu ke mana.
          </h1>
          <p className="text-muted text-base md:text-lg mb-8 max-w-md">
            Inventory Apps mencatat setiap barang masuk, keluar, transfer antar gudang,
            dan penyesuaian stok sebagai satu jejak yang bisa ditelusuri —
            bukan cuma angka stok yang berubah diam-diam.
          </p>
          <div className="flex flex-wrap items-center gap-4">
            <Link
              to="/login"
              className="inline-flex items-center gap-2 rounded-full bg-accent px-7 py-3 font-medium hover:bg-accent-soft transition-colors"
            >
              Masuk ke Akun
              <ArrowRight size={16} />
            </Link>
            <a href="#modul" className="text-sm text-muted hover:text-ink transition-colors">
              Lihat semua modul →
            </a>
          </div>
        </div>

        {/* Signature: pita pergerakan stok, terinspirasi conveyor gudang */}
        <div className="relative">
          <div className="relative rounded-2xl bg-surface border border-border shadow-2xl p-6 -rotate-2">
            <div className="flex items-center justify-between mb-4">
              <p className="font-mono text-xs text-muted tracking-widest">STOCK_MOVEMENTS</p>
              <span className="flex items-center gap-1.5 text-xs text-success">
                <span className="w-1.5 h-1.5 rounded-full bg-success" />
                live
              </span>
            </div>
            <div className="overflow-hidden -mx-2 Inventory Apps-ticker-mask">
              <div className="flex Inventory Apps-ticker-track w-max py-1">
                {doubledTicker.map((entry, i) => (
                  <TickerCard key={i} entry={entry} />
                ))}
              </div>
            </div>
            <p className="text-xs text-muted mt-4 flex items-center gap-1.5">
              <History size={13} />
              Setiap baris di atas = satu transaksi, satu database transaction.
            </p>
          </div>
        </div>
      </section>

      {/* Tentang Kami */}
      <section id="tentang" className="px-6 md:px-12 py-16 border-t border-border">
        <div className="max-w-6xl mx-auto grid md:grid-cols-3 gap-10">
          <div>
            <p className="font-mono text-xs text-accent tracking-widest mb-2">TENTANG KAMI</p>
            <h2 className="font-display text-2xl md:text-3xl font-semibold">
              Dibangun dari masalah gudang yang nyata.
            </h2>
          </div>
          <div className="md:col-span-2 space-y-4 text-muted text-sm md:text-base">
            <p>
              Inventory Apps lahir dari pertanyaan sederhana: kalau stok berkurang, kenapa kita sering
              tidak tahu barangnya kemana? Kebanyakan pencatatan manual cuma menyimpan angka
              terakhir, bukan ceritanya — siapa yang mengeluarkan, kapan, dari gudang mana.
            </p>
            <p>
              Sistem ini dirancang supaya setiap angka stok punya jejak yang bisa dipertanggungjawabkan:
              dari barang masuk pertama kali, pindah antar gudang, sampai penyesuaian karena selisih
              fisik. React di depan cuma mengirim niat transaksi; Go di belakang yang memutuskan dan
              mencatatnya dengan disiplin.
            </p>
          </div>
        </div>
      </section>

      {/* Visi & Misi */}
      <section id="visi-misi" className="px-6 md:px-12 py-16 border-t border-border">
        <div className="max-w-6xl mx-auto grid md:grid-cols-2 gap-10">
          <div className="bg-surface border border-border rounded-2xl p-7">
            <div className="w-10 h-10 rounded-lg bg-surface-soft flex items-center justify-center text-accent mb-4">
              <Compass size={18} />
            </div>
            <p className="font-mono text-xs text-accent tracking-widest mb-2">VISI</p>
            <p className="font-display text-lg font-semibold leading-snug">
              Jadi sistem pencatatan stok yang datanya bisa dipercaya oleh siapa pun yang menyentuhnya —
              dari admin gudang sampai yang baru Inventory Apps dan React.
            </p>
          </div>

          <div className="bg-surface border border-border rounded-2xl p-7">
            <div className="w-10 h-10 rounded-lg bg-surface-soft flex items-center justify-center text-accent mb-4">
              <Target size={18} />
            </div>
            <p className="font-mono text-xs text-accent tracking-widest mb-3">MISI</p>
            <ul className="space-y-2.5">
              {misi.map((item, i) => (
                <li key={i} className="flex items-start gap-2.5 text-sm text-muted">
                  <span className="mt-1.5 w-1 h-1 rounded-full bg-accent shrink-0" />
                  {item}
                </li>
              ))}
            </ul>
          </div>
        </div>
      </section>

      {/* Modul: Master Data, Transaksi & Laporan (mencerminkan struktur nyata sidebar) */}
      <section id="modul" className="px-6 md:px-12 py-16 border-t border-border">
        <div className="max-w-6xl mx-auto">
          <p className="font-mono text-xs text-accent tracking-widest mb-2">MODUL</p>
          <h2 className="font-display text-2xl md:text-3xl font-semibold mb-10 max-w-xl">
            Tiga sisi sistem: data yang stabil, transaksi yang bergerak, laporan yang bicara.
          </h2>

          <div className="grid md:grid-cols-3 gap-8">
            <div>
              <h3 className="text-sm font-mono text-muted tracking-widest mb-4">MASTER DATA</h3>
              <div className="space-y-3">
                {masterData.map((item) => {
                  const Icon = item.icon
                  return (
                    <div key={item.name} className="flex items-start gap-3 bg-surface border border-border rounded-xl p-4">
                      <div className="shrink-0 w-9 h-9 rounded-lg bg-surface-soft flex items-center justify-center text-accent">
                        <Icon size={16} />
                      </div>
                      <div>
                        <p className="font-medium text-sm mb-0.5">{item.name}</p>
                        <p className="text-xs text-muted">{item.desc}</p>
                      </div>
                    </div>
                  )
                })}
              </div>
            </div>

            <div>
              <h3 className="text-sm font-mono text-muted tracking-widest mb-4">TRANSAKSI</h3>
              <div className="space-y-3">
                {transaksi.map((item) => {
                  const Icon = item.icon
                  return (
                    <div key={item.name} className="flex items-start gap-3 bg-surface border border-border rounded-xl p-4">
                      <div className="shrink-0 w-9 h-9 rounded-lg bg-surface-soft flex items-center justify-center text-accent">
                        <Icon size={16} />
                      </div>
                      <div>
                        <p className="font-medium text-sm mb-0.5">{item.name}</p>
                        <p className="text-xs text-muted">{item.desc}</p>
                      </div>
                    </div>
                  )
                })}
              </div>
            </div>

            <div>
              <h3 className="text-sm font-mono text-muted tracking-widest mb-4">LAPORAN</h3>
              <div className="space-y-3">
                {laporan.map((item) => {
                  const Icon = item.icon
                  return (
                    <div key={item.name} className="flex items-start gap-3 bg-surface border border-border rounded-xl p-4">
                      <div className="shrink-0 w-9 h-9 rounded-lg bg-surface-soft flex items-center justify-center text-accent">
                        <Icon size={16} />
                      </div>
                      <div>
                        <p className="font-medium text-sm mb-0.5">{item.name}</p>
                        <p className="text-xs text-muted">{item.desc}</p>
                      </div>
                    </div>
                  )
                })}
              </div>
            </div>
          </div>
        </div>
      </section>

      {/* Alur kerja — proses nyata, sekuensial */}
      <section id="alur" className="px-6 md:px-12 py-16 border-t border-border">
        <div className="max-w-6xl mx-auto">
          <p className="font-mono text-xs text-accent tracking-widest mb-2">ALUR KERJA</p>
          <h2 className="font-display text-2xl md:text-3xl font-semibold mb-10 max-w-lg">
            Dari klik di React sampai baris di database.
          </h2>
          <div className="grid md:grid-cols-4 gap-8">
            {alurKerja.map((item) => (
              <div key={item.no}>
                <p className="font-mono text-accent text-sm mb-3">{item.no}</p>
                <h3 className="font-display text-base font-semibold mb-2">{item.title}</h3>
                <p className="text-muted text-sm">{item.desc}</p>
              </div>
            ))}
          </div>

          <div className="mt-10 flex items-start gap-3 bg-surface border border-border rounded-2xl p-5 max-w-2xl">
            <ShieldCheck size={18} className="text-accent shrink-0 mt-0.5" />
            <p className="text-sm text-muted">
              Frontend tidak pernah mengubah angka stok secara langsung. Semua perhitungan —
              penambahan, pengurangan, transfer, maupun selisih penyesuaian — dilakukan di backend Go
              dalam satu database transaction, supaya stok tidak pernah setengah-tersimpan.
            </p>
          </div>
        </div>
      </section>

      {/* FAQ */}
      <section id="faq" className="px-6 md:px-12 py-16 border-t border-border">
        <div className="max-w-3xl mx-auto">
          <p className="font-mono text-xs text-accent tracking-widest mb-2">FAQ</p>
          <h2 className="font-display text-2xl md:text-3xl font-semibold mb-10">
            Pertanyaan yang sering muncul soal cara kerja stoknya.
          </h2>
          <div className="space-y-3">
            {faqList.map((item, i) => (
              <FaqItem
                key={i}
                item={item}
                isOpen={openFaq === i}
                onToggle={() => setOpenFaq(openFaq === i ? -1 : i)}
              />
            ))}
          </div>
        </div>
      </section>

      {/* Kontak */}
      <section id="kontak" className="px-6 md:px-12 py-16 border-t border-border">
        <div className="max-w-6xl mx-auto grid md:grid-cols-2 gap-10 items-start">
          <div>
            <p className="font-mono text-xs text-accent tracking-widest mb-2">KONTAK</p>
            <h2 className="font-display text-2xl md:text-3xl font-semibold mb-4 max-w-md">
              Ada pertanyaan soal sistem atau mau lapor bug?
            </h2>
            <p className="text-muted text-sm md:text-base max-w-md">
              Hubungi admin sistem lewat salah satu kanal berikut. Ganti detail di bawah ini
              sesuai kontak tim kamu sebelum dipakai beneran.
            </p>
          </div>

          <div className="space-y-3">
            <a
              href="mailto:admin@Inventory Apps.local"
              className="flex items-center gap-4 bg-surface border border-border rounded-xl p-4 hover:border-accent/40 transition-colors"
            >
              <div className="shrink-0 w-9 h-9 rounded-lg bg-surface-soft flex items-center justify-center text-accent">
                <Mail size={16} />
              </div>
              <div>
                <p className="text-xs text-muted mb-0.5">Email</p>
                <p className="text-sm font-medium">admin@Inventory Apps.local</p>
              </div>
            </a>

            <a
              href="https://github.com/username/belajar_go"
              target="_blank"
              rel="noopener noreferrer"
              className="flex items-center gap-4 bg-surface border border-border rounded-xl p-4 hover:border-accent/40 transition-colors"
            >
              <div className="shrink-0 w-9 h-9 rounded-lg bg-surface-soft flex items-center justify-center text-accent">
                <Link2 size={16} />
              </div>
              <div>
                <p className="text-xs text-muted mb-0.5">Repository</p>
                <p className="text-sm font-medium">github.com/username/belajar_go</p>
              </div>
            </a>

            <div className="flex items-center gap-4 bg-surface border border-border rounded-xl p-4">
              <div className="shrink-0 w-9 h-9 rounded-lg bg-surface-soft flex items-center justify-center text-accent">
                <MapPin size={16} />
              </div>
              <div>
                <p className="text-xs text-muted mb-0.5">Lokasi</p>
                <p className="text-sm font-medium">Gudang Utama — alamat menyusul</p>
              </div>
            </div>
          </div>
        </div>
      </section>

      <footer className="px-6 md:px-12 py-8 border-t border-border text-center text-muted text-xs">
        Dibuat untuk Inventory Apps &amp; React — {new Date().getFullYear()}
      </footer>
    </div>
  )
}

export default PublicHome