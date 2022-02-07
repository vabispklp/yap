package mock

//go:generate ../../bin/mockgen -destination=service/shortener.go -package=service_mock github.com/vabispklp/yap/api/rest/handlers ShortenerExpected
//go:generate ../../bin/mockgen -destination=storage/storage.go -package=storage_mock github.com/vabispklp/yap/internal/app/storage StorageExpected
