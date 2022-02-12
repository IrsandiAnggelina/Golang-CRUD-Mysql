package main

import (
	"fmt"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//membuat type mahasiswa dengan struktur
type mahasiswa struct{
	Nim string
	Nama string
	Prodi string
	smt int
}

//membuat type respose dengan strutur
type response struct{
	Status bool
	Pesan string
	Data []mahasiswa
}

//membuat fungsi koneksi sql
//sintax -> sql.open("mysql", "user:password@tcp(host:port)/nama_Database")
//karena bawaan xampp password kosong, jadi dikosongkan saja

func koneksi() (*sql.DB, error){
	db, salah := sql.open("mysql", "root:@tcp(127.0.0.1:3306)/cloud_udb")
	if salah != nil {
		return nil, salah
	} return db, nil
}

//fungsi tampil data
func tampil(pesan string) response{
	db, salah := koneksi()
	if salah != nil{
		return response{
			Status:false,
			Pesan: "Gagal Koneksi: "+salah.Error(),
			Data[]mahasiswa{},
		}
	}
	defer db.Close()
	dataMhs, salah := db.Query("select * from mahasiswa")
	if salah != nil{
		return response{
			Status: false,
			Pesan: "Gagal Query:"+salah.Error(),
			Data:[]mahasiswa,
		}
	}
	defer dataMhs.Close()
	var hasil []mahasiswa
	for dataMhs.Next() {
		var mhs = mahasiswa{}
		var salah = dataMhs.Scan(&mhs,Nim, &mhs.Nama, &mhs.Prodi, &mhs.Smt)
		if salah !=nil {
			return response{
				Status: false,
				Pesan: "Gagal Baca:"+salah.Error(),
				Data:[]mahasiswa,
		}
	}
	hasil = append(hasil, mhs)
	}
	if salah !=nil{
		return response{
			Status: false,
			Pesan: "Kesalahan:"+salah.Error(),
			Data:[]mahasiswa(),
		}
	}
	return response{
		Status:true,
		Pesan: pesan,
		Data:hasil,
	}
}

//fungsi tampil berdasarkan nim
func getMhs(nim string) response{
	db, salah := koneksi()
	if salah != nil{
		return response{
			Status:false,
			Pesan: "Gagal Koneksi: "+salah.Error(),
			Data[]mahasiswa{},
		}
	}
	defer db.Close()
	dataMhs, salah := db.Query("select * from mahasiswa where nim=?", nim)
	if salah != nil{
		return response{
			Status: false,
			Pesan: "Gagal Query:"+salah.Error(),
			Data:[]mahasiswa{},
		}
	}
	defer dataMhs.Close()
	var hasil []mahasiswa
	for dataMhs.Next() {
		var mhs = mahasiswa{}
		var salah = dataMhs.Scan(&mhs,Nim, &mhs.Nama, &mhs.Prodi, &mhs.Smt)
		if salah !=nil {
			return response{
				Status: false,
				Pesan: "Gagal Baca:"+salah.Error(),
				Data:[]mahasiswa{},
		}
	}
	hasil = append(hasil, mhs)
	}
	if salah !=nil{
		return response{
			Status: false,
			Pesan: "Kesalahan:"+salah.Error(),
			Data:[]mahasiswa{},
		}
	}
	return response{
		Status:true,
		Pesan:"Berhasil Tampil",
		Data:hasil,
	}
}

//fungsi tambah data
func tambah (nim string, nama string, prodi string, smt string) response{
	db, salah := koneksi()
	if salah != nil{
		return response{
			Status:false,
			Pesan: "Gagal Koneksi:"+salah.Error(),
			Data:[]mahasiswa{},
		}
	}
	defer db.Close()
	_, salah = db.Exec("insert into Mahasiswa value (?, ?, ?, ?)", nim, nama, prodi, smt)
	if salah != nil{
		return response{
			Status: false,
			Pesan: "Gagal Qeury Insert: "+salah.Error(),
			Data:[]mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan: "Berhasil Tambah",
		Data:[]mahasiswa{},
	}
}

//fungsi ubah data
func ubah (nim string, nama string, prodi string, smt string) response{
	db, salah := koneksi()
	if salah != nil {
		return response{
			Status: false,
			Pesan: "Gagal Koneksi:"+salah.Error(),
			Data:[]mahasiswa,
		}
	}
	defer db.Close()
	_, salah = db.Exec("update mahasiswa set nama=?, prodi=?, smt=? where nim=?", nama, prodi, set, nim)
	if salah != nil{
		return response{
			Status:false,
			Pesan:"Gagal Query Update:"+salah.Error(),
			Data:[]mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:"Berhasil Ubah",
		Data:[]mahasiswa{}
	}
}

//fungsi hapus data
func hapus (nim string) response{
	db, salah := koneksi()
	if salah != nil {
		return response{
			Status: false,
			Pesan: "Gagal Koneksi:"+salah.Error(),
			Data:[]mahasiswa,
		}
	}
	defer db.Close()
	_, salah = db.Exec("delete from mahasiswa where nim=?", nim)
	if salah != nil{
		return response{
			Status:false,
			Pesan:"Gagal Query Delete:"+salah.Error(),
			Data:[]mahasiswa{},
		}
	}
	return response{
		Status: true,
		Pesan:"Berhasil Hapus",
		Data:[]mahasiswa{}
	}
}

func kontroler(w http.ResponseWriter, r http.Request){
	var tampilHtml, salahTampil = template.ParseFiles("template/tampil.html")
	if salahTampil != nil{
		fmt.Println(salahTampil.Error())
		return
	}
	var tambahHtml, salahTambah = template.ParseFiles("template/tambah.html")
	if salahTambah != nil{
		fmt.Println(salahTambah.Error())
		return
	}
	var ubahHtml, salahUbah = template.ParseFiles("template/ubah.html")
	if salahUbah != nil{
		fmt.Println(salahUbah.Error())
		return
	}
	var hapusHtml, salahHapus = template.ParseFiles("template/hapus.html")
	if salahHapus != nil{
		fmt.Println(salahHapus.Error())
		return
	}
}

switch r.Method {
	case "GET":
		aksi:= r.URL.Query()["aksi"]
		if (len(aksi)==0){
			tampilHtml.Execute(w, tampil("Berhasil Tampil"))
		} else if aksi[0] == "tambah"{
			tambahHtml.Execute(w, nil)
		} else if aksi[0] == "ubah" {
			nim:= r.URL.Query()["nim"]
			ubahHtml.Execute(w, getMhs(nim[0]))
		} else if aksi[0] == "hapus"{
			nim:= r.URL.Query()["nim"]
			hapusHtml.Execute(w, getMhs(nim[0]))
		} else{
			tampilHtml.Execute(w, tampil("Berhasil Tampil"))
		}
	case "POST"
		var salah = r.ParseForm():
		if salah != nil{
			fmt.Fprintln(w, "Kesalahan:", salah)
			return
		}
		var nim = r.FormValue("nim")
		var nama = r.FormValue("nama")
		var prodi = r.FormValue("prodi")
		var smt = r,FormValue("smt")
		var aksi = r.URL.Path
		if (aksi =="/tambah"){
			var hasil = tambah(nim, nama, prodi, smt)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		}else if (aksi =="/ubah"){
			var hasil = tambah(nim, nama, prodi, smt)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else if (aksi =="/hapus"){
			var hasil = hapus(nim)
			tampilHtml.Execute(w, tampil(hasil.Pesan))
		} else {
			tampilHtml.Execute(w, tampil("Berhasil Tampil"))
		}
	default:
		fmt.Println(w, "Maaf, Method yang didukung hanya GET dan POST")
	}
}

func main(){
	http.HandleFunc("/", kontroler)
	fmt.Println("Server berjalan di port 8080...")
	http.ListenAndServe(":8080", nil):
}
