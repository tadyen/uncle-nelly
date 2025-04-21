import './App.css'

import Divider, {DividerProps} from './Components/Divider'
import Navbar from './Components/Navbar'
import Footer from './Components/Footer'

const dividerProps: DividerProps = {
  start: <Navbar />,
  end: <Footer />,
}
function App() {
  
  return (
    <>
      <Navbar />
      <Footer />
    </>
  )
}

export default App
