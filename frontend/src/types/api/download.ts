import { DateTime } from "luxon"

export enum DownloadType {
    NZB = "nzb",
    TORRENT = "torrent"
}

export interface Download {
    id: number
    mediaId: number
    friendlyName: string
    identifier: string
    finalDir: string
    status: string
    started: DateTime
    bytesLeft: number
    fullSize: number
}
