import logo from './logo.svg';
import './App.css';
import React, { useEffect, useState } from 'react'
function App() {
  const [comp, setComp] = useState([])

  useEffect(() => {
    const fetchData = async () => {
      const url = "/api/test";

      const response = await fetch(url)
      if (!response.ok) {
        return null
      }

      const json = await response.json()
      setComp(json.data)
    }

    fetchData()

  }, [])

  return (
     <div className="min-h-screen bg-gradient-to-b from-gray-700 to-green-200 flex flex-col items-center p-4">
            <div className="p-6 mt-40">
                <h2 className="text-2xl font-bold text-gray-800 text-center mb-5">Spaghetti List</h2>
                <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-6 gap-4">
                    {comp.map((comp, index) => (
                        <div
                            key={index}
                            className="bg-white mx-1 px-5 text-center py-3 rounded-lg shadow transform transition-transform duration-200 hover:scale-105 cursor-pointer"
                        >
                            {comp.name}
                        </div>
                    ))}
                </div>
            </div>
        </div>
  );
}

export default App;
