import React from 'react'
import { baseIngredientsIcons, mixIngredientsIcons, effectsColors } from './assets/icons'
import { loadUncleNelly } from './unclenelly'
import { type UncleNelly, UncleNellyTables } from './unclenelly_types'
import { AppOptions, useAppContext, AppProvider } from './AppContext'
import until from './helpers/until'

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
  const UNTable = React.useRef<Record<string,any> | null>(null);

  const loadUN = async () => {
    const initUN = await loadUncleNelly();
    unelly.current = initUN();
    if (UNTable.current === null) {
      const res = unelly.current?.get_tables();
      if (res?.error || res?.response === null || res?.response === undefined) {
        console.error('Error loading Uncle Nelly tables:', res?.error);
        return;
      }
      const table: UncleNellyTables = res.response as UncleNellyTables;
      for (const [table_name,table_content] of Object.entries(table)){
        switch (table_name) {
          case 'base_ingredients':
            Object.keys(table_content).forEach((k)=>{
              table.base_ingredients[k].Icon = baseIngredientsIcons[k];
            })
            break;
          case 'mix_ingredients':
            Object.keys(table_content).forEach((k)=>{
              table.mix_ingredients[k].Icon = mixIngredientsIcons[k];
            })
            break;
          case 'effects':
            Object.keys(table_content).forEach((k)=>{
              table.effects[k].Color = effectsColors[k];
            })
            break;
          default:
            break;
        }
      }
      UNTable.current = table;
    }
    return true;
  }

  React.useEffect(()=>{
    (async () => {
      await loadUN();
      appContext.setUncleNelly(unelly.current);
      appContext.setUNtables(UNTable.current);
      appContext.setAppOption(AppOptions.cookingSim); // default
    })();
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
