import { computed, shallowRef, type Ref } from 'vue'
import { defineStore } from 'pinia'
import type { Game } from '@/models/game'

export const useGameStore = defineStore('game', () => {
    const _game: Ref<Game | null> = shallowRef(null)
    const game = computed(() => _game.value)
    const isEmpty = computed(() => _game.value === null || _game.value === undefined)

    function set(newGame: Game) {
        _game.value = newGame
    }
    return { game, isEmpty, set }
})
