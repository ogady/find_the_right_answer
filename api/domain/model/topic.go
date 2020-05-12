package model

type Topic struct {
	StartChar  StartChar  // 「startChar」で始まる
	TopicPiece TopicPiece // 「topicPiece」は？
	NumOfLikes NumOfLikes // いいねの数
}

type NumOfLikes struct {
	NumOfLikes int
}
