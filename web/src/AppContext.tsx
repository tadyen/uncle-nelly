import React from 'react'

// Options
export const AppOptions = {
  cookingSim: "cookingSim",
  recipeReverse: "recipeReverse",
  recipeOptimiser: "recipeOptimiser",
}

// App state Provider
interface AppContextInterface {
  appOption: string|null; setAppOption: React.Dispatch<React.SetStateAction<string|null>>;
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
  const [appOption, setAppOption] = React.useState<string|null>(null)
  return (
    <AppContext.Provider value={{
      appOption, setAppOption,
    }}>
      {children}
    </AppContext.Provider>
  )
}

export default AppContext
