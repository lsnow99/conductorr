export enum RipType {
  BDRip = "BDRip",
  WEBRip = "WEBRip",
  WEBDL = "WEBDL",
  WEBCap = "WEBCap",
  HDTV = "HDTV",
  DVDRip = "DVDRip",
  HDRip = "HDRip",
  SCR = "SCR",
  CAM = "CAM",
  TELESYNC = "TELESYNC"
}

export const RIP_TYPES = Object.values(RipType)

export enum ResolutionType {
  R352p = "352p",
  R360p = "360p",
  R480i = "480i",
  R480p = "480p",
  R576i = "576i",
  R576p = "576p",
  R720p = "720p",
  R1080p = "1080p",
  R2160p = "2160p"
}

export const RESOLUTION_TYPES = Object.values(ResolutionType)

export enum EncodingType {
  Xvid = "Xvid",
  x264 = "x264",
  x265 = "x265",
  VP9 = "VP9"
}

export const ENCODING_TYPES = Object.values(EncodingType)

export const DOCUMENTATION_URL = "https://conductorr.xyz/"
