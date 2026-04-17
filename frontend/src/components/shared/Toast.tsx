import { useUIStore } from '@/stores/uiStore'

export function Toast() {
  const { toast } = useUIStore()

  return (
    <div
      className="fixed bottom-5 right-5 z-[999] bg-green text-white rounded flex items-center gap-2 font-sans transition-all duration-300"
      style={{
        padding: '10px 16px',
        fontSize: 12,
        boxShadow: '0 4px 16px rgba(0,0,0,0.15)',
        transform: toast.visible ? 'translateY(0)' : 'translateY(60px)',
        opacity: toast.visible ? 1 : 0,
        pointerEvents: toast.visible ? 'auto' : 'none',
      }}
    >
      {toast.message}
    </div>
  )
}
