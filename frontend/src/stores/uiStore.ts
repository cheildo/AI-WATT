import { create } from 'zustand'

export type ToastVariant = 'success' | 'error' | 'info'

interface ToastState {
  message: string
  visible: boolean
  variant: ToastVariant
  key: number           // increments each show — lets Toast reset its progress bar
}

interface UIStore {
  toast: ToastState
  showToast: (message: string, variant?: ToastVariant) => void
  hideToast: () => void
}

let timer: ReturnType<typeof setTimeout> | null = null
const DURATION = 4000

export const useUIStore = create<UIStore>((set) => ({
  toast: { message: '', visible: false, variant: 'success', key: 0 },

  showToast: (message, variant = 'success') => {
    if (timer) clearTimeout(timer)
    set((s) => ({ toast: { message, visible: true, variant, key: s.toast.key + 1 } }))
    timer = setTimeout(
      () => set((s) => ({ toast: { ...s.toast, visible: false } })),
      DURATION,
    )
  },

  hideToast: () => {
    if (timer) clearTimeout(timer)
    set((s) => ({ toast: { ...s.toast, visible: false } }))
  },
}))
