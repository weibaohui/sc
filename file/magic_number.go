package file

// MagicType 魔数类型
type MagicType struct {
	Name   string // 类型
	Magic  string // 魔数
	Suffix string // 后缀
	Skip   bool   // 跳过
}

// Types 魔数类型
var Types = []MagicType{
	{Name: "binary", Magic: "CFFAEDFE", Skip: true},
	{Name: "linux binary", Magic: "7F454C46", Skip: true},
	{Name: "ZIP", Magic: "504B0304", Suffix: "zip;jar", Skip: true},
	{Name: "MS Office/OLE2", Magic: "D0CF11E0", Suffix: "doc;xls;dot;ppt;xla;ppa;pps;pot;msi;sdw;db", Skip: true},
	{Name: "GZ Archive", Magic: "1F8B08", Suffix: "gz;tgz", Skip: true},
	{Name: "RUST", Magic: "72757374", Suffix: "rmeta", Skip: true},
	{Name: "RUST", Magic: "213C6172", Suffix: "rlib", Skip: true},
	{Name: "dia", Magic: "444941470", Suffix: "dia", Skip: true},
	{Name: "rust bin", Magic: "52534943", Suffix: "bin", Skip: true},
	{Name: "dia", Magic: "444941470", Suffix: "dia", Skip: true},
	{Name: "font ttf", Magic: "0001000000", Suffix: "ttf", Skip: true},
	{Name: "GZ Archive", Magic: "1F8B08", Suffix: "gz;tgz", Skip: true},
	{Name: "BMP", Magic: "424D", Skip: true},
	{Name: "CAD", Magic: "41433130", Suffix: "dwg", Skip: true},
	{Name: "PSD", Magic: "38425053", Skip: true},
	{Name: "RTF", Magic: "7B5C727466", Skip: true},
	{Name: "XML", Magic: "3C3F786D6C"},
	{Name: "HTML", Magic: "68746D6C3E"},
	{Name: "EML", Magic: "44656C69766572792D646174653A", Skip: true},
	{Name: "DBX", Magic: "CFAD12FEC5FD746F", Skip: true},
	{Name: "PST", Magic: "2142444E", Skip: true},
	{Name: "XLS", Magic: "D0CF11E0", Skip: true},
	{Name: "MDB", Magic: "5374616E64617264204A", Skip: true},
	{Name: "WPD", Magic: "FF575043", Skip: true},
	{Name: "EPS", Magic: "252150532D41646F6265", Skip: true},
	{Name: "PDF", Magic: "255044462D312E", Skip: true},
	{Name: "QDF", Magic: "AC9EBD8F", Skip: true},
	{Name: "PWL", Magic: "E3828596", Skip: true},
	{Name: "RAR", Magic: "52617221", Skip: true},
	{Name: "WAV", Magic: "57415645", Skip: true},
	{Name: "AVI", Magic: "41564920", Skip: true},
	{Name: "RAM", Magic: "2E7261FD", Skip: true},
	{Name: "RM", Magic: "2E524D46", Skip: true},
	{Name: "MPG", Magic: "000001BA", Skip: true},
	{Name: "MOV", Magic: "6D6F6F76", Skip: true},
	{Name: "ASF", Magic: "3026B2758E66CF11", Skip: true},
	{Name: "MID", Magic: "4D546864", Skip: true},
	{Name: "TIFF", Magic: "49492A00", Skip: true},
	{Name: "GIF", Magic: "47494638", Suffix: "gif", Skip: true},
	{Name: "PC Paintbrush", Magic: "0A050108", Suffix: "pcx", Skip: true},
	{Name: "PC", Suffix: "pcx", Magic: "0A050108", Skip: true},
	{Name: "Graphics Metafile", Suffix: "wmf", Magic: "D7CDC69A", Skip: true},
	{Name: "Graphics Metafile", Suffix: "wmf", Magic: "01000900", Skip: true},
	{Name: "Graphics Metafile", Suffix: "wmf", Magic: "02000900", Skip: true},
	{Name: "Enhanced Metafile", Suffix: "emf", Magic: "0A050108", Skip: true},
	{Name: "Java-Class", Suffix: "class", Magic: "CAFEBASE", Skip: true},
	{Name: "PC", Suffix: "pcx", Magic: "0A050108", Skip: true},
	{Name: "PC", Suffix: "pcx", Magic: "0A050108", Skip: true},
	{Name: "AOL ART", Magic: "4A47040E", Suffix: "art", Skip: true},
	{Name: "AOL ART", Magic: "4A47030E", Suffix: "art", Skip: true},
	{Name: "TIFF", Magic: "49492A00", Suffix: "tif;tiff", Skip: true},
	{Name: "TIFF", Magic: "4D4D002A", Suffix: "tif;tiff", Skip: true},
	{Name: "PNG", Magic: "89504E47", Suffix: "png", Skip: true},
	{Name: "JEPG", Magic: "FFD8FF", Suffix: "jpg;jpeg", Skip: true},
}
