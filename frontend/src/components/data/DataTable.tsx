import { ReactNode } from 'react'

export interface TableColumn {
  key: string
  header: string
  render?: (row: Record<string, unknown>) => ReactNode
}

interface DataTableProps {
  columns: TableColumn[]
  rows: Record<string, unknown>[]
  onRowClick?: (row: Record<string, unknown>) => void
}

export function DataTable({ columns, rows, onRowClick }: DataTableProps) {
  return (
    <div style={{ overflowX: 'auto' }}>
      <table style={{ width: '100%', borderCollapse: 'collapse', fontSize: 12 }}>
        <thead>
          <tr className="bg-bg-2">
            {columns.map((col) => (
              <th
                key={col.key}
                className="text-left text-text-3 uppercase tracking-[.08em] border-b border-border font-medium"
                style={{ padding: '8px 14px', fontSize: 10 }}
              >
                {col.header}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((row, i) => (
            <tr
              key={i}
              className="border-b border-border hover:bg-bg cursor-pointer"
              onClick={() => onRowClick?.(row)}
            >
              {columns.map((col) => (
                <td
                  key={col.key}
                  className="text-text-2 font-mono"
                  style={{ padding: '10px 14px' }}
                >
                  {col.render ? col.render(row) : String(row[col.key] ?? '')}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
