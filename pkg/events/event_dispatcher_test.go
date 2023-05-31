package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event          TestEvent
	event2         TestEvent
	handler        TestEventHandler
	handler2       TestEventHandler
	handler3       TestEventHandler
	eventDispacher *EventDispacher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispacher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}
	suite.event = TestEvent{Name: "test", Payload: "test"}
	suite.event2 = TestEvent{Name: "test2", Payload: "test2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	suite.Equal(&suite.handler, suite.eventDispacher.handlers[suite.event.GetName()][0])
	suite.Equal(&suite.handler2, suite.eventDispacher.handlers[suite.event.GetName()][1])

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event2.GetName()]))

	suite.Equal(2, len(suite.eventDispacher.handlers))

	suite.eventDispacher.Clear()
	suite.Equal(0, len(suite.eventDispacher.handlers))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	suite.True(suite.eventDispacher.Has(suite.event.GetName(), &suite.handler))
	suite.True(suite.eventDispacher.Has(suite.event.GetName(), &suite.handler2))
	suite.False(suite.eventDispacher.Has(suite.event.GetName(), &suite.handler3))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)
	suite.eventDispacher.Register(suite.event.GetName(), eh)
	suite.eventDispacher.Dispatch(&suite.event)
	eh.AssertExpectations(suite.T())
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := suite.eventDispacher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	err = suite.eventDispacher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event2.GetName()]))

	suite.eventDispacher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(1, len(suite.eventDispacher.handlers[suite.event.GetName()]))
	suite.Equal(&suite.handler2, suite.eventDispacher.handlers[suite.event.GetName()][0])

	suite.eventDispacher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(0, len(suite.eventDispacher.handlers[suite.event.GetName()]))

	suite.eventDispacher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.eventDispacher.handlers[suite.event2.GetName()]))

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
