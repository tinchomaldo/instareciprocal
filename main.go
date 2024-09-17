package main

import (
	"fmt"
	"html/template"
	"os"

	"github.com/tinchomaldo/instareciprocal/internal/domain"
	"github.com/tinchomaldo/instareciprocal/internal/infrastructure"
)

const htmlTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Instagram Reciprocal Followers</title>
    <style>
        table {
            border-collapse: collapse;
            width: 100%;
        }
        th, td {
            border: 1px solid #ddd;
            padding: 8px;
            text-align: left;
        }
        th {
            background-color: #f2f2f2;
        }
    </style>
</head>
<body>
    <h1>Instagram Reciprocal Followers</h1>
    
    <h2>Users you follow but don't follow you back:</h2>
    <table>
        <tr><th>Username</th></tr>
        {{range .NotFollowingBack}}
        <tr><td>{{.}}</td></tr>
        {{end}}
    </table>

    <h2>Users who follow you but you don't follow back:</h2>
    <table>
        <tr><th>Username</th></tr>
        {{range .NotFollowedBack}}
        <tr><td>{{.}}</td></tr>
        {{end}}
    </table>
</body>
</html>
`

type templateData struct {
	NotFollowingBack []domain.User
	NotFollowedBack  []domain.User
}

func main() {
	repo := &infrastructure.JSONRepository{
		FollowersFile: "followers.json",
		FollowingFile: "following.json",
	}

	followers, err := repo.GetFollowers()
	if err != nil {
		handleError(err)
		return
	}

	following, err := repo.GetFollowing()
	if err != nil {
		handleError(err)
		return
	}

	comparator := domain.NewComparator(followers, following)
	notFollowingBack := comparator.NotFollowingBack()
	notFollowedBack := comparator.NotFollowedBack()

	tmpl, err := template.New("result").Parse(htmlTemplate)
	if err != nil {
		handleError(err)
		return
	}

	data := templateData{
		NotFollowingBack: notFollowingBack,
		NotFollowedBack:  notFollowedBack,
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		handleError(err)
		return
	}
}

func handleError(err error) {
	fmt.Printf("Content-Type: text/html\n\n")
	fmt.Printf("<html><body><p>Error: %v</p></body></html>", err)
}
