import { ref, type Ref } from 'vue'

class Timer {
    private intervalID: number

    private paused: Ref<boolean>
    public isPaused(): boolean {
        return this.paused.value
    }

    private stoped: Ref<boolean>
    public isStoped(): boolean {
        return this.stoped.value
    }

    private initialSeconds: number
    private initialDate: Date
    private lastUpdate: Date

    private remainingMinutes: Ref<number>
    private remainingSeconds: Ref<number>

    constructor(seconds: number) {
        this.intervalID = 0
        this.paused = ref(true)
        this.stoped = ref(false)
        this.initialSeconds = seconds
        this.initialDate = new Date()
        this.lastUpdate = new Date()
        this.remainingMinutes = ref(0)
        this.remainingSeconds = ref(0)

        this.remainingMinutes.value = Math.floor(this.initialSeconds / 60)
        this.remainingSeconds.value = Math.floor(this.initialSeconds % 60)
    }

    public init() {
        this.initialDate = new Date()
        this.initialDate.setMinutes(
            this.initialDate.getMinutes() + this.remainingMinutes.value,
            this.initialDate.getSeconds() + this.remainingSeconds.value,
            500
        )
    }

    private update() {
        this.lastUpdate = new Date()
        const now = this.lastUpdate.getTime()
        const distance = this.initialDate.getTime() - now

        this.setRemainingTime(distance)
    }

    public setRemainingTime(ms: number) {
        const minutes = Math.floor((ms % (1000 * 60 * 60)) / (1000 * 60))
        const seconds = Math.floor((ms % (1000 * 60)) / 1000)
        const milliseconds = Math.floor((ms % 1000) / 1000)

        this.remainingMinutes.value = minutes
        this.remainingSeconds.value = seconds

        if (minutes === 0 && seconds === 0 && milliseconds == 0) {
            this.stoped.value = true
            clearInterval(this.intervalID)
        }
    }

    public pause() {
        this.paused.value = true
        clearInterval(this.intervalID)
    }

    public resume() {
        // This update is to keep track of the countdown even when changing turns.
        const now = new Date().getTime()
        const distance = this.lastUpdate.getTime() - now
        this.initialDate.setTime(this.initialDate.getTime() + distance)

        this.paused.value = false
        this.intervalID = setInterval(this.update.bind(this), 500)
    }

    public toString(): string {
        const remainingMinutes = this.remainingMinutes.value
        const remainingSeconds = this.remainingSeconds.value

        const minutes = remainingMinutes < 10 ? `0${remainingMinutes}` : `${remainingMinutes}`
        const seconds = remainingSeconds < 10 ? `0${remainingSeconds}` : `${remainingSeconds}`

        return `${minutes}:${seconds}`
    }
}

export { Timer }
