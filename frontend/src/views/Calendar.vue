<template>
  <page-wrapper>
    <section>
      <div class="flex flex-col">
        <div class="flex flex-col items-center justify-between sm:flex-row">
          <div class="table-row">
            <div class="table-cell">
              <o-button variant="primary" @click="resetDate">Today</o-button>
            </div>
            <div class="table-cell">
              <a href="#" class="ml-3"
                ><o-button
                  tabindex="-1"
                  variant="primary"
                  class="table-cell"
                  icon-left="rss"
                ></o-button
              ></a>
            </div>
          </div>
          <radio-group
            v-model="viewType"
            name="viewType"
            :options="[
              {
                text: 'Monthly',
                value: 'monthly',
              },
              {
                text: 'Weekly',
                value: 'weekly',
              },
              {
                text: 'Daily',
                value: 'daily',
              },
            ]"
          />
        </div>
        <div class="flex flex-row justify-between mt-4">
          <o-button
            variant="primary"
            @click="goPages(-1)"
            :aria-label="`Previous ${computedRangeType}`"
          >
            <vue-fontawesome icon="chevron-left" />
          </o-button>
          <div class="text-2xl font-semibold text-center">
            {{ niceDateRange }}
          </div>
          <o-button
            variant="primary"
            @click="goPages(1)"
            :aria-label="`Next ${computedRangeType}`"
          >
            <vue-fontawesome icon="chevron-right" />
          </o-button>
        </div>
      </div>
      <div :class="calWrapperClass" class="relative mt-4">
        <div
          v-for="(date, index) in dates"
          :key="index"
          class="flex flex-col overflow-hidden date"
          :class="
            (date.today ? 'date-today' : '') +
            ' ' +
            calDateClass +
            ' ' +
            (date.gray ? 'opacity-80' : '')
          "
        >
          <div class="inline-block py-px text-center bg-gray-800 w-7">
            {{ date.day }}
          </div>
          <div class="overflow-x-hidden overflow-y-auto height-full">
            <div
              v-for="(event, index) in eventMap[date.dayStart.toISOTime()]"
              :key="index"
              :class="eventClass(event)"
              @click="selectEvent(event, $event)"
              @keydown.enter="selectEvent(event, $event)"
              @keydown.space="selectEvent(event, $event)"
              role="button"
              tabindex="0"
              class="flex flex-col cursor-pointer focus:bg-opacity-30 focus:outline-white lg:block"
            >
              <span
                class="flex-col flex-1 w-full min-h-full p-1 text-sm font-bold bg-gray-900 bg-opacity-30 lg:inline-block lg:w-4-5r"
                >{{ event.time.slice(0, -3)
                }}<span class="text-xs">{{ event.time.slice(-3) }}</span></span
              >
              <span class="ml-2 text-sm font-semibold break-words">{{
                eventTitle(event)
              }}</span>
            </div>
          </div>
        </div>
        <o-loading :active="loading" :full-page="false"></o-loading>
      </div>
      <div class="flex flex-col mt-2 text-lg md:flex-row">
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 mr-1 bg-green-600" />
          Available
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 mr-1 bg-purple-600" />
          Unaired/Monitored
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 mr-1 bg-gray-500" />
          Unmonitored
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 mr-1 bg-red-600" />
          Missing/Monitored
        </div>
      </div>
    </section>

    <Modal
      v-model="isMediaEventSelected"
      :title="mediaEventTitle"
      @close="closeMediaEventModal"
    >
      <div class="flex flex-row justify-between min-w-[400px]">
        <div></div>
        <div v-if="mediaEvent">
          <SearchActions :mediaID="mediaEvent.mediaID" />
          <o-button @click="goToMediaEvent">{{ computedGoToBtnText }}</o-button>
        </div>
      </div>
      <p>{{ computedMediaEventDescription }}</p>
    </Modal>
  </page-wrapper>
</template>

<style scoped>
.date {
  @apply h-[8rem];
  @apply bg-gray-700;
}

.date-today {
  @apply border-gray-300;
  @apply border-solid;
  @apply border-2;
}

.monthly-wrapper {
  @apply grid;
  @apply grid-cols-7;
}

