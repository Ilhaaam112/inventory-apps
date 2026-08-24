import { Menu } from 'lucide-react'

function Topbar({ title, user, onMenuClick }) {
  const initial = user?.nama_lengkap?.charAt(0)?.toUpperCase() || '?'

  return (
    <header className="sticky top-0 z-20 flex items-center justify-between gap-4 px-6 py-4 bg-canvas/90 backdrop-blur border-b border-border">
      <div className="flex items-center gap-4">
        <button className="md:hidden text-muted" onClick={onMenuClick}>
          <Menu size={22} />
        </button>
        <h1 className="font-display text-lg font-semibold">{title}</h1>
      </div>

      <div className="flex items-center gap-3">
        <span className="hidden sm:block text-sm text-muted">{user?.nama_lengkap}</span>
        <div className="w-9 h-9 rounded-full bg-accent/20 text-accent border border-accent/30 flex items-center justify-center font-display font-semibold text-sm">
          {initial}
        </div>
      </div>
    </header>
  )
}

export default Topbar