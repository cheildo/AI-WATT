import { Routes, Route } from 'react-router-dom'
import { TopBar }    from '@/components/layout/TopBar'
import { Sidebar }   from '@/components/layout/Sidebar'
import { Toast }     from '@/components/shared/Toast'
import { Buy }        from '@/pages/Buy'
import { Stake }      from '@/pages/Stake'
import { Borrow }     from '@/pages/Borrow'
import { Bridge }     from '@/pages/Bridge'
import { Dashboard }  from '@/pages/Dashboard'
import { Veriflow }   from '@/pages/Veriflow'
import { Portfolio }  from '@/pages/Portfolio'
import { Governance } from '@/pages/Governance'

export function App() {
  return (
    <div className="h-full flex flex-col overflow-hidden">
      <TopBar />
      <div className="flex flex-1 overflow-hidden">
        <Sidebar />
        <main
          className="flex-1 overflow-y-auto"
          style={{ padding: '28px 36px' }}
        >
          <Routes>
            <Route path="/"           element={<Buy />} />
            <Route path="/stake"      element={<Stake />} />
            <Route path="/borrow"     element={<Borrow />} />
            <Route path="/bridge"     element={<Bridge />} />
            <Route path="/dashboard"  element={<Dashboard />} />
            <Route path="/veriflow"   element={<Veriflow />} />
            <Route path="/portfolio"  element={<Portfolio />} />
            <Route path="/governance" element={<Governance />} />
          </Routes>
        </main>
      </div>
      <Toast />
    </div>
  )
}
