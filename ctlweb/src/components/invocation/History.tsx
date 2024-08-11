import { Invocation } from '@/gql/types'
import { Box, DataList, Heading, Separator } from '@radix-ui/themes'
import { HistoryHeaders } from './HistoryHeaders'
import { HistoryBody } from './HistoryBody'
import { HistorySectionTitle } from './HistorySectionTitle'

type Props = {
  invocation: Invocation
}
export const History = ({ invocation }: Props) => {
  return (
    <>
      <DataList.Root m='1' my='5'>
        <DataList.Item align='center'>
          <DataList.Label minWidth='88px'>status</DataList.Label>
          <DataList.Value>{invocation.status}</DataList.Value>
        </DataList.Item>

        <DataList.Item align='center'>
          <DataList.Label minWidth='88px'>method</DataList.Label>
          <DataList.Value>{invocation.method}</DataList.Value>
        </DataList.Item>

        <DataList.Item align='center'>
          <DataList.Label minWidth='88px'>url</DataList.Label>
          <DataList.Value>{invocation.url}</DataList.Value>
        </DataList.Item>

        <DataList.Item align='center'>
          <DataList.Label minWidth='88px'>started</DataList.Label>
          <DataList.Value>{invocation.started}</DataList.Value>
        </DataList.Item>
        
        <DataList.Item align='center'>
          <DataList.Label minWidth='88px'>received</DataList.Label>
          <DataList.Value>{invocation.received}</DataList.Value>
        </DataList.Item>
      </DataList.Root>

      <Box my='7' />

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
