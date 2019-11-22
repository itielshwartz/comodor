import { WebSocketLink } from 'apollo-link-ws';
import { ApolloClient } from 'apollo-client';
import { InMemoryCache } from 'apollo-cache-inmemory';

// Create a WebSocket link:
const link = new WebSocketLink({
  uri: `ws://localhost:8080/v1alpha1/graphql`,
  options: {
    reconnect: true,
    connectionParams: {
      headers: {
        // "x-hasura-admin-secret: "mylongsecretkey"
      }
    }
  }
})
const cache = new InMemoryCache();

export const apolloClient = new ApolloClient({
  link,
  cache
});

