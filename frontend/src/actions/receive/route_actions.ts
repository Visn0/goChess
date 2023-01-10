import { ReceiveAction } from './receive_action'

class RouteActions {
    private receiveActions: Map<string, ReceiveAction>

    constructor(actions: Map<string, ReceiveAction>) {
        this.receiveActions = actions
    }

    public route(): (event: MessageEvent) => void {
        const receiveActions = this.receiveActions
        return (event: MessageEvent) => {
            const body = event.data
            class Params {
                action: string
            }

            const p: Params = JSON.parse(body)
            const action = receiveActions.get(p.action)

            action?.Invoke(body)
        }
    }
}

export { RouteActions }