.weekly-wrapper {
  @apply grid;
  @apply grid-cols-7;
}

.daily-date {
  height: 32rem;
}

.weekly-date {
  height: 32rem;
}
</style>

<script setup lang="ts">
import PageWrapper from "../components/PageWrapper.vue";
import { DateTime } from "luxon";
import APIUtil from "../util/APIUtil";
import SearchActions from "../components/SearchActions.vue";
import RadioGroup from "../components/RadioGroup.vue";
import Modal from "../components/Modal.vue";
import { computed, onMounted, ref, WritableComputedRef } from "vue";
import { useTabSaver } from "@/util";
import { ContentType } from "@/types/api/media";
import { useRouter } from "vue-router";
import { MediaEvent } from "@/types/api/media";

enum ViewType {
  MONTHLY = "monthly",
  WEEKLY = "weekly",
  DAILY = "daily",
}

interface CalEvent extends Omit<MediaEvent, "timestamp"> {
  timestamp: DateTime;
  time: string;
}

const now = DateTime.now();

const selectedDate = ref<DateTime>(now);
const viewType = ref<ViewType>(ViewType.MONTHLY);
const eventMap = ref<Partial<Record<string, CalEvent[]>>>({});
const mediaEvent = ref<CalEvent | null>(null);
const loading = ref(false);

const { lastButton, restoreFocus } = useTabSaver();

const closeMediaEventModal = () => {
  isMediaEventSelected.value = false;
  restoreFocus();
};

const goPages = (pages: number) => {
  switch (viewType.value) {
    case ViewType.MONTHLY:
      selectedDate.value = selectedDate.value.plus({ months: pages });
      break;
    case ViewType.WEEKLY:
      selectedDate.value = selectedDate.value.plus({ weeks: pages });
    case ViewType.DAILY:
      selectedDate.value = selectedDate.value.plus({ days: pages });
    default:
      console.error(`unimplemented view type ${viewType.value}`);
  }
};

const router = useRouter();

const goToMediaEvent = () => {
  const id =
    mediaEvent.value?.contentType === ContentType.EPISODE
      ? mediaEvent.value.seriesID
      : mediaEvent.value?.mediaID;
  if (id) router.push({ name: "media", params: { media_id: id } });
};

const selectEvent = (event: CalEvent, $event: Event) => {
  mediaEvent.value = event;
  lastButton.value = $event.currentTarget as HTMLElement;
};

const resetDate = () => (selectedDate.value = DateTime.now());

const eventClass = (event: CalEvent) => {
  if (event.pathOk) {
    return `bg-green-600`;
  } else if (!event.monitoring) {
    return `bg-gray-500`;
  } else if (event.timestamp > now) {
    return `bg-purple-600`;
  } else if (event.timestamp <= now) {
    return `bg-red-600`;
  }
};

const eventTitle = (event: CalEvent | null): string => {
  if (event?.contentType === ContentType.MOVIE) {
    return event.title;
  } else if (event?.contentType === ContentType.EPISODE) {
    let numStr = "x" + event.episodeNum;
    if (numStr.length === 2) {
      numStr = "x0" + event.episodeNum;
    }
    return `${event.seriesTitle} ${event.seasonNum}${numStr}`;
  }
  return "";
};

onMounted(async () => {
  const screenWidth = window.innerWidth;
  if (screenWidth < 768) {
    viewType.value = ViewType.DAILY;
  } else if (screenWidth < 1024) {
    viewType.value = ViewType.WEEKLY;
  } else {
    viewType.value = ViewType.MONTHLY;
  }
  loading.value = true;
  try {
    const events = await APIUtil.getSchedule();
    for (let i = 0; i < events.length; i++) {
      const calEvent : CalEvent = {
        ...events[i],
        timestamp: DateTime.fromJSDate(new Date(events[i].timestamp)),
        time: events[i].timestamp.toLocaleString(DateTime.TIME_SIMPLE),
      };
      const dayStart = calEvent.timestamp.startOf("day").toISOTime();
      eventMap.value[dayStart] = eventMap.value[dayStart]
        ? [...eventMap.value[dayStart]!, calEvent]
        : [calEvent];
    }
  } finally {
    loading.value = false;
  }
});

