package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	Time    time.Time
	PayLoad interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayLoad() interface{} {
	return e.PayLoad
}

func (e *TestEvent) GetDateTime() time.Time {
	return e.Time
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event Event, wg *sync.WaitGroup) {
}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	EventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetUpTest() {
	suite.EventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{
		ID: 1,
	}
	suite.handler2 = TestEventHandler{
		ID: 2,
	}
	suite.handler3 = TestEventHandler{
		ID: 3,
	}

	suite.event = TestEvent{Name: "teste", PayLoad: "teste"}
	suite.event = TestEvent{Name: "teste2", PayLoad: "teste2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	suite.SetUpTest()
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_WithSameHandler() {
	suite.SetUpTest()
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	suite.SetUpTest()
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))

	suite.EventDispatcher.Clear()
	suite.Equal(0, len(suite.EventDispatcher.handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has() {
	suite.SetUpTest()
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	assert.True(suite.T(), suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler))
	assert.True(suite.T(), suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler2))
	assert.False(suite.T(), suite.EventDispatcher.Has(suite.event.GetName(), &suite.handler3))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	suite.SetUpTest()
	err := suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	err = suite.EventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(1, len(suite.EventDispatcher.handlers[suite.event.GetName()]))
	assert.Equal(suite.T(), &suite.handler, suite.EventDispatcher.handlers[suite.event.GetName()][0])

	suite.EventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Equal(0, len(suite.EventDispatcher.handlers[suite.event.GetName()]))

	suite.EventDispatcher.Remove(suite.event2.GetName(), &suite.handler3)
	suite.Equal(0, len(suite.EventDispatcher.handlers[suite.event2.GetName()]))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event Event, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispatch_Dispatch() {
	suite.SetUpTest()
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)

	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event2)

	eh3 := &MockHandler{}
	eh3.On("Handle", &suite.event2)

	suite.EventDispatcher.Register(suite.event.GetName(), eh)
	suite.EventDispatcher.Register(suite.event2.GetName(), eh2)
	suite.EventDispatcher.Register(suite.event2.GetName(), eh3)

	suite.EventDispatcher.Dispatch(&suite.event)
	suite.EventDispatcher.Dispatch(&suite.event2)

	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())
	eh3.AssertExpectations(suite.T())

	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh3.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func TestSuit(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
