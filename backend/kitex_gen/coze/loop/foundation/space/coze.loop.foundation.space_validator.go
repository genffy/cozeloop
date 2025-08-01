// Code generated by Validator v0.2.6. DO NOT EDIT.

package space

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *GetSpaceRequest) IsValid() error {
	if p.Base != nil {
		if err := p.Base.IsValid(); err != nil {
			return fmt.Errorf("field Base not valid, %w", err)
		}
	}
	return nil
}
func (p *GetSpaceResponse) IsValid() error {
	if p.Space != nil {
		if err := p.Space.IsValid(); err != nil {
			return fmt.Errorf("field Space not valid, %w", err)
		}
	}
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("field BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *ListUserSpaceRequest) IsValid() error {
	if p.PageSize != nil {
		if *p.PageSize <= int32(0) {
			return fmt.Errorf("field PageSize gt rule failed, current value: %v", *p.PageSize)
		}
		if *p.PageSize > int32(100) {
			return fmt.Errorf("field PageSize le rule failed, current value: %v", *p.PageSize)
		}
	}
	if p.PageNumber != nil {
		if *p.PageNumber <= int32(0) {
			return fmt.Errorf("field PageNumber gt rule failed, current value: %v", *p.PageNumber)
		}
	}
	if p.Base != nil {
		if err := p.Base.IsValid(); err != nil {
			return fmt.Errorf("field Base not valid, %w", err)
		}
	}
	return nil
}
func (p *ListUserSpaceResponse) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("field BaseResp not valid, %w", err)
		}
	}
	return nil
}
