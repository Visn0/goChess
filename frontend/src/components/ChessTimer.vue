<script setup lang="ts">
import type { Timer } from '@/models/timer'
import { watch } from 'vue'

const props = defineProps<{
    timer: Timer
}>()

const stylePaused = 'bg-dark text-white'
const styleRunning = 'bg-light text-dark'
let timerStyle = props.timer.isPaused() ? stylePaused : styleRunning

watch(props.timer.isPaused.bind(props.timer), () => {
    timerStyle = props.timer.isPaused() ? stylePaused : styleRunning
})
</script>

<template>
    <p class="my-0 text-size badge" :class="timerStyle">
        {{ timer.toString() }}
    </p>
</template>

<style scoped>
.text-size {
    font-size: 0.8em;
}

@media (max-width: 576px) {
    .text-size {
        font-size: 0.55em;
    }
}

@media (min-width: 576px) and (max-width: 768px) {
    .text-size {
        font-size: 0.6em;
    }
}

@media (min-width: 768px) and (max-width: 992px) {
    .text-size {
        font-size: 0.7em;
    }
}
</style>
