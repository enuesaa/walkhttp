import { Invocation } from '@/graph/types'
import { Code, DataList, Heading } from '@radix-ui/themes'

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
          <DataList.Label minWidth='88px'>created</DataList.Label>
          <DataList.Value>{invocation.created}</DataList.Value>
        </DataList.Item>
      </DataList.Root>

      <Heading m='1' mt='5' size='7'>
        Request
      </Heading>
      <Heading m='3' size='3'>
        Headers
      </Heading>
      <Heading m='1' mt='5' size='7'>
        Response
      </Heading>
      <Heading m='3' size='3'>
        Headers
      </Heading>
      <Heading m='3' size='3'>
        Body
      </Heading>
      <Code style={{ display: 'block' }}>
        <pre>{invocation.responseBody}</pre>
      </Code>
    </>
  )
}
