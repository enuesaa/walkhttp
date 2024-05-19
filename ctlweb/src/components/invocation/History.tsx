import { useGetInvocation } from '@/lib/state'
import { Code, DataList, Heading, Skeleton } from '@radix-ui/themes'

export const History = () => {
  const invocation = useGetInvocation()
  if (invocation === undefined) {
    return <Skeleton />
  }

  return (
    <>
      <DataList.Root m='5'>
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
