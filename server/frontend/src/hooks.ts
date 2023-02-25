import { Person } from './mal-types'

const jikanBaseUrl = 'https://api.jikan.moe/v4'
const jikanPeopleUrl = `${jikanBaseUrl}/people`

export const useJikan = () => ({
    getPerson: async (actorId: string): Promise<Person> => {
        const person = await (
            await fetch(`${jikanPeopleUrl}/${actorId}/full`)
        ).json()

        return person.data
    },
})

const baseUrl = `http://localhost:3333`

export const useMAL = () => ({
    getUserAnimeList: async (userName: string): Promise<any> => {
        const list = await (
            await fetch(`${baseUrl}`, {
                method: 'POST',
                body: JSON.stringify({ userName }),
            })
        ).json()

        console.log('Got anime list:', list)
    },
})
