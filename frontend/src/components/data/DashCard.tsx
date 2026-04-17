import { useState, useEffect, useRef, ReactNode } from 'react'
import { cn } from '@/lib/formatters'

function AreaChart({ data, stroke, fill }: { data: number[]; stroke: string; fill: string }) {
  const ref = useRef<HTMLDivElement>(null)
  const [w, setW] = useState(420)
  const h = 160

  useEffect(() => {
    const el = ref.current
    if (!el) return
    const update = () => setW(el.clientWidth || 420)
    update()
    const ro = new ResizeObserver(update)
    ro.observe(el)
    return () => ro.disconnect()
  }, [])

  const max = Math.max(...data)
  const pts = data.map((v, i) => {
    const x = (i / (data.length - 1)) * w
    const y = h - (v / max) * (h - 10) - 5
    return [x, y]
  })
  const pathD = pts.map((p, i) => `${i === 0 ? 'M' : 'L'}${p[0].toFixed(1)} ${p[1].toFixed(1)}`).join(' ')
  const areaD = `${pathD} L${w} ${h} L0 ${h} Z`
  const gradId = `g${stroke.replace('#', '')}`

  return (
    <div ref={ref} className="w-full h-full">
      <svg width={w} height={h} viewBox={`0 0 ${w} ${h}`} style={{ overflow: 'visible' }}>
        <defs>
          <linearGradient id={gradId} x1="0" y1="0" x2="0" y2="1">
            <stop offset="0%" stopColor={fill} stopOpacity="0.5"/>
            <stop offset="100%" stopColor={fill} stopOpacity="0.05"/>
          </linearGradient>
        </defs>
        <path d={areaD} fill={`url(#${gradId})`}/>
        <path d={pathD} fill="none" stroke={stroke} strokeWidth="1.5" strokeLinejoin="round"/>
      </svg>
    </div>
  )
}

const TIME_BTNS = ['1W', '1M', '3M', '1Y', 'ALL']

interface DashCardProps {
  label: ReactNode
  value: string
  subValue?: ReactNode
  chartData: number[]
  chartStroke: string
  chartFill: string
}

export function DashCard({ label, value, subValue, chartData, chartStroke, chartFill }: DashCardProps) {
  const [activeTime, setActiveTime] = useState('ALL')

  return (
    <div className="bg-white border border-border rounded overflow-hidden" style={{ padding: 20 }}>
      <div className="text-text-3 mb-[6px] flex items-center gap-[5px]" style={{ fontSize: 11 }}>
        {label}
      </div>
      <div className="font-serif mb-3" style={{ fontSize: 32 }}>
        {value}
        {subValue && <span className="ml-2 text-teal font-sans font-medium" style={{ fontSize: 14 }}>{subValue}</span>}
      </div>
      <div style={{ height: 160, position: 'relative', overflow: 'hidden' }}>
        <AreaChart data={chartData} stroke={chartStroke} fill={chartFill} />
      </div>
      <div className="flex gap-1 mt-[10px]">
        {TIME_BTNS.map((t) => (
          <button
            key={t}
            onClick={() => setActiveTime(t)}
            className={cn(
              'border rounded cursor-pointer transition-all font-semibold',
              activeTime === t
                ? 'bg-green text-white border-green'
                : 'bg-transparent border-border text-text-3 hover:text-text-2'
            )}
            style={{ padding: '3px 8px', fontSize: 10 }}
          >
            {t}
          </button>
        ))}
      </div>
    </div>
  )
}
