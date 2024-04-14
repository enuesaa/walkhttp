import { atom, useAtomValue } from 'jotai'
import { Invocation } from './graphtypes'

const wsEndpoint = 'ws://localhost:3000/ws'

const messages = atom<Invocation[]>([])
messages.onMount = (setMessages) => {
  const ws = new WebSocket(wsEndpoint)
  ws.addEventListener('message', (event: MessageEvent<string>) => {
    try {
      const invocation = JSON.parse(event.data) as Invocation
      // TODO: remove. this is for debug
      setMessages(values => [invocation, invocation, invocation, invocation, invocation, invocation, invocation, invocation, invocation, invocation, invocation, invocation, invocation, invocation, invocation, ...values])
    } catch (e) {
      console.error('failed to parse', e)
    }
  })

  return () => ws.close()
}

export const useMessgaes = () => useAtomValue(messages)
