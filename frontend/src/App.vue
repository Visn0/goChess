<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
    <nav class="navbar fixed-top navbar-expand-lg navbar-dark bg-dark">
        <div class="container-fluid">
            <RouterLink :to="{ path: '/' }" class="navbar-brand">GoChess</RouterLink>
            <button
                class="navbar-toggler"
                type="button"
                data-bs-toggle="collapse"
                data-bs-target="#navbarNav"
                aria-controls="navbarNav"
                aria-expanded="false"
                aria-label="Toggle navigation"
            >
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav">
                    <li class="nav-item">
                        <RouterLink :to="{ name: 'choose-nickname' }" class="nav-link">Choose nickname</RouterLink>
                    </li>
                    <li class="nav-item">
                        <RouterLink :to="{ name: 'rooms' }" class="nav-link">Rooms</RouterLink>
                    </li>
                </ul>
            </div>
        </div>
    </nav>

    <RouterView />
</template>
