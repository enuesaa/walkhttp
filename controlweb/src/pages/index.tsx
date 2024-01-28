import { Header } from '@/components/common/Header'
import { Main } from '@/components/common/Main'
import { Box, Flex } from '@radix-ui/themes'

export default function TopPage() {
  return (
    <>
      <Header />
      <Main>
        <Flex gap='3' style={{ flexBasis: '47%' }}>
          <Box grow='0' shrink='0' style={{ width: '100px', border: 'solid 1px #fafafa', height: '500px' }}>
            histories
          </Box>
          <Box grow='1' shrink='0' style={{ border: 'solid 1px #fafafa', height: '500px' }}>
            request
          </Box>
          <Box grow='1' shrink='0' style={{ border: 'solid 1px #fafafa', height: '500px' }}>
            response
          </Box>
        </Flex>
      </Main>
    </>
  )
}
