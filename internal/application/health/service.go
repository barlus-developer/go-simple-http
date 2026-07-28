package health

import (
	"math/rand"

	"github.com/barlus-developer/go-simple-http/internal/domain/health"
)

type Service interface {
	Status() health.Status
}

type service struct{}

func NewService() Service {
	return &service{}
}

var memeMessages = []string{
	"It's not a bug, it's a feature.",
	"200 OK, 0 worries.",
	"This server runs on coffee and hope.",
	"Works on my machine.",
	"99 little bugs in the code, take one down, patch it around, 127 little bugs in the code.",
	"I would love to change the world, but they won't give me the source code.",
	"Ping received. Existential crisis postponed.",
	"There's no place like 127.0.0.1.",
	"Status: alive, unlike my will to debug this.",
	"To err is human. To really foul things up requires a computer.",
	"Server's up. Coffee's down.",
	"Not all heroes wear capes, some just return 200.",
	"Compiles on the first try? Suspicious.",
	"This endpoint is powered by pure vibes.",
	"Uptime: yes. Downtime: also yes, eventually.",
	"I'm not saying it's magic, but nobody knows how it works.",
	"Health check passed. Existential dread: pending.",
	"Running smoothly, unlike my sleep schedule.",
	"404? Never heard of her.",
	"This service has more nines than my patience.",
	"Deployed on a Friday and still standing.",
	"CTRL+C, CTRL+V, and a prayer.",
	"It works, don't ask me why.",
	"Server status: caffeinated and functional.",
	"Green checkmarks make the world go round.",
	"I speak fluent JSON and sarcasm.",
	"Alive and pinging, just like your ex texting at 2am.",
	"Rebooted my confidence along with the server.",
	"Everything's fine, this is fine.",
	"Request received. Judging your API design silently.",
}

func (s *service) Status() health.Status {
	return health.Status{
		Status:  "ok",
		Message: memeMessages[rand.Intn(len(memeMessages))],
	}
}
