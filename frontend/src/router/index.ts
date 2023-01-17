import { createRouter, createWebHistory } from 'vue-router'
import GameView from '../views/GameView.vue'
import RoomsView from '../views/RoomsView.vue'
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
            path: '/rooms',
            name: 'rooms',
            component: RoomsView
        },
        {
            path: '/game',
            name: 'game',
            component: GameView
        }
    ]
})

export default router
