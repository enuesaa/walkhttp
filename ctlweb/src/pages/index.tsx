import { MakeForm } from '@/components/request/MakeForm'
import { Histories } from '@/components/invocation/Histories'
import { Box, Flex } from '@radix-ui/themes'

export default function Page() {
  return (
    <Flex gap='5'>
      <Box flexGrow='0' flexShrink='0' width='300px'>
        <Histories />
      </Box>
      <Box flexGrow='1' flexShrink='1'>
        <MakeForm />
      </Box>
    </Flex>
  )
}
