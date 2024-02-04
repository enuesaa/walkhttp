import { Flex } from '@radix-ui/themes'
import { Histories } from './Histories'
import { RequestBox } from './RequestBox'
import { ResponseBox } from './ResponseBox'

export const Invocations = () => {
  return (
    <Flex gap='3' style={{ flexBasis: '47%' }}>
      <Histories />
      <RequestBox />
      <ResponseBox />
    </Flex>
  )
}