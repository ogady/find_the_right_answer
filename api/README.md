# FTRA GraphQL API

## お題作成 Query

```
query{
  topic{
    startChar{
      startChar
    }
    topicPiece{
      topicPiece
    }
    numOfLikes
  }
}
```

## お題追加 Mutation

```
mutation addTopicPiece {
  addTopicPiece(input: {topicPiece:"尖っているもの"}) {
    topicPiece
  }
}
```
