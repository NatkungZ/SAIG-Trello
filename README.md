Trello webhook &amp; proxy
Trello Go for fb-chat proxy
docker run -d -p 8080:8080 --name go-chat-proxy -e GO_PORT=9000 -e GO_HOST="http://chat.com" trello-proxy
