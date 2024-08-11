import { Header } from '@/gql/types'
import { Table } from '@radix-ui/themes'

type Props = {
  headers: Header[]
}
export const HistoryRequestHeaders = ({ headers }: Props) => {
  return (
    <Table.Root variant='surface'>
      <Table.Body>
        {headers.map((v,i) => (
          <Table.Row key={i}>
            <Table.RowHeaderCell>{v.name}</Table.RowHeaderCell>
            <Table.Cell>{v.value}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  )
}
