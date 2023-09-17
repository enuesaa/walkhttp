import { useQuery } from '@apollo/client'
import { QueryFileinfo } from '@/graph'
import { FileDrawer } from '@/components/FileDrawer'
import { Heading } from '@chakra-ui/react'
import { Button } from "@/components/ui/button"

export default function Page() {
  const { data } = useQuery(QueryFileinfo)
  console.log(data)

  return (
    <>
      <Heading>walkin</Heading>
      <Button variant='secondary'>aa</Button>
      <FileDrawer />
    </>
  )
}
