import { useState } from 'react'
import { AppOptions, useAppContext } from '../AppContext'

import '../App.css'
import nelly_image from '../assets/uncle_nelly.png'

export default function Navbar(){
  const [activeTab, setActiveTab] = useState(0);
  const appContext = useAppContext();

  const navTabs = {
    "0": {
      name: "Cooking Simulator",
      onClick: ()=>{setActiveTab(0), appContext.setAppOption(AppOptions.cookingSim)},
    },
    "1": {
      name: "Recipe Reverse (tbd)",
      onClick: ()=>{setActiveTab(1), appContext.setAppOption(AppOptions.recipeReverse)},
    },
    "2": {
      name: "Optimiser (tbd)",
      onClick: ()=>{setActiveTab(2), appContext.setAppOption(AppOptions.recipeOptimiser)},
    }
  }

  return (<>
  <div className="navbar bg-base-100 shadow-sm">
    <div className="navbar-start">
      <img src={nelly_image} className="max-h-24"/>
      <a className="btn btn-ghost text-xl">Uncle Nelly</a>
    </div>
    <div className="navbar-center hidden lg:flex">

      <div role="tablist" className="tabs tabs-border">
        { Object.entries(navTabs).map((entry)=> {
            const index = entry[0];
            const tab = entry[1];
            const tabActive = (num: Number): string=>{
              const isActiveTab = (num === Number(index));
              return (isActiveTab ? " tab-active" : " ");
            }
            return (
              <div role="tab"
                className = {"tab" + tabActive(activeTab)}
                onClick={tab.onClick}
              >{tab.name}</div>
            )
          })
        }
      </div>
    </div>
    <div className="navbar-end">
    </div>
  </div>
  </>)
}
