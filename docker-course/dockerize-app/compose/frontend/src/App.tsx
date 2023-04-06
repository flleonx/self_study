import { useState, useEffect } from 'react'
import axios from 'axios';
import reactLogo from './assets/react.svg'
import './App.css'

function App() {

  const [dummyReq, setDummyReq] = useState('');
  const [count, setCount] = useState(0)

  const getReqData = async () => {
    const resp = await axios.get(`${import.meta.env.VITE_API_URL}`);
    setDummyReq(resp.data);
  };

  useEffect(() => {
    getReqData();
  }, [])

  return (
    <div className="App">
      <div>
        <a href="https://vitejs.dev" target="_blank">
          <img src="/vite.svg" className="logo" alt="Vite logo" />
        </a>
        <a href="https://reactjs.org" target="_blank">
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Vite + React + { dummyReq }</h1>
      <div className="card">
        <button onClick={() => setCount((count) => count + 1)}>
          count is {count}
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </div>
  )
}

export default App
