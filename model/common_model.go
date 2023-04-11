package model

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Error  interface{} `json:"errors"`
}

type Pagination struct {
	TotalPage   uint `json:"total_page"`
	CurrentPage uint `json:"current_page"`
}

type GeneralError struct {
	General string `json:"general"`
}

type NotValidImage struct {
	Image string `json:"image"`
}

var ErrorCodeMap = map[string]interface{}{
	"AGENT_BLOCKED":                "Akun ini telah terblokir. Silahkan hubungi customer service",
	"AGENT_INACTIVE":               "Akun ini tidak aktif. Silahkan hubungi customer service",
	"AGENT_NOT_FOUND":              "Data user tidak ditemukan",
	"EXTENSION_NOT_ALLOWED":        "Hanya dapat mengirim ekstensi gambar .jpg, .jpeg, .png",
	"NOT_BLANK":                    "tidak boleh kosong",
	"MUST_NUMBER":                  "hanya menerima angka",
	"MUST_STRING":                  "hanya menerima karakter ",
	"NOT_VALID":                    "tidak valid",
	"NOT_FOUND":                    "tidak ditemukan",
	"MIN":                          "minimal harus [x] karakter",
	"MAX":                          "maksimal harus [x] karakter",
	"ALREADY_EXIST":                "telah ada",
	"AUTHENTICATION_FAILURE":       "Login gagal",
	"UNAUTHORIZATION":              "Data login tidak valid",
	"UNAUTHORIZED":                 "Data login tidak valid",
	"STOCK_EXCEEDES_LIMIT":         "Pesanan melebihi batas stok",
	"USERNAME_OR_PASSWORD_INVALID": "Username atau password tidak valid",
	"PRODUCT_NOT_FOUND":            "Product ini tidak ditemukan",
	"CART_NOT_FOUND":               "Keranjang ini tidak ditemukan",
	"CANNOT_BE_NEGATIVE_NUMBER":    "Data tidak boleh bernilai negatif",
	"ORDER_NOT_FOUND":              "Data pesanan tidak ada",
	"SALES_NOT_FOUND":              "Data sales tidak ada",
	"PASSWORD_INVALID":             "Password lama tidak valid",
	"OLD_PASSWORD_CANNOT_BE_MATCH": "Password baru tidak boleh sama dengan yang lama",
	"CONFIRM_PASSWORD_NOT_MATCH":   "Konfirmasi password tidak sama",
	"OUTLET_NOT_FOUND":             "Data outlet tidak ditemukan",
	"NO_CITIES_FOUND":              "Data kota tidak ditemukan",
	"NO_DISTRICTS_FOUND":           "Data kecamatan tidak ditemukan",
	"NO_OUTLET_CATEGORIES_FOUND":   "Data kategori outlet tidak ditemukan",
	"OWNER_IS_UNDERAGE":            "Pemilik outlet belum berumur 17 tahun",
	"OUTLET_NOT_VALID":             "Perubahan outlet belum di setujui",
	"OUTLET_NAME":                  "Nama outlet",
	"OWNER_NAME":                   "Nama pemilik",
	"PHONE_NUMBER":                 "Nomor telepon",
	"ADDRESS":                      "Alamat",
	"COORDINATE":                   "Koordinat",
}
