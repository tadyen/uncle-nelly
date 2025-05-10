import React from 'react'
import { loadUncleNelly, type UncleNelly } from './unclenelly'
import { AppOptions, useAppContext, AppProvider } from './AppContext'

import './App.css'

import Navbar from './Components/Navbar'
import Footer from './Components/Footer'

import CookingSim from './SubPages/CookingSim'
import RecipeReverse from './SubPages/RecipeReverse'
import RecipeOptimiser from './SubPages/RecipeOptimiser'

function SubApp(){
  const appContext = useAppContext();
  const UN = React.useRef<UncleNelly | null>(null);
  const reloadUN = React.useRef(true);

  const loadUN = React.useCallback(()=>{
    if(! reloadUN.current){
      return;
    }
    loadUncleNelly()
      .then((initUN)=>{
        UN.current = initUN();
        appContext.setUncleNelly(UN.current);
      })
      .catch( e => console.log(e))
    reloadUN.current = false;
  },[])

  React.useEffect(()=>{
    loadUN();
    // Set default app option
    appContext.setAppOption(AppOptions.cookingSim);
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
