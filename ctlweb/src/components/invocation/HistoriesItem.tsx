import { Invocation } from '@/graph/types'
import { Badge, Button, Text, Dialog, Card } from '@radix-ui/themes'
import { History } from './History'
import styles from './HistoriesItem.css'

type Props = {
  invocation: Invocation
}
export const HistroiesItem = ({ invocation }: Props) => {
  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='ghost' className={styles.main}>
          <Badge color='green'>{invocation.method}</Badge> {invocation.url}
        </Button>
      </Dialog.Trigger>
      <Dialog.Content maxWidth='1000px'>
        <History invocation={invocation} />
      </Dialog.Content>
    </Dialog.Root>
  )
}
