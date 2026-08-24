import { Link, useLocation } from 'react-router-dom'
import {
  LayoutDashboard, Package, Tags, Truck, BarChart3, User, LogOut, Boxes, X, Ruler, Warehouse,
  ArrowDownToLine, ArrowUpFromLine, SlidersHorizontal, ArrowLeftRight, History,
  FileText, FileInput, FileOutput, Activity,
} from 'lucide-react'

const menuSections = [
  {
    label: 'Utama',
    items: [{ name: 'Dashboard', icon: LayoutDashboard, path: '/dashboard' }],
  },
  {
    label: 'Master Data',
    items: [
      { name: 'Data Barang', icon: Package, path: '/barang' },
      { name: 'Kategori Barang', icon: Tags, path: '/kategori' },
      { name: 'Satuan', icon: Ruler, path: '/satuan' },
      { name: 'Supplier', icon: Truck, path: '/supplier' },
      { name: 'Lokasi / Gudang', icon: Warehouse, path: '/lokasi' },
    ],
  },
  {
    label: 'Transaksi',
    items: [
      { name: 'Barang Masuk', icon: ArrowDownToLine, path: '/barang-masuk' },
      { name: 'Barang Keluar', icon: ArrowUpFromLine, path: '/barang-keluar' },
      { name: 'Penyesuaian Stok', icon: SlidersHorizontal, path: '/penyesuaian' },
      { name: 'Transfer Gudang', icon: ArrowLeftRight, path: '/transfer' },
    ],
  },
  {
    label: 'Laporan',
    items: [
      { name: 'Laporan Stok', icon: BarChart3, path: '/laporan/stok' },
      { name: 'Kartu Stok', icon: History, path: '/laporan/kartu-stok' },
      { name: 'Barang Masuk', icon: FileInput, path: '/laporan/barang-masuk' },
      { name: 'Barang Keluar', icon: FileOutput, path: '/laporan/barang-keluar' },
      { name: 'Pergerakan Stok', icon: Activity, path: '/laporan/pergerakan' },
    ],
  },
{
  label: 'Akun',
  items: [{ name: 'Profil', icon: User, path: '/profil' }], // hapus soon: true
},
]

function Sidebar({ mobileOpen, setMobileOpen, onLogout }) {
  const location = useLocation()

  return (
    <>
      {mobileOpen && (
        <div className="fixed inset-0 bg-black/60 z-30 md:hidden" onClick={() => setMobileOpen(false)} />
      )}

      <aside
        className={`fixed md:sticky top-0 left-0 h-screen w-64 bg-surface border-r border-border flex flex-col z-40 transition-transform duration-200 ${
          mobileOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0'
        }`}
      >
        <div className="flex items-center justify-between px-6 py-6 border-b border-border">
          <span className="font-display text-lg font-semibold flex items-center gap-2">
            <Boxes size={20} className="text-accent" />
            belajar<span className="text-accent">Go</span>.
          </span>
          <button className="md:hidden text-muted" onClick={() => setMobileOpen(false)}>
            <X size={20} />
          </button>
        </div>

        <nav className="flex-1 overflow-y-auto px-4 py-6 space-y-8">
          {menuSections.map((section) => (
            <div key={section.label}>
              <p className="px-2 mb-2 text-xs font-mono text-muted tracking-widest uppercase">
                {section.label}
              </p>
              <div className="space-y-1">
                {section.items.map((item) => {
                  const Icon = item.icon
                  const active = location.pathname === item.path

                  if (item.soon) {
                    return (
                      <div
                        key={item.name}
                        className="flex items-center justify-between px-3 py-2.5 rounded-lg text-muted/60 cursor-not-allowed"
                      >
                        <span className="flex items-center gap-3 text-sm">
                          <Icon size={17} />
                          {item.name}
                        </span>
                        <span className="text-[10px] font-mono border border-border rounded-full px-2 py-0.5">
                          SEGERA
                        </span>
                      </div>
                    )
                  }

                  return (
                    <Link
                      key={item.name}
                      to={item.path}
                      onClick={() => setMobileOpen(false)}
                      className={`flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm transition-colors border ${
                        active
                          ? 'bg-accent/10 text-accent border-accent/30'
                          : 'text-muted hover:text-ink hover:bg-surface-soft border-transparent'
                      }`}
                    >
                      <Icon size={17} />
                      {item.name}
                    </Link>
                  )
                })}
              </div>
            </div>
          ))}
        </nav>

        <div className="px-4 py-4 border-t border-border">
          <button
            onClick={onLogout}
            className="w-full flex items-center gap-3 px-3 py-2.5 rounded-lg text-sm text-muted hover:text-accent hover:bg-surface-soft transition-colors"
          >
            <LogOut size={17} />
            Keluar
          </button>
        </div>
      </aside>
    </>
  )
}

export default Sidebar
