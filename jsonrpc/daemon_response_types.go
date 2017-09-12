package jsonrpc

import (
	"encoding/json"
	"github.com/go-errors/errors"
	"reflect"

	lbryschema "github.com/lbryio/lbryschema.go/pb"
)

type Support struct {
	Amount float64 `json:"amount"`
	Nout   int     `json:"nout"`
	Txid   string  `json:"txid"`
}

type Claim struct {
	Address          string           `json:"address"`
	Amount           float64          `json:"amount"`
	ClaimID          string           `json:"claim_id"`
	ClaimSequence    int              `json:"claim_sequence"`
	DecodedClaim     bool             `json:"decoded_claim"`
	Depth            int              `json:"depth"`
	EffectiveAmount  float64          `json:"effective_amount"`
	Height           int              `json:"height"`
	Hex              string           `json:"hex"`
	Name             string           `json:"name"`
	Nout             int              `json:"nout"`
	Supports         []Support        `json:"supports"`
	Txid             string           `json:"txid"`
	ValidAtHeight    int              `json:"valid_at_height"`
	Value            lbryschema.Claim `json:"value"`
	Error            *string          `json:"error,omitempty"`
	ChannelName      *string          `json:"channel_name,omitempty"`
	HasSignature     *bool            `json:"has_signature,omitempty"`
	SignatureIsValid *bool            `json:"signature_is_valid,omitempty"`
}

type File struct {
	ClaimID           string            `json:"claim_id"`
	Completed         bool              `json:"completed"`
	DownloadDirectory string            `json:"download_directory"`
	DownloadPath      string            `json:"download_path"`
	FileName          string            `json:"file_name"`
	Key               string            `json:"key"`
	Message           string            `json:"message"`
	Metadata          *lbryschema.Claim `json:"metadata"`
	MimeType          string            `json:"mime_type"`
	Name              string            `json:"name"`
	Outpoint          string            `json:"outpoint"`
	PointsPaid        float64           `json:"points_paid"`
	SdHash            string            `json:"sd_hash"`
	Stopped           bool              `json:"stopped"`
	StreamHash        string            `json:"stream_hash"`
	StreamName        string            `json:"stream_name"`
	SuggestedFileName string            `json:"suggested_file_name"`
	TotalBytes        uint64            `json:"total_bytes"`
	WrittenBytes      uint64            `json:"written_bytes"`
	ChannelName       *string           `json:"channel_name,omitempty"`
	HasSignature      *bool             `json:"has_signature,omitempty"`
	SignatureIsValid  *bool             `json:"signature_is_valid,omitempty"`
}

func getEnumVal(enum map[string]int32, data interface{}) (int32, error) {
	s, ok := data.(string)
	if !ok {
		return 0, errors.New("expected a string")
	}
	val, ok := enum[s]
	if !ok {
		return 0, errors.New("invalid enum key")
	}
	return val, nil
}

