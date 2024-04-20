import { Histories } from '@/components/invocation/Histories'
import { History } from '@/components/invocation/History'
import { Box, Flex } from '@radix-ui/themes'
import { Graph } from '@/components/invocation/Graph'

export default function TopPage() {
  return (
    <>
      <Graph />
      <Flex gap='5'>
        <Box flexGrow='0' flexShrink='0' width='300px'>
          <Histories />
        </Box>
        <Box flexGrow='1' flexShrink='1'>
          <History />
        </Box>
      </Flex>
    </>
  )
}
