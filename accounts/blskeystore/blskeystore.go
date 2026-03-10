// Package blskeystore implements encrypted storage for BLS secret keys.
//
// Keys are encrypted using the Web3 Secret Storage format (scrypt + AES-128-CTR),
// reusing the encryption primitives from the ECDSA keystore package.
package blskeystore

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/google/uuid"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls"
	"github.com/prysmaticlabs/prysm/v5/crypto/bls/common"
)

const version = 1

// BLSKey represents a decrypted BLS key in memory.
type BLSKey struct {
	Id        uuid.UUID
	PublicKey common.PublicKey
	SecretKey common.SecretKey
}

// KeyInfo holds metadata about a stored BLS key.
type KeyInfo struct {
	PubKey string // hex-encoded compressed public key (48 bytes)
	Path   string // file path
	Id     string // uuid
}

// encryptedBLSKeyJSON is the on-disk encrypted format.
type encryptedBLSKeyJSON struct {
	PubKey  string              `json:"pubkey"`
	Crypto  keystore.CryptoJSON `json:"crypto"`
	Id      string              `json:"id"`
	Version int                 `json:"version"`
}

// EncryptBLSKey encrypts a BLS key with the given password.
func EncryptBLSKey(key *BLSKey, auth string, scryptN, scryptP int) ([]byte, error) {
	secretBytes := key.SecretKey.Marshal()
	cryptoStruct, err := keystore.EncryptDataV3(secretBytes, []byte(auth), scryptN, scryptP)
	if err != nil {
		return nil, err
	}
	encJSON := encryptedBLSKeyJSON{
		PubKey:  hex.EncodeToString(key.PublicKey.Marshal()),
		Crypto:  cryptoStruct,
		Id:      key.Id.String(),
		Version: version,
	}
	return json.Marshal(encJSON)
}

// DecryptBLSKey decrypts a BLS key from its JSON blob.
func DecryptBLSKey(keyjson []byte, auth string) (*BLSKey, error) {
	k := new(encryptedBLSKeyJSON)
	if err := json.Unmarshal(keyjson, k); err != nil {
		return nil, err
	}
	if k.Version != version {
		return nil, fmt.Errorf("unsupported BLS key version: %d", k.Version)
	}
	keyBytes, err := keystore.DecryptDataV3(k.Crypto, auth)
	if err != nil {
		return nil, err
	}
	defer func() {
		// wipe keyBytes
		for i := range keyBytes {
			keyBytes[i] = 0
		}
	}()
	secretKey, err := bls.SecretKeyFromBytes(keyBytes)
	if err != nil {
		return nil, err
	}
	id, err := uuid.Parse(k.Id)
	if err != nil {
		return nil, err
	}
	return &BLSKey{
		Id:        id,
		PublicKey: secretKey.PublicKey(),
		SecretKey: secretKey,
	}, nil
}

// ReadBLSPubKey reads a BLS keystore JSON file and extracts the public key without decryption.
func ReadBLSPubKey(filename string) ([]byte, string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", err
	}
	k := new(encryptedBLSKeyJSON)
	if err := json.Unmarshal(data, k); err != nil {
		return nil, "", err
	}
	pubBytes, err := hex.DecodeString(k.PubKey)
	if err != nil {
		return nil, "", err
	}
	return pubBytes, k.Id, nil
}

// StoreKey encrypts a BLS secret key and writes it to dir.
func StoreKey(dir string, secretKey common.SecretKey, auth string, scryptN, scryptP int) (*BLSKey, string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, "", err
	}
	key := &BLSKey{
		Id:        id,
		PublicKey: secretKey.PublicKey(),
		SecretKey: secretKey,
	}
	keyjson, err := EncryptBLSKey(key, auth, scryptN, scryptP)
	if err != nil {
		return nil, "", err
	}
	filename := filepath.Join(dir, keyFileName(key.PublicKey))
	if err := writeKeyFile(filename, keyjson); err != nil {
		return nil, "", err
	}
	// Verify we can decrypt what we wrote
	verifyData, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", err
	}
	if _, err := DecryptBLSKey(verifyData, auth); err != nil {
		return nil, "", fmt.Errorf("keystore verification failed: %v", err)
	}
	return key, filename, nil
}

// UpdateKey re-encrypts a BLS key file with a new password.
// The old password is used for decryption, and the new password for encryption.
func UpdateKey(filename, oldAuth, newAuth string, scryptN, scryptP int) error {
	keyjson, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	key, err := DecryptBLSKey(keyjson, oldAuth)
	if err != nil {
		return err
	}
	newJSON, err := EncryptBLSKey(key, newAuth, scryptN, scryptP)
	if err != nil {
		return err
	}
	return writeKeyFile(filename, newJSON)
}

// ListKeys lists all BLS key files in the given directory.
func ListKeys(dir string) ([]KeyInfo, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var keys []KeyInfo
	for _, entry := range entries {
		if entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		pubBytes, id, err := ReadBLSPubKey(path)
		if err != nil {
			continue
		}
		keys = append(keys, KeyInfo{
			PubKey: hex.EncodeToString(pubBytes),
			Path:   path,
			Id:     id,
		})
	}
	return keys, nil
}

// FindKeyByPubKey finds a key file in the directory by public key.
func FindKeyByPubKey(dir, pubKeyHex string) (string, error) {
	pubKeyHex = strings.TrimPrefix(pubKeyHex, "0x")
	pubKeyHex = strings.ToLower(pubKeyHex)

	entries, err := os.ReadDir(dir)
	if err != nil {
		return "", err
	}
	for _, entry := range entries {
		if entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		pubBytes, _, err := ReadBLSPubKey(path)
		if err != nil {
			continue
		}
		fileHex := hex.EncodeToString(pubBytes)
		if strings.HasPrefix(fileHex, pubKeyHex) || fileHex == pubKeyHex {
			return path, nil
		}
	}
	return "", fmt.Errorf("no BLS key found matching pubkey prefix %q", pubKeyHex)
}

// keyFileName implements the naming convention for keyfiles:
// UTC--<created_at UTC ISO8601>-<pubkey hex>
func keyFileName(pubKey common.PublicKey) string {
	ts := time.Now().UTC()
	pubHex := hex.EncodeToString(pubKey.Marshal())
	return fmt.Sprintf("UTC--%s--%s", toISO8601(ts), pubHex)
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}

func writeKeyFile(file string, content []byte) error {
	const dirPerm = 0700
	if err := os.MkdirAll(filepath.Dir(file), dirPerm); err != nil {
		return err
	}
	f, err := os.CreateTemp(filepath.Dir(file), "."+filepath.Base(file)+".tmp")
	if err != nil {
		return err
	}
	if _, err := f.Write(content); err != nil {
		f.Close()
		os.Remove(f.Name())
		return err
	}
	f.Close()
	return os.Rename(f.Name(), file)
}
