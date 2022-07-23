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
    | "released_at"
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
