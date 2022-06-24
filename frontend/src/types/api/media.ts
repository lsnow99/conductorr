export interface Media {
    id: number
    title: string
    description: string
    poster: string
    released_at: string
    content_type: 'movie' | 'series' | 'season' | 'episode'
    monitoring: boolean
    path_ok: boolean
    path: string
    imdb_rating: number
    imdb_id: string
    children: Media[]
    path_id: number
    profile_id: number
}
