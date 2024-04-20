import { ScrollArea } from '@radix-ui/themes'
import { useListInvocations } from '@/lib/graph'
import { HistroiesItem } from './HistoriesItem'

export const Histories = () => {
  const { loading, error, data } = useListInvocations()

  if (loading || error != undefined) {
    return <></>
  }

  return (
    <ScrollArea type='hover' scrollbars='vertical' style={{maxHeight: '700px', padding: '0 15px 0 0'}}>
      {data?.invocations.map((v,i) => <HistroiesItem key={i} invocation={v} />)}
    </ScrollArea>
  )
}