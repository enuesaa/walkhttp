import { atom, useAtomValue, useSetAtom } from 'jotai'
import { Invocation } from './api'

const messages = atom<Invocation[]>([])
messages.onMount = (setMessages) => {
  const ws = new WebSocket('ws://localhost:3000/ws')
  const onMessage = (event: MessageEvent<string>) => {
    console.log(event.data)
    try {
      const invocation = JSON.parse(event.data) as Invocation
      setMessages(values => [...values, invocation])
    } catch (e) {
      console.log(`parse error: ${e}`)
    }
  }
  ws.addEventListener('message', onMessage)

  // return () => {
  //   ws.close()
  //   ws.removeEventListener('message', onMessage)
  // }
}

export const useMessgaes = () => useAtomValue(messages)
export const useAddMessage = () => {
  const setMessages = useSetAtom(messages)
  return (value: Invocation) => setMessages(values => [...values, value])
}
