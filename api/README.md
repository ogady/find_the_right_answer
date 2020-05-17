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

## お題にいいね

```
mutation likeTopic {
  likeTopic(input: {
    startChar:{
      startChar:"あ"
    }
  	topicPiece:{
      topicPiece:"尖っているもの"
    }
    numOfLikes:{
      numOfLikes:0
    }
  }){
    startChar{
      startChar
    }
    topicPiece{
      topicPiece
    }
    numOfLikes{
      numOfLikes
    }
  }
}
  
```