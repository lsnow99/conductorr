export interface Release {
    size?: number
    title: string
    sizeUnit?: string
    runtime?: number
    age?: number
    indexer?: string
    downloadType?: 'nzb' | 'torrent'
    seeders?: number
    encoding: string
    ripType: string
    contentType: string
    resolution: string
}
