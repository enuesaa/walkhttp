import { Invocation } from '@/gql/types'
import { Code, DataList, Heading } from '@radix-ui/themes'
import { HistoryRequestHeaders } from './HistoryRequestHeaders'
import { HistoryResponseHeaders } from './HistoryResponseHeaders'

type Props = {
  invocation: Invocation
}
export const History = ({ invocation }: Props) => {
  return (
    <>
      <DataList.Root my='2' mx='5'>
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
      </DataList.Root>

      <Heading m='1' mt='5' size='3'>
        Request Headers
      </Heading>
      <HistoryRequestHeaders headers={invocation.requestHeaders} />

      <Heading m='1' mt='5' size='3'>
        Response Headers
      </Heading>
      <HistoryResponseHeaders headers={invocation.responseHeaders} />

      <Heading m='1' mt='5' size='3'>
        Response Body
      </Heading>
      <Code style={{ display: 'block' }}>
        <pre>{invocation.responseBody}</pre>
      </Code>
    </>
  )
}
