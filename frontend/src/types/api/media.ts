export enum ContentType {
  MOVIE = "movie",
  SERIES = "series",
  SEASON = "season",
  EPISODE = "episode",
}

export interface Media {
  id: number;
  title: string;
  description: string;
  poster: string;
  released_at: string;
  contentType: ContentType;
  monitoring: boolean;
  pathOk: boolean;
  path: string;
  imdbRating: number;
  imdbId: string;
  children: Media[];
  pathId: number;
  profileId: number;
}

export interface MediaSearchResult extends Media {
  searchId: string;
}

export interface MediaEvent
  extends Omit<
    Media,
    | "imdbRating"
    | "children"
    | "pathId"
    | "poster"
    | "id"
    | "released_at"
    | "path"
    | "imdbId"
    | "profileId"
  > {
  timestamp: Date;
  seriesTitle?: string;
  seriesId?: string;
  seasonNum?: number;
  episodeNum?: number;
  mediaId: number;
}
