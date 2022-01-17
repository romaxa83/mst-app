package config

import (
	"github.com/romaxa83/mst-app/pkg/tracing"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	type env struct {
		appEnv         string
		httpHost       string
		httpPort       string
		dbHost         string
		dbPort         string
		dbUser         string
		dbPassword     string
		dbName         string
		dbSSLMode      string
		minioEndpoint  string
		minioBucket    string
		minioAccessKey string
		minioSecretKey string
	}

	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		os.Setenv("APP_ENV", env.appEnv)
		os.Setenv("HTTP_HOST", env.httpHost)
		os.Setenv("HTTP_PORT", env.httpPort)
		os.Setenv("DB_HOST", env.dbHost)
		os.Setenv("DB_PORT", env.dbPort)
		os.Setenv("DB_USER", env.dbUser)
		os.Setenv("DB_PASSWORD", env.dbPassword)
		os.Setenv("DB_NAME", env.dbName)
		os.Setenv("DB_SSL_MODE", env.dbSSLMode)
		os.Setenv("MINIO_ENDPOINT", env.minioEndpoint)
		os.Setenv("MINIO_BUCKET", env.minioBucket)
		os.Setenv("MINIO_ACCESS_KEY", env.minioAccessKey)
		os.Setenv("MINIO_SECRET_KEY", env.minioSecretKey)
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				path: "fixtures",
				env: env{
					appEnv:         "local",
					httpHost:       "http://127.0.0.1",
					httpPort:       "8880",
					dbHost:         "192.168.141.1",
					dbPort:         "54321",
					dbUser:         "root",
					dbPassword:     "root",
					dbName:         "db",
					dbSSLMode:      "disable",
					minioEndpoint:  "192.168.141.1:9000",
					minioBucket:    "lib",
					minioAccessKey: "admin",
					minioSecretKey: "password",
				},
			},
			want: &Config{
				Environment: "local",
				HTTP: HTTPConfig{
					Host:               "http://127.0.0.1",
					Port:               "8880",
					ReadTimeout:        time.Second * 20,
					WriteTimeout:       time.Second * 10,
					MaxHeaderMegabytes: 1,
				},
				Limiter: LimiterConfig{
					RPS:   10,
					Burst: 20,
					TTL:   time.Minute * 10,
				},
				Postgres: PostgresConfig{
					Host:     "192.168.141.1",
					Port:     "54321",
					Username: "root",
					Password: "root",
					DBName:   "db",
					SSLMode:  "disable",
				},
				FileStorage: FileStorageConfig{
					Endpoint:  "192.168.141.1:9000",
					Bucket:    "lib",
					AccessKey: "admin",
					SecretKey: "password",
				},
				Locale: Locale{
					Default: "en",
				},
				CacheTTL: time.Second * 10,
				Jaeger: &tracing.Config{
					ServiceName: "library_service",
					HostPort:    "localhost:6831",
					Enable:      true,
					LogSpans:    false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setEnv(tt.args.env)

			got, err := Init(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() got = %v, want %v", got, tt.want)
			}
		})
	}
}
