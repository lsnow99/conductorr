export enum DownloaderType {
  nzbget = "nzbget",
  transmission = "transmission"
}

export enum FileAction {
  move = "move",
  copy = "copy"
}

export type NZBGetConfig = {
  baseUrl: string
  username: string
  password: string
}
export type TransmissionConfig = NZBGetConfig

export type BaseDownloader = {
  id: number;
  name: string;
  fileAction: FileAction
}

export type Downloader = BaseDownloader & ({ downloaderType: DownloaderType.transmission, config: TransmissionConfig } | { downloaderType: DownloaderType.nzbget, config: NZBGetConfig })

export type FlatDownloader = BaseDownloader & (NZBGetConfig | TransmissionConfig)
