import { useState } from 'react'

const App = () => {
    const [count, setCount] = useState(0)

    const handleCounterPress = () => {
        setCount((count) => count + 1)
    }

    return (
        <div>
            <h1 className="text-4xl font-bold">Wow very cool</h1>
            <button 
                className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
                onClick={handleCounterPress}>Count is: {count}</button>
        </div>
    )
}

export default App
