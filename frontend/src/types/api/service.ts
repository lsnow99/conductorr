import { DownloadType } from "./download"

export type ServiceIdentifier = string

export type EditServiceMode = "edit" | "new" | ""

export interface Service {
    name: string
    value: ServiceIdentifier
}

export interface ConfigurableService {
    name: string
    config: any
}

export interface ConfigurableIndexer extends ConfigurableService {
    config: {
        userId: number
        baseUrl: string
        forMovies: boolean
        forSeries: boolean
        downloadType: DownloadType
    }
}
