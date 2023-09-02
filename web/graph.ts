import { gql } from '@apollo/client'

export const QueryFileinfo = gql`
query {
  fileinfo(name: "aa") {
    name
    description
  }
}
`;