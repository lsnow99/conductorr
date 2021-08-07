<template>
  <page-wrapper>
    <section>
      <div class="flex flex-col">
        <div class="flex flex-col items-center sm:flex-row justify-between">
          <div>
            <o-button variant="primary" @click="resetDate">Today</o-button>
          </div>
          <div class="overflow-hidden rounded-lg inline-block mt-4 sm:mt-0">
            <o-radio v-model="viewType" name="viewType" native-value="monthly"
              >Monthly</o-radio
            >
            <o-radio v-model="viewType" name="viewType" native-value="weekly"
              >Weekly</o-radio
            >
            <o-radio v-model="viewType" name="viewType" native-value="daily"
              >Daily</o-radio
            >
          </div>
        </div>
        <div class="flex flex-row justify-between mt-4">
          <o-button
            variant="primary"
            @click="goPrev"
            :aria-label="`Previous ${computedRangeType}`"
          >
            <vue-fontawesome icon="chevron-left" />
          </o-button>
          <div class="text-2xl font-semibold text-center">
            {{ niceDateRange }}
          </div>
          <o-button
            variant="primary"
            @click="goNext"
            :aria-label="`Next ${computedRangeType}`"
          >
            <vue-fontawesome icon="chevron-right" />
          </o-button>
        </div>
      </div>
      <div :class="calWrapperClass" class="mt-4">
        <div
          v-for="date in dates"
          :key="date"
          class="date overflow-hidden flex flex-col"
          :class="calDateClass + ' ' + (date.gray ? 'opacity-80' : '')"
        >
          <div class="bg-gray-600 inline-block py-px w-7 text-center">
            {{ date.day }}
          </div>
          <div class="height-full overflow-y-auto">
            <div
              v-for="(event, index) in eventMap[date.dayStart]"
              :key="index"
              :class="eventClass(event)"
              @click="selectedMediaEvent = event"
              @keydown.enter="selectedMediaEvent = event"
              @keydown.space="selectedMediaEvent = event"
              aria-role="button"
              tabindex="0"
              class="cursor-pointer focus:bg-opacity-30 focus:outline-white"
            >
              <span
                class="
                  p-1
                  bg-gray-900 bg-opacity-30
                  text-sm
                  font-bold
                  inline-block
                  min-h-full
                "
                style="width: 70px"
                >{{ event.time.slice(0, -3)
                }}<span class="text-xs">{{ event.time.slice(-3) }}</span></span
              >
              <span class="text-sm font-semibold ml-2 break-words">{{
                eventTitle(event)
              }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>

    <o-modal v-model:active="mediaEvent">
      <header v-if="mediaEvent" class="modal-card-header">
        <p class="modal-card-title">
          {{
            eventTitle(mediaEvent) +
            (mediaEvent.content_type == "episode"
              ? " - " + mediaEvent.title
              : "")
          }}
        </p>
      </header>
      <section v-if="mediaEvent" class="modal-card-content">
        {{ mediaEvent.description }}
      </section>
      <footer v-if="mediaEvent" class="modal-card-footer"></footer>
    </o-modal>
  </page-wrapper>
</template>

<style scoped>
.date {
  height: 12rem;

  @apply bg-gray-700;
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

<script>
import PageWrapper from "../components/PageWrapper.vue";
import { DateTime } from "luxon";
import APIUtil from "../util/APIUtil";

const EVENT_BACKGROUNDS = [
  "bg-red-500",
  "bg-yellow-500",
  "bg-green-500",
  "bg-blue-500",
  "bg-purple-500",
];

export default {
  data() {
    return {
      selectedDate: DateTime.now(),
      viewType: "monthly",
      events: [],
      eventMap: {},
      mediaEvent: null,
    };
  },
  components: {
    PageWrapper,
  },
  methods: {
    goPrev() {
      if (this.viewType === "monthly") {
        this.selectedDate = this.selectedDate.minus({ months: 1 });
      } else if (this.viewType === "weekly") {
        this.selectedDate = this.selectedDate.minus({ weeks: 1 });
      } else if (this.viewType === "daily") {
        this.selectedDate = this.selectedDate.minus({ days: 1 });
      }
    },
    goNext() {
      if (this.viewType === "monthly") {
        this.selectedDate = this.selectedDate.plus({ months: 1 });
      } else if (this.viewType === "weekly") {
        this.selectedDate = this.selectedDate.plus({ weeks: 1 });
      } else if (this.viewType === "daily") {
        this.selectedDate = this.selectedDate.plus({ days: 1 });
      }
    },
    resetDate() {
      this.selectedDate = DateTime.now();
    },
    eventClass(event) {
      return `bg-red-600`;
    },
    eventTitle(event) {
      if (event.content_type == "movie") {
        return event.title;
      } else if (event.content_type == "episode") {
        let numStr = "x" + event.episode_num;
        if (numStr.length == 2) {
          numStr = "x0" + event.episode_num;
        }
        return event.series_title + " " + event.season_num + numStr;
      }
    },
  },
  mounted() {
    const screenWidth = window.innerWidth;
    if (screenWidth < 768) {
      this.viewType = "daily";
    } else if (screenWidth < 1024) {
      this.viewType = "weekly";
    } else {
      this.viewType = "monthly";
    }
    APIUtil.getSchedule().then((events) => {
      for (let i = 0; i < events.length; i++) {
        events[i].timestamp = DateTime.fromJSDate(
          new Date(events[i].timestamp)
        );
        events[i].time = events[i].timestamp.toLocaleString(
          DateTime.TIME_SIMPLE
        );
        const dayStart = events[i].timestamp.startOf("day");
        if (!this.eventMap[dayStart]) {
          this.eventMap[dayStart] = [events[i]];
        } else {
          this.eventMap[dayStart].push(events[i]);
        }
      }
      this.events = events;
    });
  },
  computed: {
    selectedMediaEvent: {
      get() {
        return this.mediaEvent;
      },
      set(newVal) {
        if (!newVal) {
          this.mediaEvent = null;
        } else {
          this.mediaEvent = newVal;
        }
      },
    },
    dates() {
      let dates = [];
      if (this.viewType === "monthly") {
        const monthStart = this.selectedDate.startOf("month");
        const monthStartDoW = monthStart.weekday;
        const monthEnd = this.selectedDate.endOf("month");
        const monthEndDoW = monthEnd.weekday;
        const daysInMonth = monthStart.daysInMonth;

        for (let i = 0; i < monthStartDoW; i++) {
          const curDate = monthStart.minus({ days: i + 1 });
          dates.unshift({
            day: curDate.day,
            dayStart: curDate.startOf("day"),
            gray: true,
          });
        }
        for (let i = 0; i < daysInMonth; i++) {
          const curDate = monthStart.plus({ days: i });
          dates.push({
            day: curDate.day,
            dayStart: curDate.startOf("day"),
            gray: false,
          });
        }
        for (let i = monthEndDoW; i < 6; i++) {
          const curDate = monthEnd.plus({ days: i - monthEndDoW + 1 });
          dates.push({
            day: curDate.day,
            dayStart: curDate.startOf("day"),
            gray: true,
          });
        }
      } else if (this.viewType === "weekly") {
        const weekStart = this.selectedDate.startOf("week");

        for (let i = 0; i < 7; i++) {
          const curDate = weekStart.plus({ days: i });
          dates.push({
            day: curDate.day,
            dayStart: curDate.startOf("day"),
            gray: false,
          });
        }
      } else if (this.viewType === "daily") {
        dates.push({
          day: this.selectedDate.day,
          dayStart: this.selectedDate.startOf("day"),
          gray: false,
        });
      }
      return dates;
    },
    calWrapperClass() {
      return this.viewType + "-wrapper";
    },
    calDateClass() {
      return this.viewType + "-date";
    },
    niceDateRange() {
      if (this.viewType === "monthly") {
        return this.selectedDate.toLocaleString({
          month: "long",
          year: "numeric",
        });
      } else if (this.viewType === "weekly") {
        const weekStart = this.selectedDate.startOf("week");
        const weekEnd = this.selectedDate.endOf("week");

        return (
          weekStart.toLocaleString(DateTime.DATE_FULL) +
          " to " +
          weekEnd.toLocaleString(DateTime.DATE_FULL)
        );
      } else if (this.viewType === "daily") {
        return this.selectedDate.toLocaleString(DateTime.DATE_FULL);
      }
    },
    computedRangeType() {
      switch (this.viewType) {
        case "monthly":
          return "monnth";
        case "weekly":
          return "week";
        case "daily":
          return "day";
      }
      return "";
    },
  },
};
</script>
