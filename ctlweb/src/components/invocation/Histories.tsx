import { ScrollArea } from '@radix-ui/themes'
import { useSubscribeInvocations } from '@/graph/invocations-subscribe'
import { HistroiesItem } from './HistoriesItem'

export const Histories = () => {
  const invocations = useSubscribeInvocations()
  if (invocations.fetching || invocations.error !== undefined) {
    return <></>
  }

  return (
    <ScrollArea type='hover' scrollbars='vertical' style={{maxHeight: '700px', padding: '0 15px 0 0'}}>
      {invocations.data?.invocations.map((v,i) => <HistroiesItem key={i} invocation={v} />)}
    </ScrollArea>
  )
}
