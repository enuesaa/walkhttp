import { newMessages } from '@/graph/subscribe';
import { useSubscription } from 'urql'

export const Subscribe = () => {
  const [res] = useSubscription({ query: newMessages });
  console.log(res.data)

  return (<></>)
}
