package handlers

import (
	"context"
	"net/http"
	"youtube-clone/database/pbs/helper"
	user_pb "youtube-clone/database/pbs/user-pb"
	"youtube-clone/gateway/client"
	"youtube-clone/gateway/helpers"

	"github.com/go-chi/chi"
)

// func GetUserLike(w http.ResponseWriter, r *http.Request) {

// }

////// media likes

func SetMediaLike(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(authUser("user")).(*user_pb.CurrentUserData)
	// currentUser := GetAuthUser(r)
	url := chi.URLParam(r, "url")
	var body map[string]any
	helpers.ReadJson(r, &body)
	helpers.ValidateAllowedParams(body, "is_like")

	isLike := helpers.ValidateBool(body["is_like"], "is_like")

	res, err := client.MediaService.CreateLikeMedia(context.Background(), &helper.LikeReq{
		IsLike: isLike,
		Url:    url,
		UserId: currentUser.Id,
	})
	if err != nil {
		panic(err)
	}
	PanicIfIsError(res.GetErr())
	if res.GetEmpty() != nil {
		helpers.WriteJsonMessage(w, "Like or dislike created.", 201)
		return
	}
	panic("LikeMedia should return empty or httpError!!!")
}

func DeleteMediaLike(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(authUser("user")).(*user_pb.CurrentUserData)
	// currentUser := GetAuthUser(r)
	url := chi.URLParam(r, "url")
	res, err := client.MediaService.DeleteLikeMedia(context.Background(), &helper.LikeReq{
		Url:    url,
		UserId: currentUser.Id,
	})
	if err != nil {
		panic(err)
	}
	PanicIfIsError(res.GetErr())
	if res.GetEmpty() != nil {
		helpers.WriteJsonMessage(w, "Like or dislike deleted.", 200)
		return
	}
	panic("DislikeMedia should return empty or httpError!!!")
}

////// comments likes

func SetCommentLike(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(authUser("user")).(*user_pb.CurrentUserData)
	// currentUser := GetAuthUser(r)
	url := chi.URLParam(r, "url")
	var body map[string]any
	helpers.ReadJson(r, &body)
	helpers.ValidateAllowedParams(body, "is_like")
	isLike := helpers.ValidateBool(body["is_like"], "is_like")

	res, err := client.CommentService.CreateLikeComment(context.Background(), &helper.LikeReq{
		IsLike: isLike,
		Url:    url,
		UserId: currentUser.Id,
	})
	if err != nil {
		panic(err)
	}
	PanicIfIsError(res.GetErr())
	if res.GetEmpty() != nil {
		helpers.WriteJsonMessage(w, "Like or dislike created.", 201)
		return
	}
	panic("CreateLikeComment should return empty or httpError!!!")
}

func DeleteCommentLike(w http.ResponseWriter, r *http.Request) {
	currentUser := r.Context().Value(authUser("user")).(*user_pb.CurrentUserData)
	// currentUser := GetAuthUser(r)
	url := chi.URLParam(r, "url")
	res, err := client.CommentService.DeleteLikeComment(context.Background(), &helper.LikeReq{
		Url:    url,
		UserId: currentUser.Id,
	})
	if err != nil {
		panic(err)
	}
	PanicIfIsError(res.GetErr())
	if res.GetEmpty() != nil {
		helpers.WriteJsonMessage(w, "Like or dislike deleted.", 200)
		return
	}
	panic("DeleteLikeComment should return empty or httpError!!!")
}
