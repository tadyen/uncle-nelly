import React from 'react'
import { loadUncleNelly, type UncleNelly } from './unclenelly'
import { AppOptions, useAppContext, AppProvider } from './AppContext'

import './App.css'

import Navbar from './Components/Navbar'
import Footer from './Components/Footer'

import CookingSim from './SubPages/CookingSim'
import RecipeReverse from './SubPages/RecipeReverse'
import RecipeOptimiser from './SubPages/RecipeOptimiser'
import CopilotCooker from './SubPages/CopilotCooker'

function SubApp(){
  const appContext = useAppContext();
  const unelly = React.useRef<UncleNelly | null>(null);
  const unLoader = React.useRef< () => UncleNelly | null>(null);

  const loadUN = React.useCallback(()=>{
    loadUncleNelly()
      .then((initUN)=>{
        unLoader.current = initUN;
        unelly.current = initUN();
        appContext.setUncleNelly(unelly.current);
      })
  },[])

  React.useEffect(()=>{
    loadUN();
    // Set default app option
    appContext.setAppOption(AppOptions.copilotCooker);
  },[])

  return <></>

}

function AppSelector(){
  const appContext = useAppContext();
  switch ( appContext.appOption ) {
    case AppOptions.cookingSim:
      return <CookingSim />
    case AppOptions.recipeReverse:
      return <RecipeReverse />
    case AppOptions.recipeOptimiser:
      return <RecipeOptimiser />
    case AppOptions.copilotCooker:
      return <CopilotCooker />
    default:
      return <><p className="mx-auto">tbd</p></>
  }
}

function App() {
  return (
    <AppProvider>
      <SubApp />
      <div className="min-h-[10vh]">
        <Navbar />
      </div>
      <div className="flex min-h-[80vh] border border-red-500">
        <AppSelector />
      </div>
      <div className="min-h-[10vh]">
        <Footer />
      </div>
    </AppProvider>
  )
}

export default App
