/**
* (C) Copyright IBM Corp. 2023, 2025.
* (C) Copyright Rocket Software, Inc. 2023-2025.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package cache

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/IBM/RMF/grafana/rmf-app/pkg/plugin/frame"
	"github.com/VictoriaMetrics/fastcache"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type ChannelCache struct {
	cache *fastcache.Cache
}

type Channel struct {
	Resource  string
	TimeRange backend.TimeRange
	Absolute  bool
	Step      time.Duration
	Interval  time.Duration
	Span      time.Duration
	Fields    frame.SeriesFields
	Mintime   time.Duration
	sentTimes map[int64]struct{}

	//do not serialize it
	sender *backend.StreamSender `json:"-"`
}

func (c *Channel) SetSender(sender *backend.StreamSender) {
	c.sender = sender
}

func (c *Channel) Send(dataFrame *data.Frame) error {
	if c.sender == nil {
		return errors.New("sender is not set")
	}
	err := c.sender.SendFrame(dataFrame, data.IncludeAll)
	if err == nil {
		frame.ForEachTime(dataFrame, func(index int, t time.Time) bool {
			c.MarkSent(t)
			return true
		})
	}
	return err
}

func (c *Channel) MarkSent(t time.Time) {
	if c.sentTimes == nil {
		c.sentTimes = make(map[int64]struct{})
	}
	c.sentTimes[t.UnixMilli()] = struct{}{}
}

func (c *Channel) HasSent(t time.Time) bool {
	if c.sentTimes == nil {
		return false
	}
	_, exists := c.sentTimes[t.UnixMilli()]
	return exists
}

func NewChannelCache(size int) *ChannelCache {
	return &ChannelCache{cache: fastcache.New(size * 1024 * 1024)}
}

func (cc *ChannelCache) Reset() {
	cc.cache.Reset()
}

func (cc *ChannelCache) Get(path string) (*Channel, error) {
	var c Channel
	b := cc.cache.Get(nil, []byte(path))
	err := json.Unmarshal(b, &c)
	return &c, err
}

func (cc *ChannelCache) Set(path string, c *Channel) error {
	b, err := json.Marshal(*c)
	if err != nil {
		return err
	}
	cc.cache.Set([]byte(path), b)
	return nil
}
