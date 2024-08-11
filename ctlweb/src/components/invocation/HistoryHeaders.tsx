import { Header } from '@/gql/types'
import { Box, Table } from '@radix-ui/themes'

type Props = {
  headers: Header[]
}
export const HistoryHeaders = ({ headers }: Props) => {
  if (headers.length === 0) {
    return <Box mb='4' />
  }

  return (
    <Table.Root variant='surface' mb='4'>
      <Table.Body>
        {headers.map((v, i) => (
          <Table.Row key={i}>
            <Table.RowHeaderCell>{v.name}</Table.RowHeaderCell>
            <Table.Cell>{v.value}</Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  )
}
