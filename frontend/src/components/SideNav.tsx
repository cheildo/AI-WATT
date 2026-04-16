import { NavLink } from 'react-router-dom'

const links = [
  { to: '/buy',        label: 'Buy',        icon: '⚡' },
  { to: '/stake',      label: 'Stake',      icon: '🔒' },
  { to: '/borrow',     label: 'Borrow',     icon: '🏦' },
  { to: '/portfolio',  label: 'Portfolio',  icon: '📊' },
  { to: '/veriflow',   label: 'Veriflow',   icon: '🖥️' },
  { to: '/activity',   label: 'Activity',   icon: '📋' },
  { to: '/governance', label: 'Governance', icon: '🗳️' },
  { to: '/bridge',     label: 'Bridge',     icon: '🌉' },
]

export function SideNav() {
  return (
    <nav className="flex w-52 shrink-0 flex-col border-r border-surface-border bg-surface pt-4">
      {links.map(({ to, label, icon }) => (
        <NavLink
          key={to}
          to={to}
          className={({ isActive }) =>
            `flex items-center gap-3 px-5 py-2.5 text-sm transition-colors ${
              isActive
                ? 'border-r-2 border-brand bg-brand-muted text-brand'
                : 'text-text-secondary hover:bg-surface-hover hover:text-text-primary'
            }`
          }
        >
          <span className="text-base leading-none">{icon}</span>
          {label}
        </NavLink>
      ))}
    </nav>
  )
}
