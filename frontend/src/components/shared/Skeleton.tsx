interface SkeletonProps {
  width?: number | string
  height?: number
  radius?: number
  className?: string
}

export function Skeleton({ width, height = 14, radius = 3, className }: SkeletonProps) {
  return (
    <div
      className={className}
      style={{
        width: width ?? '100%',
        height,
        borderRadius: radius,
        background: 'linear-gradient(90deg, #DDE1EC 25%, #EAECF3 50%, #DDE1EC 75%)',
        backgroundSize: '200% 100%',
        animation: 'skeleton-shimmer 1.5s infinite',
        flexShrink: 0,
      }}
    />
  )
}
