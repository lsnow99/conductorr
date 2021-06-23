export default {
  methods: {
    mediaYear(media) {
      return new Date(media.released_at).getUTCFullYear();
    },
  },
};
