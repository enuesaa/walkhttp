import { useQuery } from '@apollo/client'
import { QueryFileinfo } from '@/graph'

export default function Page() {
  const { data } = useQuery(QueryFileinfo)
  console.log(data)

  return (<></>)
}
