package middleware

import (
	"net/http"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/repository"
)

func HasPermission(code string, permissionRepository repository.PermissionRepository, accountPermissionRepository repository.AccountPermissionRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			subject, ok := GetSubject(r.Context())
			if !ok {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			permission, err := permissionRepository.FindByCode(r.Context(), code)
			if err != nil {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			_, err = accountPermissionRepository.FindByAccountIdAndPermissionId(r.Context(), subject, permission.GetId())
			if err != nil {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
