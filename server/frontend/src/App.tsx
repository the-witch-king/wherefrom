import './App.css'
import { useEffect, useState } from 'react'
import { useJikan, useMAL } from './hooks'
import { Person, Voice } from './types'
import { Results } from './results'

const USERNAME_KEY = 'userName'

function App() {
    const { getPerson } = useJikan()
    const { getUserAnimeList } = useMAL()

    const [saveUserName, setSaveUserName] = useState<boolean>(false)
    const [actorId, setActorId] = useState<string>('')
    const [malUserName, setMalUserName] = useState<string>('')
    const [person, setPerson] = useState<Person>()
    const [seenIn, setSeenIn] = useState<Voice[]>([])
    const [collapsed, setCollapsed] = useState<boolean>(false)

    useEffect(() => {
        const savedName = localStorage.getItem(USERNAME_KEY)

        if (!savedName) return

        setMalUserName(savedName)
        setSaveUserName(true)
    }, [malUserName, setMalUserName])

    const work = async (actorId: string) => {
        saveUserName
            ? localStorage.setItem(USERNAME_KEY, malUserName)
            : localStorage.removeItem(USERNAME_KEY)

        const [voiceActor, userAnimeList] = await Promise.all([
            getPerson(actorId),
            getUserAnimeList(malUserName),
        ])

        setPerson(voiceActor)

        const userSeenIn = voiceActor.voices.filter(
            (voice) => userAnimeList[voice.anime.mal_id]
        )

        setSeenIn(userSeenIn)
        setCollapsed(true)

        console.group('Work done!')
        console.log({ voiceActor, userAnimeList, userSeenIn })
        console.groupEnd()
    }

    return (
        <div className="App">
            <header className="App-header"></header>
            {!collapsed && (
                <form
                    onSubmit={(e) => {
                        e.preventDefault()
                        work(actorId)
                    }}
                    style={{
                        display: 'flex',
                        flexDirection: 'column',
                        alignItems: 'center',
                    }}
                >
                    <div className="input-wrapper">
                        <div
                            style={{
                                display: 'flex',
                                justifyContent: 'space-between',
                            }}
                        >
                            <label
                                style={{ display: 'block' }}
                                htmlFor="malUserName"
                            >
                                MAL User Name:
                            </label>
                            <label
                                style={{
                                    display: 'flex',
                                    alignItems: 'center',
                                }}
                            >
                                Save?
                                <input
                                    style={{
                                        margin: '0 8px',
                                    }}
                                    type="checkbox"
                                    name="save-username"
                                    checked={saveUserName}
                                    onChange={({ target: { checked } }) =>
                                        setSaveUserName(checked)
                                    }
                                />
                            </label>
                        </div>
                        <input
                            id="malUserName"
                            name="malUserName"
                            value={malUserName}
                            placeholder="xXx_shad0wKillerNaruto_xXx"
                            onChange={(e) => setMalUserName(e.target.value)}
                        />
                    </div>

                    <div className="input-wrapper">
                        <label
                            style={{ display: 'block' }}
                            htmlFor="malActorId"
                        >
                            MAL Voice Actor ID:
                        </label>
                        <input
                            id="malActorId"
                            name="malActorId"
                            value={actorId}
                            placeholder="420"
                            onChange={(e) => setActorId(e.target.value)}
                        />
                    </div>

                    <button type="submit">WHERE FROM??</button>
                </form>
            )}

            {person && (
                <div
                    style={{
                        display: 'flex',
                        width: '100vw',
                        padding: 4,
                        alignItems: 'center',
                        justifyContent: 'center',
                        marginTop: collapsed ? 0 : 8,
                        backgroundColor: 'var(--orange)',
                    }}
                    onClick={() => setCollapsed(!collapsed)}
                >
                    <span
                        style={{
                            display: 'block',
                            fontSize: '3rem',
                            height: '1.75rem',
                            color: 'var(--purple)',
                            transform: `rotateX(${collapsed ? '180deg' : '0'})`,
                        }}
                    >
                        ^
                    </span>
                </div>
            )}

            {person && <Results person={person} seenIn={seenIn} />}
        </div>
    )
}

export default App
