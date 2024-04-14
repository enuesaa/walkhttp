import { Histories } from '@/components/invocation/Histories'
import { History } from '@/components/invocation/History'
import { Box, Flex } from '@radix-ui/themes'
import { ApolloClient, InMemoryCache, ApolloProvider, gql } from '@apollo/client';
import { Graph } from '@/components/invocation/Graph';

const client = new ApolloClient({
  uri: 'http://localhost:3000/graph',
  cache: new InMemoryCache(),
});

export default function TopPage() {
  return (
    <ApolloProvider client={client}>
      <Graph />
      <Flex gap='5'>
        <Box flexGrow='0' flexShrink='0' width='300px'>
          {/* <Histories /> */}
        </Box>
        <Box flexGrow='1' flexShrink='1'>
          {/* <History /> */}
        </Box>
      </Flex>
    </ApolloProvider>
  )
}
