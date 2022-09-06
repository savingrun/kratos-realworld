package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)

func (s *RealWorldService) ListArticles(ctx context.Context, request *v1.ListArticlesQueryRequest) (reply *v1.ArticlesReply, err error) {
	return &v1.ArticlesReply{}, nil
}

func (s *RealWorldService) FeedArticles(ctx context.Context, request *v1.FeedArticlesQueryRequest) (reply *v1.ArticlesReply, err error) {
	return &v1.ArticlesReply{}, nil
}

func (s *RealWorldService) GetArticle(ctx context.Context, request *v1.SlugPathVariableRequest) (reply *v1.ArticleReply, err error) {
	return &v1.ArticleReply{Article: &v1.ArticleReply_Data{Title: "saving"}}, nil
}

func (s *RealWorldService) CreateArticle(ctx context.Context, request *v1.CreateArticleRequest) (reply *v1.ArticleReply, err error) {
	return &v1.ArticleReply{Article: &v1.ArticleReply_Data{Title: "saving"}}, nil
}

func (s *RealWorldService) UpdateArticle(ctx context.Context, request *v1.UpdateArticleRequest) (reply *v1.ArticleReply, err error) {
	return &v1.ArticleReply{Article: &v1.ArticleReply_Data{Title: "saving"}}, nil
}

func (s *RealWorldService) DeleteArticle(ctx context.Context, request *v1.SlugPathVariableRequest) (reply *v1.ArticleReply, err error) {
	return &v1.ArticleReply{Article: &v1.ArticleReply_Data{Title: "saving"}}, nil
}

func (s *RealWorldService) AddComments(ctx context.Context, request *v1.AddCommentsRequest) (reply *v1.CommentReply, err error) {
	return &v1.CommentReply{Comment: &v1.CommentReply_Data{Body: "saving"}}, nil
}

func (s *RealWorldService) GetComments(ctx context.Context, request *v1.SlugPathVariableRequest) (reply *v1.CommentsReply, err error) {
	return &v1.CommentsReply{Comments: []*v1.CommentsReply_Data{&v1.CommentsReply_Data{Body: "saving"}}}, nil
}

func (s *RealWorldService) DeleteComment(ctx context.Context, request *v1.DeleteCommentPathVariableRequest) (reply *v1.ArticleReply, err error) {
	return &v1.ArticleReply{Article: &v1.ArticleReply_Data{Title: "saving"}}, nil
}

func (s *RealWorldService) FavoriteArticle(ctx context.Context, request *v1.SlugPathVariableRequest) (reply *v1.ArticleReply, err error) {
	return &v1.ArticleReply{Article: &v1.ArticleReply_Data{Title: "saving"}}, nil
}

func (s *RealWorldService) UnfavoriteArticle(ctx context.Context, request *v1.SlugPathVariableRequest) (reply *v1.ArticleReply, err error) {
	return &v1.ArticleReply{Article: &v1.ArticleReply_Data{Title: "saving"}}, nil
}

func (s *RealWorldService) GetTags(ctx context.Context, request *v1.GetTagsRequest) (reply *v1.TagsReply, err error) {
	return &v1.TagsReply{Tags: []string{"saving"}}, nil
}
