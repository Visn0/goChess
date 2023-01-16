import { createRouter, createWebHistory } from 'vue-router'
import GameView from '../views/GameView.vue'
import ChooseNickname from '../views/ChooseNicknameView.vue'

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'choose-nickname',
            component: ChooseNickname
        },
        {
            path: '/game',
            name: 'game',
            component: GameView
        }
    ]
})

export default router
