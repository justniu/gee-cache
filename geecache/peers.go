package geecache


//PeerPicker is the interface that must be implemented to locate
//the peer that owns a specfic key.
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

type PeerGetter interface {
	Get(group, key string) ([]byte, error) 
}


