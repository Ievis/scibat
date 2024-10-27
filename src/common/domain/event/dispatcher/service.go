package dispatcher

import "scibat/src/common/domain/event"

type Service struct {
	EventList []event.DomainEvent
}

func (s *Service) Record(domainEvent event.DomainEvent) {
	s.EventList = append(s.EventList, domainEvent)
}

func (s *Service) Dispatch() {
	eventList := s.EventList
	s.Clear()
	for i := 0; i < len(eventList); i++ {
		eventList[i].Command.Execute()
	}
}

func (s *Service) Clear() {
	s.EventList = make([]event.DomainEvent, 0)
}
