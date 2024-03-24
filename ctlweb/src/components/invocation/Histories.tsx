import { useMessgaes } from '@/lib/ws'
import { Badge, Card, ScrollArea, Text } from '@radix-ui/themes'
import { HistoryDialog } from './HistoryDialog'

export const Histories = () => {
  const messages = useMessgaes()

  return (
    <ScrollArea type='hover' scrollbars='vertical' style={{maxHeight: '700px', padding: '0 15px 0 0'}}>
      {messages.map((v,i) => (
        <Card key={i} my='2' style={{position: 'relative'}}>
          <Text as='div'>
            <Badge color='green'>{v.method}</Badge> {v.url}
          </Text>
          <Text color='gray' size='2' style={{position:'absolute', top: '0', right: '0'}}>
            200
          </Text>
          <Text color='gray' size='2' style={{position:'absolute', bottom: '0', right: '0'}}>
            date
          </Text>
        </Card>
      ))}
    </ScrollArea>
  )
}