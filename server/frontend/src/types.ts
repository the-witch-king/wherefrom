/**
 * MAL Types (Go)
 */
export interface MALAnime {
    node: {
        id: string
        title: string
        main_picture: {
            medium: string
            large: string
        }
    }
}

/**
 * Jikan Types
 */
export interface Image {
    image_url: string
    large_image_url?: string
    small_image_url?: string
}

export interface Anime {
    images: {
        jpg: Image
    }
    mal_id: string
    title: string
    url: string
}

export interface Character {
    images: { jpg: Image }
    mal_id: string
    name: string
    url: string
}

export interface Voice {
    anime: Anime
    character: Character
}

export interface Person {
    name: string
    images: { jpg: Image }
    url: string
    voices: Voice[]
}
