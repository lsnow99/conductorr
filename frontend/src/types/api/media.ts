export type ContentType = 'movie' | 'series' | 'season' | 'episode'

export interface Media {
    id: number
    title: string
    description: string
    poster: string
    released_at: string
    contentType: ContentType
    monitoring: boolean
    pathOk: boolean
    path: string
    imdbRating: number
    imdbId: string
    children: Media[]
    pathId: number
    profileId: number
}

export interface MediaSearchResult extends Media {
    searchId: string
}
