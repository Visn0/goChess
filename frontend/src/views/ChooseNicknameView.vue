<script setup lang="ts">
import router from '@/router'
import { usePlayerIDStore } from '@/stores/playerID'
import { ref, watch } from 'vue'

const MaxPlayerIDLength = 15
const playerIDStore = usePlayerIDStore()
let playerID = ref(playerIDStore.id)

function confirmNickname() {
    console.log('confirmNickname: ', playerID.value)
    playerIDStore.set(playerID.value)
    router.push({ name: 'rooms' })
}

watch(playerID, () => {
    if (playerID.value.length > MaxPlayerIDLength) {
        playerID.value = playerID.value.slice(0, MaxPlayerIDLength)
    }
})
</script>

<template>
    <main>
        <form class="container h-auto position-absolute top-40 start-50 translate-middle choose-nickname-form">
            <h1 class="row header text-light d-flex justify-content-center mb-3 p-0">
                <span class="p-0">Welcome to GoChess</span>
            </h1>
            <div class="row d-flex justify-content-center form-body">
                <div class="rounded border border-green py-2">
                    <div class="mb-3">
                        <label for="input-nickname" class="form-label text-light">Choose your nickname</label>
                        <input
                            type="text"
                            class="form-control"
                            id="input-nickname"
                            placeholder="DarkGreatMaster"
                            v-model="playerID"
                        />
                    </div>
                    <button type="button" class="btn btn-green w-100" @click="confirmNickname">Confirm</button>
                </div>
            </div>
        </form>
    </main>
</template>

<style scoped>
.top-40 {
    top: 40% !important;
}

.choose-nickname-form {
    width: 48vmin !important;
    max-width: 48vmin !important;
}

.header {
    width: 110%;
    justify-content: center;
    margin-left: -5%;
}

.form-body {
    width: 75%;
    margin: auto;
    justify-content: center;
}

@media (max-width: 450px) {
    .choose-nickname-form {
        width: 100vmin !important;
        max-width: 100vmin !important;
        font-size: 0.7em;
    }

    .form-body {
        width: 100%;
    }

    .btn,
    .form-control {
        font-size: 0.8em;
    }
}

@media (min-width: 450px) and (max-width: 576px) {
    .choose-nickname-form {
        width: 96vmin !important;
        max-width: 96vmin !important;
        font-size: 0.75em;
    }

    .form-body {
        width: 80%;
    }

    .btn,
    .form-control {
        font-size: 0.85em;
    }
}

@media (min-width: 576px) and (max-width: 680px) {
    .choose-nickname-form {
        width: 80vmin !important;
        max-width: 80vmin !important;
        font-size: 0.85em;
    }

    .form-body {
        width: 80%;
    }

    .btn,
    .form-control {
        font-size: 0.85em;
    }
}

@media (min-width: 680px) and (max-width: 768px) {
    .choose-nickname-form {
        width: 76vmin !important;
        max-width: 76vmin !important;
        font-size: 0.85em;
    }

    .btn,
    .form-control {
        font-size: 0.85em;
    }
}

@media (min-width: 768px) and (max-width: 882px) {
    .choose-nickname-form {
        width: 69.6vmin !important;
        max-width: 69.6vmin !important;
        font-size: 0.87em;
    }

    .btn,
    .form-control {
        font-size: 0.87em;
    }
}

@media (min-width: 882px) and (max-width: 992px) {
    .choose-nickname-form {
        width: 61.6vmin !important;
        max-width: 61.6vmin !important;
        font-size: 0.87em;
    }

    .btn,
    .form-control {
        font-size: 0.87em;
    }
}

@media (min-width: 992px) and (max-width: 1200px) {
    .choose-nickname-form {
        width: 56.6vmin !important;
        max-width: 56.6vmin !important;
        font-size: 0.9em;
    }

    .btn,
    .form-control {
        font-size: 0.91em;
    }
}

@media (min-width: 1200px) and (max-width: 1400px) {
    .choose-nickname-form {
        width: 54.6vmin !important;
        max-width: 54.6vmin !important;
        font-size: 0.87em;
    }

    .header {
        font-size: 2.65em;
    }

    .btn,
    .form-control {
        font-size: 0.91em;
    }
}

@media (min-width: 1400px) {
    .choose-nickname-form {
        font-size: 0.91em;
    }
    .header {
        width: 200%;
        justify-content: center;
        margin-left: -50%;
        font-size: 2.7em;
    }

    .btn,
    .form-control {
        font-size: 0.91em;
    }
}
</style>
