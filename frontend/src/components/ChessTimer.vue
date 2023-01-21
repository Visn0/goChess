<script setup lang="ts">
import type { Timer } from '@/models/timer'
import { ref, watch } from 'vue'

const props = defineProps<{
    timer: Timer
}>()

const stylePaused = 'bg-dark text-white'
const styleRunning = 'bg-light text-dark'
let timerStyle = ref('')
timerStyle.value = props.timer.isPaused() ? stylePaused : styleRunning

watch(props.timer.isPaused.bind(props.timer), () => {
    timerStyle.value = props.timer.isPaused() ? stylePaused : styleRunning
})

watch(props.timer.isStoped.bind(props.timer), () => {
    if (props.timer.isStoped()) {
        timerStyle.value = stylePaused
    }
})
</script>

<template>
    <div>
        <p class="badge my-0 p-1 text-size" :class="timerStyle">
            {{ timer.toString() }}
        </p>
    </div>
</template>

<style scoped>
.text-size {
    font-size: 0.7em;
}

@media (max-width: 576px) {
    .text-size {
        font-size: 0.45em;
    }
}

@media (min-width: 576px) and (max-width: 768px) {
    .text-size {
        font-size: 0.55em;
    }
}

@media (min-width: 768px) and (max-width: 992px) {
    .text-size {
        font-size: 0.65em;
    }
}
</style>
