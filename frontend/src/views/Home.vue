<template>
  <PageWrapper>
    <section class="h-[29rem]">
      <h1 class="text-2xl">Recently Added to Library</h1>
      <div class="flex flex-wrap justify-between p-2">
        <o-carousel ref="carousel" autoplay :interval="10000" repeat>
          <template v-slot:arrow="props">
            <div
              v-if="props.hasPrev"
              @click="carousel?.prev"
              class="absolute w-10 h-10 -translate-y-1/2 bg-gray-600 rounded-full shadow-md cursor-pointer top-1/2 left-2"
            >
              <div class="relative w-full h-full">
                <o-icon
                  class="absolute text-2xl text-yellow-500 -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
                  icon="chevron-left"
                ></o-icon>
              </div>
            </div>
            <div
              v-if="props.hasNext"
              @click="carousel?.next"
              class="absolute w-10 h-10 -translate-y-1/2 bg-gray-600 rounded-full shadow-md cursor-pointer top-1/2 right-2"
            >
              <div class="relative w-full h-full">
                <o-icon
                  class="absolute text-2xl text-yellow-500 -translate-x-1/2 -translate-y-1/2 top-1/2 left-1/2"
                  icon="chevron-right"
                ></o-icon>
              </div>
            </div>
          </template>
          <o-carousel-item v-for="media in recentMedia" :key="media.id">
            <div
              class="flex flex-row justify-center py-4 md:justify-start md:px-16"
            >
              <div class="flex flex-row justify-center flex-1">
                <MediaCard :media="media" @click="gotoMedia"></MediaCard>
              </div>
              <div class="flex-col p-4 md:flex max-h-[22rem] overflow-hidden relative">
                <span class="text-3xl">{{ media.title }}</span>
                <span class="mt-8 text-lg">{{ media.description }}</span>
                <div class="absolute bottom-0 h-10 left-0 right-0 bg-darkgray from-current to-transparent"></div>
              </div>
            </div>
          </o-carousel-item>
        </o-carousel>
      </div>
    </section>
    <section class="flex-1">
      <DownloadStatusViewer wrapperClass="h-96" />
    </section>
  </PageWrapper>
</template>

<script setup lang="ts">
import PageWrapper from "../components/PageWrapper.vue";
import DownloadStatusViewer from "../components/DownloadStatusViewer.vue";
import APIUtil from "../util/APIUtil";
import MediaCard from "../components/MediaCard.vue";
import { onMounted, ref } from "vue";
import { Media } from "@/types/api/media";
import { useRouter } from "vue-router";

const recentMedia = ref<Media[]>([]);
const carousel = ref<(Element & { prev: () => void; next: () => void }) | null>(
  null
);

const router = useRouter();
const gotoMedia = (media: Media) => {
  router.push({ name: "media", params: { media_id: media.id } });
};

onMounted(async () => {
  try {
    const media = await APIUtil.getRecentMedia();
    recentMedia.value = media;
  } catch {}
});
</script>
