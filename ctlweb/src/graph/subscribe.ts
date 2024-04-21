import { useSubscription } from 'urql'
import { gql } from 'urql'

export const newMessages = gql`
  subscription {
    invocations {
      id,
      url
    }
  }
`;

// const handleSubscription = (messages = [], response) => {
//   return [response.newMessages, ...messages];
// };

// const Messages = () => {
//   const [res] = useSubscription({ query: newMessages }, handleSubscription);

//   if (!res.data) {
//     return <p>No new messages</p>;
//   }