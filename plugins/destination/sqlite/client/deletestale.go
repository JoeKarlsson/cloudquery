package client

import (
	"context"
	"strings"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
)

func (c *Client) deleteStale(_ context.Context, table *schema.Table, source string, syncTime time.Time) error {
	var sb strings.Builder
	sb.WriteString("delete from ")
	sb.WriteString(`"` + table.Name + `"`)
	sb.WriteString(" where ")
	sb.WriteString(`"` + schema.CqSourceNameColumn.Name + `"`)
	sb.WriteString(" = $1 and datetime(")
	sb.WriteString(schema.CqSyncTimeColumn.Name)
	sb.WriteString(") < datetime($2)")
	sql := sb.String()
	if _, err := c.db.Exec(sql, source, syncTime); err != nil {
		return err
	}
	return nil
}
