package helpers
import (
	"encoding/json"
)

func jsonify_err(v any) (map[string]any, error) {
    b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	out := map[string]any{}
	err = json.Unmarshal(b, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func Jsonify(v any) map[string]any {
	b, err := jsonify_err(v)
	if err != nil {
		return nil
	}
	return b
}
