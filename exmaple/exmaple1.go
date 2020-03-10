package main

//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
//
// 	"github.com/smoreg/simpleCash/fullCash"
// )
//
// func main() {
// 	withCash()
// }
//
// func simpleNoCash() {
// 	var client SomeWebClient
//
// 	// No cash
// 	client = RealWebClient{}
// 	// Request for post every time
// 	client.GetPost()
// 	client.GetPost()
// 	client.GetPost()
// }
// func withCash() {
// 	client := NewCashWebClient(RealWebClient{})
// 	client.GetPost()
// 	client.GetPost()
// 	client.GetPost()
//
// }
//
// type SomeWebClient interface {
// 	GetPostUserID() (int, error)
// 	GetPostID() (int, error)
// 	GetPost() (Post, error)
// }
//
// type Post struct {
// 	UserID int    `json:"userId"`
// 	ID     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Body   string `json:"body"`
// }
//
// type RealWebClient struct{}
//
// func (r RealWebClient) GetPost() (Post, error) {
// 	post, err := r.getPost()
// 	if err != nil {
// 		return Post{}, err
// 	}
// 	return post, nil
// }
// func (r RealWebClient) GetPostUserID() (int, error) {
// 	post, err := r.getPost()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return post.UserID, nil
// }
// func (r RealWebClient) GetPostID() (int, error) {
// 	post, err := r.getPost()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return post.ID, nil
// }
// func (r RealWebClient) getPost() (Post, error) {
// 	url := "https://jsonplaceholder.typicode.com/posts/1"
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		return Post{}, err
// 	}
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		return Post{}, err
// 	}
// 	defer res.Body.Close()
// 	post := Post{}
// 	err = json.NewDecoder(res.Body).Decode(&post)
// 	return post, err
// }
//
// // AUTO_GENERATED ====================================
//
// func NewCashWebClient(source SomeWebClient) *CashWebClient {
// 	return &CashWebClient{
// 		source:     source,
// 		cashSource: fullCash.NewResSerializer(),
// 	}
// }
//
// type CashWebClient struct {
// 	source     SomeWebClient
// 	cashSource fullCash.ResSerializer
// }
//
// func (c CashWebClient) GetPostUserID() (int, error) {
// 	panic("implement me")
// }
//
// func (c CashWebClient) GetPostID() (int, error) {
// 	panic("implement me")
// }
//
// func (c CashWebClient) GetPost() (Post, error) {
// 	names, err := c.cashSource.GetNames(2)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(123) // TODO (Smoreg) : REMOVE IT BAKA
// 	// c.GetPost()
// 	_ = names
// 	return Post{}, err
// }
