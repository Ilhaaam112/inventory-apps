import { useState } from 'react'
import Sidebar from './Sidebar'
import Topbar from './Topbar'

function Layout({ title, user, onLogout, children }) {
  const [mobileOpen, setMobileOpen] = useState(false)

  return (
    <div className="min-h-screen bg-canvas text-ink md:flex">
      <Sidebar mobileOpen={mobileOpen} setMobileOpen={setMobileOpen} onLogout={onLogout} />
      <div className="flex-1 min-w-0">
        <Topbar title={title} user={user} onMenuClick={() => setMobileOpen(true)} />
        <main className="p-6 md:p-8">{children}</main>
      </div>
    </div>
  )
}

export default Layout