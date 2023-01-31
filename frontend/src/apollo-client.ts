import { ApolloClient, InMemoryCache } from "@apollo/client";

const client = new ApolloClient({
    uri: "http://localhost:8080/query", // TODO: Don't hardcode this!!
    cache: new InMemoryCache(),
});

export default client;