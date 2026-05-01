import { useUIStore, ToastVariant } from '@/stores/uiStore'

const VARIANTS: Record<ToastVariant, { border: string; icon: string; iconBg: string }> = {
  success: { border: '#0A7068', icon: '✓', iconBg: '#EAF5F3' },
  error:   { border: '#8B2020', icon: '✗', iconBg: '#FBF0F0' },
  info:    { border: '#2563EB', icon: 'i', iconBg: '#EEF4FF' },
}

export function Toast() {
  const { toast, hideToast } = useUIStore()
  const c = VARIANTS[toast.variant]

  return (
    <>
      {/* keyframe injected once via this component */}
      <style>{`
        @keyframes toast-bar {
          from { width: 100% }
          to   { width: 0% }
        }
      `}</style>

      <div
        role="status"
        aria-live="polite"
        className="fixed z-[9999] font-sans transition-all duration-300"
        style={{
          bottom: 24, right: 24,
          background: '#FDFCF9',
          border: '1px solid #D8D3C6',
          borderLeft: `3px solid ${c.border}`,
          borderRadius: 4,
          boxShadow: '0 4px 24px rgba(0,0,0,0.12)',
          minWidth: 280, maxWidth: 380,
          overflow: 'hidden',
          transform: toast.visible ? 'translateY(0) scale(1)' : 'translateY(70px) scale(0.97)',
          opacity: toast.visible ? 1 : 0,
          pointerEvents: toast.visible ? 'auto' : 'none',
        }}
      >
        {/* Body */}
        <div className="flex items-start gap-[10px]" style={{ padding: '13px 14px 12px' }}>
          {/* Icon badge */}
          <span
            className="flex-shrink-0 flex items-center justify-center rounded-full font-bold"
            style={{
              width: 18, height: 18, fontSize: 9, marginTop: 1,
              background: c.iconBg,
              color: c.border,
              border: `1.5px solid ${c.border}40`,
            }}
          >{c.icon}</span>

          {/* Message */}
          <span className="flex-1 text-text-1 leading-snug" style={{ fontSize: 12 }}>
            {toast.message}
          </span>

          {/* Dismiss */}
          <button
            onClick={hideToast}
            className="flex-shrink-0 text-text-3 hover:text-text-1 cursor-pointer transition-colors"
            style={{ fontSize: 16, background: 'none', border: 'none', padding: '0 0 0 6px', lineHeight: 1 }}
            aria-label="Dismiss"
          >×</button>
        </div>

        {/* Auto-dismiss progress bar */}
        <div style={{ height: 2, background: '#EAECF3' }}>
          <div
            key={toast.key}
            style={{
              height: '100%',
              background: c.border,
              width: toast.visible ? '100%' : '0%',
              animation: toast.visible ? 'toast-bar 4s linear forwards' : 'none',
            }}
          />
        </div>
      </div>
    </>
  )
}
