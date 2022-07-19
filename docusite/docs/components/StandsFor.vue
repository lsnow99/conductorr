<template>
It stands for {{ standsFor }} (<a @click="orDoesIt" href="javascript:void(0);">or does it?</a>).
</template>

<script setup>
import { ref } from 'vue';

const STANDS_FOR = new Set([
  "Custom Scripting Language",
  "Content Searching Language",
  "Conductorr-Specific Language",
  "Custom Sorting Language",
  "Characteristic Searching Language",
  "Custom Shitty Lisp",
])

const pickRandom = (current) => {
  const toSubtract = new Set([current])
  const remaining = new Set([...STANDS_FOR].filter(x => !toSubtract.has(x)))

  if (remaining.size < 1) {
    return current
  }
  const pick = Array.from(remaining)[Math.floor(Math.random() * remaining.size)]
  return pick
}

const standsFor = ref(pickRandom(""))

const orDoesIt = () => {
    standsFor.value = pickRandom(standsFor.value)
}
</script>