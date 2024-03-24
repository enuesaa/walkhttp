import { Invocation } from '@/lib/api'
import { useSetInvocation } from '@/lib/state';
import { Badge, Card, Text } from '@radix-ui/themes'
import { MouseEventHandler } from 'react';

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
    <Card my='2' style={{position: 'relative'}} onClick={handleClick}>
      <Text as='div'>
        <Badge color='green'>{invocation.method}</Badge> {invocation.url}
      </Text>
      <Text color='gray' size='2' style={{position:'absolute', top: '0', right: '0'}}>
        200
      </Text>
      <Text color='gray' size='2' style={{position:'absolute', bottom: '0', right: '0'}}>
        date
      </Text>
    </Card>
  )
}
