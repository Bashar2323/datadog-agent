// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build oracle

package oracle

import (
	"fmt"
	"strings"

	"github.com/DataDog/datadog-agent/pkg/util/log"
	"github.com/jmoiron/sqlx"
)

func selectWrapper[T any](c *Check, s T, sql string, binds ...interface{}) error {
	err := c.db.Select(s, sql, binds...)
	err = handleError(c, &c.db, err)
	return err
}

func getWrapper[T any](c *Check, s T, sql string, binds ...interface{}) error {
	err := c.db.Get(s, sql, binds...)
	err = handleError(c, &c.db, err)
	return err
}

func handleError(c *Check, db **sqlx.DB, err error) error {
	if err == nil {
		return err
	}
	isPrivilegeError, err := handlePrivilegeError(c, err)
	if err != nil && isPrivilegeError {
		return err
	}
	reconnectOnConnectionError(c, db, err)
	return err
}

func handlePrivilegeError(c *Check, err error) (bool, error) {
	var isPrivilegeError bool
	if err == nil {
		return isPrivilegeError, err
	}
	if !strings.Contains(err.Error(), "ORA-00942") {
		return isPrivilegeError, err
	}
	var link string
	if c.isRDS {
		link = "https://docs.datadoghq.com/database_monitoring/setup_oracle/rds/#grant-permissions"
	} else if c.isOracleCloud {
		link = "https://docs.datadoghq.com/database_monitoring/setup_oracle/autonomous_database/#grant-permissions"
	} else {
		link = "https://docs.datadoghq.com/database_monitoring/setup_oracle/selfhosted/#grant-permissions"
	}
	isPrivilegeError = true
	return isPrivilegeError, fmt.Errorf("Some privileges are missing. Execute the `grant` commands from %s . Error: %w", link, err)
}

func isConnectionError(err error) bool {
	if err == nil {
		return false
	}
	connectionErrors := []string{"ORA-00028", "ORA-01012", "ORA-06413", "database is closed"}
	for _, e := range connectionErrors {
		if strings.Contains(err.Error(), e) {
			return true
		}
	}
	return false
}

func reconnectOnConnectionError(c *Check, db **sqlx.DB, err error) {
	if !isConnectionError(err) {
		return
	}
	log.Debugf("Reconnecting")
	if *db != nil {
		closeDatabase(c, *db)
	}
	*db, err = c.Connect()
	if err != nil {
		log.Errorf("failed to reconnect %s", err)
		closeDatabase(c, *db)
	}
}