func fixDecodeProto(src, dest reflect.Type, data interface{}) (interface{}, error) {
	switch dest {
	case reflect.TypeOf(uint64(0)):
		if n, ok := data.(json.Number); ok {
			val, err := n.Int64()
			if err != nil {
				return nil, err
			} else if val < 0 {
				return nil, errors.New("must be unsigned int")
			}
			return uint64(val), nil
		}
	case reflect.TypeOf([]byte{}):
		if s, ok := data.(string); ok {
			return []byte(s), nil
		}
	case reflect.TypeOf(lbryschema.Metadata_Version(0)):
		val, err := getEnumVal(lbryschema.Metadata_Version_value, data)
		return lbryschema.Metadata_Version(val), err
	case reflect.TypeOf(lbryschema.Metadata_Language(0)):
		val, err := getEnumVal(lbryschema.Metadata_Language_value, data)
		return lbryschema.Metadata_Language(val), err

	case reflect.TypeOf(lbryschema.Stream_Version(0)):
		val, err := getEnumVal(lbryschema.Stream_Version_value, data)
		return lbryschema.Stream_Version(val), err

	case reflect.TypeOf(lbryschema.Claim_Version(0)):
		val, err := getEnumVal(lbryschema.Claim_Version_value, data)
		return lbryschema.Claim_Version(val), err
	case reflect.TypeOf(lbryschema.Claim_ClaimType(0)):
		val, err := getEnumVal(lbryschema.Claim_ClaimType_value, data)
		return lbryschema.Claim_ClaimType(val), err

	case reflect.TypeOf(lbryschema.Fee_Version(0)):
		val, err := getEnumVal(lbryschema.Fee_Version_value, data)
		return lbryschema.Fee_Version(val), err
	case reflect.TypeOf(lbryschema.Fee_Currency(0)):
		val, err := getEnumVal(lbryschema.Fee_Currency_value, data)
		return lbryschema.Fee_Currency(val), err

	case reflect.TypeOf(lbryschema.Source_Version(0)):
		val, err := getEnumVal(lbryschema.Source_Version_value, data)
		return lbryschema.Source_Version(val), err
	case reflect.TypeOf(lbryschema.Source_SourceTypes(0)):
		val, err := getEnumVal(lbryschema.Source_SourceTypes_value, data)
		return lbryschema.Source_SourceTypes(val), err

	case reflect.TypeOf(lbryschema.KeyType(0)):
		val, err := getEnumVal(lbryschema.KeyType_value, data)
		return lbryschema.KeyType(val), err

	case reflect.TypeOf(lbryschema.Signature_Version(0)):
		val, err := getEnumVal(lbryschema.Signature_Version_value, data)
		return lbryschema.Signature_Version(val), err

	case reflect.TypeOf(lbryschema.Certificate_Version(0)):
		val, err := getEnumVal(lbryschema.Certificate_Version_value, data)
		return lbryschema.Certificate_Version(val), err
	}

	return data, nil
}

type CommandsResponse []string

type WalletBalanceResponse float64

type VersionResponse struct {
	Build             string `json:"build"`
	LbrynetVersion    string `json:"lbrynet_version"`
	LbryschemaVersion string `json:"lbryschema_version"`
	LbryumVersion     string `json:"lbryum_version"`
	OsRelease         string `json:"os_release"`
	OsSystem          string `json:"os_system"`
	Platform          string `json:"platform"`
	Processor         string `json:"processor"`
	PythonVersion     string `json:"python_version"`
}
type StatusResponse struct {
	BlockchainStatus struct {
		BestBlockhash string `json:"best_blockhash"`
		Blocks        int    `json:"blocks"`
		BlocksBehind  int    `json:"blocks_behind"`
	} `json:"blockchain_status"`
	BlocksBehind     int `json:"blocks_behind"`
	ConnectionStatus struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"connection_status"`
	InstallationID string `json:"installation_id"`
	IsFirstRun     bool   `json:"is_first_run"`
	IsRunning      bool   `json:"is_running"`
	LbryID         string `json:"lbry_id"`
	StartupStatus  struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"startup_status"`
}

type ClaimListResponse struct {
	Claims                []Claim   `json:"claims"`
	LastTakeoverHeight    int       `json:"last_takeover_height"`
	SupportsWithoutClaims []Support `json:"supports_without_claims"`
}

type ClaimShowResponse Claim

type PeerListResponsePeer struct {
	IP          string
	Port        uint
	IsAvailable bool
}
type PeerListResponse []PeerListResponsePeer

type BlobGetResponse struct {
	Blobs []struct {
		BlobHash string `json:"blob_hash,omitempty"`
		BlobNum  int    `json:"blob_num"`
		IV       string `json:"iv"`
		Length   int    `json:"length"`
	} `json:"blobs"`
	Key               string `json:"key"`
	StreamHash        string `json:"stream_hash"`
	StreamName        string `json:"stream_name"`
	StreamType        string `json:"stream_type"`
	SuggestedFileName string `json:"suggested_file_name"`
}

type StreamCostEstimateResponse *float64

type GetResponse File
type FileListResponse []File

type ResolveResponse map[string]ResolveResponseItem
type ResolveResponseItem struct {
	Certificate     *Claim  `json:"certificate,omitempty"`
	Claim           *Claim  `json:"claim,omitempty"`
	ClaimsInChannel *uint64 `json:"claims_in_channel,omitempty"`
}
