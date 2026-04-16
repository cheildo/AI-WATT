import { useState } from 'react'
import { ActivityTable } from '@/components/ActivityTable'
import { useActivity } from '@/hooks/api/useActivity'
import { ActionButton } from '@/components/ActionButton'

export function Activity() {
  const [page, setPage] = useState(1)
  const { data, isLoading } = useActivity(page)

  return (
    <div className="px-4 py-8">
      <h1 className="mb-6 text-2xl font-bold text-text-primary">Activity</h1>

      {isLoading ? (
        <div className="flex justify-center py-12">
          <span className="h-6 w-6 animate-spin rounded-full border-2 border-brand border-t-transparent" />
        </div>
      ) : (
        <ActivityTable events={data?.events ?? []} />
      )}

      <div className="mt-4 flex items-center justify-between">
        <p className="text-sm text-text-secondary">
          Total: {data?.total ?? 0} events
        </p>
        <div className="flex gap-2">
          <ActionButton variant="secondary" disabled={page === 1} onClick={() => setPage((p) => p - 1)}>
            ← Prev
          </ActionButton>
          <ActionButton variant="secondary" onClick={() => setPage((p) => p + 1)}>
            Next →
          </ActionButton>
        </div>
      </div>
    </div>
  )
}
