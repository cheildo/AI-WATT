import { useNavigate, useLocation } from 'react-router-dom'
import { cn } from '@/lib/formatters'

const TABS = [
  { label: 'Buy',    path: '/'        },
  { label: 'Stake',  path: '/stake'   },
  { label: 'Borrow', path: '/borrow'  },
  { label: 'Bridge', path: '/bridge'  },
]

export function TabBar() {
  const navigate = useNavigate()
  const { pathname } = useLocation()

  return (
    <div
      className="grid border border-border overflow-hidden mb-5"
      style={{
        gridTemplateColumns: 'repeat(4, 1fr)',
        borderRadius: 2,
        maxWidth: 820,
      }}
    >
      {TABS.map((tab) => {
        const active = pathname === tab.path
        return (
          <button
            key={tab.path}
            onClick={() => navigate(tab.path)}
            className={cn(
              'text-center border-r border-border font-sans font-semibold uppercase tracking-[.1em] cursor-pointer transition-colors',
              'last:border-r-0',
              active
                ? 'bg-green text-white'
                : 'bg-bg-2 text-text-3 hover:bg-bg-3 hover:text-text-2'
            )}
            style={{ padding: '12px 10px', fontSize: 11 }}
          >
            {tab.label}
          </button>
        )
      })}
    </div>
  )
}
