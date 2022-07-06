export interface Path {
    id: number
    path: string
    moviesDefault: boolean
    seriesDefault: boolean
}

export type NewPath = Omit<Path, "id">
