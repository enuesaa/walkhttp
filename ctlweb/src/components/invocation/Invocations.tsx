import { Flex } from '@radix-ui/themes'
import { RequestBox } from './RequestBox'
import { ResponseBox } from './ResponseBox'

export const Invocations = () => {
  return (
    <Flex gap='3' style={{ flexBasis: '47%' }}>
      <RequestBox />
      <ResponseBox />
    </Flex>
  )
}
