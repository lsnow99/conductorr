<template>
  <page-wrapper>
    <section>
      <div class="flex flex-col">
        <div class="flex flex-col items-center sm:flex-row justify-between">
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
          :class="(date.today?'date-today':'') + ' ' + calDateClass + ' ' + (date.gray ? 'opacity-80' : '')"
        >
          <div class="bg-gray-800 inline-block py-px w-7 text-center">
            {{ date.day }}
          </div>
          <div class="height-full overflow-y-auto overflow-x-hidden">
            <div
              v-for="(event, index) in eventMap[date.dayStart]"
              :key="index"
              :class="eventClass(event)"
              @click="selectEvent(event, $event)"
              @keydown.enter="selectEvent(event, $event)"
              @keydown.space="selectEvent(event, $event)"
              role="button"
              tabindex="0"
              class="cursor-pointer focus:bg-opacity-30 focus:outline-white flex flex-col lg:block"
            >
              <span
                class="
                  p-1
                  bg-gray-900 bg-opacity-30
                  text-sm
                  font-bold
                  lg:inline-block
                  lg:w-4-5r
                  flex-col
                  flex-1
                  w-full
                  min-h-full
                "
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
      <div class="flex flex-col md:flex-row text-lg mt-2">
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 bg-green-600 mr-1" />
          Available
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 bg-purple-600 mr-1" />
          Unaired/Monitored
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 bg-gray-500 mr-1" />
          Unmonitored
        </div>
        <div class="flex flex-row items-center mx-4">
          <div class="w-6 h-2 bg-red-600 mr-1" />
          Missing/Monitored
        </div>
      </div>
    </section>

    <modal
      v-model="isMediaEventSelected"
      :title="mediaEventTitle"
      @close="restoreFocus"
    >
      <search-actions :mediaID="mediaEvent && mediaEvent.media_id" />
      <p>{{ mediaEvent && mediaEvent.description }}</p>
    </modal>
  </page-wrapper>
</template>

<style scoped>
.date {
  height: 8rem;

  @apply bg-gray-700;
}

.date-today {
  @apply bg-gray-500;
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

<script>
import PageWrapper from "../components/PageWrapper.vue";
import { DateTime } from "luxon";
import APIUtil from "../util/APIUtil";
import SearchActions from "../components/SearchActions.vue";
import RadioGroup from "../components/RadioGroup.vue";
import Modal from "../components/Modal.vue";
import TabSaver from "../util/TabSaver";

const now = DateTime.now();

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
    SearchActions,
    RadioGroup,
    Modal,
  },
  mixins: [TabSaver],
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
    selectEvent(event, $event) {
      this.mediaEvent = event;
      this.lastButton = $event.currentTarget;
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
      if (event.path_ok) {
        return `bg-green-600`
      } else if (!event.monitoring) {
        return `bg-gray-500`
      } else if (event.timestamp > now) {
        return `bg-purple-600`
      } else if (event.timestamp <= now) {
        return `bg-red-600`
      }
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
    isTodayClass(date) {
      if (date.dayStart.toISODate() === DateTime.local().toISODate()) {
        return `bg-red-600`
      }
      return ''
    }
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
    mediaEventTitle() {
      if (!this.isMediaEventSelected) {
        return "";
      }
      return (
        this.eventTitle(this.mediaEvent) +
        (this.mediaEvent.content_type == "episode"
          ? " - " + this.mediaEvent.title
          : "")
      );
    },
    isMediaEventSelected: {
      get() {
        if (this.mediaEvent) {
          return true;
        }
        return false;
      },
      set(v) {
        if (!v) {
          this.mediaEvent = null;
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
            today: curDate.toISODate() === DateTime.local().toISODate()
          });
        }
        for (let i = 0; i < daysInMonth; i++) {
          const curDate = monthStart.plus({ days: i });
          dates.push({
            day: curDate.day,
            dayStart: curDate.startOf("day"),
            gray: false,
            today: curDate.toISODate() === DateTime.local().toISODate()
          });
        }
        for (let i = monthEndDoW; i < 6; i++) {
          const curDate = monthEnd.plus({ days: i - monthEndDoW + 1 });
          dates.push({
            day: curDate.day,
            dayStart: curDate.startOf("day"),
            gray: true,
            today: curDate.toISODate() === DateTime.local().toISODate()
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
            today: curDate.toISODate() === DateTime.local().toISODate()
          });
        }
      } else if (this.viewType === "daily") {
        dates.push({
          day: this.selectedDate.day,
          dayStart: this.selectedDate.startOf("day"),
          gray: false,
          today: this.selectedDate.toISODate() === DateTime.local().toISODate()
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
