import './App.css'

import Navbar from './Components/Navbar'
import Footer from './Components/Footer'
import CookingSim from './SubPages/CookingSim'

function App() {
  
  return (
    <> 
      <Navbar />
      <div className="flex min-h-80 min-w-20 border-red-500">
        <CookingSim />
      </div>
      <Footer />
    </>
  )
}

export default App
