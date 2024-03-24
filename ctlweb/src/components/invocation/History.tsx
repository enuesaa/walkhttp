import { DataList } from '@radix-ui/themes'

export const History = () => {
  return (
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
          GET
        </DataList.Value>
      </DataList.Item>
      
      <DataList.Item align='center'>
        <DataList.Label minWidth='88px'>path</DataList.Label>
        <DataList.Value>
          /api/aaa
        </DataList.Value>
      </DataList.Item>
    </DataList.Root>
  )
}