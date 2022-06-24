export type ContentType = 'movie' | 'series' | 'season' | 'episode'

export interface Media {
    id: number
    title: string
    description: string
    poster: string
    released_at: string
    content_type: ContentType
    monitoring: boolean
    path_ok: boolean
    path: string
    imdb_rating: number
    imdb_id: string
    children: Media[]
    path_id: number
    profile_id: number
}

export interface MediaSearchResult extends Media {
    search_id: string
}
