export enum ContentType {
  MOVIE = "movie",
  SERIES = "series",
  SEASON = "season",
  EPISODE = "episode",
  ALL = "all"
}

export interface Media {
  id: number;
  title: string;
  description: string;
  poster: string;
  releasedAt: string;
  contentType: ContentType;
  monitoring: boolean;
  pathOk: boolean;
  path: string;
  imdbRating: number;
  imdbID: string;
  children: Media[];
  pathID: number;
  profileID: number;
}

export interface MediaSearchResult extends Media {
  searchID: string;
}

export interface MediaEvent
  extends Omit<
  Media,
  | "imdbRating"
  | "children"
  | "pathID"
  | "poster"
  | "id"
  | "releasedAt"
  | "path"
  | "imdbID"
  | "profileID"
  > {
  timestamp: Date;
  seriesTitle?: string;
  seriesID?: string;
  seasonNum?: number;
  episodeNum?: number;
  mediaID: number;
}
