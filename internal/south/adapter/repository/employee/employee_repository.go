package employee

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/sync/singleflight"

	"github.com/andys920605/hr-system/internal/domain/model/employee"
	"github.com/andys920605/hr-system/internal/south/adapter/repository/dao/employee/mysql"
	"github.com/andys920605/hr-system/internal/south/port/repository"
	"github.com/andys920605/hr-system/pkg/errors"
)

var _ repository.EmployeeRepository = (*EmployeeRepository)(nil)

type EmployeeRepository struct {
	sg          singleflight.Group
	employeeDao *mysql.EmployeeDao
	redisClient *redis.ClusterClient
}

func NewEmployeeRepository(
	employeeDao *mysql.EmployeeDao,
	redisClient *redis.ClusterClient,
) *EmployeeRepository {
	return &EmployeeRepository{
		sg:          singleflight.Group{},
		employeeDao: employeeDao,
		redisClient: redisClient,
	}
}

func (r *EmployeeRepository) Save(ctx context.Context, employee *employee.Employee) error {
	if employee.AggregateRoot.IsNew() {
		if err := r.employeeDao.Create(ctx, employee); err != nil {
			return errors.Wrap(err, "create")
		}
	}

	// todo: employee.IsDomainEventsNotEmpty can trigger events

	return nil
}

func (r *EmployeeRepository) GetByID(ctx context.Context, id int64) (*employee.Employee, error) {
	cacheKey := fmt.Sprintf("hr-system:employee:{%d}", id)

	data, err := r.redisClient.Get(ctx, cacheKey).Bytes()
	if err == nil {
		var emp employee.Employee
		if err := json.Unmarshal(data, &emp); err == nil {
			return &emp, nil
		}
	}

	v, err, _ := r.sg.Do(strconv.FormatInt(id, 10), func() (interface{}, error) {
		emp, err := r.employeeDao.GetByID(ctx, id)
		if err != nil {
			return nil, err
		}

		jsonData, err := json.Marshal(emp)
		if err == nil {
			_ = r.redisClient.Set(ctx, cacheKey, jsonData, 3*time.Minute).Err()
		}

		return emp, nil
	})

	if err != nil {
		return nil, err
	}

	return v.(*employee.Employee), nil
}
