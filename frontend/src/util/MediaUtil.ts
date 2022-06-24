import { Media } from "@/types/api/media"

export default () => {
  const mediaYear = (media: Media) => {
    return new Date(media.released_at).getUTCFullYear();
  }
  return {
    mediaYear
  }
};
