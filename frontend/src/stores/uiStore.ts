import { create } from 'zustand'

interface UIStore {
  toast: { message: string; visible: boolean }
  showToast: (message: string) => void
}

let timer: ReturnType<typeof setTimeout> | null = null

export const useUIStore = create<UIStore>((set) => ({
  toast: { message: '', visible: false },
  showToast: (message) => {
    if (timer) clearTimeout(timer)
    set({ toast: { message, visible: true } })
    timer = setTimeout(
      () => set((s) => ({ toast: { ...s.toast, visible: false } })),
      3000
    )
  },
}))
