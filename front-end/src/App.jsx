import { useState } from 'react'
import './App.css'
import IndoorMap from "./IndoorMap.jsx";

function App() {
  const [count, setCount] = useState(0)

  return (
      <>
          <IndoorMap/>
      </>
  )
}

export default App
