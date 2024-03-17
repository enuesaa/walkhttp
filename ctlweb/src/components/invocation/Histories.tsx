import { Invocation } from '@/lib/api'
import { Box, Table } from '@radix-ui/themes'
import { HistoryDialog } from './HistoryDialog'

type Props = {
  messages: Invocation[]
}
export const HistoriesTable = ({ messages }: Props) => {
  return (
    <Table.Root>
      <Table.Header>
        <Table.Row>
          <Table.ColumnHeaderCell>Time</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Method</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell>Path</Table.ColumnHeaderCell>
          <Table.ColumnHeaderCell></Table.ColumnHeaderCell>
        </Table.Row>
      </Table.Header>

      <Table.Body>
        {messages.map((v,i) => (
          <Table.Row key={i}>
            <Table.RowHeaderCell>todo</Table.RowHeaderCell>
            <Table.Cell>{v.method}</Table.Cell>
            <Table.Cell>{v.url}</Table.Cell>
            <Table.Cell><HistoryDialog invocation={v} /></Table.Cell>
          </Table.Row>
        ))}
      </Table.Body>
    </Table.Root>
  )
}
