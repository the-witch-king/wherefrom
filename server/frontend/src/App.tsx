import './App.css'
import { useState } from 'react'
import { useJikan } from './hooks'
import { Person } from './mal-types'

function App() {
    // const [data, setData] = useState<string[]>([])
    const [actorId, setActorId] = useState<string>('')
    const [person, setPerson] = useState<Person>()

    // const getVoiceActor = () => {
    //     fetch('http://localhost:3333', {
    //         method: 'POST',
    //         headers: {
    //             'Content-Type': 'application/json',
    //         },
    //     })
    //         .then((r) => r.json())
    //         .then((response: { items: string[] }) => setData(response['items']))
    // }
  //
    const { getPerson } = useJikan()
    const doIt = async (actorId: string) => {
        const person = await getPerson(actorId)
        setPerson(person)
        console.log('Got: ', person.name)
    }

    return (
        <div className="App">
            <header className="App-header"></header>
            <h2>{person?.name}</h2>
            <input onChange={(e) => setActorId(e.target.value)} />
            <button onClick={() => doIt(actorId)}>HEY, CLICK ME</button>
        </div>
    )
}

export default App
