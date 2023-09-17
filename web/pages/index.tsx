import { useQuery } from '@apollo/client'
import { QueryFileinfo } from '@/graph'
import { Button } from "@/components/ui/button"

export default function Page() {
  const { data } = useQuery(QueryFileinfo)
  console.log(data)

  return (
    <>
      <Button variant='destructive'>aa</Button>
    </>
  )
}
