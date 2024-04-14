import { useGetInvocations } from '@/lib/graph'

export const Graph = () => {
  const { loading, error, data } = useGetInvocations()

  if (loading) {
    return (<>loading</>)
  }
  console.log(data)

  return (
    <>
      {data?.invocations.map((v,i) => (
        <div key={i}>{v.id}</div>
      ))}
    </>
  )
}
