// Copyright 2015 Osipov Konstantin <k.osipov.msk@gmail.com>. All rights reserved.
// license that can be found in the LICENSE file.

// This file is part of the application source code leveldb-cli
// This software provides a console interface to leveldb.

package cliutil

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/TomiHiltunen/geohash-golang"
	"gopkg.in/mgo.v2/bson"
	"math"
)

// Converts data to a string
func ToString(format string, value []byte) string {
	switch format {
	case "bson":
		return bsonToString(value)
	case "geohash":
		return geohashToString(value)
	case "int64":
		return int64ToString(value)
	case "float64":
		return float64ToString(value)
	case "hex":
		return hexToString(value)
	case "bin":
		return binToString(value)
	case "raw":
	default:
	}

	return string(value)
}

// Converts data from bson type to a string
func bsonToString(value []byte) string {
	var dst interface{}
	err := bson.Unmarshal(value, &dst)
	if err != nil {
		return "Error converting!"
	}

	return fmt.Sprintf("%#v", dst)
}

// Converts data from geohash type to a string
func geohashToString(value []byte) string {
	position := geohash.Decode(string(value))

	return fmt.Sprintf("lat: %f lng: %f", position.Center().Lat(), position.Center().Lng())
}

// Converts data from int64 type to a string
func int64ToString(value []byte) string {
	return fmt.Sprintf("%d", binary.BigEndian.Uint64(value))
}

// Converts data from float64 type to a string
func float64ToString(value []byte) string {
	return fmt.Sprintf("%f", math.Float64frombits(
		binary.LittleEndian.Uint64(value),
	))
}

// Converts data from hex type to a string
func hexToString(value []byte) string {
	return fmt.Sprintf("%s", hex.Dump(value))
}

// Converts data from bin type to a string
func binToString(value []byte) string {
    return fmt.Sprintf("%08b", value)
}
