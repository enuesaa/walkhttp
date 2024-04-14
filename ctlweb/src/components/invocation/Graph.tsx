import { useListInvocations } from '@/lib/graph'

export const Graph = () => {
  const { loading, error, data } = useListInvocations()

  if (loading) {
    return (<>loading</>)
  }

  return (
    <>
      {data?.invocations.map((v,i) => (
        <div key={i}>{v.id}</div>
      ))}
    </>
  )
}
