interface Props { score: number; size?: 'sm' | 'md' }

function color(score: number) {
  if (score >= 80) return 'text-yield bg-yield/10 border-yield/30'
  if (score >= 60) return 'text-warn bg-warn/10 border-warn/30'
  return 'text-danger bg-danger/10 border-danger/30'
}

export function HealthBadge({ score, size = 'md' }: Props) {
  const sz = size === 'sm' ? 'px-2 py-0.5 text-xs' : 'px-3 py-1 text-sm'
  return (
    <span className={`inline-flex items-center gap-1.5 rounded-full border font-medium ${color(score)} ${sz}`}>
      <span className="h-1.5 w-1.5 rounded-full bg-current" />
      {score.toFixed(1)}
    </span>
  )
}
