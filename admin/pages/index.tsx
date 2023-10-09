import { useQuery } from '@apollo/client'
import { QueryFileinfo } from 'admin/graph'
import { Button } from 'admin/components/ui/button'
import { Card } from 'admin/components/ui/card'

export default function Page() {
  const { data } = useQuery(QueryFileinfo)
  console.log(data)

  return (
    <>
      <h2 className='text-3xl font-bold tracking-tight'>walkin</h2>
      <Card>aa</Card>
      <Button variant='destructive'>aa</Button>
    </>
  )
}