import { atom, useAtomValue, useSetAtom } from 'jotai'

const messages = atom<string[]>([])
messages.onMount = (setMessages) => {
  const ws = new WebSocket('ws://localhost:3000/ws')
  const onMessage = (event: MessageEvent<string>) => {
    console.log(event.data)
    setMessages(values => [...values, event.data])
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
  return (value: string) => setMessages(values => [...values, value])
}
