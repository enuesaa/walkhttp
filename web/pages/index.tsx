import { useQuery, gql } from '@apollo/client'

const GET_FILEINFO = gql`
query {
  fileinfo(name: "aa") {
    name
    description
  }
}
`;

export default function Page() {
  const { loading, error, data } = useQuery(GET_FILEINFO)
  if (loading) return (<p>loading..</p>);
  if (error) return (<p>error {error.message}</p>);
  console.log(data)

  return (
    <>
      top page
    </>
  )
}
