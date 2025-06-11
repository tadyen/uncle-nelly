import React from 'react'
import { baseIngredientsIcons, mixIngredientsIcons, effectsColors } from './assets/icons'
import { loadUncleNelly } from './unclenelly'
import { type UncleNelly, UncleNellyTables } from './unclenelly_types'
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
  const UNLoader = React.useRef< () => UncleNelly | null>(null);
  const UNTable = React.useRef<Record<string,any> | null>(null);

  const loadUN = React.useCallback(()=>{
    loadUncleNelly()
      .then((initUN)=>{
        unelly.current = initUN();
        UNLoader.current = initUN;
        if (UNTable.current === null) {
          const table = unelly.current?.get_tables() as UncleNellyTables;
          for (const [table_name,table_content] of Object.entries(table)){
            switch (table_name) {
              case 'base_ingredients':
                Object.keys(table_content).forEach((k)=>{
                  table.base_ingredients[k].IconRelpath = baseIngredientsIcons[k] || '';
                })
                break;
              case 'mix_ingredients':
                Object.keys(table_content).forEach((k)=>{
                  table.mix_ingredients[k].IconRelpath = mixIngredientsIcons[k] || '';
                })
                break;
              case 'effects':
                Object.keys(table_content).forEach((k)=>{
                  table.effects[k].Color = effectsColors[k] || '#000000';
                })
                break;
              default:
                break;
            }
          }
          UNTable.current = table;
        }
      })
  },[])

  React.useEffect(()=>{
    loadUN();
    // Set default app option
    appContext.setUNellyLoader(UNLoader.current);
    appContext.setAppOption(AppOptions.cookingSim);
    appContext.setUncleNelly(unelly.current);
    appContext.setUNtables(UNTable.current);
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
