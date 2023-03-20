import { Media } from "@/types/api/media"

export default () => {
  const mediaYear = (media: Media) => {
    return new Date(media.releasedAt).getUTCFullYear();
  }
  return {
    mediaYear
  }
};
