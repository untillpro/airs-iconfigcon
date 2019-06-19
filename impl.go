/*
 * Copyright (c) 2019-present unTill Pro, Ltd. and Contributors
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package iconfigcon

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
)

// KVPair is used to represent a single K/V entry
type ConsulEntry struct {
	// Key is the name of the key. It is also part of the URL path when accessed
	// via the API.
	Key string

	// CreateIndex holds the index corresponding the creation of this KVPair. This
	// is a read-only field.
	CreateIndex uint64

	// ModifyIndex is used for the Check-And-Set operations and can also be fed
	// back into the WaitIndex of the QueryOptions in order to perform blocking
	// queries.
	ModifyIndex uint64

	// LockIndex holds the index corresponding to a lock on this key, if any. This
	// is a read-only field.
	LockIndex uint64

	// Flags are any user-defined flags on the key. It is up to the implementer
	// to check these values, since Consul does not treat them specially.
	Flags uint64

	// Value is the value for the key. This can be any value, but it will be
	// base64 encoded upon transport.
	Value []byte

	// Session is a string representing the ID of the session. Any other
	// interactions with this key over the same session must specify the same
	// session ID.
	Session string
}

func getConfig(ctx context.Context, configName string, config interface{}) (ok bool, err error) {
	rv := reflect.ValueOf(config)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return false, fmt.Errorf("%s must be a pointer", reflect.ValueOf(config))
	}
	ok, err = getService(ctx).getConfigHTTP(config, configName)
	return ok, err
}

func putConfig(ctx context.Context, configName string, config interface{}) error {
	if reflect.ValueOf(config).IsNil() {
		return errors.New("config must not be nil")
	}
	err := getService(ctx).putConfigHTTP(config, configName)
	return err
}

func (s *Service) getConfigHTTP(config interface{}, prefix string) (ok bool, err error) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/v1/kv/%s", s.Host, s.Port, prefix))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 404 {
		return false, nil
	} else if resp.StatusCode != 200 {
		return false, fmt.Errorf("unexpected response code: %d", resp.StatusCode)
	}
	err = decodeBody(resp, config)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) putConfigHTTP(config interface{}, prefix string) error {
	body, err := encodeBody(config)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("http://%s:%d/v1/kv/%s", s.Host, s.Port,
		prefix), body)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func decodeBody(resp *http.Response, value interface{}) error {
	var entry []*ConsulEntry
	dec := json.NewDecoder(resp.Body)
	err := dec.Decode(&entry)
	if err != nil {
		return err
	}
	err = json.Unmarshal(entry[0].Value, &value)
	if err != nil {
		return err
	}
	return nil
}

func encodeBody(value interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(value); err != nil {
		return nil, err
	}
	return buf, nil
}
