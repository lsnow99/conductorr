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
          <o-button variant="primary" @click="goPrev">
            <vue-fontawesome icon="chevron-left" />
          </o-button>
          <div class="text-2xl font-semibold text-center">
            {{ niceDateRange }}
          </div>
          <o-button variant="primary" @click="goNext">
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
          <div class="bg-gray-600 inline-block py-1 w-8 text-center">
            {{ date.day }}
          </div>
          <div class="overflow-y-auto height-full">
            <div
              v-for="(event, index) in date.events"
              :key="index"
              :class="eventClass(index)"
              class="text-md p-1"
            >
              <span class="text-sm font-bold float-left">{{ event.time }}</span>
              <span class="font-semibold ml-2">{{ event.title }}</span>
            </div>
          </div>
        </div>
      </div>
    </section>
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
import APIUtil from '../util/APIUtil';

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
      events: []
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
    eventClass(index) {
      return EVENT_BACKGROUNDS[index % EVENT_BACKGROUNDS.length];
    },
    getEventsForDatetime(datetime) {
      let evts = []
      for(let i = 0; i < this.events.length; i++) {
        const evt = this.events[i]
        const evtDatetime = DateTime.fromJSDate(new Date(evt.timestamp))
        if (evtDatetime.hasSame(datetime, 'day')) {
          evts.push({
            time: evtDatetime.toLocaleString(DateTime.TIME_SIMPLE),
            title: evt.title
          })
        }
      }
      return evts
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
    APIUtil.getSchedule().then(events => {
      this.events = events
    })
  },
  computed: {
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
            gray: true,
            events: [],
          });
        }
        for (let i = 0; i < daysInMonth; i++) {
          const curDate = monthStart.plus({ days: i });
          dates.push({
            day: curDate.day,
            gray: false,
            events: this.getEventsForDatetime(curDate),
          });
        }
        for (let i = monthEndDoW; i < 6; i++) {
          const curDate = monthEnd.plus({ days: i - monthEndDoW + 1 });
          dates.push({
            day: curDate.day,
            gray: true,
            events: [],
          });
        }
      } else if (this.viewType === "weekly") {
        const weekStart = this.selectedDate.startOf("week");

        for (let i = 0; i < 7; i++) {
          const curDate = weekStart.plus({ days: i });
          dates.push({
            day: curDate.day,
            gray: false,
            events: [
              {
                time: "5:00pm",
                title: "Westworld",
              },
              {
                time: "7:00pm",
                title: "Family Guy",
              },
              {
                time: "9:00pm",
                title: "Bloodline",
              },
              {
                time: "10:00pm",
                title: "Invincible",
              },
              {
                time: "10:00pm",
                title: "True Detective",
              },
              {
                time: "11:00pm",
                title: "Watchmen",
              },
            ],
          });
        }
      } else if (this.viewType === "daily") {
        dates.push({
          day: this.selectedDate.day,
          gray: false,
          events: [],
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
  },
};
</script>
