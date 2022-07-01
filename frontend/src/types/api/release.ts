export class Release {
    size?: number
    title: string
    sizeUnit?: string
    runtime?: number
    age?: number
    indexer?: string
    download_type?: 'nzb' | 'torrent'
    seeders?: number
    encoding: string
    rip_type: string
    content_type: string
    resolution: string

    constructor(title, encoding, rip_type, content_type, resolution) {
        this.title = title
        this.encoding = encoding
        this.rip_type = rip_type
    }
}