import React from 'react'
import { loadUncleNelly, type UncleNelly } from './unclenelly'
import { AppOptions, useAppContext, AppProvider } from './AppContext'

import './App.css'

import Navbar from './Components/Navbar'
import Footer from './Components/Footer'
import CookingSim from './SubPages/CookingSim'

function SubApp(){
  const appContext = useAppContext();
  switch ( appContext.appOption ) {
    case AppOptions.cookingSim:
      return <CookingSim />
    default:
      return <><p className="mx-auto">tbd</p></>
  }
}

function App() {
  const UN = React.useRef<UncleNelly | null>(null);
  React.useEffect(()=>{
    loadUncleNelly()
      .then((initUN)=>{
        UN.current = initUN();
      })
      .catch( e => console.log(e))
  },[])

  return (
    <AppProvider>
      <div className="min-h-[10vh]">
        <Navbar />
      </div>
      <div className="flex min-h-[80vh] border border-red-500">
        <SubApp />
      </div>
      <div className="min-h-[10vh]">
        <Footer />
      </div>
    </AppProvider>
  )
}

export default App
