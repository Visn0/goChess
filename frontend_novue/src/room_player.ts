class RoomPlayer {
    id: string

    constructor(id?: string) {
        this.id = id === null || id === undefined ? '' : id
    }

    public toDivHTML(): string {
        const htmlStr = `<li class="list-group-item">${this.id}</li>`
        return htmlStr
    }

    public static createFromJSON(src: any): RoomPlayer {
        const dst: RoomPlayer = Object.assign(new RoomPlayer(), src)
        return dst
    }
}

export { RoomPlayer }
