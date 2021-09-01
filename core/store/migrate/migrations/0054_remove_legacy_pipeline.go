package migrations

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/pressly/goose/v3"
)

const up54 = `
ALTER TABLE log_broadcasts DROP COLUMN job_id;
DROP TABLE service_agreements;
DROP TABLE eth_task_run_txes;
DROP TABLE task_runs;
DROP TABLE task_specs;
DROP TABLE flux_monitor_round_stats;
DROP TABLE job_runs;
DROP TABLE job_spec_errors;
DROP TABLE initiators;
DROP TABLE job_specs;

DROP TABLE run_results;
DROP TABLE run_requests;
DROP TABLE sync_events;

ALTER TABLE log_broadcasts RENAME COLUMN job_id_v2 TO job_id;
ALTER TABLE job_spec_errors_v2 RENAME TO job_spec_errors;
`

func init() {
	goose.AddMigration(Up54, Down54)
}

func Up54(tx *sql.Tx) error {
	if err := checkNoLegacyJobs(tx); err != nil {
		return err
	}
	if _, err := tx.Exec(up54); err != nil {
		return err
	}
	return nil
}

func Down54(tx *sql.Tx) error {
	return errors.New("irreversible migration")
}

func checkNoLegacyJobs(tx *sql.Tx) error {
	var count int
	if err := tx.QueryRow(`SELECT COUNT(*) FROM job_specs WHERE deleted_at IS NULL`).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return errors.Errorf("cannot migrate; this release removes support for legacy job specs but there are still %d in the database. Please migrate these job specs to the V2 pipeline and make sure all V1 job_specs are deleted or archived, then run the migration again. Migration instructions found here: https://docs.chain.link/docs/jobs/migration-v1-v2/", count)
	}
	return nil
}
