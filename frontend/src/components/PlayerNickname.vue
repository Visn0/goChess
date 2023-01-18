<script setup lang="ts">
defineProps<{
    nickname: string
}>()
</script>

<template>
    <p class="my-0 text-size badge text-start bg-dark text-white">
        {{ nickname }}
    </p>
</template>

<style scoped>
.text-size {
    font-size: 0.75em;
}

@media (max-width: 576px) {
    .text-size {
        font-size: 0.5em;
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
