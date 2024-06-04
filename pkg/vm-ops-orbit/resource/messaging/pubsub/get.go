package messaging

import (
	"github.com/taubyte/go-sdk/errno"
	messaging "github.com/taubyte/tau/core/services/substrate/components/pubsub"
)

func (f *MessagingPubSub) GetCaller(resourceId uint32) (messaging.Channel, errno.Error) {
	resource, err := f.GetResource(resourceId)
	if err != 0 {
		return nil, err
	}

	f.callersLock.Lock()
	defer f.callersLock.Unlock()

	message, ok := f.callers[resourceId]
	if !ok {
		message, ok = resource.Caller.(messaging.Channel)
		if !ok {
			return nil, errno.SmartOpErrorResourceNotFound
		}

		f.callers[resourceId] = message
	}

	return message, 0
}