const isMediaEventSelected: WritableComputedRef<boolean> = computed({
  get: () => !!mediaEvent.value,
  set(v: boolean) {
    if (!v) {
      mediaEvent.value = null;
    }
  },
});

const mediaEventTitle = computed(() => {
  if (!mediaEvent.value) {
    return "";
  }
  return `${eventTitle(mediaEvent.value)}${
    mediaEvent.value.contentType === ContentType.EPISODE && ` - ${mediaEvent.value.title}`
  }`;
});

const dates = computed(() => {
  let dates = [];
  if (viewType.value === "monthly") {
    const monthStart = selectedDate.value.startOf("month");
    const monthStartDoW = monthStart.weekday;
    const monthEnd = selectedDate.value.endOf("month");
    const monthEndDoW = monthEnd.weekday;
    const daysInMonth = monthStart.daysInMonth;

    for (let i = 0; i < monthStartDoW; i++) {
      const curDate = monthStart.minus({ days: i + 1 });
      dates.unshift({
        day: curDate.day,
        dayStart: curDate.startOf("day"),
        gray: true,
        today: curDate.toISODate() === DateTime.local().toISODate(),
      });
    }
    for (let i = 0; i < daysInMonth; i++) {
      const curDate = monthStart.plus({ days: i });
      dates.push({
        day: curDate.day,
        dayStart: curDate.startOf("day"),
        gray: false,
        today: curDate.toISODate() === DateTime.local().toISODate(),
      });
    }
    for (let i = monthEndDoW; i < 6; i++) {
      const curDate = monthEnd.plus({ days: i - monthEndDoW + 1 });
      dates.push({
        day: curDate.day,
        dayStart: curDate.startOf("day"),
        gray: true,
        today: curDate.toISODate() === DateTime.local().toISODate(),
      });
    }
  } else if (viewType.value === "weekly") {
    const weekStart = selectedDate.value.startOf("week");

    for (let i = 0; i < 7; i++) {
      const curDate = weekStart.plus({ days: i });
      dates.push({
        day: curDate.day,
        dayStart: curDate.startOf("day"),
        gray: false,
        today: curDate.toISODate() === DateTime.local().toISODate(),
      });
    }
  } else if (viewType.value === "daily") {
    dates.push({
      day: selectedDate.value.day,
      dayStart: selectedDate.value.startOf("day"),
      gray: false,
      today: selectedDate.value.toISODate() === DateTime.local().toISODate(),
    });
  }
  return dates;
});

const calWrapperClass = computed(() => `${viewType.value}-wrapper`);
const calDateClass = computed(() => `${viewType.value}-date`);

const niceDateRange = computed(() => {
  switch (viewType.value) {
    case ViewType.MONTHLY:
      return selectedDate.value.toLocaleString({
        month: "long",
        year: "numeric",
      });
    case ViewType.WEEKLY:
      const weekStart = selectedDate.value.startOf("week");
      const weekEnd = selectedDate.value.endOf("week");
      return `${weekStart.toLocaleString(
        DateTime.DATE_FULL
      )} to ${weekEnd.toLocaleString(DateTime.DATE_FULL)}`;
    case ViewType.DAILY:
      return selectedDate.value.toLocaleString(DateTime.DATE_FULL);
    default:
      return "";
  }
});

const computedRangeType = computed(() => {
  switch (viewType.value) {
    case ViewType.MONTHLY:
      return "month";
    case ViewType.WEEKLY:
      return "week";
    case ViewType.DAILY:
      return "day";
    default:
      return "range";
  }
});

const computedMediaEventDescription = computed(() => {
  if (!mediaEvent.value) return "";
  return mediaEvent.value.description ?? "No description available";
});

const computedGoToBtnText = computed(() => {
  if (!mediaEvent.value) {
    return "";
  }
  if (mediaEvent.value.contentType === ContentType.EPISODE) {
    return "Go to Series";
  } else if (mediaEvent.value.contentType === ContentType.MOVIE) {
    return "Go to Movie";
  } else {
    console.error("unrecognized content type");
  }
});
</script>
