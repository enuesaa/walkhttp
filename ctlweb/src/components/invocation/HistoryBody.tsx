import { Box, Code } from '@radix-ui/themes'

type Props = {
  body: string
}
export const HistoryBody = ({ body }: Props) => {
  if (body.length === 0) {
    return <Box mb='4' />
  }

  return (
    <Code mb='4' style={{ display: 'block' }}>
      <pre>{body}</pre>
    </Code>
  )
}
