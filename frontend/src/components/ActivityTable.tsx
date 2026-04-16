import type { ChainEvent } from '@/hooks/api/useActivity'

function shortHash(h: string) { return `${h.slice(0, 8)}…${h.slice(-6)}` }

export function ActivityTable({ events }: { events: ChainEvent[] }) {
  if (events.length === 0) {
    return <p className="py-8 text-center text-sm text-text-secondary">No activity yet.</p>
  }
  return (
    <div className="overflow-x-auto rounded-xl border border-surface-border">
      <table className="w-full text-sm">
        <thead>
          <tr className="border-b border-surface-border bg-surface-card">
            <th className="px-4 py-3 text-left font-medium text-text-secondary">Event</th>
            <th className="px-4 py-3 text-left font-medium text-text-secondary">Block</th>
            <th className="px-4 py-3 text-left font-medium text-text-secondary">Tx</th>
            <th className="px-4 py-3 text-left font-medium text-text-secondary">Time</th>
          </tr>
        </thead>
        <tbody className="divide-y divide-surface-border">
          {events.map((e) => (
            <tr key={e.id} className="hover:bg-surface-hover transition-colors">
              <td className="px-4 py-3">
                <span className="rounded bg-brand-muted px-1.5 py-0.5 text-xs font-mono text-brand">
                  {e.eventName}
                </span>
              </td>
              <td className="px-4 py-3 font-mono text-text-secondary">{e.blockNumber}</td>
              <td className="px-4 py-3 font-mono text-brand">{shortHash(e.txHash)}</td>
              <td className="px-4 py-3 text-text-secondary">
                {new Date(e.createdAt).toLocaleString()}
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
