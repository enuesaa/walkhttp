import { useGetInvocation } from '@/lib/state'
import { Code, DataList, Heading, Skeleton } from '@radix-ui/themes'

export const History = () => {
  const invocation = useGetInvocation()
  if (invocation === undefined) {
    return (<Skeleton />)
  }

  return (
    <>
      <DataList.Root>
        <DataList.Item align='center'>
          <DataList.Label minWidth='88px'>status</DataList.Label>
          <DataList.Value>
            200
          </DataList.Value>
        </DataList.Item>
        
        <DataList.Item align='center'>
          <DataList.Label minWidth='88px'>method</DataList.Label>
          <DataList.Value>
            {invocation.method}
          </DataList.Value>
        </DataList.Item>
        
        <DataList.Item align='center'>
          <DataList.Label minWidth='88px'>url</DataList.Label>
          <DataList.Value>
            {invocation.url}
          </DataList.Value>
        </DataList.Item>
      </DataList.Root>

      <Heading>Request Headers</Heading>
      <Code style={{display: 'block'}}>
        {JSON.stringify(invocation.requestHeaders, null, '  ')}
      </Code>

      <Heading>Response Headers</Heading>
      <Code style={{display: 'block'}}>
        {JSON.stringify(invocation.responseHeaders, null, '  ')}
      </Code>

      <Heading>Response Body</Heading>
      <Code style={{display: 'block'}}>
        {invocation.responseBody}
      </Code>
    </>
  )
}