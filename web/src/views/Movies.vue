<template>
    <page-wrapper>
    <section>
      <div>
        <o-field label="Search Movies">
          <o-input type="text" placeholder="Search for a movie"/>
        </o-field>
      </div>
      <div
        class="
          mt-10
          gap-y-10
          grid
          justify-items-center
          grid-flow-row grid-cols-1
          sm:grid-cols-2
          lg:grid-cols-3
          xl:grid-cols-4
          2xl:grid-cols-5
        "
      >
        <media-card v-for="media in medias" :key="media.title" :media="media" />
      </div>
    </section>
    </page-wrapper>
</template>

<script>
import MediaCard from "../components/MediaCard.vue";
import PageWrapper from "../components/PageWrapper.vue";

export default {
  data() {
    return {
      medias: [
        {
          title: "Godzilla vs Kong",
          poster:
            "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/pgqgaUx1cJb5oZQQ5v0tNARCeBp.jpg",
        },
        {
          title: "Avengers Age of Ultron",
          poster:
            "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4ssDuvEDkSArWEdyBl2X5EHvYKU.jpg",
        },
        {
          title: "Avengers Age of Ultron",
          poster:
            "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4ssDuvEDkSArWEdyBl2X5EHvYKU.jpg",
        },
        {
          title: "Avengers Age of Ultron",
          poster:
            "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4ssDuvEDkSArWEdyBl2X5EHvYKU.jpg",
        },
        {
          title: "Avengers Age of Ultron",
          poster:
            "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4ssDuvEDkSArWEdyBl2X5EHvYKU.jpg",
        },
        {
          title: "Avengers Age of Ultron",
          poster:
            "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4ssDuvEDkSArWEdyBl2X5EHvYKU.jpg",
        },
        {
          title: "Avengers Age of Ultron",
          poster:
            "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4ssDuvEDkSArWEdyBl2X5EHvYKU.jpg",
        },
        {
          title: "Avengers Age of Ultron",
          poster:
            "https://www.themoviedb.org/t/p/w600_and_h900_bestv2/4ssDuvEDkSArWEdyBl2X5EHvYKU.jpg",
        },
      ],
    };
  },
  components: { PageWrapper, MediaCard },
};
</script>
