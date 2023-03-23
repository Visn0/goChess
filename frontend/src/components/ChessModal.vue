<script setup lang="ts">
defineProps<{
    showModal: boolean
    title: string
    buttons?: Array<{
        text: string
        action: Function
    }>
}>()

const emit = defineEmits(['close'])

function closeModal() {
    emit('close')
}
</script>

<template>
    <div v-if="showModal" class="modal" tabindex="-1" style="display: block" @click.self="closeModal()">
        <div class="modal-dialog modal-dialog-centered">
            <div class="modal-content bg-dark text-light">
                <div class="modal-body">
                    <h5 class="modal-title">{{ title }}</h5>
                    <button
                        v-for="(b, idx) in buttons"
                        :key="idx"
                        type="button"
                        class="mt-2 w-100 btn btn-sm btn-green"
                        @click="b.action"
                    >
                        {{ b.text }}
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
