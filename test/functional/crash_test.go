package functional

import (
	"syscall"
	"testing"

	api "github.com/capsule8/api/v0"

	"github.com/golang/glog"
)

type crashTest struct {
	testContainer   *container
	err             error
	containerID     string
	containerExited bool
	processID       string
	processExited   bool
}

func (ct *crashTest) buildContainer(t *testing.T) {
	c := newContainer(t, "crash")
	err := c.build()
	if err != nil {
		t.Error(err)
	} else {
		ct.testContainer = c
	}
}

func (ct *crashTest) runContainer(t *testing.T) {
	err := ct.testContainer.start()
	if err != nil {
		t.Error(err)
		return
	}

	// We assume that the container will return an error, so ignore that one
	ct.testContainer.wait()
}

func (ct *crashTest) createSubscription(t *testing.T) *api.Subscription {
	containerEvents := []*api.ContainerEventFilter{
		&api.ContainerEventFilter{
			Type: api.ContainerEventType_CONTAINER_EVENT_TYPE_CREATED,
		},
		&api.ContainerEventFilter{
			Type: api.ContainerEventType_CONTAINER_EVENT_TYPE_RUNNING,
		},
		&api.ContainerEventFilter{
			Type: api.ContainerEventType_CONTAINER_EVENT_TYPE_EXITED,
		},
	}

	processEvents := []*api.ProcessEventFilter{
		&api.ProcessEventFilter{
			Type: api.ProcessEventType_PROCESS_EVENT_TYPE_FORK,
		},

		&api.ProcessEventFilter{
			Type: api.ProcessEventType_PROCESS_EVENT_TYPE_EXEC,
		},

		&api.ProcessEventFilter{
			Type: api.ProcessEventType_PROCESS_EVENT_TYPE_EXIT,
		},
	}

	eventFilter := &api.EventFilter{
		ContainerEvents: containerEvents,
		ProcessEvents:   processEvents,
	}

	sub := &api.Subscription{
		EventFilter: eventFilter,
	}

	return sub
}

func (ct *crashTest) handleTelemetryEvent(t *testing.T, telemetryEvent *api.TelemetryEvent) bool {
	glog.V(2).Infof("%+v", telemetryEvent)

	switch event := telemetryEvent.Event.Event.(type) {
	case *api.Event_Container:
		if event.Container.Type == api.ContainerEventType_CONTAINER_EVENT_TYPE_CREATED {
			if event.Container.ImageName == ct.testContainer.imageID {
				if len(ct.containerID) > 0 {
					t.Error("Already saw container created")
					return false
				}

				ct.containerID = telemetryEvent.Event.ContainerId
				glog.V(1).Infof("containerID = %s", ct.containerID)
			}
		} else if len(ct.containerID) > 0 &&
			telemetryEvent.Event.ContainerId == ct.containerID &&
			event.Container.Type == api.ContainerEventType_CONTAINER_EVENT_TYPE_EXITED {

			if event.Container.ExitCode != 139 {
				t.Errorf("Expected ExitStatus %d, got %d",
					139, event.Container.ExitStatus)
				return false
			}

			ct.containerExited = true
			glog.V(1).Infof("containerExited = true")
		}

	case *api.Event_Process:
		if event.Process.Type == api.ProcessEventType_PROCESS_EVENT_TYPE_EXEC {
			if event.Process.ExecFilename == "/main" &&
				telemetryEvent.Event.ContainerId == ct.containerID {
				if len(ct.processID) > 0 {
					t.Error("Already saw process exec")
					return false
				}

				ct.processID = telemetryEvent.Event.ProcessId
				glog.V(1).Infof("processID = %s", ct.processID)
			}
		} else if len(ct.processID) > 0 &&
			telemetryEvent.Event.ProcessId == ct.processID &&
			event.Process.Type == api.ProcessEventType_PROCESS_EVENT_TYPE_EXIT {

			if event.Process.ExitStatus != 0 {
				t.Errorf("Expected ExitStatus %d, got %d",
					0, event.Process.ExitStatus)
				return false
			}

			if event.Process.ExitSignal != uint32(syscall.SIGSEGV) {
				t.Errorf("Expected ExitSignal %d, got %d",
					syscall.SIGSEGV, event.Process.ExitSignal)
				return false
			}

			if event.Process.ExitCoreDumped != true {
				t.Errorf("Expected ExitCoreDumped %v, got %v",
					true, event.Process.ExitCoreDumped)
				return false
			}

			ct.processExited = true
			glog.V(1).Infof("processExited = true")
		}
	}

	return !(ct.containerExited && ct.processExited)
}

func TestCrash(t *testing.T) {
	ct := &crashTest{}
	tt := newTelemetryTest(ct)
	tt.runTest(t)
}