import { computed, ref } from 'vue'
import { defineStore } from 'pinia'

export const usePlayerIDStore = defineStore('playerID', () => {
    const _id = ref('')
    const id = computed(() => {
        if (_id.value !== '') {
            return _id.value
        }

        const playerID = localStorage.getItem('playerID')
        if (playerID) {
            return playerID
        }

        return ''
    })

    const isEmpty = computed(() => id.value === '')

    function set(newID: string) {
        _id.value = newID
        localStorage.setItem('playerID', newID)
    }

    return { id, isEmpty, set }
})
