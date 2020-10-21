package simulator

import (
	"github.com/Gimulator/protobuf/go/api"
	"github.com/sirupsen/logrus"
)

type Channel struct {
	Ch       chan *api.Message
	IsClosed bool
}

type watcher struct {
	key     *api.Key
	channel *Channel
}

type spreader struct {
	watchers []watcher
	log      *logrus.Entry
}

func NewSpreader() *spreader {
	return &spreader{
		watchers: make([]watcher, 0),
		log:      logrus.WithField("entity", "spreader"),
	}
}

func (s *spreader) AddWatcher(key *api.Key, ch *Channel) error {
	s.watchers = append(s.watchers, watcher{
		key:     key,
		channel: ch,
	})

	return nil
}

func (s *spreader) Spread(message *api.Message) {
	for i := 0; i < len(s.watchers); i++ {
		w := s.watchers[i]

		if w.channel.IsClosed {
			s.watchers[i] = s.watchers[len(s.watchers)-1]
			s.watchers = s.watchers[:len(s.watchers)-1]
			i--
			continue
		}

		if s.match(w.key, message.Key) {
			select {
			case w.channel.Ch <- message:
			default:
			}
		}
	}
}

func (s *spreader) match(base, check *api.Key) bool {
	if base.Type != "" && base.Type != check.Type {
		return false
	}
	if base.Name != "" && base.Name != check.Name {
		return false
	}
	if base.Namespace != "" && base.Namespace != check.Namespace {
		return false
	}
	return true
}
