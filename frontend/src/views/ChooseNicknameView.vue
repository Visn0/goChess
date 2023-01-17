<script setup lang="ts">
import router from '@/router'
import { usePlayerIDStore } from '@/stores/playerID'

const playerIDStore = usePlayerIDStore()

let playerID = playerIDStore.id

function confirmNickname() {
    console.log('confirmNickname: ', playerID)
    playerIDStore.set(playerID)
    router.push({ name: 'rooms' })
}
</script>

<template>
    <main>
        <form class="container h-auto position-absolute top-40 start-50 translate-middle">
            <h1 class="row text-light d-flex justify-content-center mb-3">
                <span class="col-12">Welcome to GoChess</span>
            </h1>
            <div class="row d-flex justify-content-center">
                <div class="col-sm-12 col-lg-6 rounded border border-green py-2">
                    <div class="mb-3">
                        <label for="input-nickname" class="form-label text-light">Choose your nickname</label>
                        <input
                            type="text"
                            class="form-control m-auto"
                            id="input-nickname"
                            placeholder="DarkGreatMaster"
                            v-model="playerID"
                        />
                    </div>
                    <button type="button" class="btn bg-green w-100" @click="confirmNickname">Confirm</button>
                </div>
            </div>
        </form>
    </main>
</template>

<style scoped>
.top-40 {
    top: 40% !important;
}
</style>
