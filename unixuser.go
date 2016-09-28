// This package implements the "bluemix-spark-utils system-check disk-full" subcommand.
package unixuser

import (
  "os/user"
)

func Lookup(username string) (*user.User, error) {
  return user.Lookup(username)
}
