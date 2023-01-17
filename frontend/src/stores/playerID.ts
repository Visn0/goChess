import { computed } from 'vue'
import { defineStore } from 'pinia'

export const usePlayerIDStore = defineStore('playerID', () => {
    const id = computed(() => {
        const playerID = localStorage.getItem('playerID')
        if (playerID) {
            return playerID
        }
        return ''
    })

    function set(newID: string) {
        localStorage.setItem('playerID', newID)
    }
    return { id, set }
})
