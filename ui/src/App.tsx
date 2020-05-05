import React from "react";
import Topic from "./components/CreateTopic";
// import Topic from "./components/CreateTopic_copy";
import Header from "./components/Header";
import { Container } from "@material-ui/core";
import ApolloClient from "apollo-boost";
import { ApolloProvider } from "@apollo/react-hooks";

const client = new ApolloClient({
  // uri: process.env.REACT_APP_API_URL,
  uri: "https://ogady.net/api",
});

function App() {
  return (
    <ApolloProvider client={client}>
      <Container maxWidth="sm">
        <div>
          <Header />
          <Topic />
        </div>
      </Container>
    </ApolloProvider>
  );
}

export default App;
