import { Invocation } from '@/graph/types'
import { useSetInvocation } from '@/lib/state'
import { Badge, Card, Text } from '@radix-ui/themes'
import { MouseEventHandler } from 'react'
import styles from './HistoriesItem.css'

type Props = {
  invocation: Invocation;
}
export const HistroiesItem = ({ invocation }: Props) => {
  const setInvocation = useSetInvocation()

  const handleClick: MouseEventHandler<HTMLDivElement> = (e) => {
    e.preventDefault()
    setInvocation(invocation)
  }

  return (
    <Card my='2' onClick={handleClick} className={styles.main}>
      <Text as='div'>
        <Badge color='green'>{invocation.method}</Badge> {invocation.url}
      </Text>
      <Text color='gray' size='2' style={{position:'absolute', top: '0', right: '0'}}>
        {invocation.status}
      </Text>
      <Text color='gray' size='2' style={{position:'absolute', bottom: '0', right: '0'}}>
        date
      </Text>
    </Card>
  )
}
