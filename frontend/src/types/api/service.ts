import { DownloadType } from "./download"
import { Indexer } from "./indexer"

export type ServiceIdentifier = string

export type EditServiceMode = "edit" | "new" | ""

export interface Service {
    name: string
    value: ServiceIdentifier
}

export type ConfigurableService = {
    [key: symbol | string]: unknown
}
