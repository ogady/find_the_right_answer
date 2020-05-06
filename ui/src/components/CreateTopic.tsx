import React from "react";
import { useLazyQuery } from "@apollo/react-hooks";
import gql from "graphql-tag";

import { Grid, Button } from "@material-ui/core";

type data = {
  topic: Topic;
};

type Topic = {
  startChar: StartChar;
  topicPiece: TopicPiece;
  numOfLikes: number;
};

type StartChar = {
  startChar: string;
};

type TopicPiece = {
  topicPiece: string;
};

const CREATE_TOPIC = gql`
  {
    topic {
      startChar {
        startChar
      }
      topicPiece {
        topicPiece
      }
      numOfLikes
    }
  }
`;

const CreateTopic: React.FC = () => {
  const [createTopic, { called, loading, error, data, refetch }] = useLazyQuery<
    data
  >(CREATE_TOPIC);

  if (called && loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error?.message}</p>;

  if (!called) {
    return (
      <div key={data?.topic.startChar.startChar}>
        <Grid container direction="column" justify="center" alignItems="center">
          <Button
            variant="contained"
            color="primary"
            onClick={() => createTopic()}
          >
            出題
          </Button>
        </Grid>
      </div>
    );
  } else {
    return (
      <div key={data?.topic.startChar.startChar}>
        <div>
          <Grid
            container
            direction="column"
            justify="center"
            alignItems="center"
          >
            <Button
              variant="contained"
              color="primary"
              onClick={() => refetch()}
            >
              再出題
            </Button>
          </Grid>
        </div>
        <br />
        <Grid container direction="column" justify="center" alignItems="center">
          <div>
            <b>{data && data.topic.startChar.startChar}</b> で始まる、
            <b>{data && data.topic.topicPiece.topicPiece}</b>は？
          </div>
        </Grid>
      </div>
    );
  }
};

const Topic: React.FC = () => {
  return (
    <div>
      <CreateTopic />
    </div>
  );
};

export default Topic;
