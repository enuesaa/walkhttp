import Link from 'next/link'
import { Flex, Box, Container } from '@radix-ui/themes'
import styles from '@/components/common/Header.css'

export const Header = () => {
  return (
    <header className={styles.main}>
      <Container size='4'>
        <Flex>
          <Box flexGrow='1'>
            <Link href='/' className={styles.heading}>
              walkin
            </Link>
          </Box>

          <Link href='/invoke' className={styles.invoke}>
            invoke
          </Link>
        </Flex>
      </Container>
    </header>
  )
}
