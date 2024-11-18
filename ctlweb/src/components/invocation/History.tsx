import { Invocation } from '@/gql/types'
import { HistoryHeaders } from './HistoryHeaders'
import { HistoryBody } from './HistoryBody'
import { HistorySectionTitle } from './HistorySectionTitle'

type Props = {
  invocation: Invocation
}
export const History = ({ invocation }: Props) => {
  return (
    <>
      <HistorySectionTitle title='Request Headers' />
      <HistoryHeaders headers={invocation.requestHeaders} />

      <HistorySectionTitle title='Request Body' />
      <HistoryBody body={invocation.requestBody} />

      <HistorySectionTitle title='Response Headers' />
      <HistoryHeaders headers={invocation.responseHeaders} />

      <HistorySectionTitle title='Response Body' />
      <HistoryBody body={invocation.responseBody} />
    </>
  )
}
