import { ScrollArea } from '@radix-ui/themes'
import { useSubscribeInvocations } from '@/gql/queries/subscribeInvocations'
import { HistroiesItem } from './HistoriesItem'

export const Histories = () => {
  const invocations = useSubscribeInvocations()
  if (invocations.fetching || invocations.error !== undefined) {
    return <></>
  }

  return (
    <ScrollArea type='hover' scrollbars='vertical'>
      {invocations.data?.subscribeInvocations.map((v, i) => <HistroiesItem key={i} invocation={v} />)}
    </ScrollArea>
  )
}
