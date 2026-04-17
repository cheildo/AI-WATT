import { ReactNode } from 'react'
import { Link, useLocation } from 'react-router-dom'
import * as Tooltip from '@radix-ui/react-tooltip'
import { cn } from '@/lib/formatters'

interface NavIconProps {
  to: string
  tip: string
  children: ReactNode
}

export function NavIcon({ to, tip, children }: NavIconProps) {
  const { pathname } = useLocation()
  const active = pathname === to || (to !== '/' && pathname.startsWith(to))

  return (
    <Tooltip.Provider delayDuration={0}>
      <Tooltip.Root>
        <Tooltip.Trigger asChild>
          <Link to={to} className="block my-[1px]">
            <div
              className={cn(
                'w-9 h-9 rounded-md flex items-center justify-center cursor-pointer transition-colors',
                active
                  ? 'bg-green text-white'
                  : 'text-text-3 hover:bg-bg-2 hover:text-text-1'
              )}
            >
              {children}
            </div>
          </Link>
        </Tooltip.Trigger>
        <Tooltip.Portal>
          <Tooltip.Content
            side="right"
            sideOffset={10}
            className="bg-green text-white rounded font-sans font-medium whitespace-nowrap z-50"
            style={{ padding: '4px 9px', fontSize: 11 }}
          >
            {tip}
          </Tooltip.Content>
        </Tooltip.Portal>
      </Tooltip.Root>
    </Tooltip.Provider>
  )
}
