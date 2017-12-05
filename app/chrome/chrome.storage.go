package chrome

import (
	"app/chrome/protocol"
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

/*
Storage - https://chromedevtools.github.io/devtools-protocol/tot/Storage/
EXPERIMENTAL
*/
type Storage struct{}

/*
ClearDataForOrigin clears storage for origin.
*/
func (Storage) ClearDataForOrigin(socket *Socket, params *storage.ClearDataForOriginParams) error {
	command := &protocol.Command{
		method: "storage.clearDataForOrigin",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
GetUsageAndQuota returns usage and quota in bytes.
*/
func (Storage) GetUsageAndQuota(socket *Socket, params *storage.GetUsageAndQuotaParams) error {
	command := &protocol.Command{
		method: "storage.getUsageAndQuota",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
TrackCacheStorageForOrigin registers origin to be notified when an update occurs to its cache
storage list.
*/
func (Storage) TrackCacheStorageForOrigin(socket *Socket, params *storage.TrackCacheStorageForOriginParams) error {
	command := &protocol.Command{
		method: "storage.trackCacheStorageForOrigin",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
TrackIndexedDBForOrigin registers origin to be notified when an update occurs to its IndexedDB.
*/
func (Storage) TrackIndexedDBForOrigin(socket *Socket, params *storage.TrackIndexedDBForOriginParams) error {
	command := &protocol.Command{
		method: "storage.trackIndexedDBForOrigin",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
UntrackCacheStorageForOrigin unregisters origin from receiving notifications for cache storage.
*/
func (Storage) UntrackCacheStorageForOrigin(socket *Socket, params *storage.UntrackCacheStorageForOriginParams) error {
	command := &protocol.Command{
		method: "storage.untrackCacheStorageForOrigin",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
UntrackIndexedDBForOrigin unregisters origin from receiving notifications for IndexedDB.
*/
func (Storage) UntrackIndexedDBForOrigin(socket *Socket, params *storage.UntrackIndexedDBForOriginParams) error {
	command := &protocol.Command{
		method: "storage.untrackIndexedDBForOrigin",
		params: params,
	}
	socket.SendCommand(command)
	return command.Err
}

/*
OnCacheStorageContentUpdated adds a handler to the Storage.cacheStorageContentUpdated event.
Storage.cacheStorageContentUpdated fires when a cache's contents have been modified.
*/
func (Storage) OnCacheStorageContentUpdated(socket *Socket, callback func(event *storage.CacheStorageContentUpdatedEvent)) error {
	handler := protocol.NewEventHandler(
		"Storage.cacheStorageContentUpdated",
		func(name string, params []byte) {
			event := &storage.CacheStorageContentUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}

/*
OnCacheStorageListUpdated adds a handler to the Storage.cacheStorageListUpdated event.
Storage.cacheStorageListUpdated fires when cache has been added/deleted.
*/
func (Storage) OnCacheStorageListUpdated(socket *Socket, callback func(event *storage.CacheStorageListUpdatedEvent)) error {
	handler := protocol.NewEventHandler(
		"Storage.cacheStorageListUpdated",
		func(name string, params []byte) {
			event := &storage.CacheStorageListUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}

/*
OnIndexedDBContentUpdated adds a handler to the Storage.indexedDBContentUpdated event.
Storage.indexedDBContentUpdated fires when the origin's IndexedDB object store has been modified.
*/
func (Storage) OnIndexedDBContentUpdated(socket *Socket, callback func(event *storage.IndexedDBContentUpdatedEvent)) error {
	handler := protocol.NewEventHandler(
		"Storage.indexedDBContentUpdated",
		func(name string, params []byte) {
			event := &storage.IndexedDBContentUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}

/*
OnIndexedDBListUpdated adds a handler to the Storage.indexedDBListUpdated event.
Storage.indexedDBListUpdated fires when the origin's IndexedDB database list has been modified.
*/
func (Storage) OnIndexedDBListUpdated(socket *Socket, callback func(event *storage.IndexedDBListUpdatedEvent)) error {
	handler := protocol.NewEventHandler(
		"Storage.indexedDBListUpdated",
		func(name string, params []byte) {
			event := &storage.IndexedDBListUpdatedEvent{}
			if err := json.Unmarshal(params, event); err != nil {
				log.Error(err)
			} else {
				callback(event)
			}
		},
	)
	socket.AddEventHandler(handler)
	return command.Err
}