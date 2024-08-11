import { Invocation } from '@/gql/types'
import { Badge, Button, Dialog, Flex } from '@radix-ui/themes'
import { History } from './History'
import styles from './HistoriesItem.css'

type Props = {
  invocation: Invocation
}
export const HistroiesItem = ({ invocation }: Props) => {
  const datefmt = new Intl.DateTimeFormat('en-US', {
    hour: 'numeric',
    hour12: false,
    minute: 'numeric',
    second: 'numeric',
  })
  const created = datefmt.format(new Date(invocation.created))

  return (
    <Dialog.Root>
      <Dialog.Trigger>
        <Button variant='ghost' className={styles.main}>
          <Badge color='green'>{invocation.method}</Badge> {invocation.url}
          <div className={styles.created}>{created}</div>
        </Button>
      </Dialog.Trigger>
      <Dialog.Content maxWidth='1000px' aria-description='Invocation' style={{ position: 'relative' }}>
        <Dialog.Title>Invocation</Dialog.Title>
        <Dialog.Close>
          <Button variant='soft' color='gray' style={{ position: 'absolute', top: '20px', right: '20px' }}>
            Close
          </Button>
        </Dialog.Close>
        <History invocation={invocation} />
      </Dialog.Content>
    </Dialog.Root>
  )
}
