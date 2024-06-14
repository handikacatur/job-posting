package gorm

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/handikacatur/jobs-api/internal/job/model/entity"
	"github.com/handikacatur/jobs-api/internal/job/model/request"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestJobRepository_GetJobs(t *testing.T) {
	jobID := uuid.NewString()

	type mockExec struct {
		data []entity.JobToCompany
		err  error
	}
	type fields struct {
		db *gorm.DB
	}
	type args struct {
		ctx context.Context
		req request.GetJobsRequest
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     []entity.JobToCompany
		wantErr  bool
		mockExec mockExec
	}{
		{
			name: "best case",
			args: args{
				ctx: context.Background(),
				req: request.GetJobsRequest{
					Keyword:     "Backend",
					CompanyName: "",
				},
			},
			want: []entity.JobToCompany{
				{
					JobID:       jobID,
					Company:     "Google",
					Title:       "Backend Engineer",
					Description: "Lorem ipsum dolor sit amet!",
				},
			},
			mockExec: mockExec{
				data: []entity.JobToCompany{
					{
						JobID:       jobID,
						Company:     "Google",
						Title:       "Backend Engineer",
						Description: "Lorem ipsum dolor sit amet!",
					},
				},
			},
		},
		{
			name: "error case",
			args: args{
				ctx: context.Background(),
				req: request.GetJobsRequest{
					Keyword:     "Backend Engineer",
					CompanyName: "",
				},
			},
			wantErr: true,
			mockExec: mockExec{
				err: errors.New("error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, m, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			gormdb, err := gorm.Open(postgres.New(postgres.Config{
				Conn: db,
			}), &gorm.Config{
				Logger: logger.Default.LogMode(logger.Info),
			})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			r := NewJobRepository(gormdb)

			q := `SELECT jobs.id as job_id, companies.name as company, jobs.title as title, jobs.description as description, jobs.created_at as created_at FROM \"jobs\" left join companies on companies.id = jobs.company_id WHERE (jobs.title @@ to_tsquery($1) OR jobs.description @@ to_tsquery($2))`

			mockExpectExec := m.ExpectQuery(q).WithArgs(tt.args.req.Keyword, tt.args.req.Keyword)
			if tt.mockExec.err != nil {
				mockExpectExec.WillReturnError(tt.mockExec.err)
			} else {
				row := sqlmock.NewRows([]string{"job_id", "company", "title",
					"description", "created_at"})

				rowData := tt.mockExec.data
				row.AddRow(rowData[0].JobID, rowData[0].Company, rowData[0].Title,
					rowData[0].Description, rowData[0].CreatedAt)

				mockExpectExec.WillReturnRows(row)
			}

			got, err := r.GetJobs(tt.args.ctx, tt.args.req)
			assert.Equal(t, tt.mockExec.err, err)
			assert.Equal(t, tt.mockExec.data, got)
		})
	}
}
