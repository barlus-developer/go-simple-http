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
	"It's not a bug, it's a feature. 😏",
	"200 OK, 0 worries, 1000 regrets. 😂",
	"This server runs on coffee, hope, and copium. ☕😩",
	"Works on my machine. Emphasis on MY. 🤷‍♂️",
	"99 little bugs in the code... jk it's like 4000. 🐛🐛🐛",
	"I would love to change the world, but I forgot my password. 🔐😅",
	"Ping received. Please hold, having an existential crisis. 🫠",
	"There's no place like 127.0.0.1, mostly because I got lost everywhere else. 🏠💻",
	"Status: alive. Will to debug: deceased. 💀",
	"To err is human, to really foul things up requires prod access. 🙃🔥",
	"Server's up, coffee's down, dignity: TBD. ☕📉",
	"Not all heroes wear capes, some just return 200 and go back to sleep. 🦸😴",
	"Compiled on the first try. Who ratted me out. 🕵️‍♂️",
	"This endpoint runs on vibes, duct tape, and one Stack Overflow answer from 2013. 🧵✨",
	"Uptime: yes. Downtime: also yes, ask again later. ⏳🤡",
	"I'm not saying it's magic, I'm saying I deleted the logs so nobody can prove otherwise. 🪄🙈",
	"Health check passed. My personal health, questionable. 🩺😬",
	"Running smoothly, unlike my sleep schedule or my life choices. 😵‍💫💤",
	"404? Never heard of her, we don't talk anymore. 🚫👀",
	"This service has more nines than my patience with your API calls. 9️⃣😤",
	"Deployed on a Friday and somehow still standing, unlike my social life. 🎉💀",
	"CTRL+C, CTRL+V, a prayer, and mild regret. 🙏😬",
	"It works, don't ask me why, I will lie to you. 🤥",
	"Server status: caffeinated, functional, mildly unhinged. ☕🤪",
	"Green checkmarks make the world go round, red ones make me go home. ✅🏃‍♂️",
	"I speak fluent JSON, sarcasm, and passive-aggressive commit messages. 🗣️💬",
	"Alive and pinging, just like your ex texting 'u up' at 2am. 📱🌙",
	"Rebooted my confidence along with the server, both still loading. 🔄😐",
	"Everything's fine, this is fine, the server room is on fire but it's FINE. 🔥🐶☕",
	"Request received. Silently judging your API design and your life choices. 👀🤫",
	"200 OK but emotionally I'm a 500. 😭",
	"I ate a whole bug for breakfast and still deployed on time. 🐛🍳",
	"Server's fine. It's the developers who are held together by duct tape. 🧑‍💻🩹",
	"Latency: low. Existential dread: high. 📉😱",
	"This is not a drill, it's just Tuesday. 🚨📅",
}

func (s *service) Status() health.Status {
	return health.Status{
		Status:  "ok",
		Message: memeMessages[rand.Intn(len(memeMessages))],
	}
}
