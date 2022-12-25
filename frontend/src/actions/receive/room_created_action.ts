class RoomCreatedAction {
    public Invoke(body: string) {
        console.log('=> Room created: ', body)
    }
}

export { RoomCreatedAction }
