import { Navigate, Route, Routes } from 'react-router-dom'
import { NavBar } from '@/components/NavBar'
import { SideNav } from '@/components/SideNav'
import { ToastProvider } from '@/components/ToastProvider'
import { Buy }        from '@/pages/Buy'
import { Stake }      from '@/pages/Stake'
import { Borrow }     from '@/pages/Borrow'
import { Portfolio }  from '@/pages/Portfolio'
import { Veriflow }   from '@/pages/Veriflow'
import { Activity }   from '@/pages/Activity'
import { Governance } from '@/pages/Governance'
import { Bridge }     from '@/pages/Bridge'

export function App() {
  return (
    <ToastProvider>
      <div className="flex h-screen flex-col bg-surface font-sans text-text-primary">
        <NavBar />
        <div className="flex flex-1 overflow-hidden">
          <SideNav />
          <main className="flex-1 overflow-y-auto">
            <Routes>
              <Route path="/"          element={<Navigate to="/buy" replace />} />
              <Route path="/buy"        element={<Buy />} />
              <Route path="/stake"      element={<Stake />} />
              <Route path="/borrow"     element={<Borrow />} />
              <Route path="/portfolio"  element={<Portfolio />} />
              <Route path="/veriflow"   element={<Veriflow />} />
              <Route path="/activity"   element={<Activity />} />
              <Route path="/governance" element={<Governance />} />
              <Route path="/bridge"     element={<Bridge />} />
              <Route path="*"           element={<Navigate to="/buy" replace />} />
            </Routes>
          </main>
        </div>
      </div>
    </ToastProvider>
  )
}
