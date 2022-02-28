// Modified package from https://github.com/veggiedefender/torrent-client used for
// parsing torrent files and computing info hashes.
package torrentfile

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"

	"github.com/jackpal/bencode-go"
)

// Port to listen on
const Port uint16 = 6881

// TorrentFile encodes the metadata from a .torrent file
type TorrentFile struct {
	Announce    string
	InfoHash    []byte
	PieceHashes [][20]byte
	PieceLength int
	Length      int
	Name        string
}

type bencodeInfo struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type BencodeTorrent struct {
	Announce string      `bencode:"announce"`
	Info     bencodeInfo `bencode:"info"`
}

func Unmarshal(r io.Reader) (BencodeTorrent, error) {
	bto := BencodeTorrent{}
	err := bencode.Unmarshal(r, &bto)
	if err != nil {
		return BencodeTorrent{}, err
	}
	return bto, nil
}

func (i *bencodeInfo) Hash() ([]byte, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, i.Pieces)
	if err != nil {
		return []byte{}, err
	}
	h := sha1.Sum(buf.Bytes())
	return h[:], nil
}

func (i *bencodeInfo) splitPieceHashes() ([][20]byte, error) {
	hashLen := 20 // Length of SHA-1 hash
	buf := []byte(i.Pieces)
	if len(buf)%hashLen != 0 {
		err := fmt.Errorf("received malformed pieces of length %d", len(buf))
		return nil, err
	}
	numHashes := len(buf) / hashLen
	hashes := make([][20]byte, numHashes)

	for i := 0; i < numHashes; i++ {
		copy(hashes[i][:], buf[i*hashLen:(i+1)*hashLen])
	}
	return hashes, nil
}

func (bto *BencodeTorrent) ToTorrentFile() (TorrentFile, error) {
	infoHash, err := bto.Info.Hash()
	if err != nil {
		return TorrentFile{}, err
	}
	pieceHashes, err := bto.Info.splitPieceHashes()
	if err != nil {
		return TorrentFile{}, err
	}
	t := TorrentFile{
		Announce:    bto.Announce,
		InfoHash:    infoHash,
		PieceHashes: pieceHashes,
		PieceLength: bto.Info.PieceLength,
		Length:      bto.Info.Length,
		Name:        bto.Info.Name,
	}
	return t, nil
}
