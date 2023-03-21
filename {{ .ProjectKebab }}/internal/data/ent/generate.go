// Package ent contains the ent schema definition and the entgo.io/ent/cmd/ent
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./schema
