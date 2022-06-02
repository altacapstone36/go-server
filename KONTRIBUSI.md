# Cara Kontribusi
untuk berkontribusi dalam projek ini, silahkan clone repositori dan buat branch baru dengan nama *branch* `feature/<fitur_yang_dikerjakan>`. lakukan *pull request* ke branch `dev`, dan tunggu hingga proses *Code Integration* selesai untuk melakukan *merging*. setelah kode dari fitur baru di *merge* ke *branch* `dev`, *update* kode pada branch `dev` di repositori lokal anda.

## Struktur Direktori
- `/config` berisi berkas konfigurasi yang dimuat oleh [viper](https://github.com/) dan ditulis dalam ekstensi `yaml`. setelah repositori di*clone*, ubah nama `config.example.yaml` menjadi `config.yaml` sebelum aplikasi dijalankan.
- `/docs` berisi berkas dokumentasi dan konfigurasi `swagger` yang di*generate* oleh [swaggo](https://github.com/swaggo).
- `/internal` berisi semua kode program.
- `/internal/core/entity/models` berisi struct untuk model yang berhubungan dengan basis data.
- `/internal/core/entity/request` berisi struct untuk model dengan field yang berbeda dari `models` yang digunakan untuk request data.
- `/internal/core/entity/response` berisi struct untuk model dengan field yang berbeda dari `models` yang digunakan untuk merespon data.
- `/internal/core/repository` berisi *interface* yang digunakan untuk mengakses basis data, juga digunakan pada `service`.
- `/internal/core/service` berisi logika dari aplikasi.
- `/internal/framework/database` berisi konfigurasi basis data.
- `/internal/framework/database/seeds` berisi *dummy data* yang secara otomatis dimasukkan saat pertama kali basis data dibuat.
- `/internal/framework/repository` berisi implementasi dari *interface* pada direktori `/internal/core/repository`.
- `/internal/routes` berisi daftar registrasi *endpoint* dari API yang dibuat.
- `/internal/transport/controller` berisi *handler* dari API yang dibuat.
- `/internal/transport/middleware` berisi *middleware* dari API yang dibuat.
- `/internal/utils` berisi fungsi tambahan untuk membantu konfigurasi aplikasi.
- `/internal/utils/errors` berisi kustom error untuk membantu pengecekan eror yang mengembalikan nilai *http request* dan pesan eror.

## Aturan Membuat Fitur Baru
untuk membuat fitur baru, dapat dilakukan dengan menambahkan interface didalam direktori `/internal/core/repository`. setelah itu tambahkan didalam *struct* `Repository` di berkas `Repository.go` nama dari *interface* yang dibuat, agar mempermudah saat melakukan *dependency injection* untuk akses ke dalam fungsi pada *interface*. contoh jika menambahkan `UserRepository`.
```go
type Repository struct {
	Auth AuthRepository
	User UserRepository
}
```
setelah itu buat *struct* dan berkas baru untuk fitur baru yang dibuat, lalu tambahkan ke *struct* `Service` pada berkas `Service.go` didalam direktori `/internal/core/service`. contohnya dengan `UserService`.
```go
// berkas di internal/core/service/UserService.go
type UserService struct {
	repo repository.AuthRepository // import dari internal/core/repository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// berkas di internal/core/service/Service.go
type Service struct {
	Auth *AuthService
	User *UserService
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(r.Auth),
		User: NewUserService(r.User),
	}
}
```
buat juga *struct* didalam direktori `/internal/framework/repository/` untuk mengimplementasikan *interface* yang dibuat sebelumnya. lalu masukkan fungsi ke dalam berkas `Repository.go` di direktori yang sama. contohnya dengan `UserRepository`.
```go
// berkas di internal/framework/repository/UserRepository.go
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

// berkas di internal/framework/repository/Repository.go
func NewRepository(db *gorm.DB) *repository.Repository {
	// kembalikan struct dari internal/core/repository/Repository.go
	return &repository.Repository{
		User: NewUserRepository(db),
	}
}
```
buat *struct* didalam direktori `/internal/framework/transport/controller` untuk handler dari fitur yang dibuat. lalu masukkan fungsi ke dalam berkas `Controller.go` di direktori yang sama. contohnya dengan `UserController`.
```go
// berkas di internal/framework/transport/controller/UserController.go
type UserController struct {
	srv *service.UserService
}

func NewAuthController(srv *service.UserService) *UserController {
	return &UserController{srv}
}


// berkas di internal/framework/repository/Repository.go
type Controller struct {
	Auth *AuthController
	User *UserController
}

// import Service dari internal/core/service/Service.go
func NewController(srv *service.Service) *Controller {
	return &Controller{
		Auth: NewAuthController(srv.Auth),
		User: NewUserController(srv.User),
	}
}
```
terakhir untuk registrasi endpoint bisa dengan menambahkan fungsi baru di `/internal/framework/routes`. contoh.
```go
func NewUserRoutes(e *echo.Group, ctrl *controller.UserController, middleware ...echo.MiddlewareFunc) {
	newAdminUserRoutes(e.Group("/admin"), ctrl, middleware...)
}

// fungsi private untuk memisahkan endpoint berdasarkan role
func newAdminUserRoutes(e *echo.Group, ctrl *controller.UserController, middleware ...echo.MiddlewareFunc) {
	middleware = append(middleware, mw.AdminPermission) // mw (import alias) dari internal/framework/transport/middleware
	group := e.Group("/user", middleware...)
	group.GET("", ctrl.ReadUser)
	group.POST("", ctrl.CreateUser)
	group.GET("/:id", ctrl.ReadUserByID)
	group.PUT("/:id/update", ctrl.UpdateUser)
	group.DELETE("/:id/delete", ctrl.DeleteUser)
}
```
