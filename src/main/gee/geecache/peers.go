package geecache

type PeerPicker interface {
	PickPeer(string) (PeerGetter, bool)
}

type PeerGetter interface {
	Get(string, string) ([]byte, error)
}
