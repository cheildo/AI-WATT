interface AttestationStripProps {
  assetCount: number
  lastAttestation: string
  blockNumber: string
}

export function AttestationStrip({ assetCount, lastAttestation, blockNumber }: AttestationStripProps) {
  return (
    <div
      className="bg-green-bg border border-green-border rounded flex items-center justify-between mb-4"
      style={{ padding: '10px 14px', fontSize: 12, borderRadius: 4 }}
    >
      <div className="flex items-center gap-2 text-green-mid">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none">
          <path d="M7 1L13 4v4c0 3-2.5 5-6 6-3.5-1-6-3-6-6V4L7 1z" stroke="currentColor" strokeWidth="1.3" strokeLinejoin="round"/>
          <path d="M4.5 7l2 2 3-3" stroke="currentColor" strokeWidth="1.3" strokeLinecap="round" strokeLinejoin="round"/>
        </svg>
        Agent-Core live · {assetCount} assets certified
      </div>
      <div className="font-mono text-text-3" style={{ fontSize: 11 }}>
        Last attestation: {lastAttestation} · XDC block #{blockNumber}
      </div>
    </div>
  )
}
