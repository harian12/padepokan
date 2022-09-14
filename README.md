# padepokan
menggunakan golang versi 1.19

silahkan buat file **.env** terlebih dahulu detail isi file **.env** ada di file **.env.example** dan lengkapi data koneksi ke DB

jalankan command **go mod vendor** untuk vendoring 3rd party library yang digunakan:
> go mod vendor

jalankan perintah dibawah untuk init testing:
> cd test
> go test

kemudian jalankan perintah berikut untuk membuat table ke dalam DB:
> go run main.go db:migrate

dan jalankan perintah berikut untuk menjalankan sistem:
> go run main.go



dokumentasi API ada di folder postman silahkan import ke postman atau yg lainnya