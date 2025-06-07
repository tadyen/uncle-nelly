import React from 'react'
import { type UncleNelly } from './unclenelly'

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
  appOption: TOrNull<string>; setAppOption: ReactSetState<TOrNull<string>>;
  uncleNelly: TOrNull<UncleNelly>; setUncleNelly: ReactSetState<TOrNull<UncleNelly>>;

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
  const [appOption, setAppOption] = React.useState<string|null>(null);
  const [uncleNelly, setUncleNelly] = React.useState<UncleNelly|null>(null);

  return (
    <AppContext.Provider value={{
      appOption, setAppOption,
      uncleNelly, setUncleNelly
    }}>
      {children}
    </AppContext.Provider>
  )
}

export default AppContext
