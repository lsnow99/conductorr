import { DownloadType } from "./download";

export interface Indexer {
    id: number,
    name: string,
    baseUrl: string,
    apiKey: string,
    userID: number,
    forMovies: boolean,
    forSeries: boolean,
    downloadType: DownloadType
}
