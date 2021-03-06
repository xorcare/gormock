// Copyright © 2018-2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build gormock

package linker

import (
	"github.com/xorcare/gormock/comp"
)

// DB link to mock.
type DB = comp.DB

// Association type alias.
type Association = comp.Association

// Callback type alias.
type Callback = comp.Callback

// CallbackProcessor type alias.
type CallbackProcessor = comp.CallbackProcessor

// Dialect type alias.
type Dialect = comp.Dialect

// Field how to handle many2many relations.
type Field = comp.Field

// JoinTableHandlerInterface type alias.
type JoinTableHandlerInterface = comp.JoinTableHandlerInterface

// Scope type alias.
type Scope = comp.Scope

// SQLCommon type alias.
type SQLCommon = comp.SQLCommon
