import React from 'react'
import { type UncleNelly } from './unclenelly_types'

// Options
export const AppOptions = {
  cookingSim: "cookingSim",
  recipeReverse: "recipeReverse",
  recipeOptimiser: "recipeOptimiser",
  copilotCooker: "copilotCooker",
}

type ReactSetState<T> = React.Dispatch<React.SetStateAction<T>>
type TOrNull<T> = T | null

// App state Provider
interface AppContextInterface {
  uncleNelly: TOrNull<UncleNelly>; setUncleNelly: ReactSetState<TOrNull< UncleNelly >>;
  appOption: TOrNull<string>; setAppOption: ReactSetState<TOrNull< string >>;
  UNtables: TOrNull<Record<string,any>>; setUNtables: ReactSetState<TOrNull< Record<string,any> >>;
}

const AppContext = React.createContext<AppContextInterface|undefined>(undefined)

export const useAppContext = () => {
  const context = React.useContext(AppContext);
  if (context === undefined) {
    throw new Error('useAppContext must be used within an AppProvider');
  }
  return context;
}

export function AppProvider({children}: {children: React.ReactNode}){
  const [uncleNelly, setUncleNelly] = React.useState<TOrNull< UncleNelly >>(null);
  const [appOption, setAppOption] = React.useState<TOrNull< string >>(null);
  const [UNtables, setUNtables] = React.useState<TOrNull< Record<string,any> >>(null);

  return (
    <AppContext.Provider value={{
      uncleNelly, setUncleNelly,
      appOption, setAppOption,
      UNtables, setUNtables
    }}>
      {children}
    </AppContext.Provider>
  )
}

export default AppContext
