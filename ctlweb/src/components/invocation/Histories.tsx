import { useMessgaes } from '@/lib/ws'
import { Badge, Card, ScrollArea, Text } from '@radix-ui/themes'
import { HistroiesItem } from './HistoriesItem'

export const Histories = () => {
  const messages = useMessgaes()

  return (
    <ScrollArea type='hover' scrollbars='vertical' style={{maxHeight: '700px', padding: '0 15px 0 0'}}>
      {messages.map((v,i) => <HistroiesItem key={i} invocation={v} />)}
    </ScrollArea>
  )
}