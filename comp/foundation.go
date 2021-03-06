// Copyright © 2018-2019, Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package comp

// Foundation an interface describing methods that implement all the basic
// processing of called methods.
type Foundation interface {
	// Handler handles calls that do not require return values.
	Handler(
		receiver interface{},
		methodName string,
		arguments ...interface{},
	)
	// HandlerOther handles calls that return a single result value.
	HandlerOther(
		receiver interface{},
		methodName string,
		outcome interface{},
		arguments ...interface{},
	) (result interface{})
	// HandlerOthers handles calls that return two result values.
	HandlerOthers(
		receiver interface{},
		methodName string,
		outcomeFirst, outcomeSecond interface{},
		arguments ...interface{},
	) (resultFirst, resultSecond interface{})
}
