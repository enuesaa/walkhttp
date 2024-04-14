import { gql, useQuery } from '@apollo/client';

const GET_LOCATIONS = gql`
  query {
    invocations(status: 404) {
      id
    }
  }
`;

export const Graph = () => {
  const { loading, error, data } = useQuery(GET_LOCATIONS);

  if (loading) {
    return (<>loading</>)
  }
  console.log(data)

  return (
    <>
      {data.invocations.map((v,i) => (
        <div key={i}>{v.id}</div>
      ))}
    </>
  )
}
