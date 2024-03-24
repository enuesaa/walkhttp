import { Histories } from '@/components/invocation/Histories'
import { History } from '@/components/invocation/History'
import { Box, Flex } from '@radix-ui/themes'

export default function TopPage() {
  return (
    <Flex gap='5'>
      <Box width='300px'>
        <Histories />
      </Box>
      <Box flexGrow='1' flexShrink='1'>
        <History />
      </Box>
    </Flex>
  )
}
