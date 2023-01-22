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

    private _durationMs: number
    public set durationMs(ms: number) {
        this._durationMs = ms
        this.setRemainingTime(ms)
    }
    private timeConsumedMs: number
    private lastClockTime: Date
    private remainingMinutes: Ref<number>
    private remainingSeconds: Ref<number>

    constructor(milliseconds: number) {
        this.intervalID = 0
        this.paused = ref(true)
        this.stoped = ref(false)
        this._durationMs = milliseconds + 500
        this.timeConsumedMs = 0
        this.lastClockTime = new Date()
        this.remainingMinutes = ref(0)
        this.remainingSeconds = ref(0)
    }

    public start() {
        this.lastClockTime = new Date()
        this.paused.value = false

        this.intervalID = setInterval(() => this.update(), 1000)
    }

    public pause() {
        this.timeConsumedMs += Date.now() - this.lastClockTime.getTime()
        this.lastClockTime = new Date()
        this.paused.value = true

        clearInterval(this.intervalID)
    }

    public setRemainingTime(ms: number) {
        this.timeConsumedMs = this.durationMs - ms

        this.remainingMinutes.value = Math.floor((ms % (1000 * 60 * 60)) / (1000 * 60))
        this.remainingSeconds.value = Math.floor((ms % (1000 * 60)) / 1000)

        if (ms === 0) {
            this.paused.value = true
            this.stoped.value = true
            clearInterval(this.intervalID)
            return
        }
    }

    private update() {
        this.timeConsumedMs += Date.now() - this.lastClockTime.getTime()
        this.lastClockTime = new Date()

        const remainingMs = this.durationMs - this.timeConsumedMs
        this.setRemainingTime(remainingMs)
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
