import React from 'react'
import { loadUncleNelly, type UncleNelly } from './unclenelly'

import './App.css'

import Navbar from './Components/Navbar'
import Footer from './Components/Footer'
import CookingSim from './SubPages/CookingSim'

function App() {
  const UN = React.useRef<UncleNelly | null>(null)
  React.useEffect(()=>{
    loadUncleNelly()
      .then((initUN)=>{
        UN.current = initUN()
      })
      .catch( e => console.log(e))
  },[])

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
