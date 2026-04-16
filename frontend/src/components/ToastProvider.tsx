import * as Toast from '@radix-ui/react-toast'
import { useTxStore } from '@/stores/txStore'

export function ToastProvider({ children }: { children: React.ReactNode }) {
  const { transactions } = useTxStore()
  const recent = transactions.slice(0, 3)

  return (
    <Toast.Provider swipeDirection="right">
      {children}
      {recent.map((tx) => (
        <Toast.Root
          key={tx.id}
          open
          className="flex items-start gap-3 rounded-xl border border-surface-border bg-surface-card p-4 shadow-xl data-[state=closed]:animate-hide data-[state=open]:animate-slideIn"
        >
          <span
            className={`mt-0.5 h-2 w-2 shrink-0 rounded-full ${
              tx.status === 'success' ? 'bg-yield' : tx.status === 'error' ? 'bg-danger' : 'bg-warn animate-pulse'
            }`}
          />
          <div className="min-w-0 flex-1">
            <Toast.Title className="text-sm font-medium text-text-primary">{tx.description}</Toast.Title>
            {tx.hash && (
              <Toast.Description className="mt-0.5 truncate text-xs text-text-secondary font-mono">
                {tx.hash.slice(0, 10)}…
              </Toast.Description>
            )}
          </div>
        </Toast.Root>
      ))}
      <Toast.Viewport className="fixed bottom-4 right-4 z-50 flex w-96 flex-col gap-2" />
    </Toast.Provider>
  )
}
